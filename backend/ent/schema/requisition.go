package schema

import (
	

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Requisition holds the schema definition for the Requisition entity.
type Requisition struct {
	ent.Schema
}

// Fields of the Requisition.
func (Requisition) Fields() []ent.Field {
	return []ent.Field{
		
		field.Time("added_time"),
	}
}

// Edges of the Requisition.
func (Requisition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("requisitions").Unique(),
		edge.From("registerstore", RegisterStore.Type).Ref("requisitions").Unique(),
		edge.From("drug", Drug.Type).Ref("requisitions").Unique(),
	}
}
