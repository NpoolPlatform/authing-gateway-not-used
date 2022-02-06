// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appauth"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
)

// AppAuthDelete is the builder for deleting a AppAuth entity.
type AppAuthDelete struct {
	config
	hooks    []Hook
	mutation *AppAuthMutation
}

// Where appends a list predicates to the AppAuthDelete builder.
func (aad *AppAuthDelete) Where(ps ...predicate.AppAuth) *AppAuthDelete {
	aad.mutation.Where(ps...)
	return aad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aad *AppAuthDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aad.hooks) == 0 {
		affected, err = aad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aad.mutation = mutation
			affected, err = aad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aad.hooks) - 1; i >= 0; i-- {
			if aad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aad *AppAuthDelete) ExecX(ctx context.Context) int {
	n, err := aad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aad *AppAuthDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appauth.FieldID,
			},
		},
	}
	if ps := aad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aad.driver, _spec)
}

// AppAuthDeleteOne is the builder for deleting a single AppAuth entity.
type AppAuthDeleteOne struct {
	aad *AppAuthDelete
}

// Exec executes the deletion query.
func (aado *AppAuthDeleteOne) Exec(ctx context.Context) error {
	n, err := aado.aad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appauth.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aado *AppAuthDeleteOne) ExecX(ctx context.Context) {
	aado.aad.ExecX(ctx)
}
