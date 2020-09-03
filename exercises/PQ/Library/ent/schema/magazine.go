package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Magazine holds the schema definition for the Magazine entity.
type Magazine struct {
	ent.Schema
}

func (Magazine) Config() ent.Config {
	return ent.Config{
		Table: "Magazine",
	}
}

// Fields of the Magazine.
func (Magazine) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
	}
}

// Edges of the Magazine.
func (Magazine) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "relatedMaterial" of type `Material`
		// and reference it to the "Material" edge (in Material schema)
		// explicitly using the `Ref` method.
		edge.From("relatedMaterial", Material.Type).
			Ref("Magazine").
			// setting the edge to unique, ensure
			// that a magazine can have only one material.
			Unique(),
		edge.To("Section", Section.Type),
	}
}
