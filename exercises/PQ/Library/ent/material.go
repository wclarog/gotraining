// Code generated by entc, DO NOT EDIT.

package ent

import (
	"excercise-library/ent/book"
	"excercise-library/ent/magazine"
	"excercise-library/ent/material"
	"excercise-library/ent/newspaper"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// Material is the model entity for the Material schema.
type Material struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UniqueCode holds the value of the "uniqueCode" field.
	UniqueCode string `json:"uniqueCode,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DateOfEmission holds the value of the "dateOfEmission" field.
	DateOfEmission time.Time `json:"dateOfEmission,omitempty"`
	// NumberOfPages holds the value of the "numberOfPages" field.
	NumberOfPages uint `json:"numberOfPages,omitempty"`
	// MaterialType holds the value of the "materialType" field.
	MaterialType int `json:"materialType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MaterialQuery when eager-loading is set.
	Edges MaterialEdges `json:"edges"`
}

// MaterialEdges holds the relations/edges for other nodes in the graph.
type MaterialEdges struct {
	// Book holds the value of the Book edge.
	Book *Book
	// Newspaper holds the value of the Newspaper edge.
	Newspaper *Newspaper
	// Magazine holds the value of the Magazine edge.
	Magazine *Magazine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MaterialEdges) BookOrErr() (*Book, error) {
	if e.loadedTypes[0] {
		if e.Book == nil {
			// The edge Book was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: book.Label}
		}
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "Book"}
}

// NewspaperOrErr returns the Newspaper value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MaterialEdges) NewspaperOrErr() (*Newspaper, error) {
	if e.loadedTypes[1] {
		if e.Newspaper == nil {
			// The edge Newspaper was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: newspaper.Label}
		}
		return e.Newspaper, nil
	}
	return nil, &NotLoadedError{edge: "Newspaper"}
}

// MagazineOrErr returns the Magazine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MaterialEdges) MagazineOrErr() (*Magazine, error) {
	if e.loadedTypes[2] {
		if e.Magazine == nil {
			// The edge Magazine was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: magazine.Label}
		}
		return e.Magazine, nil
	}
	return nil, &NotLoadedError{edge: "Magazine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Material) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // uniqueCode
		&sql.NullString{}, // name
		&sql.NullTime{},   // dateOfEmission
		&sql.NullInt64{},  // numberOfPages
		&sql.NullInt64{},  // materialType
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Material fields.
func (m *Material) assignValues(values ...interface{}) error {
	if m, n := len(values), len(material.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uniqueCode", values[0])
	} else if value.Valid {
		m.UniqueCode = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		m.Name = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field dateOfEmission", values[2])
	} else if value.Valid {
		m.DateOfEmission = value.Time
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field numberOfPages", values[3])
	} else if value.Valid {
		m.NumberOfPages = uint(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field materialType", values[4])
	} else if value.Valid {
		m.MaterialType = int(value.Int64)
	}
	return nil
}

// QueryBook queries the Book edge of the Material.
func (m *Material) QueryBook() *BookQuery {
	return (&MaterialClient{config: m.config}).QueryBook(m)
}

// QueryNewspaper queries the Newspaper edge of the Material.
func (m *Material) QueryNewspaper() *NewspaperQuery {
	return (&MaterialClient{config: m.config}).QueryNewspaper(m)
}

// QueryMagazine queries the Magazine edge of the Material.
func (m *Material) QueryMagazine() *MagazineQuery {
	return (&MaterialClient{config: m.config}).QueryMagazine(m)
}

// Update returns a builder for updating this Material.
// Note that, you need to call Material.Unwrap() before calling this method, if this Material
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Material) Update() *MaterialUpdateOne {
	return (&MaterialClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Material) Unwrap() *Material {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Material is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Material) String() string {
	var builder strings.Builder
	builder.WriteString("Material(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", uniqueCode=")
	builder.WriteString(m.UniqueCode)
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", dateOfEmission=")
	builder.WriteString(m.DateOfEmission.Format(time.ANSIC))
	builder.WriteString(", numberOfPages=")
	builder.WriteString(fmt.Sprintf("%v", m.NumberOfPages))
	builder.WriteString(", materialType=")
	builder.WriteString(fmt.Sprintf("%v", m.MaterialType))
	builder.WriteByte(')')
	return builder.String()
}

// Materials is a parsable slice of Material.
type Materials []*Material

func (m Materials) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}