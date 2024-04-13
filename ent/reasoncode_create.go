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
	"github.com/emoss08/trenova/ent/reasoncode"
	"github.com/google/uuid"
)

// ReasonCodeCreate is the builder for creating a ReasonCode entity.
type ReasonCodeCreate struct {
	config
	mutation *ReasonCodeMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (rcc *ReasonCodeCreate) SetBusinessUnitID(u uuid.UUID) *ReasonCodeCreate {
	rcc.mutation.SetBusinessUnitID(u)
	return rcc
}

// SetOrganizationID sets the "organization_id" field.
func (rcc *ReasonCodeCreate) SetOrganizationID(u uuid.UUID) *ReasonCodeCreate {
	rcc.mutation.SetOrganizationID(u)
	return rcc
}

// SetCreatedAt sets the "created_at" field.
func (rcc *ReasonCodeCreate) SetCreatedAt(t time.Time) *ReasonCodeCreate {
	rcc.mutation.SetCreatedAt(t)
	return rcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableCreatedAt(t *time.Time) *ReasonCodeCreate {
	if t != nil {
		rcc.SetCreatedAt(*t)
	}
	return rcc
}

// SetUpdatedAt sets the "updated_at" field.
func (rcc *ReasonCodeCreate) SetUpdatedAt(t time.Time) *ReasonCodeCreate {
	rcc.mutation.SetUpdatedAt(t)
	return rcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableUpdatedAt(t *time.Time) *ReasonCodeCreate {
	if t != nil {
		rcc.SetUpdatedAt(*t)
	}
	return rcc
}

// SetVersion sets the "version" field.
func (rcc *ReasonCodeCreate) SetVersion(i int) *ReasonCodeCreate {
	rcc.mutation.SetVersion(i)
	return rcc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableVersion(i *int) *ReasonCodeCreate {
	if i != nil {
		rcc.SetVersion(*i)
	}
	return rcc
}

// SetStatus sets the "status" field.
func (rcc *ReasonCodeCreate) SetStatus(r reasoncode.Status) *ReasonCodeCreate {
	rcc.mutation.SetStatus(r)
	return rcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableStatus(r *reasoncode.Status) *ReasonCodeCreate {
	if r != nil {
		rcc.SetStatus(*r)
	}
	return rcc
}

// SetCode sets the "code" field.
func (rcc *ReasonCodeCreate) SetCode(s string) *ReasonCodeCreate {
	rcc.mutation.SetCode(s)
	return rcc
}

// SetCodeType sets the "code_type" field.
func (rcc *ReasonCodeCreate) SetCodeType(rt reasoncode.CodeType) *ReasonCodeCreate {
	rcc.mutation.SetCodeType(rt)
	return rcc
}

// SetDescription sets the "description" field.
func (rcc *ReasonCodeCreate) SetDescription(s string) *ReasonCodeCreate {
	rcc.mutation.SetDescription(s)
	return rcc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableDescription(s *string) *ReasonCodeCreate {
	if s != nil {
		rcc.SetDescription(*s)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *ReasonCodeCreate) SetID(u uuid.UUID) *ReasonCodeCreate {
	rcc.mutation.SetID(u)
	return rcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rcc *ReasonCodeCreate) SetNillableID(u *uuid.UUID) *ReasonCodeCreate {
	if u != nil {
		rcc.SetID(*u)
	}
	return rcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (rcc *ReasonCodeCreate) SetBusinessUnit(b *BusinessUnit) *ReasonCodeCreate {
	return rcc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rcc *ReasonCodeCreate) SetOrganization(o *Organization) *ReasonCodeCreate {
	return rcc.SetOrganizationID(o.ID)
}

// Mutation returns the ReasonCodeMutation object of the builder.
func (rcc *ReasonCodeCreate) Mutation() *ReasonCodeMutation {
	return rcc.mutation
}

// Save creates the ReasonCode in the database.
func (rcc *ReasonCodeCreate) Save(ctx context.Context) (*ReasonCode, error) {
	rcc.defaults()
	return withHooks(ctx, rcc.sqlSave, rcc.mutation, rcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *ReasonCodeCreate) SaveX(ctx context.Context) *ReasonCode {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *ReasonCodeCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *ReasonCodeCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcc *ReasonCodeCreate) defaults() {
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		v := reasoncode.DefaultCreatedAt()
		rcc.mutation.SetCreatedAt(v)
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		v := reasoncode.DefaultUpdatedAt()
		rcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rcc.mutation.Version(); !ok {
		v := reasoncode.DefaultVersion
		rcc.mutation.SetVersion(v)
	}
	if _, ok := rcc.mutation.Status(); !ok {
		v := reasoncode.DefaultStatus
		rcc.mutation.SetStatus(v)
	}
	if _, ok := rcc.mutation.ID(); !ok {
		v := reasoncode.DefaultID()
		rcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcc *ReasonCodeCreate) check() error {
	if _, ok := rcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "ReasonCode.business_unit_id"`)}
	}
	if _, ok := rcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "ReasonCode.organization_id"`)}
	}
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ReasonCode.created_at"`)}
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ReasonCode.updated_at"`)}
	}
	if _, ok := rcc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ReasonCode.version"`)}
	}
	if _, ok := rcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ReasonCode.status"`)}
	}
	if v, ok := rcc.mutation.Status(); ok {
		if err := reasoncode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.status": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "ReasonCode.code"`)}
	}
	if v, ok := rcc.mutation.Code(); ok {
		if err := reasoncode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.CodeType(); !ok {
		return &ValidationError{Name: "code_type", err: errors.New(`ent: missing required field "ReasonCode.code_type"`)}
	}
	if v, ok := rcc.mutation.CodeType(); ok {
		if err := reasoncode.CodeTypeValidator(v); err != nil {
			return &ValidationError{Name: "code_type", err: fmt.Errorf(`ent: validator failed for field "ReasonCode.code_type": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "ReasonCode.business_unit"`)}
	}
	if _, ok := rcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "ReasonCode.organization"`)}
	}
	return nil
}

func (rcc *ReasonCodeCreate) sqlSave(ctx context.Context) (*ReasonCode, error) {
	if err := rcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
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
	rcc.mutation.id = &_node.ID
	rcc.mutation.done = true
	return _node, nil
}

func (rcc *ReasonCodeCreate) createSpec() (*ReasonCode, *sqlgraph.CreateSpec) {
	var (
		_node = &ReasonCode{config: rcc.config}
		_spec = sqlgraph.NewCreateSpec(reasoncode.Table, sqlgraph.NewFieldSpec(reasoncode.FieldID, field.TypeUUID))
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rcc.mutation.CreatedAt(); ok {
		_spec.SetField(reasoncode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rcc.mutation.UpdatedAt(); ok {
		_spec.SetField(reasoncode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rcc.mutation.Version(); ok {
		_spec.SetField(reasoncode.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := rcc.mutation.Status(); ok {
		_spec.SetField(reasoncode.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rcc.mutation.Code(); ok {
		_spec.SetField(reasoncode.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rcc.mutation.CodeType(); ok {
		_spec.SetField(reasoncode.FieldCodeType, field.TypeEnum, value)
		_node.CodeType = value
	}
	if value, ok := rcc.mutation.Description(); ok {
		_spec.SetField(reasoncode.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := rcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reasoncode.BusinessUnitTable,
			Columns: []string{reasoncode.BusinessUnitColumn},
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
	if nodes := rcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reasoncode.OrganizationTable,
			Columns: []string{reasoncode.OrganizationColumn},
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

// ReasonCodeCreateBulk is the builder for creating many ReasonCode entities in bulk.
type ReasonCodeCreateBulk struct {
	config
	err      error
	builders []*ReasonCodeCreate
}

// Save creates the ReasonCode entities in the database.
func (rccb *ReasonCodeCreateBulk) Save(ctx context.Context) ([]*ReasonCode, error) {
	if rccb.err != nil {
		return nil, rccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*ReasonCode, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReasonCodeMutation)
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
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *ReasonCodeCreateBulk) SaveX(ctx context.Context) []*ReasonCode {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *ReasonCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *ReasonCodeCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}
