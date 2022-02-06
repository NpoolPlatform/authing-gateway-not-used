// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/approleauth"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
)

// AppRoleAuthDelete is the builder for deleting a AppRoleAuth entity.
type AppRoleAuthDelete struct {
	config
	hooks    []Hook
	mutation *AppRoleAuthMutation
}

// Where appends a list predicates to the AppRoleAuthDelete builder.
func (arad *AppRoleAuthDelete) Where(ps ...predicate.AppRoleAuth) *AppRoleAuthDelete {
	arad.mutation.Where(ps...)
	return arad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (arad *AppRoleAuthDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(arad.hooks) == 0 {
		affected, err = arad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppRoleAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			arad.mutation = mutation
			affected, err = arad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(arad.hooks) - 1; i >= 0; i-- {
			if arad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = arad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (arad *AppRoleAuthDelete) ExecX(ctx context.Context) int {
	n, err := arad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (arad *AppRoleAuthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: approleauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approleauth.FieldID,
			},
		},
	}
	if ps := arad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, arad.driver, _spec)
}

// AppRoleAuthDeleteOne is the builder for deleting a single AppRoleAuth entity.
type AppRoleAuthDeleteOne struct {
	arad *AppRoleAuthDelete
}

// Exec executes the deletion query.
func (arado *AppRoleAuthDeleteOne) Exec(ctx context.Context) error {
	n, err := arado.arad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{approleauth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (arado *AppRoleAuthDeleteOne) ExecX(ctx context.Context) {
	arado.arad.ExecX(ctx)
}
