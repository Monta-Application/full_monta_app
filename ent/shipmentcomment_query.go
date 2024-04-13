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
	"github.com/emoss08/trenova/ent/commenttype"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipment"
	"github.com/emoss08/trenova/ent/shipmentcomment"
	"github.com/emoss08/trenova/ent/user"
	"github.com/google/uuid"
)

// ShipmentCommentQuery is the builder for querying ShipmentComment entities.
type ShipmentCommentQuery struct {
	config
	ctx               *QueryContext
	order             []shipmentcomment.OrderOption
	inters            []Interceptor
	predicates        []predicate.ShipmentComment
	withBusinessUnit  *BusinessUnitQuery
	withOrganization  *OrganizationQuery
	withShipment      *ShipmentQuery
	withCommentType   *CommentTypeQuery
	withCreatedByUser *UserQuery
	modifiers         []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentCommentQuery builder.
func (scq *ShipmentCommentQuery) Where(ps ...predicate.ShipmentComment) *ShipmentCommentQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ShipmentCommentQuery) Limit(limit int) *ShipmentCommentQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ShipmentCommentQuery) Offset(offset int) *ShipmentCommentQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ShipmentCommentQuery) Unique(unique bool) *ShipmentCommentQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ShipmentCommentQuery) Order(o ...shipmentcomment.OrderOption) *ShipmentCommentQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (scq *ShipmentCommentQuery) QueryBusinessUnit() *BusinessUnitQuery {
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
			sqlgraph.From(shipmentcomment.Table, shipmentcomment.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentcomment.BusinessUnitTable, shipmentcomment.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (scq *ShipmentCommentQuery) QueryOrganization() *OrganizationQuery {
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
			sqlgraph.From(shipmentcomment.Table, shipmentcomment.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentcomment.OrganizationTable, shipmentcomment.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShipment chains the current query on the "shipment" edge.
func (scq *ShipmentCommentQuery) QueryShipment() *ShipmentQuery {
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
			sqlgraph.From(shipmentcomment.Table, shipmentcomment.FieldID, selector),
			sqlgraph.To(shipment.Table, shipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmentcomment.ShipmentTable, shipmentcomment.ShipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCommentType chains the current query on the "comment_type" edge.
func (scq *ShipmentCommentQuery) QueryCommentType() *CommentTypeQuery {
	query := (&CommentTypeClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcomment.Table, shipmentcomment.FieldID, selector),
			sqlgraph.To(commenttype.Table, commenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmentcomment.CommentTypeTable, shipmentcomment.CommentTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCreatedByUser chains the current query on the "created_by_user" edge.
func (scq *ShipmentCommentQuery) QueryCreatedByUser() *UserQuery {
	query := (&UserClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcomment.Table, shipmentcomment.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmentcomment.CreatedByUserTable, shipmentcomment.CreatedByUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentComment entity from the query.
// Returns a *NotFoundError when no ShipmentComment was found.
func (scq *ShipmentCommentQuery) First(ctx context.Context) (*ShipmentComment, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentcomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ShipmentCommentQuery) FirstX(ctx context.Context) *ShipmentComment {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentComment ID from the query.
// Returns a *NotFoundError when no ShipmentComment ID was found.
func (scq *ShipmentCommentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentcomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ShipmentCommentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShipmentComment entity is found.
// Returns a *NotFoundError when no ShipmentComment entities are found.
func (scq *ShipmentCommentQuery) Only(ctx context.Context) (*ShipmentComment, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentcomment.Label}
	default:
		return nil, &NotSingularError{shipmentcomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ShipmentCommentQuery) OnlyX(ctx context.Context) *ShipmentComment {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentComment ID in the query.
// Returns a *NotSingularError when more than one ShipmentComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ShipmentCommentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentcomment.Label}
	default:
		err = &NotSingularError{shipmentcomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ShipmentCommentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentComments.
func (scq *ShipmentCommentQuery) All(ctx context.Context) ([]*ShipmentComment, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShipmentComment, *ShipmentCommentQuery]()
	return withInterceptors[[]*ShipmentComment](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ShipmentCommentQuery) AllX(ctx context.Context) []*ShipmentComment {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentComment IDs.
func (scq *ShipmentCommentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(shipmentcomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ShipmentCommentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ShipmentCommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ShipmentCommentQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ShipmentCommentQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ShipmentCommentQuery) Exist(ctx context.Context) (bool, error) {
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
func (scq *ShipmentCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ShipmentCommentQuery) Clone() *ShipmentCommentQuery {
	if scq == nil {
		return nil
	}
	return &ShipmentCommentQuery{
		config:            scq.config,
		ctx:               scq.ctx.Clone(),
		order:             append([]shipmentcomment.OrderOption{}, scq.order...),
		inters:            append([]Interceptor{}, scq.inters...),
		predicates:        append([]predicate.ShipmentComment{}, scq.predicates...),
		withBusinessUnit:  scq.withBusinessUnit.Clone(),
		withOrganization:  scq.withOrganization.Clone(),
		withShipment:      scq.withShipment.Clone(),
		withCommentType:   scq.withCommentType.Clone(),
		withCreatedByUser: scq.withCreatedByUser.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommentQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *ShipmentCommentQuery {
	query := (&BusinessUnitClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withBusinessUnit = query
	return scq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommentQuery) WithOrganization(opts ...func(*OrganizationQuery)) *ShipmentCommentQuery {
	query := (&OrganizationClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withOrganization = query
	return scq
}

// WithShipment tells the query-builder to eager-load the nodes that are connected to
// the "shipment" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommentQuery) WithShipment(opts ...func(*ShipmentQuery)) *ShipmentCommentQuery {
	query := (&ShipmentClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withShipment = query
	return scq
}

// WithCommentType tells the query-builder to eager-load the nodes that are connected to
// the "comment_type" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommentQuery) WithCommentType(opts ...func(*CommentTypeQuery)) *ShipmentCommentQuery {
	query := (&CommentTypeClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withCommentType = query
	return scq
}

// WithCreatedByUser tells the query-builder to eager-load the nodes that are connected to
// the "created_by_user" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentCommentQuery) WithCreatedByUser(opts ...func(*UserQuery)) *ShipmentCommentQuery {
	query := (&UserClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withCreatedByUser = query
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
//	client.ShipmentComment.Query().
//		GroupBy(shipmentcomment.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ShipmentCommentQuery) GroupBy(field string, fields ...string) *ShipmentCommentGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShipmentCommentGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = shipmentcomment.Label
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
//	client.ShipmentComment.Query().
//		Select(shipmentcomment.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (scq *ShipmentCommentQuery) Select(fields ...string) *ShipmentCommentSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ShipmentCommentSelect{ShipmentCommentQuery: scq}
	sbuild.label = shipmentcomment.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShipmentCommentSelect configured with the given aggregations.
func (scq *ShipmentCommentQuery) Aggregate(fns ...AggregateFunc) *ShipmentCommentSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ShipmentCommentQuery) prepareQuery(ctx context.Context) error {
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
		if !shipmentcomment.ValidColumn(f) {
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

func (scq *ShipmentCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShipmentComment, error) {
	var (
		nodes       = []*ShipmentComment{}
		_spec       = scq.querySpec()
		loadedTypes = [5]bool{
			scq.withBusinessUnit != nil,
			scq.withOrganization != nil,
			scq.withShipment != nil,
			scq.withCommentType != nil,
			scq.withCreatedByUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShipmentComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShipmentComment{config: scq.config}
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
			func(n *ShipmentComment, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withOrganization; query != nil {
		if err := scq.loadOrganization(ctx, query, nodes, nil,
			func(n *ShipmentComment, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withShipment; query != nil {
		if err := scq.loadShipment(ctx, query, nodes, nil,
			func(n *ShipmentComment, e *Shipment) { n.Edges.Shipment = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withCommentType; query != nil {
		if err := scq.loadCommentType(ctx, query, nodes, nil,
			func(n *ShipmentComment, e *CommentType) { n.Edges.CommentType = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withCreatedByUser; query != nil {
		if err := scq.loadCreatedByUser(ctx, query, nodes, nil,
			func(n *ShipmentComment, e *User) { n.Edges.CreatedByUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ShipmentCommentQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*ShipmentComment, init func(*ShipmentComment), assign func(*ShipmentComment, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentComment)
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
func (scq *ShipmentCommentQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*ShipmentComment, init func(*ShipmentComment), assign func(*ShipmentComment, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentComment)
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
func (scq *ShipmentCommentQuery) loadShipment(ctx context.Context, query *ShipmentQuery, nodes []*ShipmentComment, init func(*ShipmentComment), assign func(*ShipmentComment, *Shipment)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentComment)
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
func (scq *ShipmentCommentQuery) loadCommentType(ctx context.Context, query *CommentTypeQuery, nodes []*ShipmentComment, init func(*ShipmentComment), assign func(*ShipmentComment, *CommentType)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentComment)
	for i := range nodes {
		fk := nodes[i].CommentTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(commenttype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "comment_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (scq *ShipmentCommentQuery) loadCreatedByUser(ctx context.Context, query *UserQuery, nodes []*ShipmentComment, init func(*ShipmentComment), assign func(*ShipmentComment, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentComment)
	for i := range nodes {
		fk := nodes[i].CreatedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *ShipmentCommentQuery) sqlCount(ctx context.Context) (int, error) {
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

func (scq *ShipmentCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shipmentcomment.Table, shipmentcomment.Columns, sqlgraph.NewFieldSpec(shipmentcomment.FieldID, field.TypeUUID))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentcomment.FieldID)
		for i := range fields {
			if fields[i] != shipmentcomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if scq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(shipmentcomment.FieldBusinessUnitID)
		}
		if scq.withOrganization != nil {
			_spec.Node.AddColumnOnce(shipmentcomment.FieldOrganizationID)
		}
		if scq.withShipment != nil {
			_spec.Node.AddColumnOnce(shipmentcomment.FieldShipmentID)
		}
		if scq.withCommentType != nil {
			_spec.Node.AddColumnOnce(shipmentcomment.FieldCommentTypeID)
		}
		if scq.withCreatedByUser != nil {
			_spec.Node.AddColumnOnce(shipmentcomment.FieldCreatedBy)
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

func (scq *ShipmentCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(shipmentcomment.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = shipmentcomment.Columns
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
func (scq *ShipmentCommentQuery) Modify(modifiers ...func(s *sql.Selector)) *ShipmentCommentSelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ShipmentCommentGroupBy is the group-by builder for ShipmentComment entities.
type ShipmentCommentGroupBy struct {
	selector
	build *ShipmentCommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ShipmentCommentGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentCommentGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ShipmentCommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentCommentQuery, *ShipmentCommentGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ShipmentCommentGroupBy) sqlScan(ctx context.Context, root *ShipmentCommentQuery, v any) error {
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

// ShipmentCommentSelect is the builder for selecting fields of ShipmentComment entities.
type ShipmentCommentSelect struct {
	*ShipmentCommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ShipmentCommentSelect) Aggregate(fns ...AggregateFunc) *ShipmentCommentSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ShipmentCommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentCommentQuery, *ShipmentCommentSelect](ctx, scs.ShipmentCommentQuery, scs, scs.inters, v)
}

func (scs *ShipmentCommentSelect) sqlScan(ctx context.Context, root *ShipmentCommentQuery, v any) error {
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
func (scs *ShipmentCommentSelect) Modify(modifiers ...func(s *sql.Selector)) *ShipmentCommentSelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
