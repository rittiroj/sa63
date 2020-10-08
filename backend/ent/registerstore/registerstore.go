// Code generated by entc, DO NOT EDIT.

package registerstore

const (
	// Label holds the string label denoting the registerstore type in the database.
	Label = "register_store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeRequisitions holds the string denoting the requisitions edge name in mutations.
	EdgeRequisitions = "requisitions"

	// Table holds the table name of the registerstore in the database.
	Table = "register_stores"
	// RequisitionsTable is the table the holds the requisitions relation/edge.
	RequisitionsTable = "requisitions"
	// RequisitionsInverseTable is the table name for the Requisition entity.
	// It exists in this package in order to avoid circular dependency with the "requisition" package.
	RequisitionsInverseTable = "requisitions"
	// RequisitionsColumn is the table column denoting the requisitions relation/edge.
	RequisitionsColumn = "registerstore_id"
)

// Columns holds all SQL columns for registerstore fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
