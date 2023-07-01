package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Thread holds the schema definition for the Thread entity.
type Response struct {
	ent.Schema
}

// Mixin of the Thread.
func (Response) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Fields of the Thread.
func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.String("thread_id"),
		field.String("name"),
		field.Int("number"),
		field.String("content"),
	}
}

// Edges of the Thread.
func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread",Thread.Type).
			Ref("responses").
			Field("thread_id").
			Unique().
			Required(),
	}
}
