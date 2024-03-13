// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/accountingcontrol"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrganizationCreate) SetCreatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrganizationCreate) SetUpdatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetScacCode sets the "scac_code" field.
func (oc *OrganizationCreate) SetScacCode(s string) *OrganizationCreate {
	oc.mutation.SetScacCode(s)
	return oc
}

// SetDotNumber sets the "dot_number" field.
func (oc *OrganizationCreate) SetDotNumber(s string) *OrganizationCreate {
	oc.mutation.SetDotNumber(s)
	return oc
}

// SetLogoURL sets the "logo_url" field.
func (oc *OrganizationCreate) SetLogoURL(s string) *OrganizationCreate {
	oc.mutation.SetLogoURL(s)
	return oc
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableLogoURL(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetLogoURL(*s)
	}
	return oc
}

// SetOrgType sets the "org_type" field.
func (oc *OrganizationCreate) SetOrgType(ot organization.OrgType) *OrganizationCreate {
	oc.mutation.SetOrgType(ot)
	return oc
}

// SetNillableOrgType sets the "org_type" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableOrgType(ot *organization.OrgType) *OrganizationCreate {
	if ot != nil {
		oc.SetOrgType(*ot)
	}
	return oc
}

// SetTimezone sets the "timezone" field.
func (oc *OrganizationCreate) SetTimezone(o organization.Timezone) *OrganizationCreate {
	oc.mutation.SetTimezone(o)
	return oc
}

// SetNillableTimezone sets the "timezone" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableTimezone(o *organization.Timezone) *OrganizationCreate {
	if o != nil {
		oc.SetTimezone(*o)
	}
	return oc
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (oc *OrganizationCreate) SetBusinessUnitID(u uuid.UUID) *OrganizationCreate {
	oc.mutation.SetBusinessUnitID(u)
	return oc
}

// SetID sets the "id" field.
func (oc *OrganizationCreate) SetID(u uuid.UUID) *OrganizationCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableID(u *uuid.UUID) *OrganizationCreate {
	if u != nil {
		oc.SetID(*u)
	}
	return oc
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (oc *OrganizationCreate) SetBusinessUnitID(id uuid.UUID) *OrganizationCreate {
	oc.mutation.SetBusinessUnitID(id)
	return oc
}

// SetNillableBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableBusinessUnitID(id *uuid.UUID) *OrganizationCreate {
	if id != nil {
		oc = oc.SetBusinessUnitID(*id)
	}
	return oc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (oc *OrganizationCreate) SetBusinessUnit(b *BusinessUnit) *OrganizationCreate {
	return oc.SetBusinessUnitID(b.ID)
}

// SetAccountingControlID sets the "accounting_control" edge to the AccountingControl entity by ID.
func (oc *OrganizationCreate) SetAccountingControlID(id uuid.UUID) *OrganizationCreate {
	oc.mutation.SetAccountingControlID(id)
	return oc
}

// SetNillableAccountingControlID sets the "accounting_control" edge to the AccountingControl entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableAccountingControlID(id *uuid.UUID) *OrganizationCreate {
	if id != nil {
		oc = oc.SetAccountingControlID(*id)
	}
	return oc
}

// SetAccountingControl sets the "accounting_control" edge to the AccountingControl entity.
func (oc *OrganizationCreate) SetAccountingControl(a *AccountingControl) *OrganizationCreate {
	return oc.SetAccountingControlID(a.ID)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := organization.DefaultCreatedAt
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := organization.DefaultUpdatedAt
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.OrgType(); !ok {
		v := organization.DefaultOrgType
		oc.mutation.SetOrgType(v)
	}
	if _, ok := oc.mutation.Timezone(); !ok {
		v := organization.DefaultTimezone
		oc.mutation.SetTimezone(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := organization.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Organization.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Organization.updated_at"`)}
	}
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Organization.name"`)}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ScacCode(); !ok {
		return &ValidationError{Name: "scac_code", err: errors.New(`ent: missing required field "Organization.scac_code"`)}
	}
	if v, ok := oc.mutation.ScacCode(); ok {
		if err := organization.ScacCodeValidator(v); err != nil {
			return &ValidationError{Name: "scac_code", err: fmt.Errorf(`ent: validator failed for field "Organization.scac_code": %w`, err)}
		}
	}
	if _, ok := oc.mutation.DotNumber(); !ok {
		return &ValidationError{Name: "dot_number", err: errors.New(`ent: missing required field "Organization.dot_number"`)}
	}
	if v, ok := oc.mutation.DotNumber(); ok {
		if err := organization.DotNumberValidator(v); err != nil {
			return &ValidationError{Name: "dot_number", err: fmt.Errorf(`ent: validator failed for field "Organization.dot_number": %w`, err)}
		}
	}
	if _, ok := oc.mutation.OrgType(); !ok {
		return &ValidationError{Name: "org_type", err: errors.New(`ent: missing required field "Organization.org_type"`)}
	}
	if v, ok := oc.mutation.OrgType(); ok {
		if err := organization.OrgTypeValidator(v); err != nil {
			return &ValidationError{Name: "org_type", err: fmt.Errorf(`ent: validator failed for field "Organization.org_type": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Timezone(); !ok {
		return &ValidationError{Name: "timezone", err: errors.New(`ent: missing required field "Organization.timezone"`)}
	}
	if v, ok := oc.mutation.Timezone(); ok {
		if err := organization.TimezoneValidator(v); err != nil {
			return &ValidationError{Name: "timezone", err: fmt.Errorf(`ent: validator failed for field "Organization.timezone": %w`, err)}
		}
	}
	if _, ok := oc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "Organization.business_unit_id"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
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
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(organization.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oc.mutation.ScacCode(); ok {
		_spec.SetField(organization.FieldScacCode, field.TypeString, value)
		_node.ScacCode = value
	}
	if value, ok := oc.mutation.DotNumber(); ok {
		_spec.SetField(organization.FieldDotNumber, field.TypeString, value)
		_node.DotNumber = value
	}
	if value, ok := oc.mutation.LogoURL(); ok {
		_spec.SetField(organization.FieldLogoURL, field.TypeString, value)
		_node.LogoURL = value
	}
	if value, ok := oc.mutation.OrgType(); ok {
		_spec.SetField(organization.FieldOrgType, field.TypeEnum, value)
		_node.OrgType = value
	}
	if value, ok := oc.mutation.Timezone(); ok {
		_spec.SetField(organization.FieldTimezone, field.TypeEnum, value)
		_node.Timezone = value
	}
	if value, ok := oc.mutation.BusinessUnitID(); ok {
		_spec.SetField(organization.FieldBusinessUnitID, field.TypeUUID, value)
		_node.BusinessUnitID = value
	}
	if nodes := oc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.BusinessUnitTable,
			Columns: []string{organization.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_unit_organizations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.AccountingControlIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.AccountingControlTable,
			Columns: []string{organization.AccountingControlColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountingcontrol.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_accounting_control = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	err      error
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
