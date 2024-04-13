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
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/shipmenttype"
	"github.com/google/uuid"
)

// ShipmentTypeCreate is the builder for creating a ShipmentType entity.
type ShipmentTypeCreate struct {
	config
	mutation *ShipmentTypeMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (stc *ShipmentTypeCreate) SetBusinessUnitID(u uuid.UUID) *ShipmentTypeCreate {
	stc.mutation.SetBusinessUnitID(u)
	return stc
}

// SetOrganizationID sets the "organization_id" field.
func (stc *ShipmentTypeCreate) SetOrganizationID(u uuid.UUID) *ShipmentTypeCreate {
	stc.mutation.SetOrganizationID(u)
	return stc
}

// SetCreatedAt sets the "created_at" field.
func (stc *ShipmentTypeCreate) SetCreatedAt(t time.Time) *ShipmentTypeCreate {
	stc.mutation.SetCreatedAt(t)
	return stc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableCreatedAt(t *time.Time) *ShipmentTypeCreate {
	if t != nil {
		stc.SetCreatedAt(*t)
	}
	return stc
}

// SetUpdatedAt sets the "updated_at" field.
func (stc *ShipmentTypeCreate) SetUpdatedAt(t time.Time) *ShipmentTypeCreate {
	stc.mutation.SetUpdatedAt(t)
	return stc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableUpdatedAt(t *time.Time) *ShipmentTypeCreate {
	if t != nil {
		stc.SetUpdatedAt(*t)
	}
	return stc
}

// SetVersion sets the "version" field.
func (stc *ShipmentTypeCreate) SetVersion(i int) *ShipmentTypeCreate {
	stc.mutation.SetVersion(i)
	return stc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableVersion(i *int) *ShipmentTypeCreate {
	if i != nil {
		stc.SetVersion(*i)
	}
	return stc
}

// SetStatus sets the "status" field.
func (stc *ShipmentTypeCreate) SetStatus(s shipmenttype.Status) *ShipmentTypeCreate {
	stc.mutation.SetStatus(s)
	return stc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableStatus(s *shipmenttype.Status) *ShipmentTypeCreate {
	if s != nil {
		stc.SetStatus(*s)
	}
	return stc
}

// SetCode sets the "code" field.
func (stc *ShipmentTypeCreate) SetCode(s string) *ShipmentTypeCreate {
	stc.mutation.SetCode(s)
	return stc
}

// SetDescription sets the "description" field.
func (stc *ShipmentTypeCreate) SetDescription(s string) *ShipmentTypeCreate {
	stc.mutation.SetDescription(s)
	return stc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableDescription(s *string) *ShipmentTypeCreate {
	if s != nil {
		stc.SetDescription(*s)
	}
	return stc
}

// SetID sets the "id" field.
func (stc *ShipmentTypeCreate) SetID(u uuid.UUID) *ShipmentTypeCreate {
	stc.mutation.SetID(u)
	return stc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (stc *ShipmentTypeCreate) SetNillableID(u *uuid.UUID) *ShipmentTypeCreate {
	if u != nil {
		stc.SetID(*u)
	}
	return stc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (stc *ShipmentTypeCreate) SetBusinessUnit(b *BusinessUnit) *ShipmentTypeCreate {
	return stc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (stc *ShipmentTypeCreate) SetOrganization(o *Organization) *ShipmentTypeCreate {
	return stc.SetOrganizationID(o.ID)
}

// Mutation returns the ShipmentTypeMutation object of the builder.
func (stc *ShipmentTypeCreate) Mutation() *ShipmentTypeMutation {
	return stc.mutation
}

// Save creates the ShipmentType in the database.
func (stc *ShipmentTypeCreate) Save(ctx context.Context) (*ShipmentType, error) {
	stc.defaults()
	return withHooks(ctx, stc.sqlSave, stc.mutation, stc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (stc *ShipmentTypeCreate) SaveX(ctx context.Context) *ShipmentType {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *ShipmentTypeCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *ShipmentTypeCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stc *ShipmentTypeCreate) defaults() {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		v := shipmenttype.DefaultCreatedAt()
		stc.mutation.SetCreatedAt(v)
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		v := shipmenttype.DefaultUpdatedAt()
		stc.mutation.SetUpdatedAt(v)
	}
	if _, ok := stc.mutation.Version(); !ok {
		v := shipmenttype.DefaultVersion
		stc.mutation.SetVersion(v)
	}
	if _, ok := stc.mutation.Status(); !ok {
		v := shipmenttype.DefaultStatus
		stc.mutation.SetStatus(v)
	}
	if _, ok := stc.mutation.ID(); !ok {
		v := shipmenttype.DefaultID()
		stc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *ShipmentTypeCreate) check() error {
	if _, ok := stc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "ShipmentType.business_unit_id"`)}
	}
	if _, ok := stc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "ShipmentType.organization_id"`)}
	}
	if _, ok := stc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ShipmentType.created_at"`)}
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ShipmentType.updated_at"`)}
	}
	if _, ok := stc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ShipmentType.version"`)}
	}
	if _, ok := stc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ShipmentType.status"`)}
	}
	if v, ok := stc.mutation.Status(); ok {
		if err := shipmenttype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ShipmentType.status": %w`, err)}
		}
	}
	if _, ok := stc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "ShipmentType.code"`)}
	}
	if v, ok := stc.mutation.Code(); ok {
		if err := shipmenttype.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "ShipmentType.code": %w`, err)}
		}
	}
	if _, ok := stc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "ShipmentType.business_unit"`)}
	}
	if _, ok := stc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "ShipmentType.organization"`)}
	}
	return nil
}

func (stc *ShipmentTypeCreate) sqlSave(ctx context.Context) (*ShipmentType, error) {
	if err := stc.check(); err != nil {
		return nil, err
	}
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
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
	stc.mutation.id = &_node.ID
	stc.mutation.done = true
	return _node, nil
}

func (stc *ShipmentTypeCreate) createSpec() (*ShipmentType, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentType{config: stc.config}
		_spec = sqlgraph.NewCreateSpec(shipmenttype.Table, sqlgraph.NewFieldSpec(shipmenttype.FieldID, field.TypeUUID))
	)
	if id, ok := stc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := stc.mutation.CreatedAt(); ok {
		_spec.SetField(shipmenttype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := stc.mutation.UpdatedAt(); ok {
		_spec.SetField(shipmenttype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := stc.mutation.Version(); ok {
		_spec.SetField(shipmenttype.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := stc.mutation.Status(); ok {
		_spec.SetField(shipmenttype.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := stc.mutation.Code(); ok {
		_spec.SetField(shipmenttype.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := stc.mutation.Description(); ok {
		_spec.SetField(shipmenttype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := stc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenttype.BusinessUnitTable,
			Columns: []string{shipmenttype.BusinessUnitColumn},
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
	if nodes := stc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmenttype.OrganizationTable,
			Columns: []string{shipmenttype.OrganizationColumn},
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
	return _node, _spec
}

// ShipmentTypeCreateBulk is the builder for creating many ShipmentType entities in bulk.
type ShipmentTypeCreateBulk struct {
	config
	err      error
	builders []*ShipmentTypeCreate
}

// Save creates the ShipmentType entities in the database.
func (stcb *ShipmentTypeCreateBulk) Save(ctx context.Context) ([]*ShipmentType, error) {
	if stcb.err != nil {
		return nil, stcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*ShipmentType, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *ShipmentTypeCreateBulk) SaveX(ctx context.Context) []*ShipmentType {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *ShipmentTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *ShipmentTypeCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}
