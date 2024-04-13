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
	"github.com/emoss08/trenova/ent/customer"
	"github.com/emoss08/trenova/ent/formulatemplate"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/shipmenttype"
	"github.com/google/uuid"
)

// FormulaTemplateUpdate is the builder for updating FormulaTemplate entities.
type FormulaTemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *FormulaTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FormulaTemplateUpdate builder.
func (ftu *FormulaTemplateUpdate) Where(ps ...predicate.FormulaTemplate) *FormulaTemplateUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetUpdatedAt sets the "updated_at" field.
func (ftu *FormulaTemplateUpdate) SetUpdatedAt(t time.Time) *FormulaTemplateUpdate {
	ftu.mutation.SetUpdatedAt(t)
	return ftu
}

// SetVersion sets the "version" field.
func (ftu *FormulaTemplateUpdate) SetVersion(i int) *FormulaTemplateUpdate {
	ftu.mutation.ResetVersion()
	ftu.mutation.SetVersion(i)
	return ftu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableVersion(i *int) *FormulaTemplateUpdate {
	if i != nil {
		ftu.SetVersion(*i)
	}
	return ftu
}

// AddVersion adds i to the "version" field.
func (ftu *FormulaTemplateUpdate) AddVersion(i int) *FormulaTemplateUpdate {
	ftu.mutation.AddVersion(i)
	return ftu
}

// SetName sets the "name" field.
func (ftu *FormulaTemplateUpdate) SetName(s string) *FormulaTemplateUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableName(s *string) *FormulaTemplateUpdate {
	if s != nil {
		ftu.SetName(*s)
	}
	return ftu
}

// SetFormulaText sets the "formula_text" field.
func (ftu *FormulaTemplateUpdate) SetFormulaText(s string) *FormulaTemplateUpdate {
	ftu.mutation.SetFormulaText(s)
	return ftu
}

// SetNillableFormulaText sets the "formula_text" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableFormulaText(s *string) *FormulaTemplateUpdate {
	if s != nil {
		ftu.SetFormulaText(*s)
	}
	return ftu
}

// SetDescription sets the "description" field.
func (ftu *FormulaTemplateUpdate) SetDescription(s string) *FormulaTemplateUpdate {
	ftu.mutation.SetDescription(s)
	return ftu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableDescription(s *string) *FormulaTemplateUpdate {
	if s != nil {
		ftu.SetDescription(*s)
	}
	return ftu
}

// ClearDescription clears the value of the "description" field.
func (ftu *FormulaTemplateUpdate) ClearDescription() *FormulaTemplateUpdate {
	ftu.mutation.ClearDescription()
	return ftu
}

// SetTemplateType sets the "template_type" field.
func (ftu *FormulaTemplateUpdate) SetTemplateType(ft formulatemplate.TemplateType) *FormulaTemplateUpdate {
	ftu.mutation.SetTemplateType(ft)
	return ftu
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableTemplateType(ft *formulatemplate.TemplateType) *FormulaTemplateUpdate {
	if ft != nil {
		ftu.SetTemplateType(*ft)
	}
	return ftu
}

// SetCustomerID sets the "customer_id" field.
func (ftu *FormulaTemplateUpdate) SetCustomerID(u uuid.UUID) *FormulaTemplateUpdate {
	ftu.mutation.SetCustomerID(u)
	return ftu
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableCustomerID(u *uuid.UUID) *FormulaTemplateUpdate {
	if u != nil {
		ftu.SetCustomerID(*u)
	}
	return ftu
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ftu *FormulaTemplateUpdate) ClearCustomerID() *FormulaTemplateUpdate {
	ftu.mutation.ClearCustomerID()
	return ftu
}

// SetShipmentTypeID sets the "shipment_type_id" field.
func (ftu *FormulaTemplateUpdate) SetShipmentTypeID(u uuid.UUID) *FormulaTemplateUpdate {
	ftu.mutation.SetShipmentTypeID(u)
	return ftu
}

// SetNillableShipmentTypeID sets the "shipment_type_id" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableShipmentTypeID(u *uuid.UUID) *FormulaTemplateUpdate {
	if u != nil {
		ftu.SetShipmentTypeID(*u)
	}
	return ftu
}

// ClearShipmentTypeID clears the value of the "shipment_type_id" field.
func (ftu *FormulaTemplateUpdate) ClearShipmentTypeID() *FormulaTemplateUpdate {
	ftu.mutation.ClearShipmentTypeID()
	return ftu
}

// SetAutoApply sets the "auto_apply" field.
func (ftu *FormulaTemplateUpdate) SetAutoApply(b bool) *FormulaTemplateUpdate {
	ftu.mutation.SetAutoApply(b)
	return ftu
}

// SetNillableAutoApply sets the "auto_apply" field if the given value is not nil.
func (ftu *FormulaTemplateUpdate) SetNillableAutoApply(b *bool) *FormulaTemplateUpdate {
	if b != nil {
		ftu.SetAutoApply(*b)
	}
	return ftu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ftu *FormulaTemplateUpdate) SetCustomer(c *Customer) *FormulaTemplateUpdate {
	return ftu.SetCustomerID(c.ID)
}

// SetShipmentType sets the "shipment_type" edge to the ShipmentType entity.
func (ftu *FormulaTemplateUpdate) SetShipmentType(s *ShipmentType) *FormulaTemplateUpdate {
	return ftu.SetShipmentTypeID(s.ID)
}

// Mutation returns the FormulaTemplateMutation object of the builder.
func (ftu *FormulaTemplateUpdate) Mutation() *FormulaTemplateMutation {
	return ftu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ftu *FormulaTemplateUpdate) ClearCustomer() *FormulaTemplateUpdate {
	ftu.mutation.ClearCustomer()
	return ftu
}

// ClearShipmentType clears the "shipment_type" edge to the ShipmentType entity.
func (ftu *FormulaTemplateUpdate) ClearShipmentType() *FormulaTemplateUpdate {
	ftu.mutation.ClearShipmentType()
	return ftu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FormulaTemplateUpdate) Save(ctx context.Context) (int, error) {
	ftu.defaults()
	return withHooks(ctx, ftu.sqlSave, ftu.mutation, ftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FormulaTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FormulaTemplateUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FormulaTemplateUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FormulaTemplateUpdate) defaults() {
	if _, ok := ftu.mutation.UpdatedAt(); !ok {
		v := formulatemplate.UpdateDefaultUpdatedAt()
		ftu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FormulaTemplateUpdate) check() error {
	if v, ok := ftu.mutation.Name(); ok {
		if err := formulatemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.name": %w`, err)}
		}
	}
	if v, ok := ftu.mutation.FormulaText(); ok {
		if err := formulatemplate.FormulaTextValidator(v); err != nil {
			return &ValidationError{Name: "formula_text", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.formula_text": %w`, err)}
		}
	}
	if v, ok := ftu.mutation.TemplateType(); ok {
		if err := formulatemplate.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.template_type": %w`, err)}
		}
	}
	if _, ok := ftu.mutation.BusinessUnitID(); ftu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FormulaTemplate.business_unit"`)
	}
	if _, ok := ftu.mutation.OrganizationID(); ftu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FormulaTemplate.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftu *FormulaTemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FormulaTemplateUpdate {
	ftu.modifiers = append(ftu.modifiers, modifiers...)
	return ftu
}

func (ftu *FormulaTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ftu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(formulatemplate.Table, formulatemplate.Columns, sqlgraph.NewFieldSpec(formulatemplate.FieldID, field.TypeUUID))
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UpdatedAt(); ok {
		_spec.SetField(formulatemplate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftu.mutation.Version(); ok {
		_spec.SetField(formulatemplate.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ftu.mutation.AddedVersion(); ok {
		_spec.AddField(formulatemplate.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ftu.mutation.Name(); ok {
		_spec.SetField(formulatemplate.FieldName, field.TypeString, value)
	}
	if value, ok := ftu.mutation.FormulaText(); ok {
		_spec.SetField(formulatemplate.FieldFormulaText, field.TypeString, value)
	}
	if value, ok := ftu.mutation.Description(); ok {
		_spec.SetField(formulatemplate.FieldDescription, field.TypeString, value)
	}
	if ftu.mutation.DescriptionCleared() {
		_spec.ClearField(formulatemplate.FieldDescription, field.TypeString)
	}
	if value, ok := ftu.mutation.TemplateType(); ok {
		_spec.SetField(formulatemplate.FieldTemplateType, field.TypeEnum, value)
	}
	if value, ok := ftu.mutation.AutoApply(); ok {
		_spec.SetField(formulatemplate.FieldAutoApply, field.TypeBool, value)
	}
	if ftu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.CustomerTable,
			Columns: []string{formulatemplate.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.CustomerTable,
			Columns: []string{formulatemplate.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftu.mutation.ShipmentTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.ShipmentTypeTable,
			Columns: []string{formulatemplate.ShipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.ShipmentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.ShipmentTypeTable,
			Columns: []string{formulatemplate.ShipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ftu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{formulatemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftu.mutation.done = true
	return n, nil
}

// FormulaTemplateUpdateOne is the builder for updating a single FormulaTemplate entity.
type FormulaTemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FormulaTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ftuo *FormulaTemplateUpdateOne) SetUpdatedAt(t time.Time) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetUpdatedAt(t)
	return ftuo
}

// SetVersion sets the "version" field.
func (ftuo *FormulaTemplateUpdateOne) SetVersion(i int) *FormulaTemplateUpdateOne {
	ftuo.mutation.ResetVersion()
	ftuo.mutation.SetVersion(i)
	return ftuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableVersion(i *int) *FormulaTemplateUpdateOne {
	if i != nil {
		ftuo.SetVersion(*i)
	}
	return ftuo
}

// AddVersion adds i to the "version" field.
func (ftuo *FormulaTemplateUpdateOne) AddVersion(i int) *FormulaTemplateUpdateOne {
	ftuo.mutation.AddVersion(i)
	return ftuo
}

// SetName sets the "name" field.
func (ftuo *FormulaTemplateUpdateOne) SetName(s string) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableName(s *string) *FormulaTemplateUpdateOne {
	if s != nil {
		ftuo.SetName(*s)
	}
	return ftuo
}

// SetFormulaText sets the "formula_text" field.
func (ftuo *FormulaTemplateUpdateOne) SetFormulaText(s string) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetFormulaText(s)
	return ftuo
}

// SetNillableFormulaText sets the "formula_text" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableFormulaText(s *string) *FormulaTemplateUpdateOne {
	if s != nil {
		ftuo.SetFormulaText(*s)
	}
	return ftuo
}

// SetDescription sets the "description" field.
func (ftuo *FormulaTemplateUpdateOne) SetDescription(s string) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetDescription(s)
	return ftuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableDescription(s *string) *FormulaTemplateUpdateOne {
	if s != nil {
		ftuo.SetDescription(*s)
	}
	return ftuo
}

// ClearDescription clears the value of the "description" field.
func (ftuo *FormulaTemplateUpdateOne) ClearDescription() *FormulaTemplateUpdateOne {
	ftuo.mutation.ClearDescription()
	return ftuo
}

// SetTemplateType sets the "template_type" field.
func (ftuo *FormulaTemplateUpdateOne) SetTemplateType(ft formulatemplate.TemplateType) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetTemplateType(ft)
	return ftuo
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableTemplateType(ft *formulatemplate.TemplateType) *FormulaTemplateUpdateOne {
	if ft != nil {
		ftuo.SetTemplateType(*ft)
	}
	return ftuo
}

// SetCustomerID sets the "customer_id" field.
func (ftuo *FormulaTemplateUpdateOne) SetCustomerID(u uuid.UUID) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetCustomerID(u)
	return ftuo
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableCustomerID(u *uuid.UUID) *FormulaTemplateUpdateOne {
	if u != nil {
		ftuo.SetCustomerID(*u)
	}
	return ftuo
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ftuo *FormulaTemplateUpdateOne) ClearCustomerID() *FormulaTemplateUpdateOne {
	ftuo.mutation.ClearCustomerID()
	return ftuo
}

// SetShipmentTypeID sets the "shipment_type_id" field.
func (ftuo *FormulaTemplateUpdateOne) SetShipmentTypeID(u uuid.UUID) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetShipmentTypeID(u)
	return ftuo
}

// SetNillableShipmentTypeID sets the "shipment_type_id" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableShipmentTypeID(u *uuid.UUID) *FormulaTemplateUpdateOne {
	if u != nil {
		ftuo.SetShipmentTypeID(*u)
	}
	return ftuo
}

// ClearShipmentTypeID clears the value of the "shipment_type_id" field.
func (ftuo *FormulaTemplateUpdateOne) ClearShipmentTypeID() *FormulaTemplateUpdateOne {
	ftuo.mutation.ClearShipmentTypeID()
	return ftuo
}

// SetAutoApply sets the "auto_apply" field.
func (ftuo *FormulaTemplateUpdateOne) SetAutoApply(b bool) *FormulaTemplateUpdateOne {
	ftuo.mutation.SetAutoApply(b)
	return ftuo
}

// SetNillableAutoApply sets the "auto_apply" field if the given value is not nil.
func (ftuo *FormulaTemplateUpdateOne) SetNillableAutoApply(b *bool) *FormulaTemplateUpdateOne {
	if b != nil {
		ftuo.SetAutoApply(*b)
	}
	return ftuo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ftuo *FormulaTemplateUpdateOne) SetCustomer(c *Customer) *FormulaTemplateUpdateOne {
	return ftuo.SetCustomerID(c.ID)
}

// SetShipmentType sets the "shipment_type" edge to the ShipmentType entity.
func (ftuo *FormulaTemplateUpdateOne) SetShipmentType(s *ShipmentType) *FormulaTemplateUpdateOne {
	return ftuo.SetShipmentTypeID(s.ID)
}

// Mutation returns the FormulaTemplateMutation object of the builder.
func (ftuo *FormulaTemplateUpdateOne) Mutation() *FormulaTemplateMutation {
	return ftuo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (ftuo *FormulaTemplateUpdateOne) ClearCustomer() *FormulaTemplateUpdateOne {
	ftuo.mutation.ClearCustomer()
	return ftuo
}

// ClearShipmentType clears the "shipment_type" edge to the ShipmentType entity.
func (ftuo *FormulaTemplateUpdateOne) ClearShipmentType() *FormulaTemplateUpdateOne {
	ftuo.mutation.ClearShipmentType()
	return ftuo
}

// Where appends a list predicates to the FormulaTemplateUpdate builder.
func (ftuo *FormulaTemplateUpdateOne) Where(ps ...predicate.FormulaTemplate) *FormulaTemplateUpdateOne {
	ftuo.mutation.Where(ps...)
	return ftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FormulaTemplateUpdateOne) Select(field string, fields ...string) *FormulaTemplateUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FormulaTemplate entity.
func (ftuo *FormulaTemplateUpdateOne) Save(ctx context.Context) (*FormulaTemplate, error) {
	ftuo.defaults()
	return withHooks(ctx, ftuo.sqlSave, ftuo.mutation, ftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FormulaTemplateUpdateOne) SaveX(ctx context.Context) *FormulaTemplate {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FormulaTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FormulaTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FormulaTemplateUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdatedAt(); !ok {
		v := formulatemplate.UpdateDefaultUpdatedAt()
		ftuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FormulaTemplateUpdateOne) check() error {
	if v, ok := ftuo.mutation.Name(); ok {
		if err := formulatemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.name": %w`, err)}
		}
	}
	if v, ok := ftuo.mutation.FormulaText(); ok {
		if err := formulatemplate.FormulaTextValidator(v); err != nil {
			return &ValidationError{Name: "formula_text", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.formula_text": %w`, err)}
		}
	}
	if v, ok := ftuo.mutation.TemplateType(); ok {
		if err := formulatemplate.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.template_type": %w`, err)}
		}
	}
	if _, ok := ftuo.mutation.BusinessUnitID(); ftuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FormulaTemplate.business_unit"`)
	}
	if _, ok := ftuo.mutation.OrganizationID(); ftuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FormulaTemplate.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftuo *FormulaTemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FormulaTemplateUpdateOne {
	ftuo.modifiers = append(ftuo.modifiers, modifiers...)
	return ftuo
}

func (ftuo *FormulaTemplateUpdateOne) sqlSave(ctx context.Context) (_node *FormulaTemplate, err error) {
	if err := ftuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(formulatemplate.Table, formulatemplate.Columns, sqlgraph.NewFieldSpec(formulatemplate.FieldID, field.TypeUUID))
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FormulaTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, formulatemplate.FieldID)
		for _, f := range fields {
			if !formulatemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != formulatemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(formulatemplate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftuo.mutation.Version(); ok {
		_spec.SetField(formulatemplate.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ftuo.mutation.AddedVersion(); ok {
		_spec.AddField(formulatemplate.FieldVersion, field.TypeInt, value)
	}
	if value, ok := ftuo.mutation.Name(); ok {
		_spec.SetField(formulatemplate.FieldName, field.TypeString, value)
	}
	if value, ok := ftuo.mutation.FormulaText(); ok {
		_spec.SetField(formulatemplate.FieldFormulaText, field.TypeString, value)
	}
	if value, ok := ftuo.mutation.Description(); ok {
		_spec.SetField(formulatemplate.FieldDescription, field.TypeString, value)
	}
	if ftuo.mutation.DescriptionCleared() {
		_spec.ClearField(formulatemplate.FieldDescription, field.TypeString)
	}
	if value, ok := ftuo.mutation.TemplateType(); ok {
		_spec.SetField(formulatemplate.FieldTemplateType, field.TypeEnum, value)
	}
	if value, ok := ftuo.mutation.AutoApply(); ok {
		_spec.SetField(formulatemplate.FieldAutoApply, field.TypeBool, value)
	}
	if ftuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.CustomerTable,
			Columns: []string{formulatemplate.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.CustomerTable,
			Columns: []string{formulatemplate.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftuo.mutation.ShipmentTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.ShipmentTypeTable,
			Columns: []string{formulatemplate.ShipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.ShipmentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.ShipmentTypeTable,
			Columns: []string{formulatemplate.ShipmentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ftuo.modifiers...)
	_node = &FormulaTemplate{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{formulatemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftuo.mutation.done = true
	return _node, nil
}
