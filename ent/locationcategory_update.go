// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/locationcategory"
	"github.com/emoss08/trenova/ent/predicate"
)

// LocationCategoryUpdate is the builder for updating LocationCategory entities.
type LocationCategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *LocationCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LocationCategoryUpdate builder.
func (lcu *LocationCategoryUpdate) Where(ps ...predicate.LocationCategory) *LocationCategoryUpdate {
	lcu.mutation.Where(ps...)
	return lcu
}

// SetUpdatedAt sets the "updated_at" field.
func (lcu *LocationCategoryUpdate) SetUpdatedAt(t time.Time) *LocationCategoryUpdate {
	lcu.mutation.SetUpdatedAt(t)
	return lcu
}

// SetVersion sets the "version" field.
func (lcu *LocationCategoryUpdate) SetVersion(i int) *LocationCategoryUpdate {
	lcu.mutation.ResetVersion()
	lcu.mutation.SetVersion(i)
	return lcu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (lcu *LocationCategoryUpdate) SetNillableVersion(i *int) *LocationCategoryUpdate {
	if i != nil {
		lcu.SetVersion(*i)
	}
	return lcu
}

// AddVersion adds i to the "version" field.
func (lcu *LocationCategoryUpdate) AddVersion(i int) *LocationCategoryUpdate {
	lcu.mutation.AddVersion(i)
	return lcu
}

// SetName sets the "name" field.
func (lcu *LocationCategoryUpdate) SetName(s string) *LocationCategoryUpdate {
	lcu.mutation.SetName(s)
	return lcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lcu *LocationCategoryUpdate) SetNillableName(s *string) *LocationCategoryUpdate {
	if s != nil {
		lcu.SetName(*s)
	}
	return lcu
}

// SetDescription sets the "description" field.
func (lcu *LocationCategoryUpdate) SetDescription(s string) *LocationCategoryUpdate {
	lcu.mutation.SetDescription(s)
	return lcu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lcu *LocationCategoryUpdate) SetNillableDescription(s *string) *LocationCategoryUpdate {
	if s != nil {
		lcu.SetDescription(*s)
	}
	return lcu
}

// ClearDescription clears the value of the "description" field.
func (lcu *LocationCategoryUpdate) ClearDescription() *LocationCategoryUpdate {
	lcu.mutation.ClearDescription()
	return lcu
}

// SetColor sets the "color" field.
func (lcu *LocationCategoryUpdate) SetColor(s string) *LocationCategoryUpdate {
	lcu.mutation.SetColor(s)
	return lcu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (lcu *LocationCategoryUpdate) SetNillableColor(s *string) *LocationCategoryUpdate {
	if s != nil {
		lcu.SetColor(*s)
	}
	return lcu
}

// ClearColor clears the value of the "color" field.
func (lcu *LocationCategoryUpdate) ClearColor() *LocationCategoryUpdate {
	lcu.mutation.ClearColor()
	return lcu
}

// Mutation returns the LocationCategoryMutation object of the builder.
func (lcu *LocationCategoryUpdate) Mutation() *LocationCategoryMutation {
	return lcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lcu *LocationCategoryUpdate) Save(ctx context.Context) (int, error) {
	lcu.defaults()
	return withHooks(ctx, lcu.sqlSave, lcu.mutation, lcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcu *LocationCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := lcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lcu *LocationCategoryUpdate) Exec(ctx context.Context) error {
	_, err := lcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcu *LocationCategoryUpdate) ExecX(ctx context.Context) {
	if err := lcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcu *LocationCategoryUpdate) defaults() {
	if _, ok := lcu.mutation.UpdatedAt(); !ok {
		v := locationcategory.UpdateDefaultUpdatedAt()
		lcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcu *LocationCategoryUpdate) check() error {
	if v, ok := lcu.mutation.Name(); ok {
		if err := locationcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LocationCategory.name": %w`, err)}
		}
	}
	if _, ok := lcu.mutation.BusinessUnitID(); lcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationCategory.business_unit"`)
	}
	if _, ok := lcu.mutation.OrganizationID(); lcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationCategory.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lcu *LocationCategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LocationCategoryUpdate {
	lcu.modifiers = append(lcu.modifiers, modifiers...)
	return lcu
}

func (lcu *LocationCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(locationcategory.Table, locationcategory.Columns, sqlgraph.NewFieldSpec(locationcategory.FieldID, field.TypeUUID))
	if ps := lcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcu.mutation.UpdatedAt(); ok {
		_spec.SetField(locationcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lcu.mutation.Version(); ok {
		_spec.SetField(locationcategory.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcu.mutation.AddedVersion(); ok {
		_spec.AddField(locationcategory.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcu.mutation.Name(); ok {
		_spec.SetField(locationcategory.FieldName, field.TypeString, value)
	}
	if value, ok := lcu.mutation.Description(); ok {
		_spec.SetField(locationcategory.FieldDescription, field.TypeString, value)
	}
	if lcu.mutation.DescriptionCleared() {
		_spec.ClearField(locationcategory.FieldDescription, field.TypeString)
	}
	if value, ok := lcu.mutation.Color(); ok {
		_spec.SetField(locationcategory.FieldColor, field.TypeString, value)
	}
	if lcu.mutation.ColorCleared() {
		_spec.ClearField(locationcategory.FieldColor, field.TypeString)
	}
	_spec.AddModifiers(lcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{locationcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lcu.mutation.done = true
	return n, nil
}

// LocationCategoryUpdateOne is the builder for updating a single LocationCategory entity.
type LocationCategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LocationCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (lcuo *LocationCategoryUpdateOne) SetUpdatedAt(t time.Time) *LocationCategoryUpdateOne {
	lcuo.mutation.SetUpdatedAt(t)
	return lcuo
}

// SetVersion sets the "version" field.
func (lcuo *LocationCategoryUpdateOne) SetVersion(i int) *LocationCategoryUpdateOne {
	lcuo.mutation.ResetVersion()
	lcuo.mutation.SetVersion(i)
	return lcuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (lcuo *LocationCategoryUpdateOne) SetNillableVersion(i *int) *LocationCategoryUpdateOne {
	if i != nil {
		lcuo.SetVersion(*i)
	}
	return lcuo
}

// AddVersion adds i to the "version" field.
func (lcuo *LocationCategoryUpdateOne) AddVersion(i int) *LocationCategoryUpdateOne {
	lcuo.mutation.AddVersion(i)
	return lcuo
}

// SetName sets the "name" field.
func (lcuo *LocationCategoryUpdateOne) SetName(s string) *LocationCategoryUpdateOne {
	lcuo.mutation.SetName(s)
	return lcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lcuo *LocationCategoryUpdateOne) SetNillableName(s *string) *LocationCategoryUpdateOne {
	if s != nil {
		lcuo.SetName(*s)
	}
	return lcuo
}

// SetDescription sets the "description" field.
func (lcuo *LocationCategoryUpdateOne) SetDescription(s string) *LocationCategoryUpdateOne {
	lcuo.mutation.SetDescription(s)
	return lcuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lcuo *LocationCategoryUpdateOne) SetNillableDescription(s *string) *LocationCategoryUpdateOne {
	if s != nil {
		lcuo.SetDescription(*s)
	}
	return lcuo
}

// ClearDescription clears the value of the "description" field.
func (lcuo *LocationCategoryUpdateOne) ClearDescription() *LocationCategoryUpdateOne {
	lcuo.mutation.ClearDescription()
	return lcuo
}

// SetColor sets the "color" field.
func (lcuo *LocationCategoryUpdateOne) SetColor(s string) *LocationCategoryUpdateOne {
	lcuo.mutation.SetColor(s)
	return lcuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (lcuo *LocationCategoryUpdateOne) SetNillableColor(s *string) *LocationCategoryUpdateOne {
	if s != nil {
		lcuo.SetColor(*s)
	}
	return lcuo
}

// ClearColor clears the value of the "color" field.
func (lcuo *LocationCategoryUpdateOne) ClearColor() *LocationCategoryUpdateOne {
	lcuo.mutation.ClearColor()
	return lcuo
}

// Mutation returns the LocationCategoryMutation object of the builder.
func (lcuo *LocationCategoryUpdateOne) Mutation() *LocationCategoryMutation {
	return lcuo.mutation
}

// Where appends a list predicates to the LocationCategoryUpdate builder.
func (lcuo *LocationCategoryUpdateOne) Where(ps ...predicate.LocationCategory) *LocationCategoryUpdateOne {
	lcuo.mutation.Where(ps...)
	return lcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lcuo *LocationCategoryUpdateOne) Select(field string, fields ...string) *LocationCategoryUpdateOne {
	lcuo.fields = append([]string{field}, fields...)
	return lcuo
}

// Save executes the query and returns the updated LocationCategory entity.
func (lcuo *LocationCategoryUpdateOne) Save(ctx context.Context) (*LocationCategory, error) {
	lcuo.defaults()
	return withHooks(ctx, lcuo.sqlSave, lcuo.mutation, lcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcuo *LocationCategoryUpdateOne) SaveX(ctx context.Context) *LocationCategory {
	node, err := lcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lcuo *LocationCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := lcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcuo *LocationCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := lcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcuo *LocationCategoryUpdateOne) defaults() {
	if _, ok := lcuo.mutation.UpdatedAt(); !ok {
		v := locationcategory.UpdateDefaultUpdatedAt()
		lcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcuo *LocationCategoryUpdateOne) check() error {
	if v, ok := lcuo.mutation.Name(); ok {
		if err := locationcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LocationCategory.name": %w`, err)}
		}
	}
	if _, ok := lcuo.mutation.BusinessUnitID(); lcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationCategory.business_unit"`)
	}
	if _, ok := lcuo.mutation.OrganizationID(); lcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LocationCategory.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lcuo *LocationCategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LocationCategoryUpdateOne {
	lcuo.modifiers = append(lcuo.modifiers, modifiers...)
	return lcuo
}

func (lcuo *LocationCategoryUpdateOne) sqlSave(ctx context.Context) (_node *LocationCategory, err error) {
	if err := lcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(locationcategory.Table, locationcategory.Columns, sqlgraph.NewFieldSpec(locationcategory.FieldID, field.TypeUUID))
	id, ok := lcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LocationCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, locationcategory.FieldID)
		for _, f := range fields {
			if !locationcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != locationcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(locationcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lcuo.mutation.Version(); ok {
		_spec.SetField(locationcategory.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcuo.mutation.AddedVersion(); ok {
		_spec.AddField(locationcategory.FieldVersion, field.TypeInt, value)
	}
	if value, ok := lcuo.mutation.Name(); ok {
		_spec.SetField(locationcategory.FieldName, field.TypeString, value)
	}
	if value, ok := lcuo.mutation.Description(); ok {
		_spec.SetField(locationcategory.FieldDescription, field.TypeString, value)
	}
	if lcuo.mutation.DescriptionCleared() {
		_spec.ClearField(locationcategory.FieldDescription, field.TypeString)
	}
	if value, ok := lcuo.mutation.Color(); ok {
		_spec.SetField(locationcategory.FieldColor, field.TypeString, value)
	}
	if lcuo.mutation.ColorCleared() {
		_spec.ClearField(locationcategory.FieldColor, field.TypeString)
	}
	_spec.AddModifiers(lcuo.modifiers...)
	_node = &LocationCategory{config: lcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{locationcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lcuo.mutation.done = true
	return _node, nil
}
