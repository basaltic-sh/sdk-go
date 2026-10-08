package basaltic_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/iam"
	"github.com/basaltic-sh/sdk-go/workspace"
)

type workspaceTransport func(*http.Request) (*http.Response, error)

func (f workspaceTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// The two policy catalogs deliberately share their paths. Their service hosts
// must remain distinct, including when no region is configured.
func TestWorkspaceAndAccountIAMStaySeparate(t *testing.T) {
	const accountID = "11111111-1111-4111-8111-111111111111"
	const roleID = "22222222-2222-4222-8222-222222222222"
	const policyID = "33333333-3333-4333-8333-333333333333"
	const userID = "44444444-4444-4444-8444-444444444444"
	var calls []string
	cfg, err := basaltic.NewConfig(context.Background(),
		basaltic.WithAccessToken("personal-token"),
		basaltic.WithAccountID("source-account"),
		basaltic.WithRegion(""),
		basaltic.WithDomain("contract.invalid"),
		basaltic.WithHTTPClient(&http.Client{Transport: workspaceTransport(func(r *http.Request) (*http.Response, error) {
			key := r.Method + " " + r.URL.Host + r.URL.Path
			calls = append(calls, key)
			if r.Header.Get("Authorization") != "Bearer personal-token" {
				t.Fatalf("missing caller credential on %s", key)
			}
			body := `{}`
			switch key {
			case "GET workspace.contract.invalid/v1/policies":
				body = `{"policies":[{"id":"org-policy","name":"OrganizationReader"}]}`
			case "GET iam.contract.invalid/v1/policies":
				body = `{"policies":[{"id":"account-policy","name":"AccountReader"}]}`
			case "POST workspace.contract.invalid/v1/roles/" + roleID + "/policies":
				var payload map[string]string
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || len(payload) != 1 || payload["policy_id"] != policyID {
					t.Fatalf("organization grant must preserve explicit policy UUID: %v, %v", payload, err)
				}
			case "POST workspace.contract.invalid/v1/accounts/" + accountID + "/role-assignments":
				var payload map[string]string
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload["principal_type"] != "user" || payload["principal_id"] != userID || payload["role_id"] != roleID {
					t.Fatalf("incorrect account assignment: %v, %v", payload, err)
				}
				body = `{"role_assignment":{"id":"assignment","account_id":"` + accountID + `","role_id":"` + roleID + `","principal_id":"` + userID + `","principal_type":"user"}}`
			case "POST iam.contract.invalid/v1/assume-role":
				body = `{"access_token":"role-token","account_id":"` + accountID + `","account_handle":"target-account","role_id":"` + roleID + `"}`
			default:
				t.Fatalf("unexpected request: %s", key)
			}
			return &http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
		})}),
		basaltic.WithoutRetry(),
	)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	org := workspace.New(cfg)
	account := iam.New(cfg)
	orgPolicies, err := org.ListPolicies(ctx, nil)
	if err != nil || len(orgPolicies.Items) != 1 || orgPolicies.Items[0].ID != "org-policy" {
		t.Fatalf("organization catalog: %+v, %v", orgPolicies, err)
	}
	accountPolicies, err := account.ListPolicies(ctx, nil)
	if err != nil || len(accountPolicies.Items) != 1 || accountPolicies.Items[0].ID != "account-policy" {
		t.Fatalf("account catalog: %+v, %v", accountPolicies, err)
	}
	if err := org.AttachRolePolicy(ctx, roleID, &workspace.OrganizationPolicyAttachRequest{PolicyID: policyID}); err != nil {
		t.Fatal(err)
	}
	assignment, err := org.AssignAccountRole(ctx, accountID, &workspace.AccountRoleAssignmentCreateRequest{PrincipalType: "user", PrincipalID: userID, RoleID: roleID})
	if err != nil || assignment.AccountID != accountID || assignment.RoleID != roleID || assignment.PrincipalID != userID {
		t.Fatalf("assignment response: %+v, %v", assignment, err)
	}
	credentials, err := account.AssumeRole(ctx, &iam.AssumeRoleRequest{Role: "crn:iam::target-account:role/Reader"})
	if err != nil || credentials.AccountID != accountID || credentials.AccountHandle != "target-account" || credentials.RoleID != roleID {
		t.Fatalf("target binding lost from role credentials: %+v, %v", credentials, err)
	}
	if len(calls) != 5 {
		t.Fatalf("got %d calls, want 5", len(calls))
	}
}
