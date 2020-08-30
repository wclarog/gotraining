package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Newspaper holds the schema definition for the Newspaper entity.
type Newspaper struct {
	ent.Schema
}

// Fields of the Newspaper.
func (Newspaper) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
	}
}

// Edges of the Newspaper.
func (Newspaper) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "relatedMaterial" of type `Material`
		// and reference it to the "Material" edge (in Material schema)
		// explicitly using the `Ref` method.
		edge.From("relatedMaterial", Material.Type).
			Ref("Material").
			// setting the edge to unique, ensure
			// that a newspaper can have only one material.
			Unique(),
	}
}