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

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
    return []ent.Field{
        field.Text("text").
            NotEmpty(),
        field.Time("created_at").
            Default(time.Now).
            Immutable(),
        field.Enum("status").
            NamedValues(
                "InProgress", "IN_PROGRESS",
                "Completed", "COMPLETED",
            ).
            Default("IN_PROGRESS"),
        field.Int("priority").
            Default(0),
    }
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("parent", Todo.Type).
            Unique().
            From("children"),
    }
}

func (Todo) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
    }
}