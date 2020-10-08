// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Admin/app/ent/registerstore"
	"github.com/facebookincubator/ent/dialect/sql"
)

// RegisterStore is the model entity for the RegisterStore schema.
type RegisterStore struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RegisterStoreQuery when eager-loading is set.
	Edges RegisterStoreEdges `json:"edges"`
}

// RegisterStoreEdges holds the relations/edges for other nodes in the graph.
type RegisterStoreEdges struct {
	// Requisitions holds the value of the requisitions edge.
	Requisitions []*Requisition
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RequisitionsOrErr returns the Requisitions value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterStoreEdges) RequisitionsOrErr() ([]*Requisition, error) {
	if e.loadedTypes[0] {
		return e.Requisitions, nil
	}
	return nil, &NotLoadedError{edge: "requisitions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RegisterStore) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RegisterStore fields.
func (rs *RegisterStore) assignValues(values ...interface{}) error {
	if m, n := len(values), len(registerstore.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rs.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		rs.Name = value.String
	}
	return nil
}

// QueryRequisitions queries the requisitions edge of the RegisterStore.
func (rs *RegisterStore) QueryRequisitions() *RequisitionQuery {
	return (&RegisterStoreClient{config: rs.config}).QueryRequisitions(rs)
}

// Update returns a builder for updating this RegisterStore.
// Note that, you need to call RegisterStore.Unwrap() before calling this method, if this RegisterStore
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *RegisterStore) Update() *RegisterStoreUpdateOne {
	return (&RegisterStoreClient{config: rs.config}).UpdateOne(rs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rs *RegisterStore) Unwrap() *RegisterStore {
	tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: RegisterStore is not a transactional entity")
	}
	rs.config.driver = tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *RegisterStore) String() string {
	var builder strings.Builder
	builder.WriteString("RegisterStore(")
	builder.WriteString(fmt.Sprintf("id=%v", rs.ID))
	builder.WriteString(", name=")
	builder.WriteString(rs.Name)
	builder.WriteByte(')')
	return builder.String()
}

// RegisterStores is a parsable slice of RegisterStore.
type RegisterStores []*RegisterStore

func (rs RegisterStores) config(cfg config) {
	for _i := range rs {
		rs[_i].config = cfg
	}
}
