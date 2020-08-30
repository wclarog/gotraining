package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
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
	return nil
}
