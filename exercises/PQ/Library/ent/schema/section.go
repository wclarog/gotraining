package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

func (Section) Config() ent.Config {
	return ent.Config{
		Table: "Section",
	}
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("content"),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "relatedMagazine" of type `Magazine`
		// and reference it to the "Magazine" edge (in Magazine schema)
		// explicitly using the `Ref` method.
		edge.From("relatedMagazine", Magazine.Type).
			Ref("Section").
			// setting the edge to unique, ensure
			// that a section can have only one magazine.
			Unique(),
	}
}
