// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/approleauth"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppRoleAuthQuery is the builder for querying AppRoleAuth entities.
type AppRoleAuthQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppRoleAuth
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppRoleAuthQuery builder.
func (araq *AppRoleAuthQuery) Where(ps ...predicate.AppRoleAuth) *AppRoleAuthQuery {
	araq.predicates = append(araq.predicates, ps...)
	return araq
}

// Limit adds a limit step to the query.
func (araq *AppRoleAuthQuery) Limit(limit int) *AppRoleAuthQuery {
	araq.limit = &limit
	return araq
}

// Offset adds an offset step to the query.
func (araq *AppRoleAuthQuery) Offset(offset int) *AppRoleAuthQuery {
	araq.offset = &offset
	return araq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (araq *AppRoleAuthQuery) Unique(unique bool) *AppRoleAuthQuery {
	araq.unique = &unique
	return araq
}

// Order adds an order step to the query.
func (araq *AppRoleAuthQuery) Order(o ...OrderFunc) *AppRoleAuthQuery {
	araq.order = append(araq.order, o...)
	return araq
}

// First returns the first AppRoleAuth entity from the query.
// Returns a *NotFoundError when no AppRoleAuth was found.
func (araq *AppRoleAuthQuery) First(ctx context.Context) (*AppRoleAuth, error) {
	nodes, err := araq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{approleauth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (araq *AppRoleAuthQuery) FirstX(ctx context.Context) *AppRoleAuth {
	node, err := araq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppRoleAuth ID from the query.
// Returns a *NotFoundError when no AppRoleAuth ID was found.
func (araq *AppRoleAuthQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = araq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{approleauth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (araq *AppRoleAuthQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := araq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppRoleAuth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppRoleAuth entity is not found.
// Returns a *NotFoundError when no AppRoleAuth entities are found.
func (araq *AppRoleAuthQuery) Only(ctx context.Context) (*AppRoleAuth, error) {
	nodes, err := araq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{approleauth.Label}
	default:
		return nil, &NotSingularError{approleauth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (araq *AppRoleAuthQuery) OnlyX(ctx context.Context) *AppRoleAuth {
	node, err := araq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppRoleAuth ID in the query.
// Returns a *NotSingularError when exactly one AppRoleAuth ID is not found.
// Returns a *NotFoundError when no entities are found.
func (araq *AppRoleAuthQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = araq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = &NotSingularError{approleauth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (araq *AppRoleAuthQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := araq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppRoleAuths.
func (araq *AppRoleAuthQuery) All(ctx context.Context) ([]*AppRoleAuth, error) {
	if err := araq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return araq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (araq *AppRoleAuthQuery) AllX(ctx context.Context) []*AppRoleAuth {
	nodes, err := araq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppRoleAuth IDs.
func (araq *AppRoleAuthQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := araq.Select(approleauth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (araq *AppRoleAuthQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := araq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (araq *AppRoleAuthQuery) Count(ctx context.Context) (int, error) {
	if err := araq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return araq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (araq *AppRoleAuthQuery) CountX(ctx context.Context) int {
	count, err := araq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (araq *AppRoleAuthQuery) Exist(ctx context.Context) (bool, error) {
	if err := araq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return araq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (araq *AppRoleAuthQuery) ExistX(ctx context.Context) bool {
	exist, err := araq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppRoleAuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (araq *AppRoleAuthQuery) Clone() *AppRoleAuthQuery {
	if araq == nil {
		return nil
	}
	return &AppRoleAuthQuery{
		config:     araq.config,
		limit:      araq.limit,
		offset:     araq.offset,
		order:      append([]OrderFunc{}, araq.order...),
		predicates: append([]predicate.AppRoleAuth{}, araq.predicates...),
		// clone intermediate query.
		sql:  araq.sql.Clone(),
		path: araq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppRoleAuth.Query().
//		GroupBy(approleauth.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (araq *AppRoleAuthQuery) GroupBy(field string, fields ...string) *AppRoleAuthGroupBy {
	group := &AppRoleAuthGroupBy{config: araq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := araq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return araq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppRoleAuth.Query().
//		Select(approleauth.FieldAppID).
//		Scan(ctx, &v)
//
func (araq *AppRoleAuthQuery) Select(fields ...string) *AppRoleAuthSelect {
	araq.fields = append(araq.fields, fields...)
	return &AppRoleAuthSelect{AppRoleAuthQuery: araq}
}

func (araq *AppRoleAuthQuery) prepareQuery(ctx context.Context) error {
	for _, f := range araq.fields {
		if !approleauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if araq.path != nil {
		prev, err := araq.path(ctx)
		if err != nil {
			return err
		}
		araq.sql = prev
	}
	return nil
}

func (araq *AppRoleAuthQuery) sqlAll(ctx context.Context) ([]*AppRoleAuth, error) {
	var (
		nodes = []*AppRoleAuth{}
		_spec = araq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppRoleAuth{config: araq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, araq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (araq *AppRoleAuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := araq.querySpec()
	_spec.Node.Columns = araq.fields
	if len(araq.fields) > 0 {
		_spec.Unique = araq.unique != nil && *araq.unique
	}
	return sqlgraph.CountNodes(ctx, araq.driver, _spec)
}

func (araq *AppRoleAuthQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := araq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (araq *AppRoleAuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   approleauth.Table,
			Columns: approleauth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approleauth.FieldID,
			},
		},
		From:   araq.sql,
		Unique: true,
	}
	if unique := araq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := araq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approleauth.FieldID)
		for i := range fields {
			if fields[i] != approleauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := araq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := araq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := araq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := araq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (araq *AppRoleAuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(araq.driver.Dialect())
	t1 := builder.Table(approleauth.Table)
	columns := araq.fields
	if len(columns) == 0 {
		columns = approleauth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if araq.sql != nil {
		selector = araq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if araq.unique != nil && *araq.unique {
		selector.Distinct()
	}
	for _, p := range araq.predicates {
		p(selector)
	}
	for _, p := range araq.order {
		p(selector)
	}
	if offset := araq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := araq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppRoleAuthGroupBy is the group-by builder for AppRoleAuth entities.
type AppRoleAuthGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aragb *AppRoleAuthGroupBy) Aggregate(fns ...AggregateFunc) *AppRoleAuthGroupBy {
	aragb.fns = append(aragb.fns, fns...)
	return aragb
}

// Scan applies the group-by query and scans the result into the given value.
func (aragb *AppRoleAuthGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aragb.path(ctx)
	if err != nil {
		return err
	}
	aragb.sql = query
	return aragb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aragb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aragb.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) StringsX(ctx context.Context) []string {
	v, err := aragb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aragb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) StringX(ctx context.Context) string {
	v, err := aragb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aragb.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) IntsX(ctx context.Context) []int {
	v, err := aragb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aragb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) IntX(ctx context.Context) int {
	v, err := aragb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aragb.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aragb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aragb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aragb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aragb.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aragb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aragb *AppRoleAuthGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aragb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aragb *AppRoleAuthGroupBy) BoolX(ctx context.Context) bool {
	v, err := aragb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aragb *AppRoleAuthGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aragb.fields {
		if !approleauth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aragb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aragb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aragb *AppRoleAuthGroupBy) sqlQuery() *sql.Selector {
	selector := aragb.sql.Select()
	aggregation := make([]string, 0, len(aragb.fns))
	for _, fn := range aragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aragb.fields)+len(aragb.fns))
		for _, f := range aragb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aragb.fields...)...)
}

// AppRoleAuthSelect is the builder for selecting fields of AppRoleAuth entities.
type AppRoleAuthSelect struct {
	*AppRoleAuthQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aras *AppRoleAuthSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aras.prepareQuery(ctx); err != nil {
		return err
	}
	aras.sql = aras.AppRoleAuthQuery.sqlQuery(ctx)
	return aras.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aras *AppRoleAuthSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aras.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aras.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aras *AppRoleAuthSelect) StringsX(ctx context.Context) []string {
	v, err := aras.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aras.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aras *AppRoleAuthSelect) StringX(ctx context.Context) string {
	v, err := aras.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aras.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aras *AppRoleAuthSelect) IntsX(ctx context.Context) []int {
	v, err := aras.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aras.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aras *AppRoleAuthSelect) IntX(ctx context.Context) int {
	v, err := aras.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aras.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aras *AppRoleAuthSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aras.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aras.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aras *AppRoleAuthSelect) Float64X(ctx context.Context) float64 {
	v, err := aras.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aras.fields) > 1 {
		return nil, errors.New("ent: AppRoleAuthSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aras *AppRoleAuthSelect) BoolsX(ctx context.Context) []bool {
	v, err := aras.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aras *AppRoleAuthSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aras.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleauth.Label}
	default:
		err = fmt.Errorf("ent: AppRoleAuthSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aras *AppRoleAuthSelect) BoolX(ctx context.Context) bool {
	v, err := aras.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aras *AppRoleAuthSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aras.sql.Query()
	if err := aras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
