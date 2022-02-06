// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appuserauth"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserAuthUpdate is the builder for updating AppUserAuth entities.
type AppUserAuthUpdate struct {
	config
	hooks    []Hook
	mutation *AppUserAuthMutation
}

// Where appends a list predicates to the AppUserAuthUpdate builder.
func (auau *AppUserAuthUpdate) Where(ps ...predicate.AppUserAuth) *AppUserAuthUpdate {
	auau.mutation.Where(ps...)
	return auau
}

// SetAppID sets the "app_id" field.
func (auau *AppUserAuthUpdate) SetAppID(u uuid.UUID) *AppUserAuthUpdate {
	auau.mutation.SetAppID(u)
	return auau
}

// SetUserID sets the "user_id" field.
func (auau *AppUserAuthUpdate) SetUserID(u uuid.UUID) *AppUserAuthUpdate {
	auau.mutation.SetUserID(u)
	return auau
}

// SetResource sets the "resource" field.
func (auau *AppUserAuthUpdate) SetResource(s string) *AppUserAuthUpdate {
	auau.mutation.SetResource(s)
	return auau
}

// SetMethod sets the "method" field.
func (auau *AppUserAuthUpdate) SetMethod(s string) *AppUserAuthUpdate {
	auau.mutation.SetMethod(s)
	return auau
}

// SetCreateAt sets the "create_at" field.
func (auau *AppUserAuthUpdate) SetCreateAt(u uint32) *AppUserAuthUpdate {
	auau.mutation.ResetCreateAt()
	auau.mutation.SetCreateAt(u)
	return auau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auau *AppUserAuthUpdate) SetNillableCreateAt(u *uint32) *AppUserAuthUpdate {
	if u != nil {
		auau.SetCreateAt(*u)
	}
	return auau
}

// AddCreateAt adds u to the "create_at" field.
func (auau *AppUserAuthUpdate) AddCreateAt(u int32) *AppUserAuthUpdate {
	auau.mutation.AddCreateAt(u)
	return auau
}

// SetUpdateAt sets the "update_at" field.
func (auau *AppUserAuthUpdate) SetUpdateAt(u uint32) *AppUserAuthUpdate {
	auau.mutation.ResetUpdateAt()
	auau.mutation.SetUpdateAt(u)
	return auau
}

// AddUpdateAt adds u to the "update_at" field.
func (auau *AppUserAuthUpdate) AddUpdateAt(u int32) *AppUserAuthUpdate {
	auau.mutation.AddUpdateAt(u)
	return auau
}

// SetDeleteAt sets the "delete_at" field.
func (auau *AppUserAuthUpdate) SetDeleteAt(u uint32) *AppUserAuthUpdate {
	auau.mutation.ResetDeleteAt()
	auau.mutation.SetDeleteAt(u)
	return auau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auau *AppUserAuthUpdate) SetNillableDeleteAt(u *uint32) *AppUserAuthUpdate {
	if u != nil {
		auau.SetDeleteAt(*u)
	}
	return auau
}

// AddDeleteAt adds u to the "delete_at" field.
func (auau *AppUserAuthUpdate) AddDeleteAt(u int32) *AppUserAuthUpdate {
	auau.mutation.AddDeleteAt(u)
	return auau
}

// Mutation returns the AppUserAuthMutation object of the builder.
func (auau *AppUserAuthUpdate) Mutation() *AppUserAuthMutation {
	return auau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auau *AppUserAuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	auau.defaults()
	if len(auau.hooks) == 0 {
		affected, err = auau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auau.mutation = mutation
			affected, err = auau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(auau.hooks) - 1; i >= 0; i-- {
			if auau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (auau *AppUserAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := auau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auau *AppUserAuthUpdate) Exec(ctx context.Context) error {
	_, err := auau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auau *AppUserAuthUpdate) ExecX(ctx context.Context) {
	if err := auau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auau *AppUserAuthUpdate) defaults() {
	if _, ok := auau.mutation.UpdateAt(); !ok {
		v := appuserauth.UpdateDefaultUpdateAt()
		auau.mutation.SetUpdateAt(v)
	}
}

func (auau *AppUserAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserauth.Table,
			Columns: appuserauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserauth.FieldID,
			},
		},
	}
	if ps := auau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldAppID,
		})
	}
	if value, ok := auau.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldUserID,
		})
	}
	if value, ok := auau.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldResource,
		})
	}
	if value, ok := auau.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldMethod,
		})
	}
	if value, ok := auau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldCreateAt,
		})
	}
	if value, ok := auau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldCreateAt,
		})
	}
	if value, ok := auau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldUpdateAt,
		})
	}
	if value, ok := auau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldUpdateAt,
		})
	}
	if value, ok := auau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldDeleteAt,
		})
	}
	if value, ok := auau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserAuthUpdateOne is the builder for updating a single AppUserAuth entity.
type AppUserAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppUserAuthMutation
}

// SetAppID sets the "app_id" field.
func (auauo *AppUserAuthUpdateOne) SetAppID(u uuid.UUID) *AppUserAuthUpdateOne {
	auauo.mutation.SetAppID(u)
	return auauo
}

// SetUserID sets the "user_id" field.
func (auauo *AppUserAuthUpdateOne) SetUserID(u uuid.UUID) *AppUserAuthUpdateOne {
	auauo.mutation.SetUserID(u)
	return auauo
}

// SetResource sets the "resource" field.
func (auauo *AppUserAuthUpdateOne) SetResource(s string) *AppUserAuthUpdateOne {
	auauo.mutation.SetResource(s)
	return auauo
}

// SetMethod sets the "method" field.
func (auauo *AppUserAuthUpdateOne) SetMethod(s string) *AppUserAuthUpdateOne {
	auauo.mutation.SetMethod(s)
	return auauo
}

// SetCreateAt sets the "create_at" field.
func (auauo *AppUserAuthUpdateOne) SetCreateAt(u uint32) *AppUserAuthUpdateOne {
	auauo.mutation.ResetCreateAt()
	auauo.mutation.SetCreateAt(u)
	return auauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auauo *AppUserAuthUpdateOne) SetNillableCreateAt(u *uint32) *AppUserAuthUpdateOne {
	if u != nil {
		auauo.SetCreateAt(*u)
	}
	return auauo
}

// AddCreateAt adds u to the "create_at" field.
func (auauo *AppUserAuthUpdateOne) AddCreateAt(u int32) *AppUserAuthUpdateOne {
	auauo.mutation.AddCreateAt(u)
	return auauo
}

// SetUpdateAt sets the "update_at" field.
func (auauo *AppUserAuthUpdateOne) SetUpdateAt(u uint32) *AppUserAuthUpdateOne {
	auauo.mutation.ResetUpdateAt()
	auauo.mutation.SetUpdateAt(u)
	return auauo
}

// AddUpdateAt adds u to the "update_at" field.
func (auauo *AppUserAuthUpdateOne) AddUpdateAt(u int32) *AppUserAuthUpdateOne {
	auauo.mutation.AddUpdateAt(u)
	return auauo
}

// SetDeleteAt sets the "delete_at" field.
func (auauo *AppUserAuthUpdateOne) SetDeleteAt(u uint32) *AppUserAuthUpdateOne {
	auauo.mutation.ResetDeleteAt()
	auauo.mutation.SetDeleteAt(u)
	return auauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auauo *AppUserAuthUpdateOne) SetNillableDeleteAt(u *uint32) *AppUserAuthUpdateOne {
	if u != nil {
		auauo.SetDeleteAt(*u)
	}
	return auauo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auauo *AppUserAuthUpdateOne) AddDeleteAt(u int32) *AppUserAuthUpdateOne {
	auauo.mutation.AddDeleteAt(u)
	return auauo
}

// Mutation returns the AppUserAuthMutation object of the builder.
func (auauo *AppUserAuthUpdateOne) Mutation() *AppUserAuthMutation {
	return auauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auauo *AppUserAuthUpdateOne) Select(field string, fields ...string) *AppUserAuthUpdateOne {
	auauo.fields = append([]string{field}, fields...)
	return auauo
}

// Save executes the query and returns the updated AppUserAuth entity.
func (auauo *AppUserAuthUpdateOne) Save(ctx context.Context) (*AppUserAuth, error) {
	var (
		err  error
		node *AppUserAuth
	)
	auauo.defaults()
	if len(auauo.hooks) == 0 {
		node, err = auauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auauo.mutation = mutation
			node, err = auauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auauo.hooks) - 1; i >= 0; i-- {
			if auauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auauo *AppUserAuthUpdateOne) SaveX(ctx context.Context) *AppUserAuth {
	node, err := auauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auauo *AppUserAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := auauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auauo *AppUserAuthUpdateOne) ExecX(ctx context.Context) {
	if err := auauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auauo *AppUserAuthUpdateOne) defaults() {
	if _, ok := auauo.mutation.UpdateAt(); !ok {
		v := appuserauth.UpdateDefaultUpdateAt()
		auauo.mutation.SetUpdateAt(v)
	}
}

func (auauo *AppUserAuthUpdateOne) sqlSave(ctx context.Context) (_node *AppUserAuth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserauth.Table,
			Columns: appuserauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserauth.FieldID,
			},
		},
	}
	id, ok := auauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserauth.FieldID)
		for _, f := range fields {
			if !appuserauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appuserauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldAppID,
		})
	}
	if value, ok := auauo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldUserID,
		})
	}
	if value, ok := auauo.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldResource,
		})
	}
	if value, ok := auauo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldMethod,
		})
	}
	if value, ok := auauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldCreateAt,
		})
	}
	if value, ok := auauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldCreateAt,
		})
	}
	if value, ok := auauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldUpdateAt,
		})
	}
	if value, ok := auauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldUpdateAt,
		})
	}
	if value, ok := auauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldDeleteAt,
		})
	}
	if value, ok := auauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldDeleteAt,
		})
	}
	_node = &AppUserAuth{config: auauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
