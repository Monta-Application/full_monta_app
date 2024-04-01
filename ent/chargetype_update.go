// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/chargetype"
	"github.com/emoss08/trenova/ent/predicate"
)

// ChargeTypeUpdate is the builder for updating ChargeType entities.
type ChargeTypeUpdate struct {
	config
	hooks     []Hook
	mutation  *ChargeTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ChargeTypeUpdate builder.
func (ctu *ChargeTypeUpdate) Where(ps ...predicate.ChargeType) *ChargeTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetUpdatedAt sets the "updated_at" field.
func (ctu *ChargeTypeUpdate) SetUpdatedAt(t time.Time) *ChargeTypeUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// SetVersion sets the "version" field.
func (ctu *ChargeTypeUpdate) SetVersion(i int) *ChargeTypeUpdate {
	ctu.mutation.ResetVersion()
	ctu.mutation.SetVersion(i)
	return ctu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ctu *ChargeTypeUpdate) SetNillableVersion(i *int) *ChargeTypeUpdate {
	if i != nil {
		ctu.SetVersion(*i)
	}
	return ctu
}

// AddVersion adds i to the "version" field.
func (ctu *ChargeTypeUpdate) AddVersion(i int) *ChargeTypeUpdate {
	ctu.mutation.AddVersion(i)
	return ctu
}

// SetStatus sets the "status" field.
func (ctu *ChargeTypeUpdate) SetStatus(c chargetype.Status) *ChargeTypeUpdate {
	ctu.mutation.SetStatus(c)
	return ctu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ctu *ChargeTypeUpdate) SetNillableStatus(c *chargetype.Status) *ChargeTypeUpdate {
	if c != nil {
		ctu.SetStatus(*c)
	}
	return ctu
}

// SetName sets the "name" field.
func (ctu *ChargeTypeUpdate) SetName(s string) *ChargeTypeUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ctu *ChargeTypeUpdate) SetNillableName(s *string) *ChargeTypeUpdate {
	if s != nil {
		ctu.SetName(*s)
	}
	return ctu
}

// SetDescription sets the "description" field.
func (ctu *ChargeTypeUpdate) SetDescription(s string) *ChargeTypeUpdate {
	ctu.mutation.SetDescription(s)
	return ctu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ctu *ChargeTypeUpdate) SetNillableDescription(s *string) *ChargeTypeUpdate {
	if s != nil {
		ctu.SetDescription(*s)
	}
	return ctu
}

// ClearDescription clears the value of the "description" field.
func (ctu *ChargeTypeUpdate) ClearDescription() *ChargeTypeUpdate {
	ctu.mutation.ClearDescription()
	return ctu
}

// Mutation returns the ChargeTypeMutation object of the builder.
func (ctu *ChargeTypeUpdate) Mutation() *ChargeTypeMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *ChargeTypeUpdate) Save(ctx context.Context) (int, error) {
	ctu.defaults()
	return withHooks(ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *ChargeTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *ChargeTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *ChargeTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *ChargeTypeUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := chargetype.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *ChargeTypeUpdate) check() error {
	if v, ok := ctu.mutation.Status(); ok {
		if err := chargetype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ChargeType.status": %w`, err)}
		}
	}
	if v, ok := ctu.mutation.Name(); ok {
		if err := chargetype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ChargeType.name": %w`, err)}
		}
	}
	if _, ok := ctu.mutation.BusinessUnitID(); ctu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChargeType.business_unit"`)
	}
	if _, ok := ctu.mutation.OrganizationID(); ctu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChargeType.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctu *ChargeTypeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChargeTypeUpdate {
	ctu.modifiers = append(ctu.modifiers, modifiers...)
	return ctu
}

func (ctu *ChargeTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chargetype.Table, chargetype.Columns, sqlgraph.NewFieldSpec(chargetype.FieldID, field.TypeUUID))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.SetField(chargetype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.Version(); ok {
		_spec.SetField(chargetype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ctu.mutation.AddedVersion(); ok {
		_spec.AddField(chargetype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ctu.mutation.Status(); ok {
		_spec.SetField(chargetype.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.SetField(chargetype.FieldName, field.TypeString, value)
	}
	if value, ok := ctu.mutation.Description(); ok {
		_spec.SetField(chargetype.FieldDescription, field.TypeString, value)
	}
	if ctu.mutation.DescriptionCleared() {
		_spec.ClearField(chargetype.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(ctu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chargetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// ChargeTypeUpdateOne is the builder for updating a single ChargeType entity.
type ChargeTypeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ChargeTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ctuo *ChargeTypeUpdateOne) SetUpdatedAt(t time.Time) *ChargeTypeUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// SetVersion sets the "version" field.
func (ctuo *ChargeTypeUpdateOne) SetVersion(i int) *ChargeTypeUpdateOne {
	ctuo.mutation.ResetVersion()
	ctuo.mutation.SetVersion(i)
	return ctuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ctuo *ChargeTypeUpdateOne) SetNillableVersion(i *int) *ChargeTypeUpdateOne {
	if i != nil {
		ctuo.SetVersion(*i)
	}
	return ctuo
}

// AddVersion adds i to the "version" field.
func (ctuo *ChargeTypeUpdateOne) AddVersion(i int) *ChargeTypeUpdateOne {
	ctuo.mutation.AddVersion(i)
	return ctuo
}

// SetStatus sets the "status" field.
func (ctuo *ChargeTypeUpdateOne) SetStatus(c chargetype.Status) *ChargeTypeUpdateOne {
	ctuo.mutation.SetStatus(c)
	return ctuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ctuo *ChargeTypeUpdateOne) SetNillableStatus(c *chargetype.Status) *ChargeTypeUpdateOne {
	if c != nil {
		ctuo.SetStatus(*c)
	}
	return ctuo
}

// SetName sets the "name" field.
func (ctuo *ChargeTypeUpdateOne) SetName(s string) *ChargeTypeUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ctuo *ChargeTypeUpdateOne) SetNillableName(s *string) *ChargeTypeUpdateOne {
	if s != nil {
		ctuo.SetName(*s)
	}
	return ctuo
}

// SetDescription sets the "description" field.
func (ctuo *ChargeTypeUpdateOne) SetDescription(s string) *ChargeTypeUpdateOne {
	ctuo.mutation.SetDescription(s)
	return ctuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ctuo *ChargeTypeUpdateOne) SetNillableDescription(s *string) *ChargeTypeUpdateOne {
	if s != nil {
		ctuo.SetDescription(*s)
	}
	return ctuo
}

// ClearDescription clears the value of the "description" field.
func (ctuo *ChargeTypeUpdateOne) ClearDescription() *ChargeTypeUpdateOne {
	ctuo.mutation.ClearDescription()
	return ctuo
}

// Mutation returns the ChargeTypeMutation object of the builder.
func (ctuo *ChargeTypeUpdateOne) Mutation() *ChargeTypeMutation {
	return ctuo.mutation
}

// Where appends a list predicates to the ChargeTypeUpdate builder.
func (ctuo *ChargeTypeUpdateOne) Where(ps ...predicate.ChargeType) *ChargeTypeUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *ChargeTypeUpdateOne) Select(field string, fields ...string) *ChargeTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated ChargeType entity.
func (ctuo *ChargeTypeUpdateOne) Save(ctx context.Context) (*ChargeType, error) {
	ctuo.defaults()
	return withHooks(ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *ChargeTypeUpdateOne) SaveX(ctx context.Context) *ChargeType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *ChargeTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *ChargeTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *ChargeTypeUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := chargetype.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *ChargeTypeUpdateOne) check() error {
	if v, ok := ctuo.mutation.Status(); ok {
		if err := chargetype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ChargeType.status": %w`, err)}
		}
	}
	if v, ok := ctuo.mutation.Name(); ok {
		if err := chargetype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ChargeType.name": %w`, err)}
		}
	}
	if _, ok := ctuo.mutation.BusinessUnitID(); ctuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChargeType.business_unit"`)
	}
	if _, ok := ctuo.mutation.OrganizationID(); ctuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ChargeType.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctuo *ChargeTypeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChargeTypeUpdateOne {
	ctuo.modifiers = append(ctuo.modifiers, modifiers...)
	return ctuo
}

func (ctuo *ChargeTypeUpdateOne) sqlSave(ctx context.Context) (_node *ChargeType, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chargetype.Table, chargetype.Columns, sqlgraph.NewFieldSpec(chargetype.FieldID, field.TypeUUID))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChargeType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chargetype.FieldID)
		for _, f := range fields {
			if !chargetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chargetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(chargetype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.Version(); ok {
		_spec.SetField(chargetype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ctuo.mutation.AddedVersion(); ok {
		_spec.AddField(chargetype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ctuo.mutation.Status(); ok {
		_spec.SetField(chargetype.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.SetField(chargetype.FieldName, field.TypeString, value)
	}
	if value, ok := ctuo.mutation.Description(); ok {
		_spec.SetField(chargetype.FieldDescription, field.TypeString, value)
	}
	if ctuo.mutation.DescriptionCleared() {
		_spec.ClearField(chargetype.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(ctuo.modifiers...)
	_node = &ChargeType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chargetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
