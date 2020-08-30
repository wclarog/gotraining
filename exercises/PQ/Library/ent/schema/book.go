package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("authorName"),
		field.String("genre"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "relatedMaterial" of type `Material`
		// and reference it to the "Material" edge (in Material schema)
		// explicitly using the `Ref` method.
		edge.From("relatedMaterial", Material.Type).
			Ref("Material").
			// setting the edge to unique, ensure
			// that a book can have only one material.
			Unique(),
	}
}
