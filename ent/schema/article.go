package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("url").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Enum("status").NamedValues("Readin Done","Not Read"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}

func (Article) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
    }
}