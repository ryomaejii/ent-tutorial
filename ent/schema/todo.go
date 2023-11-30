package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
  return []ent.Field{
        field.Text("text").
            NotEmpty().
            Annotations(
                entgql.OrderField("TEXT"),
            ),
        field.Time("created_at").
            Default(time.Now).
            Immutable().
            Annotations(
                entgql.OrderField("CREATED_AT"),
            ),
        field.Enum("status").
            NamedValues(
                "InProgress", "IN_PROGRESS",
                "Completed", "COMPLETED",
            ).
            Default("IN_PROGRESS").
            Annotations(
                entgql.OrderField("STATUS"),
            ),
        field.Int("priority").
            Default(0).
            Annotations(
                entgql.OrderField("PRIORITY"),
            ),
    }
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Todo.Type).
				Annotations(
						entgql.RelayConnection(),
						entgql.OrderField("CHILDREN_COUNT"),
				).
				From("parent").
				Unique(),
}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
			entgql.QueryField(),
			entgql.Mutations(entgql.MutationCreate()),
 			entgql.RelayConnection(),
	}
}
