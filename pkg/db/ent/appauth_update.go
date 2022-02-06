// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appauth"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppAuthUpdate is the builder for updating AppAuth entities.
type AppAuthUpdate struct {
	config
	hooks    []Hook
	mutation *AppAuthMutation
}

// Where appends a list predicates to the AppAuthUpdate builder.
func (aau *AppAuthUpdate) Where(ps ...predicate.AppAuth) *AppAuthUpdate {
	aau.mutation.Where(ps...)
	return aau
}

// SetAppID sets the "app_id" field.
func (aau *AppAuthUpdate) SetAppID(u uuid.UUID) *AppAuthUpdate {
	aau.mutation.SetAppID(u)
	return aau
}

// SetResource sets the "resource" field.
func (aau *AppAuthUpdate) SetResource(s string) *AppAuthUpdate {
	aau.mutation.SetResource(s)
	return aau
}

// SetMethod sets the "method" field.
func (aau *AppAuthUpdate) SetMethod(s string) *AppAuthUpdate {
	aau.mutation.SetMethod(s)
	return aau
}

// SetCreateAt sets the "create_at" field.
func (aau *AppAuthUpdate) SetCreateAt(u uint32) *AppAuthUpdate {
	aau.mutation.ResetCreateAt()
	aau.mutation.SetCreateAt(u)
	return aau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aau *AppAuthUpdate) SetNillableCreateAt(u *uint32) *AppAuthUpdate {
	if u != nil {
		aau.SetCreateAt(*u)
	}
	return aau
}

// AddCreateAt adds u to the "create_at" field.
func (aau *AppAuthUpdate) AddCreateAt(u int32) *AppAuthUpdate {
	aau.mutation.AddCreateAt(u)
	return aau
}

// SetUpdateAt sets the "update_at" field.
func (aau *AppAuthUpdate) SetUpdateAt(u uint32) *AppAuthUpdate {
	aau.mutation.ResetUpdateAt()
	aau.mutation.SetUpdateAt(u)
	return aau
}

// AddUpdateAt adds u to the "update_at" field.
func (aau *AppAuthUpdate) AddUpdateAt(u int32) *AppAuthUpdate {
	aau.mutation.AddUpdateAt(u)
	return aau
}

// SetDeleteAt sets the "delete_at" field.
func (aau *AppAuthUpdate) SetDeleteAt(u uint32) *AppAuthUpdate {
	aau.mutation.ResetDeleteAt()
	aau.mutation.SetDeleteAt(u)
	return aau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aau *AppAuthUpdate) SetNillableDeleteAt(u *uint32) *AppAuthUpdate {
	if u != nil {
		aau.SetDeleteAt(*u)
	}
	return aau
}

// AddDeleteAt adds u to the "delete_at" field.
func (aau *AppAuthUpdate) AddDeleteAt(u int32) *AppAuthUpdate {
	aau.mutation.AddDeleteAt(u)
	return aau
}

// Mutation returns the AppAuthMutation object of the builder.
func (aau *AppAuthUpdate) Mutation() *AppAuthMutation {
	return aau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aau *AppAuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aau.defaults()
	if len(aau.hooks) == 0 {
		affected, err = aau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aau.mutation = mutation
			affected, err = aau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aau.hooks) - 1; i >= 0; i-- {
			if aau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aau *AppAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := aau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aau *AppAuthUpdate) Exec(ctx context.Context) error {
	_, err := aau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aau *AppAuthUpdate) ExecX(ctx context.Context) {
	if err := aau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aau *AppAuthUpdate) defaults() {
	if _, ok := aau.mutation.UpdateAt(); !ok {
		v := appauth.UpdateDefaultUpdateAt()
		aau.mutation.SetUpdateAt(v)
	}
}

func (aau *AppAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appauth.Table,
			Columns: appauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appauth.FieldID,
			},
		},
	}
	if ps := aau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appauth.FieldAppID,
		})
	}
	if value, ok := aau.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appauth.FieldResource,
		})
	}
	if value, ok := aau.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appauth.FieldMethod,
		})
	}
	if value, ok := aau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldCreateAt,
		})
	}
	if value, ok := aau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldCreateAt,
		})
	}
	if value, ok := aau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldUpdateAt,
		})
	}
	if value, ok := aau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldUpdateAt,
		})
	}
	if value, ok := aau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldDeleteAt,
		})
	}
	if value, ok := aau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppAuthUpdateOne is the builder for updating a single AppAuth entity.
type AppAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppAuthMutation
}

// SetAppID sets the "app_id" field.
func (aauo *AppAuthUpdateOne) SetAppID(u uuid.UUID) *AppAuthUpdateOne {
	aauo.mutation.SetAppID(u)
	return aauo
}

// SetResource sets the "resource" field.
func (aauo *AppAuthUpdateOne) SetResource(s string) *AppAuthUpdateOne {
	aauo.mutation.SetResource(s)
	return aauo
}

// SetMethod sets the "method" field.
func (aauo *AppAuthUpdateOne) SetMethod(s string) *AppAuthUpdateOne {
	aauo.mutation.SetMethod(s)
	return aauo
}

// SetCreateAt sets the "create_at" field.
func (aauo *AppAuthUpdateOne) SetCreateAt(u uint32) *AppAuthUpdateOne {
	aauo.mutation.ResetCreateAt()
	aauo.mutation.SetCreateAt(u)
	return aauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aauo *AppAuthUpdateOne) SetNillableCreateAt(u *uint32) *AppAuthUpdateOne {
	if u != nil {
		aauo.SetCreateAt(*u)
	}
	return aauo
}

// AddCreateAt adds u to the "create_at" field.
func (aauo *AppAuthUpdateOne) AddCreateAt(u int32) *AppAuthUpdateOne {
	aauo.mutation.AddCreateAt(u)
	return aauo
}

// SetUpdateAt sets the "update_at" field.
func (aauo *AppAuthUpdateOne) SetUpdateAt(u uint32) *AppAuthUpdateOne {
	aauo.mutation.ResetUpdateAt()
	aauo.mutation.SetUpdateAt(u)
	return aauo
}

// AddUpdateAt adds u to the "update_at" field.
func (aauo *AppAuthUpdateOne) AddUpdateAt(u int32) *AppAuthUpdateOne {
	aauo.mutation.AddUpdateAt(u)
	return aauo
}

// SetDeleteAt sets the "delete_at" field.
func (aauo *AppAuthUpdateOne) SetDeleteAt(u uint32) *AppAuthUpdateOne {
	aauo.mutation.ResetDeleteAt()
	aauo.mutation.SetDeleteAt(u)
	return aauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aauo *AppAuthUpdateOne) SetNillableDeleteAt(u *uint32) *AppAuthUpdateOne {
	if u != nil {
		aauo.SetDeleteAt(*u)
	}
	return aauo
}

// AddDeleteAt adds u to the "delete_at" field.
func (aauo *AppAuthUpdateOne) AddDeleteAt(u int32) *AppAuthUpdateOne {
	aauo.mutation.AddDeleteAt(u)
	return aauo
}

// Mutation returns the AppAuthMutation object of the builder.
func (aauo *AppAuthUpdateOne) Mutation() *AppAuthMutation {
	return aauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aauo *AppAuthUpdateOne) Select(field string, fields ...string) *AppAuthUpdateOne {
	aauo.fields = append([]string{field}, fields...)
	return aauo
}

// Save executes the query and returns the updated AppAuth entity.
func (aauo *AppAuthUpdateOne) Save(ctx context.Context) (*AppAuth, error) {
	var (
		err  error
		node *AppAuth
	)
	aauo.defaults()
	if len(aauo.hooks) == 0 {
		node, err = aauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aauo.mutation = mutation
			node, err = aauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aauo.hooks) - 1; i >= 0; i-- {
			if aauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aauo *AppAuthUpdateOne) SaveX(ctx context.Context) *AppAuth {
	node, err := aauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aauo *AppAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := aauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aauo *AppAuthUpdateOne) ExecX(ctx context.Context) {
	if err := aauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aauo *AppAuthUpdateOne) defaults() {
	if _, ok := aauo.mutation.UpdateAt(); !ok {
		v := appauth.UpdateDefaultUpdateAt()
		aauo.mutation.SetUpdateAt(v)
	}
}

func (aauo *AppAuthUpdateOne) sqlSave(ctx context.Context) (_node *AppAuth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appauth.Table,
			Columns: appauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appauth.FieldID,
			},
		},
	}
	id, ok := aauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appauth.FieldID)
		for _, f := range fields {
			if !appauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appauth.FieldAppID,
		})
	}
	if value, ok := aauo.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appauth.FieldResource,
		})
	}
	if value, ok := aauo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appauth.FieldMethod,
		})
	}
	if value, ok := aauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldCreateAt,
		})
	}
	if value, ok := aauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldCreateAt,
		})
	}
	if value, ok := aauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldUpdateAt,
		})
	}
	if value, ok := aauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldUpdateAt,
		})
	}
	if value, ok := aauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldDeleteAt,
		})
	}
	if value, ok := aauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appauth.FieldDeleteAt,
		})
	}
	_node = &AppAuth{config: aauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
