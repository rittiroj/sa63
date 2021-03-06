// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeRequisitions holds the string denoting the requisitions edge name in mutations.
	EdgeRequisitions = "requisitions"

	// Table holds the table name of the user in the database.
	Table = "users"
	// RequisitionsTable is the table the holds the requisitions relation/edge.
	RequisitionsTable = "requisitions"
	// RequisitionsInverseTable is the table name for the Requisition entity.
	// It exists in this package in order to avoid circular dependency with the "requisition" package.
	RequisitionsInverseTable = "requisitions"
	// RequisitionsColumn is the table column denoting the requisitions relation/edge.
	RequisitionsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
