// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/customer"
	"github.com/emoss08/trenova/ent/formulatemplate"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/shipmenttype"
	"github.com/google/uuid"
)

// FormulaTemplateCreate is the builder for creating a FormulaTemplate entity.
type FormulaTemplateCreate struct {
	config
	mutation *FormulaTemplateMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (ftc *FormulaTemplateCreate) SetBusinessUnitID(u uuid.UUID) *FormulaTemplateCreate {
	ftc.mutation.SetBusinessUnitID(u)
	return ftc
}

// SetOrganizationID sets the "organization_id" field.
func (ftc *FormulaTemplateCreate) SetOrganizationID(u uuid.UUID) *FormulaTemplateCreate {
	ftc.mutation.SetOrganizationID(u)
	return ftc
}

// SetCreatedAt sets the "created_at" field.
func (ftc *FormulaTemplateCreate) SetCreatedAt(t time.Time) *FormulaTemplateCreate {
	ftc.mutation.SetCreatedAt(t)
	return ftc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableCreatedAt(t *time.Time) *FormulaTemplateCreate {
	if t != nil {
		ftc.SetCreatedAt(*t)
	}
	return ftc
}

// SetUpdatedAt sets the "updated_at" field.
func (ftc *FormulaTemplateCreate) SetUpdatedAt(t time.Time) *FormulaTemplateCreate {
	ftc.mutation.SetUpdatedAt(t)
	return ftc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableUpdatedAt(t *time.Time) *FormulaTemplateCreate {
	if t != nil {
		ftc.SetUpdatedAt(*t)
	}
	return ftc
}

// SetVersion sets the "version" field.
func (ftc *FormulaTemplateCreate) SetVersion(i int) *FormulaTemplateCreate {
	ftc.mutation.SetVersion(i)
	return ftc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableVersion(i *int) *FormulaTemplateCreate {
	if i != nil {
		ftc.SetVersion(*i)
	}
	return ftc
}

// SetName sets the "name" field.
func (ftc *FormulaTemplateCreate) SetName(s string) *FormulaTemplateCreate {
	ftc.mutation.SetName(s)
	return ftc
}

// SetFormulaText sets the "formula_text" field.
func (ftc *FormulaTemplateCreate) SetFormulaText(s string) *FormulaTemplateCreate {
	ftc.mutation.SetFormulaText(s)
	return ftc
}

// SetDescription sets the "description" field.
func (ftc *FormulaTemplateCreate) SetDescription(s string) *FormulaTemplateCreate {
	ftc.mutation.SetDescription(s)
	return ftc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableDescription(s *string) *FormulaTemplateCreate {
	if s != nil {
		ftc.SetDescription(*s)
	}
	return ftc
}

// SetTemplateType sets the "template_type" field.
func (ftc *FormulaTemplateCreate) SetTemplateType(ft formulatemplate.TemplateType) *FormulaTemplateCreate {
	ftc.mutation.SetTemplateType(ft)
	return ftc
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableTemplateType(ft *formulatemplate.TemplateType) *FormulaTemplateCreate {
	if ft != nil {
		ftc.SetTemplateType(*ft)
	}
	return ftc
}

// SetCustomerID sets the "customer_id" field.
func (ftc *FormulaTemplateCreate) SetCustomerID(u uuid.UUID) *FormulaTemplateCreate {
	ftc.mutation.SetCustomerID(u)
	return ftc
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableCustomerID(u *uuid.UUID) *FormulaTemplateCreate {
	if u != nil {
		ftc.SetCustomerID(*u)
	}
	return ftc
}

// SetShipmentTypeID sets the "shipment_type_id" field.
func (ftc *FormulaTemplateCreate) SetShipmentTypeID(u uuid.UUID) *FormulaTemplateCreate {
	ftc.mutation.SetShipmentTypeID(u)
	return ftc
}

// SetNillableShipmentTypeID sets the "shipment_type_id" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableShipmentTypeID(u *uuid.UUID) *FormulaTemplateCreate {
	if u != nil {
		ftc.SetShipmentTypeID(*u)
	}
	return ftc
}

// SetAutoApply sets the "auto_apply" field.
func (ftc *FormulaTemplateCreate) SetAutoApply(b bool) *FormulaTemplateCreate {
	ftc.mutation.SetAutoApply(b)
	return ftc
}

// SetNillableAutoApply sets the "auto_apply" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableAutoApply(b *bool) *FormulaTemplateCreate {
	if b != nil {
		ftc.SetAutoApply(*b)
	}
	return ftc
}

// SetID sets the "id" field.
func (ftc *FormulaTemplateCreate) SetID(u uuid.UUID) *FormulaTemplateCreate {
	ftc.mutation.SetID(u)
	return ftc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ftc *FormulaTemplateCreate) SetNillableID(u *uuid.UUID) *FormulaTemplateCreate {
	if u != nil {
		ftc.SetID(*u)
	}
	return ftc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (ftc *FormulaTemplateCreate) SetBusinessUnit(b *BusinessUnit) *FormulaTemplateCreate {
	return ftc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ftc *FormulaTemplateCreate) SetOrganization(o *Organization) *FormulaTemplateCreate {
	return ftc.SetOrganizationID(o.ID)
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ftc *FormulaTemplateCreate) SetCustomer(c *Customer) *FormulaTemplateCreate {
	return ftc.SetCustomerID(c.ID)
}

// SetShipmentType sets the "shipment_type" edge to the ShipmentType entity.
func (ftc *FormulaTemplateCreate) SetShipmentType(s *ShipmentType) *FormulaTemplateCreate {
	return ftc.SetShipmentTypeID(s.ID)
}

// Mutation returns the FormulaTemplateMutation object of the builder.
func (ftc *FormulaTemplateCreate) Mutation() *FormulaTemplateMutation {
	return ftc.mutation
}

// Save creates the FormulaTemplate in the database.
func (ftc *FormulaTemplateCreate) Save(ctx context.Context) (*FormulaTemplate, error) {
	ftc.defaults()
	return withHooks(ctx, ftc.sqlSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FormulaTemplateCreate) SaveX(ctx context.Context) *FormulaTemplate {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FormulaTemplateCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FormulaTemplateCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FormulaTemplateCreate) defaults() {
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		v := formulatemplate.DefaultCreatedAt()
		ftc.mutation.SetCreatedAt(v)
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		v := formulatemplate.DefaultUpdatedAt()
		ftc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ftc.mutation.Version(); !ok {
		v := formulatemplate.DefaultVersion
		ftc.mutation.SetVersion(v)
	}
	if _, ok := ftc.mutation.TemplateType(); !ok {
		v := formulatemplate.DefaultTemplateType
		ftc.mutation.SetTemplateType(v)
	}
	if _, ok := ftc.mutation.AutoApply(); !ok {
		v := formulatemplate.DefaultAutoApply
		ftc.mutation.SetAutoApply(v)
	}
	if _, ok := ftc.mutation.ID(); !ok {
		v := formulatemplate.DefaultID()
		ftc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FormulaTemplateCreate) check() error {
	if _, ok := ftc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "FormulaTemplate.business_unit_id"`)}
	}
	if _, ok := ftc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "FormulaTemplate.organization_id"`)}
	}
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FormulaTemplate.created_at"`)}
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FormulaTemplate.updated_at"`)}
	}
	if _, ok := ftc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "FormulaTemplate.version"`)}
	}
	if _, ok := ftc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FormulaTemplate.name"`)}
	}
	if v, ok := ftc.mutation.Name(); ok {
		if err := formulatemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.name": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.FormulaText(); !ok {
		return &ValidationError{Name: "formula_text", err: errors.New(`ent: missing required field "FormulaTemplate.formula_text"`)}
	}
	if v, ok := ftc.mutation.FormulaText(); ok {
		if err := formulatemplate.FormulaTextValidator(v); err != nil {
			return &ValidationError{Name: "formula_text", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.formula_text": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.TemplateType(); !ok {
		return &ValidationError{Name: "template_type", err: errors.New(`ent: missing required field "FormulaTemplate.template_type"`)}
	}
	if v, ok := ftc.mutation.TemplateType(); ok {
		if err := formulatemplate.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`ent: validator failed for field "FormulaTemplate.template_type": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.AutoApply(); !ok {
		return &ValidationError{Name: "auto_apply", err: errors.New(`ent: missing required field "FormulaTemplate.auto_apply"`)}
	}
	if _, ok := ftc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "FormulaTemplate.business_unit"`)}
	}
	if _, ok := ftc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "FormulaTemplate.organization"`)}
	}
	return nil
}

func (ftc *FormulaTemplateCreate) sqlSave(ctx context.Context) (*FormulaTemplate, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ftc.mutation.id = &_node.ID
	ftc.mutation.done = true
	return _node, nil
}

func (ftc *FormulaTemplateCreate) createSpec() (*FormulaTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &FormulaTemplate{config: ftc.config}
		_spec = sqlgraph.NewCreateSpec(formulatemplate.Table, sqlgraph.NewFieldSpec(formulatemplate.FieldID, field.TypeUUID))
	)
	if id, ok := ftc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ftc.mutation.CreatedAt(); ok {
		_spec.SetField(formulatemplate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ftc.mutation.UpdatedAt(); ok {
		_spec.SetField(formulatemplate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ftc.mutation.Version(); ok {
		_spec.SetField(formulatemplate.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := ftc.mutation.Name(); ok {
		_spec.SetField(formulatemplate.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ftc.mutation.FormulaText(); ok {
		_spec.SetField(formulatemplate.FieldFormulaText, field.TypeString, value)
		_node.FormulaText = value
	}
	if value, ok := ftc.mutation.Description(); ok {
		_spec.SetField(formulatemplate.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ftc.mutation.TemplateType(); ok {
		_spec.SetField(formulatemplate.FieldTemplateType, field.TypeEnum, value)
		_node.TemplateType = value
	}
	if value, ok := ftc.mutation.AutoApply(); ok {
		_spec.SetField(formulatemplate.FieldAutoApply, field.TypeBool, value)
		_node.AutoApply = value
	}
	if nodes := ftc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.BusinessUnitTable,
			Columns: []string{formulatemplate.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BusinessUnitID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ftc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   formulatemplate.OrganizationTable,
			Columns: []string{formulatemplate.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ftc.mutation.CustomerIDs(); len(nodes) > 0 {
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
		_node.CustomerID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ftc.mutation.ShipmentTypeIDs(); len(nodes) > 0 {
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
		_node.ShipmentTypeID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FormulaTemplateCreateBulk is the builder for creating many FormulaTemplate entities in bulk.
type FormulaTemplateCreateBulk struct {
	config
	err      error
	builders []*FormulaTemplateCreate
}

// Save creates the FormulaTemplate entities in the database.
func (ftcb *FormulaTemplateCreateBulk) Save(ctx context.Context) ([]*FormulaTemplate, error) {
	if ftcb.err != nil {
		return nil, ftcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FormulaTemplate, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FormulaTemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FormulaTemplateCreateBulk) SaveX(ctx context.Context) []*FormulaTemplate {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcb *FormulaTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := ftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcb *FormulaTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := ftcb.Exec(ctx); err != nil {
		panic(err)
	}
}
