// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// BusinessUnitQuery is the builder for querying BusinessUnit entities.
type BusinessUnitQuery struct {
	config
	ctx               *QueryContext
	order             []businessunit.OrderOption
	inters            []Interceptor
	predicates        []predicate.BusinessUnit
	withPrev          *BusinessUnitQuery
	withNext          *BusinessUnitQuery
	withOrganizations *OrganizationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusinessUnitQuery builder.
func (buq *BusinessUnitQuery) Where(ps ...predicate.BusinessUnit) *BusinessUnitQuery {
	buq.predicates = append(buq.predicates, ps...)
	return buq
}

// Limit the number of records to be returned by this query.
func (buq *BusinessUnitQuery) Limit(limit int) *BusinessUnitQuery {
	buq.ctx.Limit = &limit
	return buq
}

// Offset to start from.
func (buq *BusinessUnitQuery) Offset(offset int) *BusinessUnitQuery {
	buq.ctx.Offset = &offset
	return buq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (buq *BusinessUnitQuery) Unique(unique bool) *BusinessUnitQuery {
	buq.ctx.Unique = &unique
	return buq
}

// Order specifies how the records should be ordered.
func (buq *BusinessUnitQuery) Order(o ...businessunit.OrderOption) *BusinessUnitQuery {
	buq.order = append(buq.order, o...)
	return buq
}

// QueryPrev chains the current query on the "prev" edge.
func (buq *BusinessUnitQuery) QueryPrev() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: buq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := buq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessunit.Table, businessunit.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, businessunit.PrevTable, businessunit.PrevColumn),
		)
		fromU = sqlgraph.SetNeighbors(buq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNext chains the current query on the "next" edge.
func (buq *BusinessUnitQuery) QueryNext() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: buq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := buq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessunit.Table, businessunit.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, businessunit.NextTable, businessunit.NextColumn),
		)
		fromU = sqlgraph.SetNeighbors(buq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganizations chains the current query on the "organizations" edge.
func (buq *BusinessUnitQuery) QueryOrganizations() *OrganizationQuery {
	query := (&OrganizationClient{config: buq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := buq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessunit.Table, businessunit.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, businessunit.OrganizationsTable, businessunit.OrganizationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(buq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BusinessUnit entity from the query.
// Returns a *NotFoundError when no BusinessUnit was found.
func (buq *BusinessUnitQuery) First(ctx context.Context) (*BusinessUnit, error) {
	nodes, err := buq.Limit(1).All(setContextOp(ctx, buq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{businessunit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (buq *BusinessUnitQuery) FirstX(ctx context.Context) *BusinessUnit {
	node, err := buq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BusinessUnit ID from the query.
// Returns a *NotFoundError when no BusinessUnit ID was found.
func (buq *BusinessUnitQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = buq.Limit(1).IDs(setContextOp(ctx, buq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{businessunit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (buq *BusinessUnitQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := buq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BusinessUnit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BusinessUnit entity is found.
// Returns a *NotFoundError when no BusinessUnit entities are found.
func (buq *BusinessUnitQuery) Only(ctx context.Context) (*BusinessUnit, error) {
	nodes, err := buq.Limit(2).All(setContextOp(ctx, buq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{businessunit.Label}
	default:
		return nil, &NotSingularError{businessunit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (buq *BusinessUnitQuery) OnlyX(ctx context.Context) *BusinessUnit {
	node, err := buq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BusinessUnit ID in the query.
// Returns a *NotSingularError when more than one BusinessUnit ID is found.
// Returns a *NotFoundError when no entities are found.
func (buq *BusinessUnitQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = buq.Limit(2).IDs(setContextOp(ctx, buq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{businessunit.Label}
	default:
		err = &NotSingularError{businessunit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (buq *BusinessUnitQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := buq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BusinessUnits.
func (buq *BusinessUnitQuery) All(ctx context.Context) ([]*BusinessUnit, error) {
	ctx = setContextOp(ctx, buq.ctx, "All")
	if err := buq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BusinessUnit, *BusinessUnitQuery]()
	return withInterceptors[[]*BusinessUnit](ctx, buq, qr, buq.inters)
}

// AllX is like All, but panics if an error occurs.
func (buq *BusinessUnitQuery) AllX(ctx context.Context) []*BusinessUnit {
	nodes, err := buq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BusinessUnit IDs.
func (buq *BusinessUnitQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if buq.ctx.Unique == nil && buq.path != nil {
		buq.Unique(true)
	}
	ctx = setContextOp(ctx, buq.ctx, "IDs")
	if err = buq.Select(businessunit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (buq *BusinessUnitQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := buq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (buq *BusinessUnitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, buq.ctx, "Count")
	if err := buq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, buq, querierCount[*BusinessUnitQuery](), buq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (buq *BusinessUnitQuery) CountX(ctx context.Context) int {
	count, err := buq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (buq *BusinessUnitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, buq.ctx, "Exist")
	switch _, err := buq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (buq *BusinessUnitQuery) ExistX(ctx context.Context) bool {
	exist, err := buq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusinessUnitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (buq *BusinessUnitQuery) Clone() *BusinessUnitQuery {
	if buq == nil {
		return nil
	}
	return &BusinessUnitQuery{
		config:            buq.config,
		ctx:               buq.ctx.Clone(),
		order:             append([]businessunit.OrderOption{}, buq.order...),
		inters:            append([]Interceptor{}, buq.inters...),
		predicates:        append([]predicate.BusinessUnit{}, buq.predicates...),
		withPrev:          buq.withPrev.Clone(),
		withNext:          buq.withNext.Clone(),
		withOrganizations: buq.withOrganizations.Clone(),
		// clone intermediate query.
		sql:  buq.sql.Clone(),
		path: buq.path,
	}
}

// WithPrev tells the query-builder to eager-load the nodes that are connected to
// the "prev" edge. The optional arguments are used to configure the query builder of the edge.
func (buq *BusinessUnitQuery) WithPrev(opts ...func(*BusinessUnitQuery)) *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: buq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	buq.withPrev = query
	return buq
}

// WithNext tells the query-builder to eager-load the nodes that are connected to
// the "next" edge. The optional arguments are used to configure the query builder of the edge.
func (buq *BusinessUnitQuery) WithNext(opts ...func(*BusinessUnitQuery)) *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: buq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	buq.withNext = query
	return buq
}

// WithOrganizations tells the query-builder to eager-load the nodes that are connected to
// the "organizations" edge. The optional arguments are used to configure the query builder of the edge.
func (buq *BusinessUnitQuery) WithOrganizations(opts ...func(*OrganizationQuery)) *BusinessUnitQuery {
	query := (&OrganizationClient{config: buq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	buq.withOrganizations = query
	return buq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BusinessUnit.Query().
//		GroupBy(businessunit.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (buq *BusinessUnitQuery) GroupBy(field string, fields ...string) *BusinessUnitGroupBy {
	buq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BusinessUnitGroupBy{build: buq}
	grbuild.flds = &buq.ctx.Fields
	grbuild.label = businessunit.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.BusinessUnit.Query().
//		Select(businessunit.FieldCreatedAt).
//		Scan(ctx, &v)
func (buq *BusinessUnitQuery) Select(fields ...string) *BusinessUnitSelect {
	buq.ctx.Fields = append(buq.ctx.Fields, fields...)
	sbuild := &BusinessUnitSelect{BusinessUnitQuery: buq}
	sbuild.label = businessunit.Label
	sbuild.flds, sbuild.scan = &buq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BusinessUnitSelect configured with the given aggregations.
func (buq *BusinessUnitQuery) Aggregate(fns ...AggregateFunc) *BusinessUnitSelect {
	return buq.Select().Aggregate(fns...)
}

func (buq *BusinessUnitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range buq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, buq); err != nil {
				return err
			}
		}
	}
	for _, f := range buq.ctx.Fields {
		if !businessunit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if buq.path != nil {
		prev, err := buq.path(ctx)
		if err != nil {
			return err
		}
		buq.sql = prev
	}
	return nil
}

func (buq *BusinessUnitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BusinessUnit, error) {
	var (
		nodes       = []*BusinessUnit{}
		_spec       = buq.querySpec()
		loadedTypes = [3]bool{
			buq.withPrev != nil,
			buq.withNext != nil,
			buq.withOrganizations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BusinessUnit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BusinessUnit{config: buq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, buq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := buq.withPrev; query != nil {
		if err := buq.loadPrev(ctx, query, nodes, nil,
			func(n *BusinessUnit, e *BusinessUnit) { n.Edges.Prev = e }); err != nil {
			return nil, err
		}
	}
	if query := buq.withNext; query != nil {
		if err := buq.loadNext(ctx, query, nodes, nil,
			func(n *BusinessUnit, e *BusinessUnit) { n.Edges.Next = e }); err != nil {
			return nil, err
		}
	}
	if query := buq.withOrganizations; query != nil {
		if err := buq.loadOrganizations(ctx, query, nodes,
			func(n *BusinessUnit) { n.Edges.Organizations = []*Organization{} },
			func(n *BusinessUnit, e *Organization) { n.Edges.Organizations = append(n.Edges.Organizations, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (buq *BusinessUnitQuery) loadPrev(ctx context.Context, query *BusinessUnitQuery, nodes []*BusinessUnit, init func(*BusinessUnit), assign func(*BusinessUnit, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*BusinessUnit)
	for i := range nodes {
		fk := nodes[i].ParentID
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
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (buq *BusinessUnitQuery) loadNext(ctx context.Context, query *BusinessUnitQuery, nodes []*BusinessUnit, init func(*BusinessUnit), assign func(*BusinessUnit, *BusinessUnit)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*BusinessUnit)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(businessunit.FieldParentID)
	}
	query.Where(predicate.BusinessUnit(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(businessunit.NextColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (buq *BusinessUnitQuery) loadOrganizations(ctx context.Context, query *OrganizationQuery, nodes []*BusinessUnit, init func(*BusinessUnit), assign func(*BusinessUnit, *Organization)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*BusinessUnit)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(businessunit.OrganizationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.business_unit_organizations
		if fk == nil {
			return fmt.Errorf(`foreign-key "business_unit_organizations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "business_unit_organizations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (buq *BusinessUnitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := buq.querySpec()
	_spec.Node.Columns = buq.ctx.Fields
	if len(buq.ctx.Fields) > 0 {
		_spec.Unique = buq.ctx.Unique != nil && *buq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, buq.driver, _spec)
}

func (buq *BusinessUnitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(businessunit.Table, businessunit.Columns, sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID))
	_spec.From = buq.sql
	if unique := buq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if buq.path != nil {
		_spec.Unique = true
	}
	if fields := buq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessunit.FieldID)
		for i := range fields {
			if fields[i] != businessunit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if buq.withPrev != nil {
			_spec.Node.AddColumnOnce(businessunit.FieldParentID)
		}
	}
	if ps := buq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := buq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := buq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := buq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (buq *BusinessUnitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(buq.driver.Dialect())
	t1 := builder.Table(businessunit.Table)
	columns := buq.ctx.Fields
	if len(columns) == 0 {
		columns = businessunit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if buq.sql != nil {
		selector = buq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if buq.ctx.Unique != nil && *buq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range buq.predicates {
		p(selector)
	}
	for _, p := range buq.order {
		p(selector)
	}
	if offset := buq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := buq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BusinessUnitGroupBy is the group-by builder for BusinessUnit entities.
type BusinessUnitGroupBy struct {
	selector
	build *BusinessUnitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bugb *BusinessUnitGroupBy) Aggregate(fns ...AggregateFunc) *BusinessUnitGroupBy {
	bugb.fns = append(bugb.fns, fns...)
	return bugb
}

// Scan applies the selector query and scans the result into the given value.
func (bugb *BusinessUnitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bugb.build.ctx, "GroupBy")
	if err := bugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessUnitQuery, *BusinessUnitGroupBy](ctx, bugb.build, bugb, bugb.build.inters, v)
}

func (bugb *BusinessUnitGroupBy) sqlScan(ctx context.Context, root *BusinessUnitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bugb.fns))
	for _, fn := range bugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bugb.flds)+len(bugb.fns))
		for _, f := range *bugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BusinessUnitSelect is the builder for selecting fields of BusinessUnit entities.
type BusinessUnitSelect struct {
	*BusinessUnitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bus *BusinessUnitSelect) Aggregate(fns ...AggregateFunc) *BusinessUnitSelect {
	bus.fns = append(bus.fns, fns...)
	return bus
}

// Scan applies the selector query and scans the result into the given value.
func (bus *BusinessUnitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bus.ctx, "Select")
	if err := bus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessUnitQuery, *BusinessUnitSelect](ctx, bus.BusinessUnitQuery, bus, bus.inters, v)
}

func (bus *BusinessUnitSelect) sqlScan(ctx context.Context, root *BusinessUnitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bus.fns))
	for _, fn := range bus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
