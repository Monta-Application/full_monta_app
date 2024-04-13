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
	"github.com/emoss08/trenova/ent/equipmentmanufactuer"
	"github.com/emoss08/trenova/ent/equipmenttype"
	"github.com/emoss08/trenova/ent/fleetcode"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/trailer"
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/google/uuid"
)

// TrailerQuery is the builder for querying Trailer entities.
type TrailerQuery struct {
	config
	ctx                       *QueryContext
	order                     []trailer.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Trailer
	withBusinessUnit          *BusinessUnitQuery
	withOrganization          *OrganizationQuery
	withEquipmentType         *EquipmentTypeQuery
	withEquipmentManufacturer *EquipmentManufactuerQuery
	withState                 *UsStateQuery
	withRegistrationState     *UsStateQuery
	withFleetCode             *FleetCodeQuery
	modifiers                 []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TrailerQuery builder.
func (tq *TrailerQuery) Where(ps ...predicate.Trailer) *TrailerQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TrailerQuery) Limit(limit int) *TrailerQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TrailerQuery) Offset(offset int) *TrailerQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TrailerQuery) Unique(unique bool) *TrailerQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TrailerQuery) Order(o ...trailer.OrderOption) *TrailerQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (tq *TrailerQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.BusinessUnitTable, trailer.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (tq *TrailerQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.OrganizationTable, trailer.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipmentType chains the current query on the "equipment_type" edge.
func (tq *TrailerQuery) QueryEquipmentType() *EquipmentTypeQuery {
	query := (&EquipmentTypeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(equipmenttype.Table, equipmenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.EquipmentTypeTable, trailer.EquipmentTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipmentManufacturer chains the current query on the "equipment_manufacturer" edge.
func (tq *TrailerQuery) QueryEquipmentManufacturer() *EquipmentManufactuerQuery {
	query := (&EquipmentManufactuerClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(equipmentmanufactuer.Table, equipmentmanufactuer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.EquipmentManufacturerTable, trailer.EquipmentManufacturerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryState chains the current query on the "state" edge.
func (tq *TrailerQuery) QueryState() *UsStateQuery {
	query := (&UsStateClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(usstate.Table, usstate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.StateTable, trailer.StateColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegistrationState chains the current query on the "registration_state" edge.
func (tq *TrailerQuery) QueryRegistrationState() *UsStateQuery {
	query := (&UsStateClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(usstate.Table, usstate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.RegistrationStateTable, trailer.RegistrationStateColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFleetCode chains the current query on the "fleet_code" edge.
func (tq *TrailerQuery) QueryFleetCode() *FleetCodeQuery {
	query := (&FleetCodeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trailer.Table, trailer.FieldID, selector),
			sqlgraph.To(fleetcode.Table, fleetcode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, trailer.FleetCodeTable, trailer.FleetCodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Trailer entity from the query.
// Returns a *NotFoundError when no Trailer was found.
func (tq *TrailerQuery) First(ctx context.Context) (*Trailer, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{trailer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TrailerQuery) FirstX(ctx context.Context) *Trailer {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Trailer ID from the query.
// Returns a *NotFoundError when no Trailer ID was found.
func (tq *TrailerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{trailer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TrailerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Trailer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Trailer entity is found.
// Returns a *NotFoundError when no Trailer entities are found.
func (tq *TrailerQuery) Only(ctx context.Context) (*Trailer, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{trailer.Label}
	default:
		return nil, &NotSingularError{trailer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TrailerQuery) OnlyX(ctx context.Context) *Trailer {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Trailer ID in the query.
// Returns a *NotSingularError when more than one Trailer ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TrailerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{trailer.Label}
	default:
		err = &NotSingularError{trailer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TrailerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Trailers.
func (tq *TrailerQuery) All(ctx context.Context) ([]*Trailer, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Trailer, *TrailerQuery]()
	return withInterceptors[[]*Trailer](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TrailerQuery) AllX(ctx context.Context) []*Trailer {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Trailer IDs.
func (tq *TrailerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(trailer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TrailerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TrailerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TrailerQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TrailerQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TrailerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TrailerQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TrailerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TrailerQuery) Clone() *TrailerQuery {
	if tq == nil {
		return nil
	}
	return &TrailerQuery{
		config:                    tq.config,
		ctx:                       tq.ctx.Clone(),
		order:                     append([]trailer.OrderOption{}, tq.order...),
		inters:                    append([]Interceptor{}, tq.inters...),
		predicates:                append([]predicate.Trailer{}, tq.predicates...),
		withBusinessUnit:          tq.withBusinessUnit.Clone(),
		withOrganization:          tq.withOrganization.Clone(),
		withEquipmentType:         tq.withEquipmentType.Clone(),
		withEquipmentManufacturer: tq.withEquipmentManufacturer.Clone(),
		withState:                 tq.withState.Clone(),
		withRegistrationState:     tq.withRegistrationState.Clone(),
		withFleetCode:             tq.withFleetCode.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *TrailerQuery {
	query := (&BusinessUnitClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withBusinessUnit = query
	return tq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithOrganization(opts ...func(*OrganizationQuery)) *TrailerQuery {
	query := (&OrganizationClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withOrganization = query
	return tq
}

// WithEquipmentType tells the query-builder to eager-load the nodes that are connected to
// the "equipment_type" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithEquipmentType(opts ...func(*EquipmentTypeQuery)) *TrailerQuery {
	query := (&EquipmentTypeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withEquipmentType = query
	return tq
}

// WithEquipmentManufacturer tells the query-builder to eager-load the nodes that are connected to
// the "equipment_manufacturer" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithEquipmentManufacturer(opts ...func(*EquipmentManufactuerQuery)) *TrailerQuery {
	query := (&EquipmentManufactuerClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withEquipmentManufacturer = query
	return tq
}

// WithState tells the query-builder to eager-load the nodes that are connected to
// the "state" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithState(opts ...func(*UsStateQuery)) *TrailerQuery {
	query := (&UsStateClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withState = query
	return tq
}

// WithRegistrationState tells the query-builder to eager-load the nodes that are connected to
// the "registration_state" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithRegistrationState(opts ...func(*UsStateQuery)) *TrailerQuery {
	query := (&UsStateClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withRegistrationState = query
	return tq
}

// WithFleetCode tells the query-builder to eager-load the nodes that are connected to
// the "fleet_code" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrailerQuery) WithFleetCode(opts ...func(*FleetCodeQuery)) *TrailerQuery {
	query := (&FleetCodeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withFleetCode = query
	return tq
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
//	client.Trailer.Query().
//		GroupBy(trailer.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TrailerQuery) GroupBy(field string, fields ...string) *TrailerGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TrailerGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = trailer.Label
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
//	client.Trailer.Query().
//		Select(trailer.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (tq *TrailerQuery) Select(fields ...string) *TrailerSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TrailerSelect{TrailerQuery: tq}
	sbuild.label = trailer.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TrailerSelect configured with the given aggregations.
func (tq *TrailerQuery) Aggregate(fns ...AggregateFunc) *TrailerSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TrailerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !trailer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TrailerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Trailer, error) {
	var (
		nodes       = []*Trailer{}
		_spec       = tq.querySpec()
		loadedTypes = [7]bool{
			tq.withBusinessUnit != nil,
			tq.withOrganization != nil,
			tq.withEquipmentType != nil,
			tq.withEquipmentManufacturer != nil,
			tq.withState != nil,
			tq.withRegistrationState != nil,
			tq.withFleetCode != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Trailer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Trailer{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withBusinessUnit; query != nil {
		if err := tq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *Trailer, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withOrganization; query != nil {
		if err := tq.loadOrganization(ctx, query, nodes, nil,
			func(n *Trailer, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withEquipmentType; query != nil {
		if err := tq.loadEquipmentType(ctx, query, nodes, nil,
			func(n *Trailer, e *EquipmentType) { n.Edges.EquipmentType = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withEquipmentManufacturer; query != nil {
		if err := tq.loadEquipmentManufacturer(ctx, query, nodes, nil,
			func(n *Trailer, e *EquipmentManufactuer) { n.Edges.EquipmentManufacturer = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withState; query != nil {
		if err := tq.loadState(ctx, query, nodes, nil,
			func(n *Trailer, e *UsState) { n.Edges.State = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withRegistrationState; query != nil {
		if err := tq.loadRegistrationState(ctx, query, nodes, nil,
			func(n *Trailer, e *UsState) { n.Edges.RegistrationState = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withFleetCode; query != nil {
		if err := tq.loadFleetCode(ctx, query, nodes, nil,
			func(n *Trailer, e *FleetCode) { n.Edges.FleetCode = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TrailerQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
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
func (tq *TrailerQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
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
func (tq *TrailerQuery) loadEquipmentType(ctx context.Context, query *EquipmentTypeQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *EquipmentType)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
	for i := range nodes {
		fk := nodes[i].EquipmentTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(equipmenttype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TrailerQuery) loadEquipmentManufacturer(ctx context.Context, query *EquipmentManufactuerQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *EquipmentManufactuer)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
	for i := range nodes {
		if nodes[i].EquipmentManufacturerID == nil {
			continue
		}
		fk := *nodes[i].EquipmentManufacturerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(equipmentmanufactuer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_manufacturer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TrailerQuery) loadState(ctx context.Context, query *UsStateQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *UsState)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
	for i := range nodes {
		if nodes[i].StateID == nil {
			continue
		}
		fk := *nodes[i].StateID
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
			return fmt.Errorf(`unexpected foreign-key "state_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TrailerQuery) loadRegistrationState(ctx context.Context, query *UsStateQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *UsState)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
	for i := range nodes {
		if nodes[i].RegistrationStateID == nil {
			continue
		}
		fk := *nodes[i].RegistrationStateID
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
			return fmt.Errorf(`unexpected foreign-key "registration_state_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TrailerQuery) loadFleetCode(ctx context.Context, query *FleetCodeQuery, nodes []*Trailer, init func(*Trailer), assign func(*Trailer, *FleetCode)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trailer)
	for i := range nodes {
		fk := nodes[i].FleetCodeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(fleetcode.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "fleet_code_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TrailerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TrailerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(trailer.Table, trailer.Columns, sqlgraph.NewFieldSpec(trailer.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trailer.FieldID)
		for i := range fields {
			if fields[i] != trailer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(trailer.FieldBusinessUnitID)
		}
		if tq.withOrganization != nil {
			_spec.Node.AddColumnOnce(trailer.FieldOrganizationID)
		}
		if tq.withEquipmentType != nil {
			_spec.Node.AddColumnOnce(trailer.FieldEquipmentTypeID)
		}
		if tq.withEquipmentManufacturer != nil {
			_spec.Node.AddColumnOnce(trailer.FieldEquipmentManufacturerID)
		}
		if tq.withState != nil {
			_spec.Node.AddColumnOnce(trailer.FieldStateID)
		}
		if tq.withRegistrationState != nil {
			_spec.Node.AddColumnOnce(trailer.FieldRegistrationStateID)
		}
		if tq.withFleetCode != nil {
			_spec.Node.AddColumnOnce(trailer.FieldFleetCodeID)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TrailerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(trailer.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = trailer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TrailerQuery) Modify(modifiers ...func(s *sql.Selector)) *TrailerSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// TrailerGroupBy is the group-by builder for Trailer entities.
type TrailerGroupBy struct {
	selector
	build *TrailerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TrailerGroupBy) Aggregate(fns ...AggregateFunc) *TrailerGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TrailerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TrailerQuery, *TrailerGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TrailerGroupBy) sqlScan(ctx context.Context, root *TrailerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TrailerSelect is the builder for selecting fields of Trailer entities.
type TrailerSelect struct {
	*TrailerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TrailerSelect) Aggregate(fns ...AggregateFunc) *TrailerSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TrailerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TrailerQuery, *TrailerSelect](ctx, ts.TrailerQuery, ts, ts.inters, v)
}

func (ts *TrailerSelect) sqlScan(ctx context.Context, root *TrailerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TrailerSelect) Modify(modifiers ...func(s *sql.Selector)) *TrailerSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
