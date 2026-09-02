// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package database

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListBackupsParams are the optional filters and pagination controls for
// [Client.ListBackups]. A nil *ListBackupsParams sends none of them.
type ListBackupsParams struct {
	// Kind filter by backup kind. full and incr are accepted as aliases for
	// base and incremental.
	//
	// One of: "base", "incremental", "wal", "snapshot", "full", "incr".
	Kind   string
	Status string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListBackupsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Kind != "" {
		q.Set("kind", p.Kind)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	return q
}

// ListClustersParams are the optional filters and pagination controls for
// [Client.ListClusters]. A nil *ListClustersParams sends none of them.
type ListClustersParams struct {
	// EngineType one of: "postgres", "valkey".
	EngineType string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
	Name   string
	Status string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListClustersParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.EngineType != "" {
		q.Set("engine_type", p.EngineType)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListClustersParams) withMarker(marker string) *ListClustersParams {
	var out ListClustersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListParameterGroupsParams are the optional filters and pagination controls for
// [Client.ListParameterGroups]. A nil *ListParameterGroupsParams sends none of them.
type ListParameterGroupsParams struct {
	// EngineType one of: "postgres".
	EngineType string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListParameterGroupsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.EngineType != "" {
		q.Set("engine_type", p.EngineType)
	}
	return q
}

// AddReplica adds a read replica to the cluster.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AddReplica(ctx context.Context, clusterID string, body AddReplicaRequest, opts ...basaltic.RequestOption) (*Instance, error) {
	op := &basaltic.Operation{
		ID:       "addReplica",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/replicas",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Replica *Instance `json:"replica"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Replica, nil
}

// ConvertClusterToHA converts a single-node postgres cluster to Patroni-managed HA.
//
// Turns a single-node postgres cluster into a Patroni-managed one IN
// PLACE, keeping its endpoint and its data — the path from the cheap
// first choice to HA without dumping and recreating.
//
// **This restarts postgres.** A cluster created single-node runs plain
// postgres; converting installs Patroni, provisions the replication user
// and the internal security group, and brings the node back up under
// Patroni against its existing data directory. There is no replica to
// fall back on while that happens, so the restart is on the only copy of
// the data.
//
// Asynchronous. The cluster moves to `converting` and returns to
// `active` when the member reports the outcome; it becomes
// `patroni_managed` only on success. A conversion that fails leaves the
// node serving plain postgres, which is the state it was already in.
// Poll the cluster to follow it.
//
// Postgres only — valkey self-clusters through Sentinel and needs at
// least three members, so its conversion is a different shape.
//
// Afterwards, add replicas with POST /v1/clusters/{cluster_id}/replicas.
func (c *Client) ConvertClusterToHA(ctx context.Context, clusterID string, opts ...basaltic.RequestOption) (*Cluster, error) {
	op := &basaltic.Operation{
		ID:       "convertClusterToHA",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/convert-to-ha",
		PathArgs: []string{clusterID},
	}
	var out struct {
		Cluster *Cluster `json:"cluster"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

// CreateCluster creates a database cluster.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateCluster(ctx context.Context, body *CreateClusterRequest, opts ...basaltic.RequestOption) (*Cluster, error) {
	op := &basaltic.Operation{
		ID:     "createCluster",
		Method: "POST",
		Path:   "/v1/clusters",
		Body:   body,
	}
	var out struct {
		Cluster *Cluster `json:"cluster"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

// CreateDBUser creates a database user.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateDBUser(ctx context.Context, clusterID string, body *CreateDBUserRequest, opts ...basaltic.RequestOption) (*DBUser, error) {
	op := &basaltic.Operation{
		ID:       "createDBUser",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/users",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		User *DBUser `json:"user"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.User, nil
}

// CreateDatabase creates a logical database.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateDatabase(ctx context.Context, clusterID string, body *CreateDatabaseRequest, opts ...basaltic.RequestOption) (*Database, error) {
	op := &basaltic.Operation{
		ID:       "createDatabase",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/databases",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Database *Database `json:"database"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Database, nil
}

// CreateParameterGroup creates a parameter group.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateParameterGroup(ctx context.Context, body *CreateParameterGroupRequest, opts ...basaltic.RequestOption) (*ParameterGroup, error) {
	op := &basaltic.Operation{
		ID:     "createParameterGroup",
		Method: "POST",
		Path:   "/v1/parameter-groups",
		Body:   body,
	}
	var out struct {
		ParameterGroup *ParameterGroup `json:"parameter_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ParameterGroup, nil
}

// DeleteCluster deletes a database cluster.
func (c *Client) DeleteCluster(ctx context.Context, clusterID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteCluster",
		Method:   "DELETE",
		Path:     "/v1/clusters/{cluster_id}",
		PathArgs: []string{clusterID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteDBUser deletes a database user.
func (c *Client) DeleteDBUser(ctx context.Context, clusterID string, userID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteDBUser",
		Method:   "DELETE",
		Path:     "/v1/clusters/{cluster_id}/users/{user_id}",
		PathArgs: []string{clusterID, userID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteDatabase deletes a logical database.
func (c *Client) DeleteDatabase(ctx context.Context, clusterID string, databaseID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteDatabase",
		Method:   "DELETE",
		Path:     "/v1/clusters/{cluster_id}/databases/{database_id}",
		PathArgs: []string{clusterID, databaseID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteParameterGroup deletes a parameter group.
func (c *Client) DeleteParameterGroup(ctx context.Context, parameterGroupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteParameterGroup",
		Method:   "DELETE",
		Path:     "/v1/parameter-groups/{parameter_group_id}",
		PathArgs: []string{parameterGroupID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// FailoverCluster triggers a planned HA switchover (Patroni or Valkey Sentinel).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) FailoverCluster(ctx context.Context, clusterID string, body *FailoverRequest, opts ...basaltic.RequestOption) (bool, error) {
	op := &basaltic.Operation{
		ID:       "failoverCluster",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/failover",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Initiated bool `json:"initiated"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return false, err
	}
	return out.Initiated, nil
}

// GetBackup gets a backup.
func (c *Client) GetBackup(ctx context.Context, clusterID string, backupID string, opts ...basaltic.RequestOption) (*Backup, error) {
	op := &basaltic.Operation{
		ID:       "getBackup",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/backups/{backup_id}",
		PathArgs: []string{clusterID, backupID},
	}
	var out struct {
		Backup *Backup `json:"backup"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Backup, nil
}

// GetCluster gets a database cluster.
func (c *Client) GetCluster(ctx context.Context, clusterID string, opts ...basaltic.RequestOption) (*Cluster, error) {
	op := &basaltic.Operation{
		ID:       "getCluster",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}",
		PathArgs: []string{clusterID},
	}
	var out struct {
		Cluster *Cluster `json:"cluster"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

// GetDBUser gets a database user.
func (c *Client) GetDBUser(ctx context.Context, clusterID string, userID string, opts ...basaltic.RequestOption) (*DBUser, error) {
	op := &basaltic.Operation{
		ID:       "getDBUser",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/users/{user_id}",
		PathArgs: []string{clusterID, userID},
	}
	var out struct {
		User *DBUser `json:"user"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.User, nil
}

// GetDatabase gets a logical database.
func (c *Client) GetDatabase(ctx context.Context, clusterID string, databaseID string, opts ...basaltic.RequestOption) (*Database, error) {
	op := &basaltic.Operation{
		ID:       "getDatabase",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/databases/{database_id}",
		PathArgs: []string{clusterID, databaseID},
	}
	var out struct {
		Database *Database `json:"database"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Database, nil
}

// GetParameterGroup gets a parameter group.
func (c *Client) GetParameterGroup(ctx context.Context, parameterGroupID string, opts ...basaltic.RequestOption) (*ParameterGroup, error) {
	op := &basaltic.Operation{
		ID:       "getParameterGroup",
		Method:   "GET",
		Path:     "/v1/parameter-groups/{parameter_group_id}",
		PathArgs: []string{parameterGroupID},
	}
	var out struct {
		ParameterGroup *ParameterGroup `json:"parameter_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ParameterGroup, nil
}

// ListBackups lists a cluster's backups.
func (c *Client) ListBackups(ctx context.Context, clusterID string, params *ListBackupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Backup], error) {
	op := &basaltic.Operation{
		ID:       "listBackups",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/backups",
		PathArgs: []string{clusterID},
	}
	op.Query = params.query()
	var out struct {
		Items []Backup `json:"backups"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Backup]{Items: out.Items}
	return page, nil
}

// ListClusters lists database clusters.
//
// Returns one page. Use ListClustersAll to walk every page.
func (c *Client) ListClusters(ctx context.Context, params *ListClustersParams, opts ...basaltic.RequestOption) (*basaltic.Page[Cluster], error) {
	op := &basaltic.Operation{
		ID:     "listClusters",
		Method: "GET",
		Path:   "/v1/clusters",
	}
	op.Query = params.query()
	var out struct {
		Items []Cluster `json:"clusters"`
		Meta  *struct {
			Total   int    `json:"total"`
			Limit   int    `json:"limit"`
			Marker  string `json:"marker"`
			HasMore bool   `json:"has_more"`
		} `json:"meta"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Cluster]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListClustersAll walks every page of ListClusters, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListClustersAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListClustersAll(ctx context.Context, params *ListClustersParams, opts ...basaltic.RequestOption) iter.Seq2[Cluster, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Cluster], error) {
		return c.ListClusters(ctx, params.withMarker(marker), opts...)
	})
}

// ListDBUsers lists a cluster's database users.
func (c *Client) ListDBUsers(ctx context.Context, clusterID string, opts ...basaltic.RequestOption) (*basaltic.Page[DBUser], error) {
	op := &basaltic.Operation{
		ID:       "listDBUsers",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/users",
		PathArgs: []string{clusterID},
	}
	var out struct {
		Items []DBUser `json:"users"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[DBUser]{Items: out.Items}
	return page, nil
}

// ListDatabases lists a cluster's logical databases.
func (c *Client) ListDatabases(ctx context.Context, clusterID string, opts ...basaltic.RequestOption) (*basaltic.Page[Database], error) {
	op := &basaltic.Operation{
		ID:       "listDatabases",
		Method:   "GET",
		Path:     "/v1/clusters/{cluster_id}/databases",
		PathArgs: []string{clusterID},
	}
	var out struct {
		Items []Database `json:"databases"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Database]{Items: out.Items}
	return page, nil
}

// ListEngineParameters lists an engine's tunable parameters.
//
// The allowlist of settings a parameter group for this engine may carry,
// with each one's value grammar, bounds, and whether applying it needs a
// restart. Parameter groups validate against this same catalog, so
// without it the only way to learn what is tunable is to submit a group
// and read the 400. Settings the platform owns are absent and are
// refused. Static platform metadata, not account-scoped.
func (c *Client) ListEngineParameters(ctx context.Context, engineType string, opts ...basaltic.RequestOption) (*EngineParameterListResponse, error) {
	op := &basaltic.Operation{
		ID:       "listEngineParameters",
		Method:   "GET",
		Path:     "/v1/engines/{engine_type}/parameters",
		PathArgs: []string{engineType},
	}
	var out EngineParameterListResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListEngines lists supported database engines.
//
// The supported engine + major-version catalog. Users pick a major only
// — minors and patches auto-upgrade within it. Static platform
// metadata, not account-scoped.
func (c *Client) ListEngines(ctx context.Context, opts ...basaltic.RequestOption) (*basaltic.Page[Engine], error) {
	op := &basaltic.Operation{
		ID:     "listEngines",
		Method: "GET",
		Path:   "/v1/engines",
	}
	var out struct {
		Items []Engine `json:"engines"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Engine]{Items: out.Items}
	return page, nil
}

// ListParameterGroups lists parameter groups.
func (c *Client) ListParameterGroups(ctx context.Context, params *ListParameterGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[ParameterGroup], error) {
	op := &basaltic.Operation{
		ID:     "listParameterGroups",
		Method: "GET",
		Path:   "/v1/parameter-groups",
	}
	op.Query = params.query()
	var out struct {
		Items []ParameterGroup `json:"parameter_groups"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[ParameterGroup]{Items: out.Items}
	return page, nil
}

// RemoveReplica removes a read replica from the cluster.
func (c *Client) RemoveReplica(ctx context.Context, clusterID string, instanceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeReplica",
		Method:   "DELETE",
		Path:     "/v1/clusters/{cluster_id}/replicas/{instance_id}",
		PathArgs: []string{clusterID, instanceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RequestBackup requests a manual backup.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RequestBackup(ctx context.Context, clusterID string, body *RequestBackupRequest, opts ...basaltic.RequestOption) (*Backup, error) {
	op := &basaltic.Operation{
		ID:       "requestBackup",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/backups",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Backup *Backup `json:"backup"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Backup, nil
}

// RestoreCluster restores a backup into this cluster, in place.
//
// Overwrites this cluster's data from a backup, keeping its endpoint,
// security groups, IAM role and parameter groups. Restore was previously
// a create-time bootstrap only, so recovering from a bad migration meant
// creating a second cluster and cutting over by hand — every
// connection string, plus everything attached to the original. This
// keeps all of it.
//
// **Destructive and irreversible.** `confirm` must equal the cluster's
// name. A pre-restore backup is requested before anything is
// overwritten, so a mistaken restore has a way back.
//
// The source backup may belong to this cluster or to another in the same
// account; a cross-cluster restore grants this cluster read on that
// backup bucket for the duration and revokes it when the restore
// reports.
//
// Asynchronous. The cluster reports `restoring` — it is neither active
// nor building and must not be read as serving current data — and
// returns to `active` when its members report. An HA cluster is restored
// as a WHOLE: Patroni is paused, the leader is restored, and the
// replicas are reinitialised from it. A failed restore leaves the
// cluster serving with a fault recorded and the pre-restore backup still
// available.
//
// Postgres only.
func (c *Client) RestoreCluster(ctx context.Context, clusterID string, body *RestoreClusterRequest, opts ...basaltic.RequestOption) (*Cluster, error) {
	op := &basaltic.Operation{
		ID:       "restoreCluster",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/restore",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Cluster *Cluster `json:"cluster"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

// RotateDBUserPassword rotates a database user's password.
//
// Returns the new password ONCE — capture it from the response.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RotateDBUserPassword(ctx context.Context, clusterID string, userID string, opts ...basaltic.RequestOption) (*DBUser, error) {
	op := &basaltic.Operation{
		ID:       "rotateDBUserPassword",
		Method:   "POST",
		Path:     "/v1/clusters/{cluster_id}/users/{user_id}/rotate-password",
		PathArgs: []string{clusterID, userID},
	}
	var out struct {
		User *DBUser `json:"user"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.User, nil
}

// UpdateCluster updates a database cluster (description + tags only).
func (c *Client) UpdateCluster(ctx context.Context, clusterID string, body *UpdateClusterRequest, opts ...basaltic.RequestOption) (*Cluster, error) {
	op := &basaltic.Operation{
		ID:       "updateCluster",
		Method:   "PATCH",
		Path:     "/v1/clusters/{cluster_id}",
		PathArgs: []string{clusterID},
		Body:     body,
	}
	var out struct {
		Cluster *Cluster `json:"cluster"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Cluster, nil
}

// UpdateParameterGroup updates a parameter group.
func (c *Client) UpdateParameterGroup(ctx context.Context, parameterGroupID string, body *UpdateParameterGroupRequest, opts ...basaltic.RequestOption) (*ParameterGroup, error) {
	op := &basaltic.Operation{
		ID:       "updateParameterGroup",
		Method:   "PATCH",
		Path:     "/v1/parameter-groups/{parameter_group_id}",
		PathArgs: []string{parameterGroupID},
		Body:     body,
	}
	var out struct {
		ParameterGroup *ParameterGroup `json:"parameter_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ParameterGroup, nil
}
