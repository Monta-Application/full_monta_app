// Code generated by entc, DO NOT EDIT.

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
	"github.com/emoss08/trenova/ent/commenttype"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmentcomment"
	"github.com/google/uuid"
)

// CommentTypeQuery is the builder for querying CommentType entities.
type CommentTypeQuery struct {
	config
	ctx                       *QueryContext
	order                     []commenttype.OrderOption
	inters                    []Interceptor
	predicates                []predicate.CommentType
	withBusinessUnit          *BusinessUnitQuery
	withOrganization          *OrganizationQuery
	withShipmentComments      *ShipmentCommentQuery
	modifiers                 []func(*sql.Selector)
	withNamedShipmentComments map[string]*ShipmentCommentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommentTypeQuery builder.
func (ctq *CommentTypeQuery) Where(ps ...predicate.CommentType) *CommentTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *CommentTypeQuery) Limit(limit int) *CommentTypeQuery {
	ctq.ctx.Limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *CommentTypeQuery) Offset(offset int) *CommentTypeQuery {
	ctq.ctx.Offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CommentTypeQuery) Unique(unique bool) *CommentTypeQuery {
	ctq.ctx.Unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *CommentTypeQuery) Order(o ...commenttype.OrderOption) *CommentTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (ctq *CommentTypeQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commenttype.Table, commenttype.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, commenttype.BusinessUnitTable, commenttype.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (ctq *CommentTypeQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commenttype.Table, commenttype.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, commenttype.OrganizationTable, commenttype.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShipmentComments chains the current query on the "shipment_comments" edge.
func (ctq *CommentTypeQuery) QueryShipmentComments() *ShipmentCommentQuery {
	query := (&ShipmentCommentClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commenttype.Table, commenttype.FieldID, selector),
			sqlgraph.To(shipmentcomment.Table, shipmentcomment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, commenttype.ShipmentCommentsTable, commenttype.ShipmentCommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommentType entity from the query.
// Returns a *NotFoundError when no CommentType was found.
func (ctq *CommentTypeQuery) First(ctx context.Context) (*CommentType, error) {
	nodes, err := ctq.Limit(1).All(setContextOp(ctx, ctq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commenttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CommentTypeQuery) FirstX(ctx context.Context) *CommentType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommentType ID from the query.
// Returns a *NotFoundError when no CommentType ID was found.
func (ctq *CommentTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(1).IDs(setContextOp(ctx, ctq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CommentTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommentType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommentType entity is found.
// Returns a *NotFoundError when no CommentType entities are found.
func (ctq *CommentTypeQuery) Only(ctx context.Context) (*CommentType, error) {
	nodes, err := ctq.Limit(2).All(setContextOp(ctx, ctq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commenttype.Label}
	default:
		return nil, &NotSingularError{commenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CommentTypeQuery) OnlyX(ctx context.Context) *CommentType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommentType ID in the query.
// Returns a *NotSingularError when more than one CommentType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CommentTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ctq.Limit(2).IDs(setContextOp(ctx, ctq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commenttype.Label}
	default:
		err = &NotSingularError{commenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CommentTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommentTypes.
func (ctq *CommentTypeQuery) All(ctx context.Context) ([]*CommentType, error) {
	ctx = setContextOp(ctx, ctq.ctx, "All")
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommentType, *CommentTypeQuery]()
	return withInterceptors[[]*CommentType](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CommentTypeQuery) AllX(ctx context.Context) []*CommentType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommentType IDs.
func (ctq *CommentTypeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ctq.ctx.Unique == nil && ctq.path != nil {
		ctq.Unique(true)
	}
	ctx = setContextOp(ctx, ctq.ctx, "IDs")
	if err = ctq.Select(commenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CommentTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CommentTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Count")
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*CommentTypeQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CommentTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CommentTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Exist")
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CommentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CommentTypeQuery) Clone() *CommentTypeQuery {
	if ctq == nil {
		return nil
	}
	return &CommentTypeQuery{
		config:               ctq.config,
		ctx:                  ctq.ctx.Clone(),
		order:                append([]commenttype.OrderOption{}, ctq.order...),
		inters:               append([]Interceptor{}, ctq.inters...),
		predicates:           append([]predicate.CommentType{}, ctq.predicates...),
		withBusinessUnit:     ctq.withBusinessUnit.Clone(),
		withOrganization:     ctq.withOrganization.Clone(),
		withShipmentComments: ctq.withShipmentComments.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CommentTypeQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *CommentTypeQuery {
	query := (&BusinessUnitClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withBusinessUnit = query
	return ctq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CommentTypeQuery) WithOrganization(opts ...func(*OrganizationQuery)) *CommentTypeQuery {
	query := (&OrganizationClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withOrganization = query
	return ctq
}

// WithShipmentComments tells the query-builder to eager-load the nodes that are connected to
// the "shipment_comments" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CommentTypeQuery) WithShipmentComments(opts ...func(*ShipmentCommentQuery)) *CommentTypeQuery {
	query := (&ShipmentCommentClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withShipmentComments = query
	return ctq
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
//	client.CommentType.Query().
//		GroupBy(commenttype.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *CommentTypeQuery) GroupBy(field string, fields ...string) *CommentTypeGroupBy {
	ctq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentTypeGroupBy{build: ctq}
	grbuild.flds = &ctq.ctx.Fields
	grbuild.label = commenttype.Label
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
//	client.CommentType.Query().
//		Select(commenttype.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (ctq *CommentTypeQuery) Select(fields ...string) *CommentTypeSelect {
	ctq.ctx.Fields = append(ctq.ctx.Fields, fields...)
	sbuild := &CommentTypeSelect{CommentTypeQuery: ctq}
	sbuild.label = commenttype.Label
	sbuild.flds, sbuild.scan = &ctq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentTypeSelect configured with the given aggregations.
func (ctq *CommentTypeQuery) Aggregate(fns ...AggregateFunc) *CommentTypeSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *CommentTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.ctx.Fields {
		if !commenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CommentTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommentType, error) {
	var (
		nodes       = []*CommentType{}
		_spec       = ctq.querySpec()
		loadedTypes = [3]bool{
			ctq.withBusinessUnit != nil,
			ctq.withOrganization != nil,
			ctq.withShipmentComments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommentType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommentType{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withBusinessUnit; query != nil {
		if err := ctq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *CommentType, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := ctq.withOrganization; query != nil {
		if err := ctq.loadOrganization(ctx, query, nodes, nil,
			func(n *CommentType, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := ctq.withShipmentComments; query != nil {
		if err := ctq.loadShipmentComments(ctx, query, nodes,
			func(n *CommentType) { n.Edges.ShipmentComments = []*ShipmentComment{} },
			func(n *CommentType, e *ShipmentComment) {
				n.Edges.ShipmentComments = append(n.Edges.ShipmentComments, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range ctq.withNamedShipmentComments {
		if err := ctq.loadShipmentComments(ctx, query, nodes,
			func(n *CommentType) { n.appendNamedShipmentComments(name) },
			func(n *CommentType, e *ShipmentComment) { n.appendNamedShipmentComments(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *CommentTypeQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*CommentType, init func(*CommentType), assign func(*CommentType, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CommentType)
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
func (ctq *CommentTypeQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*CommentType, init func(*CommentType), assign func(*CommentType, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*CommentType)
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
func (ctq *CommentTypeQuery) loadShipmentComments(ctx context.Context, query *ShipmentCommentQuery, nodes []*CommentType, init func(*CommentType), assign func(*CommentType, *ShipmentComment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CommentType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(shipmentcomment.FieldCommentTypeID)
	}
	query.Where(predicate.ShipmentComment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(commenttype.ShipmentCommentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CommentTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "comment_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ctq *CommentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	_spec.Node.Columns = ctq.ctx.Fields
	if len(ctq.ctx.Fields) > 0 {
		_spec.Unique = ctq.ctx.Unique != nil && *ctq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CommentTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commenttype.Table, commenttype.Columns, sqlgraph.NewFieldSpec(commenttype.FieldID, field.TypeUUID))
	_spec.From = ctq.sql
	if unique := ctq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ctq.path != nil {
		_spec.Unique = true
	}
	if fields := ctq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commenttype.FieldID)
		for i := range fields {
			if fields[i] != commenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ctq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(commenttype.FieldBusinessUnitID)
		}
		if ctq.withOrganization != nil {
			_spec.Node.AddColumnOnce(commenttype.FieldOrganizationID)
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CommentTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(commenttype.Table)
	columns := ctq.ctx.Fields
	if len(columns) == 0 {
		columns = commenttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.ctx.Unique != nil && *ctq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ctq.modifiers {
		m(selector)
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ctq *CommentTypeQuery) Modify(modifiers ...func(s *sql.Selector)) *CommentTypeSelect {
	ctq.modifiers = append(ctq.modifiers, modifiers...)
	return ctq.Select()
}

// WithNamedShipmentComments tells the query-builder to eager-load the nodes that are connected to the "shipment_comments"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ctq *CommentTypeQuery) WithNamedShipmentComments(name string, opts ...func(*ShipmentCommentQuery)) *CommentTypeQuery {
	query := (&ShipmentCommentClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ctq.withNamedShipmentComments == nil {
		ctq.withNamedShipmentComments = make(map[string]*ShipmentCommentQuery)
	}
	ctq.withNamedShipmentComments[name] = query
	return ctq
}

// CommentTypeGroupBy is the group-by builder for CommentType entities.
type CommentTypeGroupBy struct {
	selector
	build *CommentTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CommentTypeGroupBy) Aggregate(fns ...AggregateFunc) *CommentTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *CommentTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ctgb.build.ctx, "GroupBy")
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentTypeQuery, *CommentTypeGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *CommentTypeGroupBy) sqlScan(ctx context.Context, root *CommentTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommentTypeSelect is the builder for selecting fields of CommentType entities.
type CommentTypeSelect struct {
	*CommentTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *CommentTypeSelect) Aggregate(fns ...AggregateFunc) *CommentTypeSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CommentTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cts.ctx, "Select")
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentTypeQuery, *CommentTypeSelect](ctx, cts.CommentTypeQuery, cts, cts.inters, v)
}

func (cts *CommentTypeSelect) sqlScan(ctx context.Context, root *CommentTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cts *CommentTypeSelect) Modify(modifiers ...func(s *sql.Selector)) *CommentTypeSelect {
	cts.modifiers = append(cts.modifiers, modifiers...)
	return cts
}
