package basaltic

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Error is a failure reported by the platform.
//
// Every service answers a failure with the same envelope, so one type covers
// all of them:
//
//	{"error": {"code": "NOT_FOUND", "message": "...", "request_id": "..."}}
//
// RequestID identifies the call in the platform's logs. Include it verbatim
// when reporting a problem — it is the fastest way to find what happened.
type Error struct {
	// StatusCode is the HTTP status that carried the failure.
	StatusCode int
	// Code is the platform's stable error code, such as "NOT_FOUND" or
	// "QUOTA_EXCEEDED". Prefer the Is* helpers over comparing it directly:
	// a class of failure can gain new codes, and the helpers move with it.
	Code string
	// Message is the human-readable explanation. It is meant for an operator
	// reading a log, not for matching on.
	Message string
	// RequestID identifies this call in the platform's logs.
	RequestID string
	// OperationID names the API operation that failed, such as "getInstance".
	// Filled in by the SDK, not by the platform.
	OperationID string
	// Header carries the response headers, so a caller can read rate-limit
	// state off a 429 without a second request.
	Header http.Header
	// Body is the raw response body, kept for the cases the envelope does not
	// explain — a proxy error page, say, where Code is empty.
	Body []byte
}

func (e *Error) Error() string {
	var b strings.Builder
	if e.OperationID != "" {
		b.WriteString(e.OperationID)
		b.WriteString(": ")
	}
	switch {
	case e.Code != "" && e.Message != "":
		fmt.Fprintf(&b, "%s: %s", e.Code, e.Message)
	case e.Message != "":
		b.WriteString(e.Message)
	case e.Code != "":
		b.WriteString(e.Code)
	default:
		fmt.Fprintf(&b, "unexpected status %d", e.StatusCode)
	}
	fmt.Fprintf(&b, " (http %d", e.StatusCode)
	if e.RequestID != "" {
		fmt.Fprintf(&b, ", request %s", e.RequestID)
	}
	b.WriteString(")")
	return b.String()
}

// errorEnvelope is the platform's failure body.
type errorEnvelope struct {
	Error struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	} `json:"error"`
}

// parseError builds an Error from a failed response. It never returns nil:
// a body that does not parse still yields a usable error carrying the status
// and the raw bytes, because a caller that gets nil back from a 500 has no
// way to proceed.
func parseError(operationID string, resp *http.Response, body []byte) *Error {
	e := &Error{
		StatusCode:  resp.StatusCode,
		OperationID: operationID,
		Header:      resp.Header,
		Body:        body,
	}
	var env errorEnvelope
	if err := json.Unmarshal(body, &env); err == nil && (env.Error.Code != "" || env.Error.Message != "") {
		e.Code = env.Error.Code
		e.Message = env.Error.Message
		e.RequestID = env.Error.RequestID
	}
	if e.RequestID == "" {
		e.RequestID = resp.Header.Get("X-Request-Id")
	}
	return e
}

// AsError extracts the platform error from err, if there is one.
//
//	if apiErr, ok := basaltic.AsError(err); ok {
//		log.Printf("request %s failed: %s", apiErr.RequestID, apiErr.Code)
//	}
func AsError(err error) (*Error, bool) {
	var e *Error
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

// The platform names errors precisely: there are 77 distinct codes ending in
// _NOT_FOUND, one per resource kind, and as many again for conflicts. Every
// one of them carries an HTTP status chosen for it, and that status is the
// stable signal — a new resource kind adds a code but never a new status. So
// these helpers classify on the status and consult the code only where one
// status covers two situations whose remedies differ.

// IsNotFound reports whether the resource does not exist — or is not visible
// to this credential, which the platform answers the same way, deliberately.
//
// It also covers 410 Gone, for a resource that existed and was removed.
func IsNotFound(err error) bool {
	e, ok := AsError(err)
	if !ok {
		return false
	}
	return e.StatusCode == http.StatusNotFound || e.StatusCode == http.StatusGone
}

// IsUnauthorized reports whether authentication failed: no token, an expired
// one, or one the platform refused because its session was revoked.
//
// The SDK already re-exchanges once when a cached token is refused, so seeing
// this means the credential itself is the problem.
func IsUnauthorized(err error) bool {
	e, ok := AsError(err)
	return ok && e.StatusCode == http.StatusUnauthorized
}

// IsAccessDenied reports whether the credential authenticated but policy did
// not allow the action. Distinct from [IsUnauthorized]: the remedy is a policy
// change, not a new credential.
//
// A quota refusal shares this status and is excluded — see [IsQuotaExceeded],
// whose remedy is the opposite.
func IsAccessDenied(err error) bool {
	e, ok := AsError(err)
	if !ok || e.StatusCode != http.StatusForbidden {
		return false
	}
	return !isQuotaCode(e.Code)
}

// IsQuotaExceeded reports whether the account is at a quota limit.
//
// The platform answers 403 for this, the same as an authorization refusal,
// but the two need opposite responses: retrying does not clear a quota and
// raising it does, while no quota change will fix a policy that says no.
func IsQuotaExceeded(err error) bool {
	e, ok := AsError(err)
	return ok && isQuotaCode(e.Code)
}

func isQuotaCode(code string) bool {
	return code == "QUOTA_EXCEEDED" || code == "LIMIT_EXCEEDED" ||
		strings.HasSuffix(code, "_QUOTA_EXCEEDED") || strings.HasSuffix(code, "_LIMIT_EXCEEDED")
}

// IsConflict reports whether the request contradicted the resource's current
// state — a duplicate name, a resource still in use, or an action its status
// does not allow right now.
func IsConflict(err error) bool {
	e, ok := AsError(err)
	return ok && e.StatusCode == http.StatusConflict
}

// IsInvalidInput reports whether the platform rejected the request as
// malformed or semantically invalid. Retrying it unchanged will not help.
func IsInvalidInput(err error) bool {
	e, ok := AsError(err)
	if !ok {
		return false
	}
	return e.StatusCode == http.StatusBadRequest || e.StatusCode == http.StatusUnprocessableEntity
}

// IsRateLimited reports whether the request was throttled.
//
// The SDK already retries these within its budget, so seeing one means the
// budget ran out. Back off further rather than retrying immediately, and read
// [Error.RetryAfter] rather than guessing — retrying early is refused and
// extends the window.
func IsRateLimited(err error) bool {
	e, ok := AsError(err)
	return ok && e.StatusCode == http.StatusTooManyRequests
}

// IsTransient reports whether the failure was on the platform's side and may
// well succeed on a later attempt: a 429, or a 5xx.
//
// The SDK retries these itself within the configured budget, so this is for
// deciding what to do after that budget is spent.
func IsTransient(err error) bool {
	e, ok := AsError(err)
	if !ok {
		return false
	}
	return e.StatusCode == http.StatusTooManyRequests || e.StatusCode >= 500
}

// RetryAfter returns how long the platform asked the caller to wait before
// retrying, and whether it said anything at all. Read from the Retry-After
// header on a 429 or 503.
func (e *Error) RetryAfter() (retryAfter time.Duration, ok bool) {
	if e.Header == nil {
		return 0, false
	}
	return parseRetryAfter(e.Header.Get("Retry-After"))
}

// RateLimit reports the throttling state the platform attached to this
// response. Values are only present on operations that document a 429; ok is
// false when the response carried no rate-limit headers.
func (e *Error) RateLimit() (limit, remaining int, ok bool) {
	if e.Header == nil {
		return 0, 0, false
	}
	l, err1 := atoiHeader(e.Header.Get("X-RateLimit-Limit"))
	r, err2 := atoiHeader(e.Header.Get("X-RateLimit-Remaining"))
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return l, r, true
}
