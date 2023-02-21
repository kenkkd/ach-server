package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Thread holds the schema definition for the Thread entity.
type Thread struct {
	ent.Schema
}

// Mixin of the Thread.
// func (Anchor) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		TimeMixin{},
// 		DeletedMixin{},
// 	}
// }

// Fields of the Thread.
func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Thread.
func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
	}
}
