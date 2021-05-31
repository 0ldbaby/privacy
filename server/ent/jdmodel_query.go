// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/wq22304/privacy/ent/jdmodel"
	"github.com/wq22304/privacy/ent/predicate"
)

// JDModelQuery is the builder for querying JDModel entities.
type JDModelQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.JDModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (jmq *JDModelQuery) Where(ps ...predicate.JDModel) *JDModelQuery {
	jmq.predicates = append(jmq.predicates, ps...)
	return jmq
}

// Limit adds a limit step to the query.
func (jmq *JDModelQuery) Limit(limit int) *JDModelQuery {
	jmq.limit = &limit
	return jmq
}

// Offset adds an offset step to the query.
func (jmq *JDModelQuery) Offset(offset int) *JDModelQuery {
	jmq.offset = &offset
	return jmq
}

// Order adds an order step to the query.
func (jmq *JDModelQuery) Order(o ...OrderFunc) *JDModelQuery {
	jmq.order = append(jmq.order, o...)
	return jmq
}

// First returns the first JDModel entity in the query. Returns *NotFoundError when no jdmodel was found.
func (jmq *JDModelQuery) First(ctx context.Context) (*JDModel, error) {
	nodes, err := jmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jdmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jmq *JDModelQuery) FirstX(ctx context.Context) *JDModel {
	node, err := jmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JDModel id in the query. Returns *NotFoundError when no id was found.
func (jmq *JDModelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jdmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jmq *JDModelQuery) FirstIDX(ctx context.Context) int {
	id, err := jmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only JDModel entity in the query, returns an error if not exactly one entity was returned.
func (jmq *JDModelQuery) Only(ctx context.Context) (*JDModel, error) {
	nodes, err := jmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jdmodel.Label}
	default:
		return nil, &NotSingularError{jdmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jmq *JDModelQuery) OnlyX(ctx context.Context) *JDModel {
	node, err := jmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only JDModel id in the query, returns an error if not exactly one id was returned.
func (jmq *JDModelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = &NotSingularError{jdmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jmq *JDModelQuery) OnlyIDX(ctx context.Context) int {
	id, err := jmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JDModels.
func (jmq *JDModelQuery) All(ctx context.Context) ([]*JDModel, error) {
	if err := jmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jmq *JDModelQuery) AllX(ctx context.Context) []*JDModel {
	nodes, err := jmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JDModel ids.
func (jmq *JDModelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jmq.Select(jdmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jmq *JDModelQuery) IDsX(ctx context.Context) []int {
	ids, err := jmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jmq *JDModelQuery) Count(ctx context.Context) (int, error) {
	if err := jmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jmq *JDModelQuery) CountX(ctx context.Context) int {
	count, err := jmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jmq *JDModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := jmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jmq *JDModelQuery) ExistX(ctx context.Context) bool {
	exist, err := jmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jmq *JDModelQuery) Clone() *JDModelQuery {
	if jmq == nil {
		return nil
	}
	return &JDModelQuery{
		config:     jmq.config,
		limit:      jmq.limit,
		offset:     jmq.offset,
		order:      append([]OrderFunc{}, jmq.order...),
		unique:     append([]string{}, jmq.unique...),
		predicates: append([]predicate.JDModel{}, jmq.predicates...),
		// clone intermediate query.
		sql:  jmq.sql.Clone(),
		path: jmq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JDModel.Query().
//		GroupBy(jdmodel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (jmq *JDModelQuery) GroupBy(field string, fields ...string) *JDModelGroupBy {
	group := &JDModelGroupBy{config: jmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jmq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.JDModel.Query().
//		Select(jdmodel.FieldName).
//		Scan(ctx, &v)
//
func (jmq *JDModelQuery) Select(field string, fields ...string) *JDModelSelect {
	selector := &JDModelSelect{config: jmq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jmq.sqlQuery(), nil
	}
	return selector
}

func (jmq *JDModelQuery) prepareQuery(ctx context.Context) error {
	if jmq.path != nil {
		prev, err := jmq.path(ctx)
		if err != nil {
			return err
		}
		jmq.sql = prev
	}
	return nil
}

func (jmq *JDModelQuery) sqlAll(ctx context.Context) ([]*JDModel, error) {
	var (
		nodes = []*JDModel{}
		_spec = jmq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &JDModel{config: jmq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, jmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (jmq *JDModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jmq.querySpec()
	return sqlgraph.CountNodes(ctx, jmq.driver, _spec)
}

func (jmq *JDModelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := jmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (jmq *JDModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jdmodel.Table,
			Columns: jdmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jdmodel.FieldID,
			},
		},
		From:   jmq.sql,
		Unique: true,
	}
	if ps := jmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, jdmodel.ValidColumn)
			}
		}
	}
	return _spec
}

func (jmq *JDModelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(jmq.driver.Dialect())
	t1 := builder.Table(jdmodel.Table)
	selector := builder.Select(t1.Columns(jdmodel.Columns...)...).From(t1)
	if jmq.sql != nil {
		selector = jmq.sql
		selector.Select(selector.Columns(jdmodel.Columns...)...)
	}
	for _, p := range jmq.predicates {
		p(selector)
	}
	for _, p := range jmq.order {
		p(selector, jdmodel.ValidColumn)
	}
	if offset := jmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JDModelGroupBy is the builder for group-by JDModel entities.
type JDModelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jmgb *JDModelGroupBy) Aggregate(fns ...AggregateFunc) *JDModelGroupBy {
	jmgb.fns = append(jmgb.fns, fns...)
	return jmgb
}

// Scan applies the group-by query and scan the result into the given value.
func (jmgb *JDModelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jmgb.path(ctx)
	if err != nil {
		return err
	}
	jmgb.sql = query
	return jmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jmgb *JDModelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jmgb.fields) > 1 {
		return nil, errors.New("ent: JDModelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jmgb *JDModelGroupBy) StringsX(ctx context.Context) []string {
	v, err := jmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jmgb *JDModelGroupBy) StringX(ctx context.Context) string {
	v, err := jmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jmgb.fields) > 1 {
		return nil, errors.New("ent: JDModelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jmgb *JDModelGroupBy) IntsX(ctx context.Context) []int {
	v, err := jmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jmgb *JDModelGroupBy) IntX(ctx context.Context) int {
	v, err := jmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jmgb.fields) > 1 {
		return nil, errors.New("ent: JDModelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jmgb *JDModelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jmgb *JDModelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jmgb.fields) > 1 {
		return nil, errors.New("ent: JDModelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jmgb *JDModelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (jmgb *JDModelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jmgb *JDModelGroupBy) BoolX(ctx context.Context) bool {
	v, err := jmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jmgb *JDModelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jmgb.fields {
		if !jdmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jmgb *JDModelGroupBy) sqlQuery() *sql.Selector {
	selector := jmgb.sql
	columns := make([]string, 0, len(jmgb.fields)+len(jmgb.fns))
	columns = append(columns, jmgb.fields...)
	for _, fn := range jmgb.fns {
		columns = append(columns, fn(selector, jdmodel.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(jmgb.fields...)
}

// JDModelSelect is the builder for select fields of JDModel entities.
type JDModelSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (jms *JDModelSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := jms.path(ctx)
	if err != nil {
		return err
	}
	jms.sql = query
	return jms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jms *JDModelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := jms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(jms.fields) > 1 {
		return nil, errors.New("ent: JDModelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := jms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jms *JDModelSelect) StringsX(ctx context.Context) []string {
	v, err := jms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jms *JDModelSelect) StringX(ctx context.Context) string {
	v, err := jms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(jms.fields) > 1 {
		return nil, errors.New("ent: JDModelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := jms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jms *JDModelSelect) IntsX(ctx context.Context) []int {
	v, err := jms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jms *JDModelSelect) IntX(ctx context.Context) int {
	v, err := jms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(jms.fields) > 1 {
		return nil, errors.New("ent: JDModelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := jms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jms *JDModelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := jms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jms *JDModelSelect) Float64X(ctx context.Context) float64 {
	v, err := jms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(jms.fields) > 1 {
		return nil, errors.New("ent: JDModelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := jms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jms *JDModelSelect) BoolsX(ctx context.Context) []bool {
	v, err := jms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (jms *JDModelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jdmodel.Label}
	default:
		err = fmt.Errorf("ent: JDModelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jms *JDModelSelect) BoolX(ctx context.Context) bool {
	v, err := jms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jms *JDModelSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jms.fields {
		if !jdmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := jms.sqlQuery().Query()
	if err := jms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jms *JDModelSelect) sqlQuery() sql.Querier {
	selector := jms.sql
	selector.Select(selector.Columns(jms.fields...)...)
	return selector
}
