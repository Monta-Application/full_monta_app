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
	"github.com/emoss08/trenova/ent/usstate"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/emoss08/trenova/ent/workerprofile"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// WorkerProfileCreate is the builder for creating a WorkerProfile entity.
type WorkerProfileCreate struct {
	config
	mutation *WorkerProfileMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (wpc *WorkerProfileCreate) SetBusinessUnitID(u uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetBusinessUnitID(u)
	return wpc
}

// SetOrganizationID sets the "organization_id" field.
func (wpc *WorkerProfileCreate) SetOrganizationID(u uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetOrganizationID(u)
	return wpc
}

// SetCreatedAt sets the "created_at" field.
func (wpc *WorkerProfileCreate) SetCreatedAt(t time.Time) *WorkerProfileCreate {
	wpc.mutation.SetCreatedAt(t)
	return wpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableCreatedAt(t *time.Time) *WorkerProfileCreate {
	if t != nil {
		wpc.SetCreatedAt(*t)
	}
	return wpc
}

// SetUpdatedAt sets the "updated_at" field.
func (wpc *WorkerProfileCreate) SetUpdatedAt(t time.Time) *WorkerProfileCreate {
	wpc.mutation.SetUpdatedAt(t)
	return wpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableUpdatedAt(t *time.Time) *WorkerProfileCreate {
	if t != nil {
		wpc.SetUpdatedAt(*t)
	}
	return wpc
}

// SetVersion sets the "version" field.
func (wpc *WorkerProfileCreate) SetVersion(i int) *WorkerProfileCreate {
	wpc.mutation.SetVersion(i)
	return wpc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableVersion(i *int) *WorkerProfileCreate {
	if i != nil {
		wpc.SetVersion(*i)
	}
	return wpc
}

// SetWorkerID sets the "worker_id" field.
func (wpc *WorkerProfileCreate) SetWorkerID(u uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetWorkerID(u)
	return wpc
}

// SetRace sets the "race" field.
func (wpc *WorkerProfileCreate) SetRace(s string) *WorkerProfileCreate {
	wpc.mutation.SetRace(s)
	return wpc
}

// SetNillableRace sets the "race" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableRace(s *string) *WorkerProfileCreate {
	if s != nil {
		wpc.SetRace(*s)
	}
	return wpc
}

// SetSex sets the "sex" field.
func (wpc *WorkerProfileCreate) SetSex(s string) *WorkerProfileCreate {
	wpc.mutation.SetSex(s)
	return wpc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableSex(s *string) *WorkerProfileCreate {
	if s != nil {
		wpc.SetSex(*s)
	}
	return wpc
}

// SetDateOfBirth sets the "date_of_birth" field.
func (wpc *WorkerProfileCreate) SetDateOfBirth(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetDateOfBirth(pg)
	return wpc
}

// SetLicenseNumber sets the "license_number" field.
func (wpc *WorkerProfileCreate) SetLicenseNumber(s string) *WorkerProfileCreate {
	wpc.mutation.SetLicenseNumber(s)
	return wpc
}

// SetLicenseStateID sets the "license_state_id" field.
func (wpc *WorkerProfileCreate) SetLicenseStateID(u uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetLicenseStateID(u)
	return wpc
}

// SetLicenseExpirationDate sets the "license_expiration_date" field.
func (wpc *WorkerProfileCreate) SetLicenseExpirationDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetLicenseExpirationDate(pg)
	return wpc
}

// SetEndorsements sets the "endorsements" field.
func (wpc *WorkerProfileCreate) SetEndorsements(w workerprofile.Endorsements) *WorkerProfileCreate {
	wpc.mutation.SetEndorsements(w)
	return wpc
}

// SetNillableEndorsements sets the "endorsements" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableEndorsements(w *workerprofile.Endorsements) *WorkerProfileCreate {
	if w != nil {
		wpc.SetEndorsements(*w)
	}
	return wpc
}

// SetHazmatExpirationDate sets the "hazmat_expiration_date" field.
func (wpc *WorkerProfileCreate) SetHazmatExpirationDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetHazmatExpirationDate(pg)
	return wpc
}

// SetHireDate sets the "hire_date" field.
func (wpc *WorkerProfileCreate) SetHireDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetHireDate(pg)
	return wpc
}

// SetTerminationDate sets the "termination_date" field.
func (wpc *WorkerProfileCreate) SetTerminationDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetTerminationDate(pg)
	return wpc
}

// SetPhysicalDueDate sets the "physical_due_date" field.
func (wpc *WorkerProfileCreate) SetPhysicalDueDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetPhysicalDueDate(pg)
	return wpc
}

// SetMedicalCertDate sets the "medical_cert_date" field.
func (wpc *WorkerProfileCreate) SetMedicalCertDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetMedicalCertDate(pg)
	return wpc
}

// SetMvrDueDate sets the "mvr_due_date" field.
func (wpc *WorkerProfileCreate) SetMvrDueDate(pg *pgtype.Date) *WorkerProfileCreate {
	wpc.mutation.SetMvrDueDate(pg)
	return wpc
}

// SetID sets the "id" field.
func (wpc *WorkerProfileCreate) SetID(u uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetID(u)
	return wpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wpc *WorkerProfileCreate) SetNillableID(u *uuid.UUID) *WorkerProfileCreate {
	if u != nil {
		wpc.SetID(*u)
	}
	return wpc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (wpc *WorkerProfileCreate) SetBusinessUnit(b *BusinessUnit) *WorkerProfileCreate {
	return wpc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wpc *WorkerProfileCreate) SetOrganization(o *Organization) *WorkerProfileCreate {
	return wpc.SetOrganizationID(o.ID)
}

// SetWorker sets the "worker" edge to the Worker entity.
func (wpc *WorkerProfileCreate) SetWorker(w *Worker) *WorkerProfileCreate {
	return wpc.SetWorkerID(w.ID)
}

// SetStateID sets the "state" edge to the UsState entity by ID.
func (wpc *WorkerProfileCreate) SetStateID(id uuid.UUID) *WorkerProfileCreate {
	wpc.mutation.SetStateID(id)
	return wpc
}

// SetState sets the "state" edge to the UsState entity.
func (wpc *WorkerProfileCreate) SetState(u *UsState) *WorkerProfileCreate {
	return wpc.SetStateID(u.ID)
}

// Mutation returns the WorkerProfileMutation object of the builder.
func (wpc *WorkerProfileCreate) Mutation() *WorkerProfileMutation {
	return wpc.mutation
}

// Save creates the WorkerProfile in the database.
func (wpc *WorkerProfileCreate) Save(ctx context.Context) (*WorkerProfile, error) {
	if err := wpc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wpc.sqlSave, wpc.mutation, wpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wpc *WorkerProfileCreate) SaveX(ctx context.Context) *WorkerProfile {
	v, err := wpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wpc *WorkerProfileCreate) Exec(ctx context.Context) error {
	_, err := wpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpc *WorkerProfileCreate) ExecX(ctx context.Context) {
	if err := wpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wpc *WorkerProfileCreate) defaults() error {
	if _, ok := wpc.mutation.CreatedAt(); !ok {
		if workerprofile.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized workerprofile.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := workerprofile.DefaultCreatedAt()
		wpc.mutation.SetCreatedAt(v)
	}
	if _, ok := wpc.mutation.UpdatedAt(); !ok {
		if workerprofile.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized workerprofile.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := workerprofile.DefaultUpdatedAt()
		wpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wpc.mutation.Version(); !ok {
		v := workerprofile.DefaultVersion
		wpc.mutation.SetVersion(v)
	}
	if _, ok := wpc.mutation.Endorsements(); !ok {
		v := workerprofile.DefaultEndorsements
		wpc.mutation.SetEndorsements(v)
	}
	if _, ok := wpc.mutation.ID(); !ok {
		if workerprofile.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized workerprofile.DefaultID (forgotten import ent/runtime?)")
		}
		v := workerprofile.DefaultID()
		wpc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wpc *WorkerProfileCreate) check() error {
	if _, ok := wpc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "WorkerProfile.business_unit_id"`)}
	}
	if _, ok := wpc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "WorkerProfile.organization_id"`)}
	}
	if _, ok := wpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkerProfile.created_at"`)}
	}
	if _, ok := wpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WorkerProfile.updated_at"`)}
	}
	if _, ok := wpc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "WorkerProfile.version"`)}
	}
	if _, ok := wpc.mutation.WorkerID(); !ok {
		return &ValidationError{Name: "worker_id", err: errors.New(`ent: missing required field "WorkerProfile.worker_id"`)}
	}
	if _, ok := wpc.mutation.LicenseNumber(); !ok {
		return &ValidationError{Name: "license_number", err: errors.New(`ent: missing required field "WorkerProfile.license_number"`)}
	}
	if v, ok := wpc.mutation.LicenseNumber(); ok {
		if err := workerprofile.LicenseNumberValidator(v); err != nil {
			return &ValidationError{Name: "license_number", err: fmt.Errorf(`ent: validator failed for field "WorkerProfile.license_number": %w`, err)}
		}
	}
	if _, ok := wpc.mutation.LicenseStateID(); !ok {
		return &ValidationError{Name: "license_state_id", err: errors.New(`ent: missing required field "WorkerProfile.license_state_id"`)}
	}
	if v, ok := wpc.mutation.Endorsements(); ok {
		if err := workerprofile.EndorsementsValidator(v); err != nil {
			return &ValidationError{Name: "endorsements", err: fmt.Errorf(`ent: validator failed for field "WorkerProfile.endorsements": %w`, err)}
		}
	}
	if _, ok := wpc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "WorkerProfile.business_unit"`)}
	}
	if _, ok := wpc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "WorkerProfile.organization"`)}
	}
	if _, ok := wpc.mutation.WorkerID(); !ok {
		return &ValidationError{Name: "worker", err: errors.New(`ent: missing required edge "WorkerProfile.worker"`)}
	}
	if _, ok := wpc.mutation.StateID(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required edge "WorkerProfile.state"`)}
	}
	return nil
}

func (wpc *WorkerProfileCreate) sqlSave(ctx context.Context) (*WorkerProfile, error) {
	if err := wpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wpc.driver, _spec); err != nil {
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
	wpc.mutation.id = &_node.ID
	wpc.mutation.done = true
	return _node, nil
}

func (wpc *WorkerProfileCreate) createSpec() (*WorkerProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkerProfile{config: wpc.config}
		_spec = sqlgraph.NewCreateSpec(workerprofile.Table, sqlgraph.NewFieldSpec(workerprofile.FieldID, field.TypeUUID))
	)
	if id, ok := wpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wpc.mutation.CreatedAt(); ok {
		_spec.SetField(workerprofile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wpc.mutation.UpdatedAt(); ok {
		_spec.SetField(workerprofile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wpc.mutation.Version(); ok {
		_spec.SetField(workerprofile.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := wpc.mutation.Race(); ok {
		_spec.SetField(workerprofile.FieldRace, field.TypeString, value)
		_node.Race = value
	}
	if value, ok := wpc.mutation.Sex(); ok {
		_spec.SetField(workerprofile.FieldSex, field.TypeString, value)
		_node.Sex = value
	}
	if value, ok := wpc.mutation.DateOfBirth(); ok {
		_spec.SetField(workerprofile.FieldDateOfBirth, field.TypeOther, value)
		_node.DateOfBirth = value
	}
	if value, ok := wpc.mutation.LicenseNumber(); ok {
		_spec.SetField(workerprofile.FieldLicenseNumber, field.TypeString, value)
		_node.LicenseNumber = value
	}
	if value, ok := wpc.mutation.LicenseExpirationDate(); ok {
		_spec.SetField(workerprofile.FieldLicenseExpirationDate, field.TypeOther, value)
		_node.LicenseExpirationDate = value
	}
	if value, ok := wpc.mutation.Endorsements(); ok {
		_spec.SetField(workerprofile.FieldEndorsements, field.TypeEnum, value)
		_node.Endorsements = value
	}
	if value, ok := wpc.mutation.HazmatExpirationDate(); ok {
		_spec.SetField(workerprofile.FieldHazmatExpirationDate, field.TypeOther, value)
		_node.HazmatExpirationDate = value
	}
	if value, ok := wpc.mutation.HireDate(); ok {
		_spec.SetField(workerprofile.FieldHireDate, field.TypeOther, value)
		_node.HireDate = value
	}
	if value, ok := wpc.mutation.TerminationDate(); ok {
		_spec.SetField(workerprofile.FieldTerminationDate, field.TypeOther, value)
		_node.TerminationDate = value
	}
	if value, ok := wpc.mutation.PhysicalDueDate(); ok {
		_spec.SetField(workerprofile.FieldPhysicalDueDate, field.TypeOther, value)
		_node.PhysicalDueDate = value
	}
	if value, ok := wpc.mutation.MedicalCertDate(); ok {
		_spec.SetField(workerprofile.FieldMedicalCertDate, field.TypeOther, value)
		_node.MedicalCertDate = value
	}
	if value, ok := wpc.mutation.MvrDueDate(); ok {
		_spec.SetField(workerprofile.FieldMvrDueDate, field.TypeOther, value)
		_node.MvrDueDate = value
	}
	if nodes := wpc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workerprofile.BusinessUnitTable,
			Columns: []string{workerprofile.BusinessUnitColumn},
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
	if nodes := wpc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workerprofile.OrganizationTable,
			Columns: []string{workerprofile.OrganizationColumn},
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
	if nodes := wpc.mutation.WorkerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workerprofile.WorkerTable,
			Columns: []string{workerprofile.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wpc.mutation.StateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workerprofile.StateTable,
			Columns: []string{workerprofile.StateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usstate.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LicenseStateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkerProfileCreateBulk is the builder for creating many WorkerProfile entities in bulk.
type WorkerProfileCreateBulk struct {
	config
	err      error
	builders []*WorkerProfileCreate
}

// Save creates the WorkerProfile entities in the database.
func (wpcb *WorkerProfileCreateBulk) Save(ctx context.Context) ([]*WorkerProfile, error) {
	if wpcb.err != nil {
		return nil, wpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wpcb.builders))
	nodes := make([]*WorkerProfile, len(wpcb.builders))
	mutators := make([]Mutator, len(wpcb.builders))
	for i := range wpcb.builders {
		func(i int, root context.Context) {
			builder := wpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkerProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, wpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wpcb *WorkerProfileCreateBulk) SaveX(ctx context.Context) []*WorkerProfile {
	v, err := wpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wpcb *WorkerProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := wpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpcb *WorkerProfileCreateBulk) ExecX(ctx context.Context) {
	if err := wpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
