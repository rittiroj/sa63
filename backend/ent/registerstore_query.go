// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/Admin/app/ent/predicate"
	"github.com/Admin/app/ent/registerstore"
	"github.com/Admin/app/ent/requisition"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RegisterStoreQuery is the builder for querying RegisterStore entities.
type RegisterStoreQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.RegisterStore
	// eager-loading edges.
	withRequisitions *RequisitionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rsq *RegisterStoreQuery) Where(ps ...predicate.RegisterStore) *RegisterStoreQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit adds a limit step to the query.
func (rsq *RegisterStoreQuery) Limit(limit int) *RegisterStoreQuery {
	rsq.limit = &limit
	return rsq
}

// Offset adds an offset step to the query.
func (rsq *RegisterStoreQuery) Offset(offset int) *RegisterStoreQuery {
	rsq.offset = &offset
	return rsq
}

// Order adds an order step to the query.
func (rsq *RegisterStoreQuery) Order(o ...OrderFunc) *RegisterStoreQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// QueryRequisitions chains the current query on the requisitions edge.
func (rsq *RegisterStoreQuery) QueryRequisitions() *RequisitionQuery {
	query := &RequisitionQuery{config: rsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(registerstore.Table, registerstore.FieldID, rsq.sqlQuery()),
			sqlgraph.To(requisition.Table, requisition.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, registerstore.RequisitionsTable, registerstore.RequisitionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RegisterStore entity in the query. Returns *NotFoundError when no registerstore was found.
func (rsq *RegisterStoreQuery) First(ctx context.Context) (*RegisterStore, error) {
	rsSlice, err := rsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rsSlice) == 0 {
		return nil, &NotFoundError{registerstore.Label}
	}
	return rsSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *RegisterStoreQuery) FirstX(ctx context.Context) *RegisterStore {
	rs, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return rs
}

// FirstID returns the first RegisterStore id in the query. Returns *NotFoundError when no id was found.
func (rsq *RegisterStoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{registerstore.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rsq *RegisterStoreQuery) FirstXID(ctx context.Context) int {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only RegisterStore entity in the query, returns an error if not exactly one entity was returned.
func (rsq *RegisterStoreQuery) Only(ctx context.Context) (*RegisterStore, error) {
	rsSlice, err := rsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rsSlice) {
	case 1:
		return rsSlice[0], nil
	case 0:
		return nil, &NotFoundError{registerstore.Label}
	default:
		return nil, &NotSingularError{registerstore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *RegisterStoreQuery) OnlyX(ctx context.Context) *RegisterStore {
	rs, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// OnlyID returns the only RegisterStore id in the query, returns an error if not exactly one id was returned.
func (rsq *RegisterStoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = &NotSingularError{registerstore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *RegisterStoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RegisterStores.
func (rsq *RegisterStoreQuery) All(ctx context.Context) ([]*RegisterStore, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rsq *RegisterStoreQuery) AllX(ctx context.Context) []*RegisterStore {
	rsSlice, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rsSlice
}

// IDs executes the query and returns a list of RegisterStore ids.
func (rsq *RegisterStoreQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rsq.Select(registerstore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *RegisterStoreQuery) IDsX(ctx context.Context) []int {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *RegisterStoreQuery) Count(ctx context.Context) (int, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *RegisterStoreQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *RegisterStoreQuery) Exist(ctx context.Context) (bool, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *RegisterStoreQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *RegisterStoreQuery) Clone() *RegisterStoreQuery {
	return &RegisterStoreQuery{
		config:     rsq.config,
		limit:      rsq.limit,
		offset:     rsq.offset,
		order:      append([]OrderFunc{}, rsq.order...),
		unique:     append([]string{}, rsq.unique...),
		predicates: append([]predicate.RegisterStore{}, rsq.predicates...),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

//  WithRequisitions tells the query-builder to eager-loads the nodes that are connected to
// the "requisitions" edge. The optional arguments used to configure the query builder of the edge.
func (rsq *RegisterStoreQuery) WithRequisitions(opts ...func(*RequisitionQuery)) *RegisterStoreQuery {
	query := &RequisitionQuery{config: rsq.config}
	for _, opt := range opts {
		opt(query)
	}
	rsq.withRequisitions = query
	return rsq
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
//	client.RegisterStore.Query().
//		GroupBy(registerstore.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rsq *RegisterStoreQuery) GroupBy(field string, fields ...string) *RegisterStoreGroupBy {
	group := &RegisterStoreGroupBy{config: rsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
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
//	client.RegisterStore.Query().
//		Select(registerstore.FieldName).
//		Scan(ctx, &v)
//
func (rsq *RegisterStoreQuery) Select(field string, fields ...string) *RegisterStoreSelect {
	selector := &RegisterStoreSelect{config: rsq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(), nil
	}
	return selector
}

func (rsq *RegisterStoreQuery) prepareQuery(ctx context.Context) error {
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	return nil
}

func (rsq *RegisterStoreQuery) sqlAll(ctx context.Context) ([]*RegisterStore, error) {
	var (
		nodes       = []*RegisterStore{}
		_spec       = rsq.querySpec()
		loadedTypes = [1]bool{
			rsq.withRequisitions != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &RegisterStore{config: rsq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rsq.withRequisitions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*RegisterStore)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Requisition(func(s *sql.Selector) {
			s.Where(sql.InValues(registerstore.RequisitionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.registerstore_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "registerstore_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "registerstore_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Requisitions = append(node.Edges.Requisitions, n)
		}
	}

	return nodes, nil
}

func (rsq *RegisterStoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *RegisterStoreQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rsq *RegisterStoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registerstore.Table,
			Columns: registerstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registerstore.FieldID,
			},
		},
		From:   rsq.sql,
		Unique: true,
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rsq *RegisterStoreQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(registerstore.Table)
	selector := builder.Select(t1.Columns(registerstore.Columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(registerstore.Columns...)...)
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector)
	}
	if offset := rsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RegisterStoreGroupBy is the builder for group-by RegisterStore entities.
type RegisterStoreGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *RegisterStoreGroupBy) Aggregate(fns ...AggregateFunc) *RegisterStoreGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rsgb *RegisterStoreGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rsgb.path(ctx)
	if err != nil {
		return err
	}
	rsgb.sql = query
	return rsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) StringsX(ctx context.Context) []string {
	v, err := rsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) StringX(ctx context.Context) string {
	v, err := rsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) IntsX(ctx context.Context) []int {
	v, err := rsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) IntX(ctx context.Context) int {
	v, err := rsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rsgb *RegisterStoreGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rsgb *RegisterStoreGroupBy) BoolX(ctx context.Context) bool {
	v, err := rsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rsgb *RegisterStoreGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rsgb.sqlQuery().Query()
	if err := rsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rsgb *RegisterStoreGroupBy) sqlQuery() *sql.Selector {
	selector := rsgb.sql
	columns := make([]string, 0, len(rsgb.fields)+len(rsgb.fns))
	columns = append(columns, rsgb.fields...)
	for _, fn := range rsgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rsgb.fields...)
}

// RegisterStoreSelect is the builder for select fields of RegisterStore entities.
type RegisterStoreSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rss *RegisterStoreSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rss.path(ctx)
	if err != nil {
		return err
	}
	rss.sql = query
	return rss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rss *RegisterStoreSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rss *RegisterStoreSelect) StringsX(ctx context.Context) []string {
	v, err := rss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rss *RegisterStoreSelect) StringX(ctx context.Context) string {
	v, err := rss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rss *RegisterStoreSelect) IntsX(ctx context.Context) []int {
	v, err := rss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rss *RegisterStoreSelect) IntX(ctx context.Context) int {
	v, err := rss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rss *RegisterStoreSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rss *RegisterStoreSelect) Float64X(ctx context.Context) float64 {
	v, err := rss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: RegisterStoreSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rss *RegisterStoreSelect) BoolsX(ctx context.Context) []bool {
	v, err := rss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rss *RegisterStoreSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{registerstore.Label}
	default:
		err = fmt.Errorf("ent: RegisterStoreSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rss *RegisterStoreSelect) BoolX(ctx context.Context) bool {
	v, err := rss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rss *RegisterStoreSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rss.sqlQuery().Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rss *RegisterStoreSelect) sqlQuery() sql.Querier {
	selector := rss.sql
	selector.Select(selector.Columns(rss.fields...)...)
	return selector
}
