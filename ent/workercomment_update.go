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
	"github.com/emoss08/trenova/ent/commenttype"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/user"
	"github.com/emoss08/trenova/ent/worker"
	"github.com/emoss08/trenova/ent/workercomment"
	"github.com/google/uuid"
)

// WorkerCommentUpdate is the builder for updating WorkerComment entities.
type WorkerCommentUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkerCommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WorkerCommentUpdate builder.
func (wcu *WorkerCommentUpdate) Where(ps ...predicate.WorkerComment) *WorkerCommentUpdate {
	wcu.mutation.Where(ps...)
	return wcu
}

// SetUpdatedAt sets the "updated_at" field.
func (wcu *WorkerCommentUpdate) SetUpdatedAt(t time.Time) *WorkerCommentUpdate {
	wcu.mutation.SetUpdatedAt(t)
	return wcu
}

// SetVersion sets the "version" field.
func (wcu *WorkerCommentUpdate) SetVersion(i int) *WorkerCommentUpdate {
	wcu.mutation.ResetVersion()
	wcu.mutation.SetVersion(i)
	return wcu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wcu *WorkerCommentUpdate) SetNillableVersion(i *int) *WorkerCommentUpdate {
	if i != nil {
		wcu.SetVersion(*i)
	}
	return wcu
}

// AddVersion adds i to the "version" field.
func (wcu *WorkerCommentUpdate) AddVersion(i int) *WorkerCommentUpdate {
	wcu.mutation.AddVersion(i)
	return wcu
}

// SetWorkerID sets the "worker_id" field.
func (wcu *WorkerCommentUpdate) SetWorkerID(u uuid.UUID) *WorkerCommentUpdate {
	wcu.mutation.SetWorkerID(u)
	return wcu
}

// SetNillableWorkerID sets the "worker_id" field if the given value is not nil.
func (wcu *WorkerCommentUpdate) SetNillableWorkerID(u *uuid.UUID) *WorkerCommentUpdate {
	if u != nil {
		wcu.SetWorkerID(*u)
	}
	return wcu
}

// SetCommentTypeID sets the "comment_type_id" field.
func (wcu *WorkerCommentUpdate) SetCommentTypeID(u uuid.UUID) *WorkerCommentUpdate {
	wcu.mutation.SetCommentTypeID(u)
	return wcu
}

// SetNillableCommentTypeID sets the "comment_type_id" field if the given value is not nil.
func (wcu *WorkerCommentUpdate) SetNillableCommentTypeID(u *uuid.UUID) *WorkerCommentUpdate {
	if u != nil {
		wcu.SetCommentTypeID(*u)
	}
	return wcu
}

// SetUserID sets the "user_id" field.
func (wcu *WorkerCommentUpdate) SetUserID(u uuid.UUID) *WorkerCommentUpdate {
	wcu.mutation.SetUserID(u)
	return wcu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcu *WorkerCommentUpdate) SetNillableUserID(u *uuid.UUID) *WorkerCommentUpdate {
	if u != nil {
		wcu.SetUserID(*u)
	}
	return wcu
}

// SetComment sets the "comment" field.
func (wcu *WorkerCommentUpdate) SetComment(s string) *WorkerCommentUpdate {
	wcu.mutation.SetComment(s)
	return wcu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (wcu *WorkerCommentUpdate) SetNillableComment(s *string) *WorkerCommentUpdate {
	if s != nil {
		wcu.SetComment(*s)
	}
	return wcu
}

// SetWorker sets the "worker" edge to the Worker entity.
func (wcu *WorkerCommentUpdate) SetWorker(w *Worker) *WorkerCommentUpdate {
	return wcu.SetWorkerID(w.ID)
}

// SetCommentType sets the "comment_type" edge to the CommentType entity.
func (wcu *WorkerCommentUpdate) SetCommentType(c *CommentType) *WorkerCommentUpdate {
	return wcu.SetCommentTypeID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (wcu *WorkerCommentUpdate) SetUser(u *User) *WorkerCommentUpdate {
	return wcu.SetUserID(u.ID)
}

// Mutation returns the WorkerCommentMutation object of the builder.
func (wcu *WorkerCommentUpdate) Mutation() *WorkerCommentMutation {
	return wcu.mutation
}

// ClearWorker clears the "worker" edge to the Worker entity.
func (wcu *WorkerCommentUpdate) ClearWorker() *WorkerCommentUpdate {
	wcu.mutation.ClearWorker()
	return wcu
}

// ClearCommentType clears the "comment_type" edge to the CommentType entity.
func (wcu *WorkerCommentUpdate) ClearCommentType() *WorkerCommentUpdate {
	wcu.mutation.ClearCommentType()
	return wcu
}

// ClearUser clears the "user" edge to the User entity.
func (wcu *WorkerCommentUpdate) ClearUser() *WorkerCommentUpdate {
	wcu.mutation.ClearUser()
	return wcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wcu *WorkerCommentUpdate) Save(ctx context.Context) (int, error) {
	wcu.defaults()
	return withHooks(ctx, wcu.sqlSave, wcu.mutation, wcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcu *WorkerCommentUpdate) SaveX(ctx context.Context) int {
	affected, err := wcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wcu *WorkerCommentUpdate) Exec(ctx context.Context) error {
	_, err := wcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcu *WorkerCommentUpdate) ExecX(ctx context.Context) {
	if err := wcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcu *WorkerCommentUpdate) defaults() {
	if _, ok := wcu.mutation.UpdatedAt(); !ok {
		v := workercomment.UpdateDefaultUpdatedAt()
		wcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcu *WorkerCommentUpdate) check() error {
	if v, ok := wcu.mutation.Comment(); ok {
		if err := workercomment.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "WorkerComment.comment": %w`, err)}
		}
	}
	if _, ok := wcu.mutation.BusinessUnitID(); wcu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.business_unit"`)
	}
	if _, ok := wcu.mutation.OrganizationID(); wcu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.organization"`)
	}
	if _, ok := wcu.mutation.WorkerID(); wcu.mutation.WorkerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.worker"`)
	}
	if _, ok := wcu.mutation.CommentTypeID(); wcu.mutation.CommentTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.comment_type"`)
	}
	if _, ok := wcu.mutation.UserID(); wcu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wcu *WorkerCommentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkerCommentUpdate {
	wcu.modifiers = append(wcu.modifiers, modifiers...)
	return wcu
}

func (wcu *WorkerCommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workercomment.Table, workercomment.Columns, sqlgraph.NewFieldSpec(workercomment.FieldID, field.TypeUUID))
	if ps := wcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcu.mutation.UpdatedAt(); ok {
		_spec.SetField(workercomment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wcu.mutation.Version(); ok {
		_spec.SetField(workercomment.FieldVersion, field.TypeInt, value)
	}
	if value, ok := wcu.mutation.AddedVersion(); ok {
		_spec.AddField(workercomment.FieldVersion, field.TypeInt, value)
	}
	if value, ok := wcu.mutation.Comment(); ok {
		_spec.SetField(workercomment.FieldComment, field.TypeString, value)
	}
	if wcu.mutation.WorkerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workercomment.WorkerTable,
			Columns: []string{workercomment.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.WorkerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workercomment.WorkerTable,
			Columns: []string{workercomment.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcu.mutation.CommentTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.CommentTypeTable,
			Columns: []string{workercomment.CommentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commenttype.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.CommentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.CommentTypeTable,
			Columns: []string{workercomment.CommentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commenttype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.UserTable,
			Columns: []string{workercomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.UserTable,
			Columns: []string{workercomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workercomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wcu.mutation.done = true
	return n, nil
}

// WorkerCommentUpdateOne is the builder for updating a single WorkerComment entity.
type WorkerCommentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkerCommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (wcuo *WorkerCommentUpdateOne) SetUpdatedAt(t time.Time) *WorkerCommentUpdateOne {
	wcuo.mutation.SetUpdatedAt(t)
	return wcuo
}

// SetVersion sets the "version" field.
func (wcuo *WorkerCommentUpdateOne) SetVersion(i int) *WorkerCommentUpdateOne {
	wcuo.mutation.ResetVersion()
	wcuo.mutation.SetVersion(i)
	return wcuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wcuo *WorkerCommentUpdateOne) SetNillableVersion(i *int) *WorkerCommentUpdateOne {
	if i != nil {
		wcuo.SetVersion(*i)
	}
	return wcuo
}

// AddVersion adds i to the "version" field.
func (wcuo *WorkerCommentUpdateOne) AddVersion(i int) *WorkerCommentUpdateOne {
	wcuo.mutation.AddVersion(i)
	return wcuo
}

// SetWorkerID sets the "worker_id" field.
func (wcuo *WorkerCommentUpdateOne) SetWorkerID(u uuid.UUID) *WorkerCommentUpdateOne {
	wcuo.mutation.SetWorkerID(u)
	return wcuo
}

// SetNillableWorkerID sets the "worker_id" field if the given value is not nil.
func (wcuo *WorkerCommentUpdateOne) SetNillableWorkerID(u *uuid.UUID) *WorkerCommentUpdateOne {
	if u != nil {
		wcuo.SetWorkerID(*u)
	}
	return wcuo
}

// SetCommentTypeID sets the "comment_type_id" field.
func (wcuo *WorkerCommentUpdateOne) SetCommentTypeID(u uuid.UUID) *WorkerCommentUpdateOne {
	wcuo.mutation.SetCommentTypeID(u)
	return wcuo
}

// SetNillableCommentTypeID sets the "comment_type_id" field if the given value is not nil.
func (wcuo *WorkerCommentUpdateOne) SetNillableCommentTypeID(u *uuid.UUID) *WorkerCommentUpdateOne {
	if u != nil {
		wcuo.SetCommentTypeID(*u)
	}
	return wcuo
}

// SetUserID sets the "user_id" field.
func (wcuo *WorkerCommentUpdateOne) SetUserID(u uuid.UUID) *WorkerCommentUpdateOne {
	wcuo.mutation.SetUserID(u)
	return wcuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcuo *WorkerCommentUpdateOne) SetNillableUserID(u *uuid.UUID) *WorkerCommentUpdateOne {
	if u != nil {
		wcuo.SetUserID(*u)
	}
	return wcuo
}

// SetComment sets the "comment" field.
func (wcuo *WorkerCommentUpdateOne) SetComment(s string) *WorkerCommentUpdateOne {
	wcuo.mutation.SetComment(s)
	return wcuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (wcuo *WorkerCommentUpdateOne) SetNillableComment(s *string) *WorkerCommentUpdateOne {
	if s != nil {
		wcuo.SetComment(*s)
	}
	return wcuo
}

// SetWorker sets the "worker" edge to the Worker entity.
func (wcuo *WorkerCommentUpdateOne) SetWorker(w *Worker) *WorkerCommentUpdateOne {
	return wcuo.SetWorkerID(w.ID)
}

// SetCommentType sets the "comment_type" edge to the CommentType entity.
func (wcuo *WorkerCommentUpdateOne) SetCommentType(c *CommentType) *WorkerCommentUpdateOne {
	return wcuo.SetCommentTypeID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (wcuo *WorkerCommentUpdateOne) SetUser(u *User) *WorkerCommentUpdateOne {
	return wcuo.SetUserID(u.ID)
}

// Mutation returns the WorkerCommentMutation object of the builder.
func (wcuo *WorkerCommentUpdateOne) Mutation() *WorkerCommentMutation {
	return wcuo.mutation
}

// ClearWorker clears the "worker" edge to the Worker entity.
func (wcuo *WorkerCommentUpdateOne) ClearWorker() *WorkerCommentUpdateOne {
	wcuo.mutation.ClearWorker()
	return wcuo
}

// ClearCommentType clears the "comment_type" edge to the CommentType entity.
func (wcuo *WorkerCommentUpdateOne) ClearCommentType() *WorkerCommentUpdateOne {
	wcuo.mutation.ClearCommentType()
	return wcuo
}

// ClearUser clears the "user" edge to the User entity.
func (wcuo *WorkerCommentUpdateOne) ClearUser() *WorkerCommentUpdateOne {
	wcuo.mutation.ClearUser()
	return wcuo
}

// Where appends a list predicates to the WorkerCommentUpdate builder.
func (wcuo *WorkerCommentUpdateOne) Where(ps ...predicate.WorkerComment) *WorkerCommentUpdateOne {
	wcuo.mutation.Where(ps...)
	return wcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wcuo *WorkerCommentUpdateOne) Select(field string, fields ...string) *WorkerCommentUpdateOne {
	wcuo.fields = append([]string{field}, fields...)
	return wcuo
}

// Save executes the query and returns the updated WorkerComment entity.
func (wcuo *WorkerCommentUpdateOne) Save(ctx context.Context) (*WorkerComment, error) {
	wcuo.defaults()
	return withHooks(ctx, wcuo.sqlSave, wcuo.mutation, wcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcuo *WorkerCommentUpdateOne) SaveX(ctx context.Context) *WorkerComment {
	node, err := wcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wcuo *WorkerCommentUpdateOne) Exec(ctx context.Context) error {
	_, err := wcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcuo *WorkerCommentUpdateOne) ExecX(ctx context.Context) {
	if err := wcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcuo *WorkerCommentUpdateOne) defaults() {
	if _, ok := wcuo.mutation.UpdatedAt(); !ok {
		v := workercomment.UpdateDefaultUpdatedAt()
		wcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcuo *WorkerCommentUpdateOne) check() error {
	if v, ok := wcuo.mutation.Comment(); ok {
		if err := workercomment.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "WorkerComment.comment": %w`, err)}
		}
	}
	if _, ok := wcuo.mutation.BusinessUnitID(); wcuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.business_unit"`)
	}
	if _, ok := wcuo.mutation.OrganizationID(); wcuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.organization"`)
	}
	if _, ok := wcuo.mutation.WorkerID(); wcuo.mutation.WorkerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.worker"`)
	}
	if _, ok := wcuo.mutation.CommentTypeID(); wcuo.mutation.CommentTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.comment_type"`)
	}
	if _, ok := wcuo.mutation.UserID(); wcuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkerComment.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wcuo *WorkerCommentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkerCommentUpdateOne {
	wcuo.modifiers = append(wcuo.modifiers, modifiers...)
	return wcuo
}

func (wcuo *WorkerCommentUpdateOne) sqlSave(ctx context.Context) (_node *WorkerComment, err error) {
	if err := wcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workercomment.Table, workercomment.Columns, sqlgraph.NewFieldSpec(workercomment.FieldID, field.TypeUUID))
	id, ok := wcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkerComment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workercomment.FieldID)
		for _, f := range fields {
			if !workercomment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workercomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(workercomment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wcuo.mutation.Version(); ok {
		_spec.SetField(workercomment.FieldVersion, field.TypeInt, value)
	}
	if value, ok := wcuo.mutation.AddedVersion(); ok {
		_spec.AddField(workercomment.FieldVersion, field.TypeInt, value)
	}
	if value, ok := wcuo.mutation.Comment(); ok {
		_spec.SetField(workercomment.FieldComment, field.TypeString, value)
	}
	if wcuo.mutation.WorkerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workercomment.WorkerTable,
			Columns: []string{workercomment.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.WorkerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workercomment.WorkerTable,
			Columns: []string{workercomment.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcuo.mutation.CommentTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.CommentTypeTable,
			Columns: []string{workercomment.CommentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commenttype.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.CommentTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.CommentTypeTable,
			Columns: []string{workercomment.CommentTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commenttype.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.UserTable,
			Columns: []string{workercomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercomment.UserTable,
			Columns: []string{workercomment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wcuo.modifiers...)
	_node = &WorkerComment{config: wcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workercomment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wcuo.mutation.done = true
	return _node, nil
}
