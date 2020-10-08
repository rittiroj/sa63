// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Admin/app/ent/drug"
	"github.com/Admin/app/ent/predicate"
	"github.com/Admin/app/ent/registerstore"
	"github.com/Admin/app/ent/requisition"
	"github.com/Admin/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RequisitionQuery is the builder for querying Requisition entities.
type RequisitionQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Requisition
	// eager-loading edges.
	withUser          *UserQuery
	withRegisterstore *RegisterStoreQuery
	withDrug          *DrugQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *RequisitionQuery) Where(ps ...predicate.Requisition) *RequisitionQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RequisitionQuery) Limit(limit int) *RequisitionQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RequisitionQuery) Offset(offset int) *RequisitionQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *RequisitionQuery) Order(o ...OrderFunc) *RequisitionQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the user edge.
func (rq *RequisitionQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(requisition.Table, requisition.FieldID, rq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, requisition.UserTable, requisition.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegisterstore chains the current query on the registerstore edge.
func (rq *RequisitionQuery) QueryRegisterstore() *RegisterStoreQuery {
	query := &RegisterStoreQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(requisition.Table, requisition.FieldID, rq.sqlQuery()),
			sqlgraph.To(registerstore.Table, registerstore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, requisition.RegisterstoreTable, requisition.RegisterstoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDrug chains the current query on the drug edge.
func (rq *RequisitionQuery) QueryDrug() *DrugQuery {
	query := &DrugQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(requisition.Table, requisition.FieldID, rq.sqlQuery()),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, requisition.DrugTable, requisition.DrugColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Requisition entity in the query. Returns *NotFoundError when no requisition was found.
func (rq *RequisitionQuery) First(ctx context.Context) (*Requisition, error) {
	rs, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, &NotFoundError{requisition.Label}
	}
	return rs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RequisitionQuery) FirstX(ctx context.Context) *Requisition {
	r, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return r
}

// FirstID returns the first Requisition id in the query. Returns *NotFoundError when no id was found.
func (rq *RequisitionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{requisition.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rq *RequisitionQuery) FirstXID(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Requisition entity in the query, returns an error if not exactly one entity was returned.
func (rq *RequisitionQuery) Only(ctx context.Context) (*Requisition, error) {
	rs, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rs) {
	case 1:
		return rs[0], nil
	case 0:
		return nil, &NotFoundError{requisition.Label}
	default:
		return nil, &NotSingularError{requisition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RequisitionQuery) OnlyX(ctx context.Context) *Requisition {
	r, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// OnlyID returns the only Requisition id in the query, returns an error if not exactly one id was returned.
func (rq *RequisitionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = &NotSingularError{requisition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RequisitionQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Requisitions.
func (rq *RequisitionQuery) All(ctx context.Context) ([]*Requisition, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RequisitionQuery) AllX(ctx context.Context) []*Requisition {
	rs, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// IDs executes the query and returns a list of Requisition ids.
func (rq *RequisitionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(requisition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RequisitionQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RequisitionQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RequisitionQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RequisitionQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RequisitionQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RequisitionQuery) Clone() *RequisitionQuery {
	return &RequisitionQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Requisition{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RequisitionQuery) WithUser(opts ...func(*UserQuery)) *RequisitionQuery {
	query := &UserQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

//  WithRegisterstore tells the query-builder to eager-loads the nodes that are connected to
// the "registerstore" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RequisitionQuery) WithRegisterstore(opts ...func(*RegisterStoreQuery)) *RequisitionQuery {
	query := &RegisterStoreQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRegisterstore = query
	return rq
}

//  WithDrug tells the query-builder to eager-loads the nodes that are connected to
// the "drug" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RequisitionQuery) WithDrug(opts ...func(*DrugQuery)) *RequisitionQuery {
	query := &DrugQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withDrug = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Requisition.Query().
//		GroupBy(requisition.FieldAddedTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RequisitionQuery) GroupBy(field string, fields ...string) *RequisitionGroupBy {
	group := &RequisitionGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"added_time,omitempty"`
//	}
//
//	client.Requisition.Query().
//		Select(requisition.FieldAddedTime).
//		Scan(ctx, &v)
//
func (rq *RequisitionQuery) Select(field string, fields ...string) *RequisitionSelect {
	selector := &RequisitionSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *RequisitionQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RequisitionQuery) sqlAll(ctx context.Context) ([]*Requisition, error) {
	var (
		nodes       = []*Requisition{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withUser != nil,
			rq.withRegisterstore != nil,
			rq.withDrug != nil,
		}
	)
	if rq.withUser != nil || rq.withRegisterstore != nil || rq.withDrug != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, requisition.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Requisition{config: rq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Requisition)
		for i := range nodes {
			if fk := nodes[i].user_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := rq.withRegisterstore; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Requisition)
		for i := range nodes {
			if fk := nodes[i].registerstore_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(registerstore.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "registerstore_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Registerstore = n
			}
		}
	}

	if query := rq.withDrug; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Requisition)
		for i := range nodes {
			if fk := nodes[i].drug_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(drug.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "drug_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Drug = n
			}
		}
	}

	return nodes, nil
}

func (rq *RequisitionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RequisitionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *RequisitionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requisition.Table,
			Columns: requisition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: requisition.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RequisitionQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(requisition.Table)
	selector := builder.Select(t1.Columns(requisition.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(requisition.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RequisitionGroupBy is the builder for group-by Requisition entities.
type RequisitionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RequisitionGroupBy) Aggregate(fns ...AggregateFunc) *RequisitionGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *RequisitionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RequisitionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RequisitionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RequisitionGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RequisitionGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RequisitionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RequisitionGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RequisitionGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RequisitionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RequisitionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RequisitionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RequisitionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RequisitionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *RequisitionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RequisitionGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RequisitionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgb.sqlQuery().Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RequisitionGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// RequisitionSelect is the builder for select fields of Requisition entities.
type RequisitionSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *RequisitionSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RequisitionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RequisitionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RequisitionSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RequisitionSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RequisitionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RequisitionSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RequisitionSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RequisitionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RequisitionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RequisitionSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RequisitionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RequisitionSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *RequisitionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{requisition.Label}
	default:
		err = fmt.Errorf("ent: RequisitionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RequisitionSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RequisitionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *RequisitionSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
