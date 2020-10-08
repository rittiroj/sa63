package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// RegisterStore holds the schema definition for the RegisterStore entity.
type RegisterStore struct {
	ent.Schema
}

// Fields of the RegisterStore.
func (RegisterStore) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the RegisterStore.
func (RegisterStore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requisitions", Requisition.Type).StorageKey(edge.Column("registerstore_id")),
	}
}
