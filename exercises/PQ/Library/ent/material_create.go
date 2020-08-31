// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"excercise-library/ent/book"
	"excercise-library/ent/magazine"
	"excercise-library/ent/material"
	"excercise-library/ent/newspaper"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MaterialCreate is the builder for creating a Material entity.
type MaterialCreate struct {
	config
	mutation *MaterialMutation
	hooks    []Hook
}

// SetUniqueCode sets the uniqueCode field.
func (mc *MaterialCreate) SetUniqueCode(s string) *MaterialCreate {
	mc.mutation.SetUniqueCode(s)
	return mc
}

// SetName sets the name field.
func (mc *MaterialCreate) SetName(s string) *MaterialCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetDateOfEmission sets the dateOfEmission field.
func (mc *MaterialCreate) SetDateOfEmission(t time.Time) *MaterialCreate {
	mc.mutation.SetDateOfEmission(t)
	return mc
}

// SetNumberOfPages sets the numberOfPages field.
func (mc *MaterialCreate) SetNumberOfPages(i int) *MaterialCreate {
	mc.mutation.SetNumberOfPages(i)
	return mc
}

// SetMaterialType sets the materialType field.
func (mc *MaterialCreate) SetMaterialType(i int) *MaterialCreate {
	mc.mutation.SetMaterialType(i)
	return mc
}

// SetBookID sets the Book edge to Book by id.
func (mc *MaterialCreate) SetBookID(id int) *MaterialCreate {
	mc.mutation.SetBookID(id)
	return mc
}

// SetNillableBookID sets the Book edge to Book by id if the given value is not nil.
func (mc *MaterialCreate) SetNillableBookID(id *int) *MaterialCreate {
	if id != nil {
		mc = mc.SetBookID(*id)
	}
	return mc
}

// SetBook sets the Book edge to Book.
func (mc *MaterialCreate) SetBook(b *Book) *MaterialCreate {
	return mc.SetBookID(b.ID)
}

// SetNewspaperID sets the Newspaper edge to Newspaper by id.
func (mc *MaterialCreate) SetNewspaperID(id int) *MaterialCreate {
	mc.mutation.SetNewspaperID(id)
	return mc
}

// SetNillableNewspaperID sets the Newspaper edge to Newspaper by id if the given value is not nil.
func (mc *MaterialCreate) SetNillableNewspaperID(id *int) *MaterialCreate {
	if id != nil {
		mc = mc.SetNewspaperID(*id)
	}
	return mc
}

// SetNewspaper sets the Newspaper edge to Newspaper.
func (mc *MaterialCreate) SetNewspaper(n *Newspaper) *MaterialCreate {
	return mc.SetNewspaperID(n.ID)
}

// SetMagazineID sets the Magazine edge to Magazine by id.
func (mc *MaterialCreate) SetMagazineID(id int) *MaterialCreate {
	mc.mutation.SetMagazineID(id)
	return mc
}

// SetNillableMagazineID sets the Magazine edge to Magazine by id if the given value is not nil.
func (mc *MaterialCreate) SetNillableMagazineID(id *int) *MaterialCreate {
	if id != nil {
		mc = mc.SetMagazineID(*id)
	}
	return mc
}

// SetMagazine sets the Magazine edge to Magazine.
func (mc *MaterialCreate) SetMagazine(m *Magazine) *MaterialCreate {
	return mc.SetMagazineID(m.ID)
}

// Mutation returns the MaterialMutation object of the builder.
func (mc *MaterialCreate) Mutation() *MaterialMutation {
	return mc.mutation
}

// Save creates the Material in the database.
func (mc *MaterialCreate) Save(ctx context.Context) (*Material, error) {
	if err := mc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Material
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaterialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MaterialCreate) SaveX(ctx context.Context) *Material {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MaterialCreate) preSave() error {
	if _, ok := mc.mutation.UniqueCode(); !ok {
		return &ValidationError{Name: "uniqueCode", err: errors.New("ent: missing required field \"uniqueCode\"")}
	}
	if v, ok := mc.mutation.UniqueCode(); ok {
		if err := material.UniqueCodeValidator(v); err != nil {
			return &ValidationError{Name: "uniqueCode", err: fmt.Errorf("ent: validator failed for field \"uniqueCode\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := mc.mutation.DateOfEmission(); !ok {
		return &ValidationError{Name: "dateOfEmission", err: errors.New("ent: missing required field \"dateOfEmission\"")}
	}
	if _, ok := mc.mutation.NumberOfPages(); !ok {
		return &ValidationError{Name: "numberOfPages", err: errors.New("ent: missing required field \"numberOfPages\"")}
	}
	if v, ok := mc.mutation.NumberOfPages(); ok {
		if err := material.NumberOfPagesValidator(v); err != nil {
			return &ValidationError{Name: "numberOfPages", err: fmt.Errorf("ent: validator failed for field \"numberOfPages\": %w", err)}
		}
	}
	if _, ok := mc.mutation.MaterialType(); !ok {
		return &ValidationError{Name: "materialType", err: errors.New("ent: missing required field \"materialType\"")}
	}
	if v, ok := mc.mutation.MaterialType(); ok {
		if err := material.MaterialTypeValidator(v); err != nil {
			return &ValidationError{Name: "materialType", err: fmt.Errorf("ent: validator failed for field \"materialType\": %w", err)}
		}
	}
	return nil
}

func (mc *MaterialCreate) sqlSave(ctx context.Context) (*Material, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MaterialCreate) createSpec() (*Material, *sqlgraph.CreateSpec) {
	var (
		m     = &Material{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: material.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: material.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.UniqueCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: material.FieldUniqueCode,
		})
		m.UniqueCode = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: material.FieldName,
		})
		m.Name = value
	}
	if value, ok := mc.mutation.DateOfEmission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: material.FieldDateOfEmission,
		})
		m.DateOfEmission = value
	}
	if value, ok := mc.mutation.NumberOfPages(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldNumberOfPages,
		})
		m.NumberOfPages = value
	}
	if value, ok := mc.mutation.MaterialType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldMaterialType,
		})
		m.MaterialType = value
	}
	if nodes := mc.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   material.BookTable,
			Columns: []string{material.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.NewspaperIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   material.NewspaperTable,
			Columns: []string{material.NewspaperColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: newspaper.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MagazineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   material.MagazineTable,
			Columns: []string{material.MagazineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: magazine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}

// MaterialCreateBulk is the builder for creating a bulk of Material entities.
type MaterialCreateBulk struct {
	config
	builders []*MaterialCreate
}

// Save creates the Material entities in the database.
func (mcb *MaterialCreateBulk) Save(ctx context.Context) ([]*Material, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Material, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*MaterialMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MaterialCreateBulk) SaveX(ctx context.Context) []*Material {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
