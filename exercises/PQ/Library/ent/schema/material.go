package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Material holds the schema definition for the Material entity.
type Material struct {
	ent.Schema
}

// Fields of the Material.
func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.String("uniqueCode").
			Unique(),
		field.String("name"),
		field.Time("dateOfEmission"),
		field.Int("numberOfPages").
			Positive(),
		field.Int("materialType").
			NonNegative(),
	}
}

// Edges of the Material.
func (Material) Edges() []ent.Edge {
	return nil
}
