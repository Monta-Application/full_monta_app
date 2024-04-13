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
	"github.com/emoss08/trenova/ent/invoicecontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// InvoiceControlQuery is the builder for querying InvoiceControl entities.
type InvoiceControlQuery struct {
	config
	ctx              *QueryContext
	order            []invoicecontrol.OrderOption
	inters           []Interceptor
	predicates       []predicate.InvoiceControl
	withOrganization *OrganizationQuery
	withBusinessUnit *BusinessUnitQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceControlQuery builder.
func (icq *InvoiceControlQuery) Where(ps ...predicate.InvoiceControl) *InvoiceControlQuery {
	icq.predicates = append(icq.predicates, ps...)
	return icq
}

// Limit the number of records to be returned by this query.
func (icq *InvoiceControlQuery) Limit(limit int) *InvoiceControlQuery {
	icq.ctx.Limit = &limit
	return icq
}

// Offset to start from.
func (icq *InvoiceControlQuery) Offset(offset int) *InvoiceControlQuery {
	icq.ctx.Offset = &offset
	return icq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (icq *InvoiceControlQuery) Unique(unique bool) *InvoiceControlQuery {
	icq.ctx.Unique = &unique
	return icq
}

// Order specifies how the records should be ordered.
func (icq *InvoiceControlQuery) Order(o ...invoicecontrol.OrderOption) *InvoiceControlQuery {
	icq.order = append(icq.order, o...)
	return icq
}

// QueryOrganization chains the current query on the "organization" edge.
func (icq *InvoiceControlQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: icq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicecontrol.Table, invoicecontrol.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, invoicecontrol.OrganizationTable, invoicecontrol.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (icq *InvoiceControlQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: icq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := icq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := icq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoicecontrol.Table, invoicecontrol.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoicecontrol.BusinessUnitTable, invoicecontrol.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(icq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InvoiceControl entity from the query.
// Returns a *NotFoundError when no InvoiceControl was found.
func (icq *InvoiceControlQuery) First(ctx context.Context) (*InvoiceControl, error) {
	nodes, err := icq.Limit(1).All(setContextOp(ctx, icq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoicecontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (icq *InvoiceControlQuery) FirstX(ctx context.Context) *InvoiceControl {
	node, err := icq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InvoiceControl ID from the query.
// Returns a *NotFoundError when no InvoiceControl ID was found.
func (icq *InvoiceControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = icq.Limit(1).IDs(setContextOp(ctx, icq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoicecontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (icq *InvoiceControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := icq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InvoiceControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InvoiceControl entity is found.
// Returns a *NotFoundError when no InvoiceControl entities are found.
func (icq *InvoiceControlQuery) Only(ctx context.Context) (*InvoiceControl, error) {
	nodes, err := icq.Limit(2).All(setContextOp(ctx, icq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoicecontrol.Label}
	default:
		return nil, &NotSingularError{invoicecontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (icq *InvoiceControlQuery) OnlyX(ctx context.Context) *InvoiceControl {
	node, err := icq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InvoiceControl ID in the query.
// Returns a *NotSingularError when more than one InvoiceControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (icq *InvoiceControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = icq.Limit(2).IDs(setContextOp(ctx, icq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoicecontrol.Label}
	default:
		err = &NotSingularError{invoicecontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (icq *InvoiceControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := icq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InvoiceControls.
func (icq *InvoiceControlQuery) All(ctx context.Context) ([]*InvoiceControl, error) {
	ctx = setContextOp(ctx, icq.ctx, "All")
	if err := icq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InvoiceControl, *InvoiceControlQuery]()
	return withInterceptors[[]*InvoiceControl](ctx, icq, qr, icq.inters)
}

// AllX is like All, but panics if an error occurs.
func (icq *InvoiceControlQuery) AllX(ctx context.Context) []*InvoiceControl {
	nodes, err := icq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InvoiceControl IDs.
func (icq *InvoiceControlQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if icq.ctx.Unique == nil && icq.path != nil {
		icq.Unique(true)
	}
	ctx = setContextOp(ctx, icq.ctx, "IDs")
	if err = icq.Select(invoicecontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (icq *InvoiceControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := icq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (icq *InvoiceControlQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, icq.ctx, "Count")
	if err := icq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, icq, querierCount[*InvoiceControlQuery](), icq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (icq *InvoiceControlQuery) CountX(ctx context.Context) int {
	count, err := icq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (icq *InvoiceControlQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, icq.ctx, "Exist")
	switch _, err := icq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (icq *InvoiceControlQuery) ExistX(ctx context.Context) bool {
	exist, err := icq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (icq *InvoiceControlQuery) Clone() *InvoiceControlQuery {
	if icq == nil {
		return nil
	}
	return &InvoiceControlQuery{
		config:           icq.config,
		ctx:              icq.ctx.Clone(),
		order:            append([]invoicecontrol.OrderOption{}, icq.order...),
		inters:           append([]Interceptor{}, icq.inters...),
		predicates:       append([]predicate.InvoiceControl{}, icq.predicates...),
		withOrganization: icq.withOrganization.Clone(),
		withBusinessUnit: icq.withBusinessUnit.Clone(),
		// clone intermediate query.
		sql:  icq.sql.Clone(),
		path: icq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *InvoiceControlQuery) WithOrganization(opts ...func(*OrganizationQuery)) *InvoiceControlQuery {
	query := (&OrganizationClient{config: icq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	icq.withOrganization = query
	return icq
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (icq *InvoiceControlQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *InvoiceControlQuery {
	query := (&BusinessUnitClient{config: icq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	icq.withBusinessUnit = query
	return icq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InvoiceControl.Query().
//		GroupBy(invoicecontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (icq *InvoiceControlQuery) GroupBy(field string, fields ...string) *InvoiceControlGroupBy {
	icq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceControlGroupBy{build: icq}
	grbuild.flds = &icq.ctx.Fields
	grbuild.label = invoicecontrol.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//	}
//
//	client.InvoiceControl.Query().
//		Select(invoicecontrol.FieldCreatedAt).
//		Scan(ctx, &v)
func (icq *InvoiceControlQuery) Select(fields ...string) *InvoiceControlSelect {
	icq.ctx.Fields = append(icq.ctx.Fields, fields...)
	sbuild := &InvoiceControlSelect{InvoiceControlQuery: icq}
	sbuild.label = invoicecontrol.Label
	sbuild.flds, sbuild.scan = &icq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceControlSelect configured with the given aggregations.
func (icq *InvoiceControlQuery) Aggregate(fns ...AggregateFunc) *InvoiceControlSelect {
	return icq.Select().Aggregate(fns...)
}

func (icq *InvoiceControlQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range icq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, icq); err != nil {
				return err
			}
		}
	}
	for _, f := range icq.ctx.Fields {
		if !invoicecontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if icq.path != nil {
		prev, err := icq.path(ctx)
		if err != nil {
			return err
		}
		icq.sql = prev
	}
	return nil
}

func (icq *InvoiceControlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InvoiceControl, error) {
	var (
		nodes       = []*InvoiceControl{}
		withFKs     = icq.withFKs
		_spec       = icq.querySpec()
		loadedTypes = [2]bool{
			icq.withOrganization != nil,
			icq.withBusinessUnit != nil,
		}
	)
	if icq.withOrganization != nil || icq.withBusinessUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invoicecontrol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InvoiceControl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InvoiceControl{config: icq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(icq.modifiers) > 0 {
		_spec.Modifiers = icq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, icq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := icq.withOrganization; query != nil {
		if err := icq.loadOrganization(ctx, query, nodes, nil,
			func(n *InvoiceControl, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := icq.withBusinessUnit; query != nil {
		if err := icq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *InvoiceControl, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (icq *InvoiceControlQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*InvoiceControl, init func(*InvoiceControl), assign func(*InvoiceControl, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InvoiceControl)
	for i := range nodes {
		if nodes[i].organization_id == nil {
			continue
		}
		fk := *nodes[i].organization_id
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
func (icq *InvoiceControlQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*InvoiceControl, init func(*InvoiceControl), assign func(*InvoiceControl, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InvoiceControl)
	for i := range nodes {
		if nodes[i].business_unit_id == nil {
			continue
		}
		fk := *nodes[i].business_unit_id
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

func (icq *InvoiceControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := icq.querySpec()
	if len(icq.modifiers) > 0 {
		_spec.Modifiers = icq.modifiers
	}
	_spec.Node.Columns = icq.ctx.Fields
	if len(icq.ctx.Fields) > 0 {
		_spec.Unique = icq.ctx.Unique != nil && *icq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, icq.driver, _spec)
}

func (icq *InvoiceControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoicecontrol.Table, invoicecontrol.Columns, sqlgraph.NewFieldSpec(invoicecontrol.FieldID, field.TypeUUID))
	_spec.From = icq.sql
	if unique := icq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if icq.path != nil {
		_spec.Unique = true
	}
	if fields := icq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicecontrol.FieldID)
		for i := range fields {
			if fields[i] != invoicecontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := icq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := icq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := icq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := icq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (icq *InvoiceControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(icq.driver.Dialect())
	t1 := builder.Table(invoicecontrol.Table)
	columns := icq.ctx.Fields
	if len(columns) == 0 {
		columns = invoicecontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if icq.sql != nil {
		selector = icq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if icq.ctx.Unique != nil && *icq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range icq.modifiers {
		m(selector)
	}
	for _, p := range icq.predicates {
		p(selector)
	}
	for _, p := range icq.order {
		p(selector)
	}
	if offset := icq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := icq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (icq *InvoiceControlQuery) Modify(modifiers ...func(s *sql.Selector)) *InvoiceControlSelect {
	icq.modifiers = append(icq.modifiers, modifiers...)
	return icq.Select()
}

// InvoiceControlGroupBy is the group-by builder for InvoiceControl entities.
type InvoiceControlGroupBy struct {
	selector
	build *InvoiceControlQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (icgb *InvoiceControlGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceControlGroupBy {
	icgb.fns = append(icgb.fns, fns...)
	return icgb
}

// Scan applies the selector query and scans the result into the given value.
func (icgb *InvoiceControlGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, icgb.build.ctx, "GroupBy")
	if err := icgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceControlQuery, *InvoiceControlGroupBy](ctx, icgb.build, icgb, icgb.build.inters, v)
}

func (icgb *InvoiceControlGroupBy) sqlScan(ctx context.Context, root *InvoiceControlQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(icgb.fns))
	for _, fn := range icgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*icgb.flds)+len(icgb.fns))
		for _, f := range *icgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*icgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := icgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceControlSelect is the builder for selecting fields of InvoiceControl entities.
type InvoiceControlSelect struct {
	*InvoiceControlQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ics *InvoiceControlSelect) Aggregate(fns ...AggregateFunc) *InvoiceControlSelect {
	ics.fns = append(ics.fns, fns...)
	return ics
}

// Scan applies the selector query and scans the result into the given value.
func (ics *InvoiceControlSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ics.ctx, "Select")
	if err := ics.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceControlQuery, *InvoiceControlSelect](ctx, ics.InvoiceControlQuery, ics, ics.inters, v)
}

func (ics *InvoiceControlSelect) sqlScan(ctx context.Context, root *InvoiceControlQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ics.fns))
	for _, fn := range ics.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ics.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ics.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ics *InvoiceControlSelect) Modify(modifiers ...func(s *sql.Selector)) *InvoiceControlSelect {
	ics.modifiers = append(ics.modifiers, modifiers...)
	return ics
}
