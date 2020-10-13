// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Admin/app/ent/predicate"
	"github.com/Admin/app/ent/registerstore"
	"github.com/Admin/app/ent/requisition"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RegisterStoreUpdate is the builder for updating RegisterStore entities.
type RegisterStoreUpdate struct {
	config
	hooks      []Hook
	mutation   *RegisterStoreMutation
	predicates []predicate.RegisterStore
}

// Where adds a new predicate for the builder.
func (rsu *RegisterStoreUpdate) Where(ps ...predicate.RegisterStore) *RegisterStoreUpdate {
	rsu.predicates = append(rsu.predicates, ps...)
	return rsu
}

// SetName sets the name field.
func (rsu *RegisterStoreUpdate) SetName(s string) *RegisterStoreUpdate {
	rsu.mutation.SetName(s)
	return rsu
}

// AddRequisitionIDs adds the requisitions edge to Requisition by ids.
func (rsu *RegisterStoreUpdate) AddRequisitionIDs(ids ...int) *RegisterStoreUpdate {
	rsu.mutation.AddRequisitionIDs(ids...)
	return rsu
}

// AddRequisitions adds the requisitions edges to Requisition.
func (rsu *RegisterStoreUpdate) AddRequisitions(r ...*Requisition) *RegisterStoreUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsu.AddRequisitionIDs(ids...)
}

// Mutation returns the RegisterStoreMutation object of the builder.
func (rsu *RegisterStoreUpdate) Mutation() *RegisterStoreMutation {
	return rsu.mutation
}

// RemoveRequisitionIDs removes the requisitions edge to Requisition by ids.
func (rsu *RegisterStoreUpdate) RemoveRequisitionIDs(ids ...int) *RegisterStoreUpdate {
	rsu.mutation.RemoveRequisitionIDs(ids...)
	return rsu
}

// RemoveRequisitions removes requisitions edges to Requisition.
func (rsu *RegisterStoreUpdate) RemoveRequisitions(r ...*Requisition) *RegisterStoreUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsu.RemoveRequisitionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (rsu *RegisterStoreUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := rsu.mutation.Name(); ok {
		if err := registerstore.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(rsu.hooks) == 0 {
		affected, err = rsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegisterStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsu.mutation = mutation
			affected, err = rsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rsu.hooks) - 1; i >= 0; i-- {
			mut = rsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsu *RegisterStoreUpdate) SaveX(ctx context.Context) int {
	affected, err := rsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rsu *RegisterStoreUpdate) Exec(ctx context.Context) error {
	_, err := rsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsu *RegisterStoreUpdate) ExecX(ctx context.Context) {
	if err := rsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rsu *RegisterStoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registerstore.Table,
			Columns: registerstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registerstore.FieldID,
			},
		},
	}
	if ps := rsu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: registerstore.FieldName,
		})
	}
	if nodes := rsu.mutation.RemovedRequisitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   registerstore.RequisitionsTable,
			Columns: []string{registerstore.RequisitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: requisition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsu.mutation.RequisitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   registerstore.RequisitionsTable,
			Columns: []string{registerstore.RequisitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: requisition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registerstore.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RegisterStoreUpdateOne is the builder for updating a single RegisterStore entity.
type RegisterStoreUpdateOne struct {
	config
	hooks    []Hook
	mutation *RegisterStoreMutation
}

// SetName sets the name field.
func (rsuo *RegisterStoreUpdateOne) SetName(s string) *RegisterStoreUpdateOne {
	rsuo.mutation.SetName(s)
	return rsuo
}

// AddRequisitionIDs adds the requisitions edge to Requisition by ids.
func (rsuo *RegisterStoreUpdateOne) AddRequisitionIDs(ids ...int) *RegisterStoreUpdateOne {
	rsuo.mutation.AddRequisitionIDs(ids...)
	return rsuo
}

// AddRequisitions adds the requisitions edges to Requisition.
func (rsuo *RegisterStoreUpdateOne) AddRequisitions(r ...*Requisition) *RegisterStoreUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsuo.AddRequisitionIDs(ids...)
}

// Mutation returns the RegisterStoreMutation object of the builder.
func (rsuo *RegisterStoreUpdateOne) Mutation() *RegisterStoreMutation {
	return rsuo.mutation
}

// RemoveRequisitionIDs removes the requisitions edge to Requisition by ids.
func (rsuo *RegisterStoreUpdateOne) RemoveRequisitionIDs(ids ...int) *RegisterStoreUpdateOne {
	rsuo.mutation.RemoveRequisitionIDs(ids...)
	return rsuo
}

// RemoveRequisitions removes requisitions edges to Requisition.
func (rsuo *RegisterStoreUpdateOne) RemoveRequisitions(r ...*Requisition) *RegisterStoreUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rsuo.RemoveRequisitionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (rsuo *RegisterStoreUpdateOne) Save(ctx context.Context) (*RegisterStore, error) {
	if v, ok := rsuo.mutation.Name(); ok {
		if err := registerstore.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *RegisterStore
	)
	if len(rsuo.hooks) == 0 {
		node, err = rsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegisterStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rsuo.mutation = mutation
			node, err = rsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rsuo.hooks) - 1; i >= 0; i-- {
			mut = rsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rsuo *RegisterStoreUpdateOne) SaveX(ctx context.Context) *RegisterStore {
	rs, err := rsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// Exec executes the query on the entity.
func (rsuo *RegisterStoreUpdateOne) Exec(ctx context.Context) error {
	_, err := rsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *RegisterStoreUpdateOne) ExecX(ctx context.Context) {
	if err := rsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rsuo *RegisterStoreUpdateOne) sqlSave(ctx context.Context) (rs *RegisterStore, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registerstore.Table,
			Columns: registerstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registerstore.FieldID,
			},
		},
	}
	id, ok := rsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RegisterStore.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := rsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: registerstore.FieldName,
		})
	}
	if nodes := rsuo.mutation.RemovedRequisitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   registerstore.RequisitionsTable,
			Columns: []string{registerstore.RequisitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: requisition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rsuo.mutation.RequisitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   registerstore.RequisitionsTable,
			Columns: []string{registerstore.RequisitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: requisition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	rs = &RegisterStore{config: rsuo.config}
	_spec.Assign = rs.assignValues
	_spec.ScanValues = rs.scanValues()
	if err = sqlgraph.UpdateNode(ctx, rsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registerstore.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return rs, nil
}