// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"excercise-library/ent/newspaper"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// NewspaperCreate is the builder for creating a Newspaper entity.
type NewspaperCreate struct {
	config
	mutation *NewspaperMutation
	hooks    []Hook
}

// SetURL sets the url field.
func (nc *NewspaperCreate) SetURL(s string) *NewspaperCreate {
	nc.mutation.SetURL(s)
	return nc
}

// Mutation returns the NewspaperMutation object of the builder.
func (nc *NewspaperCreate) Mutation() *NewspaperMutation {
	return nc.mutation
}

// Save creates the Newspaper in the database.
func (nc *NewspaperCreate) Save(ctx context.Context) (*Newspaper, error) {
	if err := nc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Newspaper
	)
	if len(nc.hooks) == 0 {
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewspaperMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NewspaperCreate) SaveX(ctx context.Context) *Newspaper {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NewspaperCreate) preSave() error {
	if _, ok := nc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	return nil
}

func (nc *NewspaperCreate) sqlSave(ctx context.Context) (*Newspaper, error) {
	n, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	n.ID = int(id)
	return n, nil
}

func (nc *NewspaperCreate) createSpec() (*Newspaper, *sqlgraph.CreateSpec) {
	var (
		n     = &Newspaper{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: newspaper.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newspaper.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newspaper.FieldURL,
		})
		n.URL = value
	}
	return n, _spec
}

// NewspaperCreateBulk is the builder for creating a bulk of Newspaper entities.
type NewspaperCreateBulk struct {
	config
	builders []*NewspaperCreate
}

// Save creates the Newspaper entities in the database.
func (ncb *NewspaperCreateBulk) Save(ctx context.Context) ([]*Newspaper, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Newspaper, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*NewspaperMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ncb *NewspaperCreateBulk) SaveX(ctx context.Context) []*Newspaper {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
