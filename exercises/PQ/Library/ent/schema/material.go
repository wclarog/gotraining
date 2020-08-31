package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Material holds the schema definition for the Material entity.
type Material struct {
	ent.Schema
}

func (Material) Config() ent.Config {
	return ent.Config{
		Table: "Material",
	}
}

// Fields of the Material.
func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.String("uniqueCode").
			Unique(),
		field.String("name"),
		field.Time("dateOfEmission"),
		field.Uint("numberOfPages"),
		field.Int("materialType").
			NonNegative(),
	}
}

// Edges of the Material.
func (Material) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Book", Book.Type).StorageKey(edge.Column("material_id")).Unique(),
		edge.To("Newspaper", Newspaper.Type).StorageKey(edge.Column("material_id")).Unique(),
		edge.To("Magazine", Magazine.Type).StorageKey(edge.Column("material_id")).Unique(),
	}
}
