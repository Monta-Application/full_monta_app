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
	"github.com/emoss08/trenova/ent/shipment"
	"github.com/emoss08/trenova/ent/shipmentcommodity"
	"github.com/google/uuid"
)

// ShipmentCommodityQuery is the builder for querying ShipmentCommodity entities.
type ShipmentCommodityQuery struct {
	config
	ctx              *QueryContext
	order            []shipmentcommodity.OrderOption
	inters           []Interceptor
	predicates       []predicate.ShipmentCommodity
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withShipment     *ShipmentQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentCommodityQuery builder.
func (scq *ShipmentCommodityQuery) Where(ps ...predicate.ShipmentCommodity) *ShipmentCommodityQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ShipmentCommodityQuery) Limit(limit int) *ShipmentCommodityQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ShipmentCommodityQuery) Offset(offset int) *ShipmentCommodityQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ShipmentCommodityQuery) Unique(unique bool) *ShipmentCommodityQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ShipmentCommodityQuery) Order(o ...shipmentcommodity.OrderOption) *ShipmentCommodityQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (scq *ShipmentCommodityQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcommodity.Table, shipmentcommodity.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentcommodity.BusinessUnitTable, shipmentcommodity.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (scq *ShipmentCommodityQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcommodity.Table, shipmentcommodity.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentcommodity.OrganizationTable, shipmentcommodity.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShipment chains the current query on the "shipment" edge.
func (scq *ShipmentCommodityQuery) QueryShipment() *ShipmentQuery {
	query := (&ShipmentClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcommodity.Table, shipmentcommodity.FieldID, selector),
			sqlgraph.To(shipment.Table, shipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmentcommodity.ShipmentTable, shipmentcommodity.ShipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentCommodity entity from the query.
// Returns a *NotFoundError when no ShipmentCommodity was found.
func (scq *ShipmentCommodityQuery) First(ctx context.Context) (*ShipmentCommodity, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentcommodity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) FirstX(ctx context.Context) *ShipmentCommodity {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentCommodity ID from the query.
// Returns a *NotFoundError when no ShipmentCommodity ID was found.
func (scq *ShipmentCommodityQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentcommodity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentCommodity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShipmentCommodity entity is found.
// Returns a *NotFoundError when no ShipmentCommodity entities are found.
func (scq *ShipmentCommodityQuery) Only(ctx context.Context) (*ShipmentCommodity, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentcommodity.Label}
	default:
		return nil, &NotSingularError{shipmentcommodity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) OnlyX(ctx context.Context) *ShipmentCommodity {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentCommodity ID in the query.
// Returns a *NotSingularError when more than one ShipmentCommodity ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ShipmentCommodityQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentcommodity.Label}
	default:
		err = &NotSingularError{shipmentcommodity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentCommodities.
func (scq *ShipmentCommodityQuery) All(ctx context.Context) ([]*ShipmentCommodity, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShipmentCommodity, *ShipmentCommodityQuery]()
	return withInterceptors[[]*ShipmentCommodity](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) AllX(ctx context.Context) []*ShipmentCommodity {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentCommodity IDs.
func (scq *ShipmentCommodityQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(shipmentcommodity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ShipmentCommodityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ShipmentCommodityQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ShipmentCommodityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ShipmentCommodityQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentCommodityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ShipmentCommodityQuery) Clone() *ShipmentCommodityQuery {
	if scq == nil {
		return nil
	}
	return &ShipmentCommodityQuery{
		config:           scq.config,
		ctx:              scq.ctx.Clone(),
		order:            append([]shipmentcommodity.OrderOption{}, scq.order...),
		inters:           append([]Interceptor{}, scq.inters...),
		predicates:       append([]predicate.ShipmentCommodity{}, scq.predicates...),
		withBusinessUnit: scq.withBusinessUnit.Clone(),
		withOrganization: scq.withOrganization.Clone(),
		withShipment:     scq.withShipment.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommodityQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *ShipmentCommodityQuery {
	query := (&BusinessUnitClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withBusinessUnit = query
	return scq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommodityQuery) WithOrganization(opts ...func(*OrganizationQuery)) *ShipmentCommodityQuery {
	query := (&OrganizationClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withOrganization = query
	return scq
}

// WithShipment tells the query-builder to eager-load the nodes that are connected to
// the "shipment" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommodityQuery) WithShipment(opts ...func(*ShipmentQuery)) *ShipmentCommodityQuery {
	query := (&ShipmentClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withShipment = query
	return scq
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
//	client.ShipmentCommodity.Query().
//		GroupBy(shipmentcommodity.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ShipmentCommodityQuery) GroupBy(field string, fields ...string) *ShipmentCommodityGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShipmentCommodityGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = shipmentcommodity.Label
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
//	client.ShipmentCommodity.Query().
//		Select(shipmentcommodity.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (scq *ShipmentCommodityQuery) Select(fields ...string) *ShipmentCommoditySelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ShipmentCommoditySelect{ShipmentCommodityQuery: scq}
	sbuild.label = shipmentcommodity.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShipmentCommoditySelect configured with the given aggregations.
func (scq *ShipmentCommodityQuery) Aggregate(fns ...AggregateFunc) *ShipmentCommoditySelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ShipmentCommodityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !shipmentcommodity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ShipmentCommodityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShipmentCommodity, error) {
	var (
		nodes       = []*ShipmentCommodity{}
		_spec       = scq.querySpec()
		loadedTypes = [3]bool{
			scq.withBusinessUnit != nil,
			scq.withOrganization != nil,
			scq.withShipment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShipmentCommodity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShipmentCommodity{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withBusinessUnit; query != nil {
		if err := scq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *ShipmentCommodity, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withOrganization; query != nil {
		if err := scq.loadOrganization(ctx, query, nodes, nil,
			func(n *ShipmentCommodity, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withShipment; query != nil {
		if err := scq.loadShipment(ctx, query, nodes, nil,
			func(n *ShipmentCommodity, e *Shipment) { n.Edges.Shipment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ShipmentCommodityQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*ShipmentCommodity, init func(*ShipmentCommodity), assign func(*ShipmentCommodity, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentCommodity)
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
func (scq *ShipmentCommodityQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*ShipmentCommodity, init func(*ShipmentCommodity), assign func(*ShipmentCommodity, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentCommodity)
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
func (scq *ShipmentCommodityQuery) loadShipment(ctx context.Context, query *ShipmentQuery, nodes []*ShipmentCommodity, init func(*ShipmentCommodity), assign func(*ShipmentCommodity, *Shipment)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentCommodity)
	for i := range nodes {
		fk := nodes[i].ShipmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(shipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "shipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *ShipmentCommodityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ShipmentCommodityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shipmentcommodity.Table, shipmentcommodity.Columns, sqlgraph.NewFieldSpec(shipmentcommodity.FieldID, field.TypeUUID))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentcommodity.FieldID)
		for i := range fields {
			if fields[i] != shipmentcommodity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if scq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(shipmentcommodity.FieldBusinessUnitID)
		}
		if scq.withOrganization != nil {
			_spec.Node.AddColumnOnce(shipmentcommodity.FieldOrganizationID)
		}
		if scq.withShipment != nil {
			_spec.Node.AddColumnOnce(shipmentcommodity.FieldShipmentID)
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ShipmentCommodityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(shipmentcommodity.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = shipmentcommodity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range scq.modifiers {
		m(selector)
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scq *ShipmentCommodityQuery) Modify(modifiers ...func(s *sql.Selector)) *ShipmentCommoditySelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ShipmentCommodityGroupBy is the group-by builder for ShipmentCommodity entities.
type ShipmentCommodityGroupBy struct {
	selector
	build *ShipmentCommodityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ShipmentCommodityGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentCommodityGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ShipmentCommodityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentCommodityQuery, *ShipmentCommodityGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ShipmentCommodityGroupBy) sqlScan(ctx context.Context, root *ShipmentCommodityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShipmentCommoditySelect is the builder for selecting fields of ShipmentCommodity entities.
type ShipmentCommoditySelect struct {
	*ShipmentCommodityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ShipmentCommoditySelect) Aggregate(fns ...AggregateFunc) *ShipmentCommoditySelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ShipmentCommoditySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentCommodityQuery, *ShipmentCommoditySelect](ctx, scs.ShipmentCommodityQuery, scs, scs.inters, v)
}

func (scs *ShipmentCommoditySelect) sqlScan(ctx context.Context, root *ShipmentCommodityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scs *ShipmentCommoditySelect) Modify(modifiers ...func(s *sql.Selector)) *ShipmentCommoditySelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
