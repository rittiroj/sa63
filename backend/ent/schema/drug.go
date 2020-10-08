package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Drug holds the schema definition for the Drug entity.
type Drug struct {
	ent.Schema
}

// Fields of the Drug.
func (Drug) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty(),
		field.Int("value").Positive(),
	}
}

// Edges of the Drug.
func (Drug) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requisitions", Requisition.Type).StorageKey(edge.Column("drug_id")),
	}
}
