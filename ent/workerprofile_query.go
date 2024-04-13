// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/emoss08/trenova/ent/workerprofile"
	"github.com/google/uuid"
)

// WorkerProfileQuery is the builder for querying WorkerProfile entities.
type WorkerProfileQuery struct {
	config
	ctx              *QueryContext
	order            []workerprofile.OrderOption
	inters           []Interceptor
	predicates       []predicate.WorkerProfile
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withWorker       *WorkerQuery
	withState        *UsStateQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkerProfileQuery builder.
func (wpq *WorkerProfileQuery) Where(ps ...predicate.WorkerProfile) *WorkerProfileQuery {
	wpq.predicates = append(wpq.predicates, ps...)
	return wpq
}

// Limit the number of records to be returned by this query.
func (wpq *WorkerProfileQuery) Limit(limit int) *WorkerProfileQuery {
	wpq.ctx.Limit = &limit
	return wpq
}

// Offset to start from.
func (wpq *WorkerProfileQuery) Offset(offset int) *WorkerProfileQuery {
	wpq.ctx.Offset = &offset
	return wpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wpq *WorkerProfileQuery) Unique(unique bool) *WorkerProfileQuery {
	wpq.ctx.Unique = &unique
	return wpq
}

// Order specifies how the records should be ordered.
func (wpq *WorkerProfileQuery) Order(o ...workerprofile.OrderOption) *WorkerProfileQuery {
	wpq.order = append(wpq.order, o...)
	return wpq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (wpq *WorkerProfileQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: wpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workerprofile.Table, workerprofile.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workerprofile.BusinessUnitTable, workerprofile.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(wpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (wpq *WorkerProfileQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: wpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workerprofile.Table, workerprofile.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workerprofile.OrganizationTable, workerprofile.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(wpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorker chains the current query on the "worker" edge.
func (wpq *WorkerProfileQuery) QueryWorker() *WorkerQuery {
	query := (&WorkerClient{config: wpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workerprofile.Table, workerprofile.FieldID, selector),
			sqlgraph.To(worker.Table, worker.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, workerprofile.WorkerTable, workerprofile.WorkerColumn),
		)
		fromU = sqlgraph.SetNeighbors(wpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryState chains the current query on the "state" edge.
func (wpq *WorkerProfileQuery) QueryState() *UsStateQuery {
	query := (&UsStateClient{config: wpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workerprofile.Table, workerprofile.FieldID, selector),
			sqlgraph.To(usstate.Table, usstate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workerprofile.StateTable, workerprofile.StateColumn),
		)
		fromU = sqlgraph.SetNeighbors(wpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkerProfile entity from the query.
// Returns a *NotFoundError when no WorkerProfile was found.
func (wpq *WorkerProfileQuery) First(ctx context.Context) (*WorkerProfile, error) {
	nodes, err := wpq.Limit(1).All(setContextOp(ctx, wpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workerprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wpq *WorkerProfileQuery) FirstX(ctx context.Context) *WorkerProfile {
	node, err := wpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkerProfile ID from the query.
// Returns a *NotFoundError when no WorkerProfile ID was found.
func (wpq *WorkerProfileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wpq.Limit(1).IDs(setContextOp(ctx, wpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workerprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wpq *WorkerProfileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkerProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkerProfile entity is found.
// Returns a *NotFoundError when no WorkerProfile entities are found.
func (wpq *WorkerProfileQuery) Only(ctx context.Context) (*WorkerProfile, error) {
	nodes, err := wpq.Limit(2).All(setContextOp(ctx, wpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workerprofile.Label}
	default:
		return nil, &NotSingularError{workerprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wpq *WorkerProfileQuery) OnlyX(ctx context.Context) *WorkerProfile {
	node, err := wpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkerProfile ID in the query.
// Returns a *NotSingularError when more than one WorkerProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (wpq *WorkerProfileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wpq.Limit(2).IDs(setContextOp(ctx, wpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workerprofile.Label}
	default:
		err = &NotSingularError{workerprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wpq *WorkerProfileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkerProfiles.
func (wpq *WorkerProfileQuery) All(ctx context.Context) ([]*WorkerProfile, error) {
	ctx = setContextOp(ctx, wpq.ctx, "All")
	if err := wpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkerProfile, *WorkerProfileQuery]()
	return withInterceptors[[]*WorkerProfile](ctx, wpq, qr, wpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wpq *WorkerProfileQuery) AllX(ctx context.Context) []*WorkerProfile {
	nodes, err := wpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkerProfile IDs.
func (wpq *WorkerProfileQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wpq.ctx.Unique == nil && wpq.path != nil {
		wpq.Unique(true)
	}
	ctx = setContextOp(ctx, wpq.ctx, "IDs")
	if err = wpq.Select(workerprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wpq *WorkerProfileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wpq *WorkerProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wpq.ctx, "Count")
	if err := wpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wpq, querierCount[*WorkerProfileQuery](), wpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wpq *WorkerProfileQuery) CountX(ctx context.Context) int {
	count, err := wpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wpq *WorkerProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wpq.ctx, "Exist")
	switch _, err := wpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wpq *WorkerProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := wpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkerProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wpq *WorkerProfileQuery) Clone() *WorkerProfileQuery {
	if wpq == nil {
		return nil
	}
	return &WorkerProfileQuery{
		config:           wpq.config,
		ctx:              wpq.ctx.Clone(),
		order:            append([]workerprofile.OrderOption{}, wpq.order...),
		inters:           append([]Interceptor{}, wpq.inters...),
		predicates:       append([]predicate.WorkerProfile{}, wpq.predicates...),
		withBusinessUnit: wpq.withBusinessUnit.Clone(),
		withOrganization: wpq.withOrganization.Clone(),
		withWorker:       wpq.withWorker.Clone(),
		withState:        wpq.withState.Clone(),
		// clone intermediate query.
		sql:  wpq.sql.Clone(),
		path: wpq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (wpq *WorkerProfileQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *WorkerProfileQuery {
	query := (&BusinessUnitClient{config: wpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wpq.withBusinessUnit = query
	return wpq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (wpq *WorkerProfileQuery) WithOrganization(opts ...func(*OrganizationQuery)) *WorkerProfileQuery {
	query := (&OrganizationClient{config: wpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wpq.withOrganization = query
	return wpq
}

// WithWorker tells the query-builder to eager-load the nodes that are connected to
// the "worker" edge. The optional arguments are used to configure the query builder of the edge.
func (wpq *WorkerProfileQuery) WithWorker(opts ...func(*WorkerQuery)) *WorkerProfileQuery {
	query := (&WorkerClient{config: wpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wpq.withWorker = query
	return wpq
}

// WithState tells the query-builder to eager-load the nodes that are connected to
// the "state" edge. The optional arguments are used to configure the query builder of the edge.
func (wpq *WorkerProfileQuery) WithState(opts ...func(*UsStateQuery)) *WorkerProfileQuery {
	query := (&UsStateClient{config: wpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wpq.withState = query
	return wpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkerProfile.Query().
//		GroupBy(workerprofile.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wpq *WorkerProfileQuery) GroupBy(field string, fields ...string) *WorkerProfileGroupBy {
	wpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkerProfileGroupBy{build: wpq}
	grbuild.flds = &wpq.ctx.Fields
	grbuild.label = workerprofile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.WorkerProfile.Query().
//		Select(workerprofile.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (wpq *WorkerProfileQuery) Select(fields ...string) *WorkerProfileSelect {
	wpq.ctx.Fields = append(wpq.ctx.Fields, fields...)
	sbuild := &WorkerProfileSelect{WorkerProfileQuery: wpq}
	sbuild.label = workerprofile.Label
	sbuild.flds, sbuild.scan = &wpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkerProfileSelect configured with the given aggregations.
func (wpq *WorkerProfileQuery) Aggregate(fns ...AggregateFunc) *WorkerProfileSelect {
	return wpq.Select().Aggregate(fns...)
}

func (wpq *WorkerProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wpq); err != nil {
				return err
			}
		}
	}
	for _, f := range wpq.ctx.Fields {
		if !workerprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wpq.path != nil {
		prev, err := wpq.path(ctx)
		if err != nil {
			return err
		}
		wpq.sql = prev
	}
	return nil
}

func (wpq *WorkerProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkerProfile, error) {
	var (
		nodes       = []*WorkerProfile{}
		_spec       = wpq.querySpec()
		loadedTypes = [4]bool{
			wpq.withBusinessUnit != nil,
			wpq.withOrganization != nil,
			wpq.withWorker != nil,
			wpq.withState != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkerProfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkerProfile{config: wpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wpq.modifiers) > 0 {
		_spec.Modifiers = wpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wpq.withBusinessUnit; query != nil {
		if err := wpq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *WorkerProfile, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := wpq.withOrganization; query != nil {
		if err := wpq.loadOrganization(ctx, query, nodes, nil,
			func(n *WorkerProfile, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := wpq.withWorker; query != nil {
		if err := wpq.loadWorker(ctx, query, nodes, nil,
			func(n *WorkerProfile, e *Worker) { n.Edges.Worker = e }); err != nil {
			return nil, err
		}
	}
	if query := wpq.withState; query != nil {
		if err := wpq.loadState(ctx, query, nodes, nil,
			func(n *WorkerProfile, e *UsState) { n.Edges.State = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wpq *WorkerProfileQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*WorkerProfile, init func(*WorkerProfile), assign func(*WorkerProfile, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkerProfile)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wpq *WorkerProfileQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*WorkerProfile, init func(*WorkerProfile), assign func(*WorkerProfile, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkerProfile)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wpq *WorkerProfileQuery) loadWorker(ctx context.Context, query *WorkerQuery, nodes []*WorkerProfile, init func(*WorkerProfile), assign func(*WorkerProfile, *Worker)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkerProfile)
	for i := range nodes {
		fk := nodes[i].WorkerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(worker.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "worker_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wpq *WorkerProfileQuery) loadState(ctx context.Context, query *UsStateQuery, nodes []*WorkerProfile, init func(*WorkerProfile), assign func(*WorkerProfile, *UsState)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WorkerProfile)
	for i := range nodes {
		fk := nodes[i].LicenseStateID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(usstate.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "license_state_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wpq *WorkerProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wpq.querySpec()
	if len(wpq.modifiers) > 0 {
		_spec.Modifiers = wpq.modifiers
	}
	_spec.Node.Columns = wpq.ctx.Fields
	if len(wpq.ctx.Fields) > 0 {
		_spec.Unique = wpq.ctx.Unique != nil && *wpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wpq.driver, _spec)
}

func (wpq *WorkerProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workerprofile.Table, workerprofile.Columns, sqlgraph.NewFieldSpec(workerprofile.FieldID, field.TypeUUID))
	_spec.From = wpq.sql
	if unique := wpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wpq.path != nil {
		_spec.Unique = true
	}
	if fields := wpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workerprofile.FieldID)
		for i := range fields {
			if fields[i] != workerprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wpq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(workerprofile.FieldBusinessUnitID)
		}
		if wpq.withOrganization != nil {
			_spec.Node.AddColumnOnce(workerprofile.FieldOrganizationID)
		}
		if wpq.withWorker != nil {
			_spec.Node.AddColumnOnce(workerprofile.FieldWorkerID)
		}
		if wpq.withState != nil {
			_spec.Node.AddColumnOnce(workerprofile.FieldLicenseStateID)
		}
	}
	if ps := wpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wpq *WorkerProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wpq.driver.Dialect())
	t1 := builder.Table(workerprofile.Table)
	columns := wpq.ctx.Fields
	if len(columns) == 0 {
		columns = workerprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wpq.sql != nil {
		selector = wpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wpq.ctx.Unique != nil && *wpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wpq.modifiers {
		m(selector)
	}
	for _, p := range wpq.predicates {
		p(selector)
	}
	for _, p := range wpq.order {
		p(selector)
	}
	if offset := wpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wpq *WorkerProfileQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkerProfileSelect {
	wpq.modifiers = append(wpq.modifiers, modifiers...)
	return wpq.Select()
}

// WorkerProfileGroupBy is the group-by builder for WorkerProfile entities.
type WorkerProfileGroupBy struct {
	selector
	build *WorkerProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wpgb *WorkerProfileGroupBy) Aggregate(fns ...AggregateFunc) *WorkerProfileGroupBy {
	wpgb.fns = append(wpgb.fns, fns...)
	return wpgb
}

// Scan applies the selector query and scans the result into the given value.
func (wpgb *WorkerProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wpgb.build.ctx, "GroupBy")
	if err := wpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkerProfileQuery, *WorkerProfileGroupBy](ctx, wpgb.build, wpgb, wpgb.build.inters, v)
}

func (wpgb *WorkerProfileGroupBy) sqlScan(ctx context.Context, root *WorkerProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wpgb.fns))
	for _, fn := range wpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wpgb.flds)+len(wpgb.fns))
		for _, f := range *wpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkerProfileSelect is the builder for selecting fields of WorkerProfile entities.
type WorkerProfileSelect struct {
	*WorkerProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wps *WorkerProfileSelect) Aggregate(fns ...AggregateFunc) *WorkerProfileSelect {
	wps.fns = append(wps.fns, fns...)
	return wps
}

// Scan applies the selector query and scans the result into the given value.
func (wps *WorkerProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wps.ctx, "Select")
	if err := wps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkerProfileQuery, *WorkerProfileSelect](ctx, wps.WorkerProfileQuery, wps, wps.inters, v)
}

func (wps *WorkerProfileSelect) sqlScan(ctx context.Context, root *WorkerProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wps.fns))
	for _, fn := range wps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wps *WorkerProfileSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkerProfileSelect {
	wps.modifiers = append(wps.modifiers, modifiers...)
	return wps
}
