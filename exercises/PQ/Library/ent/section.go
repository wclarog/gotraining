// Code generated by entc, DO NOT EDIT.

package ent

import (
	"excercise-library/ent/section"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Content holds the value of the "content" field.
	Content          string `json:"content,omitempty"`
	magazine_section *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // code
		&sql.NullString{}, // content
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Section) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // magazine_section
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(values ...interface{}) error {
	if m, n := len(values), len(section.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field code", values[0])
	} else if value.Valid {
		s.Code = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field content", values[1])
	} else if value.Valid {
		s.Content = value.String
	}
	values = values[2:]
	if len(values) == len(section.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field magazine_section", value)
		} else if value.Valid {
			s.magazine_section = new(int)
			*s.magazine_section = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Section.
// Note that, you need to call Section.Unwrap() before calling this method, if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return (&SectionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", code=")
	builder.WriteString(s.Code)
	builder.WriteString(", content=")
	builder.WriteString(s.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section

func (s Sections) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
