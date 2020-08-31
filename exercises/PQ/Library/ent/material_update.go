// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"excercise-library/ent/book"
	"excercise-library/ent/magazine"
	"excercise-library/ent/material"
	"excercise-library/ent/newspaper"
	"excercise-library/ent/predicate"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MaterialUpdate is the builder for updating Material entities.
type MaterialUpdate struct {
	config
	hooks      []Hook
	mutation   *MaterialMutation
	predicates []predicate.Material
}

// Where adds a new predicate for the builder.
func (mu *MaterialUpdate) Where(ps ...predicate.Material) *MaterialUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *MaterialUpdate) SetName(s string) *MaterialUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetDateOfEmission sets the dateOfEmission field.
func (mu *MaterialUpdate) SetDateOfEmission(t time.Time) *MaterialUpdate {
	mu.mutation.SetDateOfEmission(t)
	return mu
}

// SetNumberOfPages sets the numberOfPages field.
func (mu *MaterialUpdate) SetNumberOfPages(i int) *MaterialUpdate {
	mu.mutation.ResetNumberOfPages()
	mu.mutation.SetNumberOfPages(i)
	return mu
}

// AddNumberOfPages adds i to numberOfPages.
func (mu *MaterialUpdate) AddNumberOfPages(i int) *MaterialUpdate {
	mu.mutation.AddNumberOfPages(i)
	return mu
}

// SetMaterialType sets the materialType field.
func (mu *MaterialUpdate) SetMaterialType(i int) *MaterialUpdate {
	mu.mutation.ResetMaterialType()
	mu.mutation.SetMaterialType(i)
	return mu
}

// AddMaterialType adds i to materialType.
func (mu *MaterialUpdate) AddMaterialType(i int) *MaterialUpdate {
	mu.mutation.AddMaterialType(i)
	return mu
}

// SetBookID sets the Book edge to Book by id.
func (mu *MaterialUpdate) SetBookID(id int) *MaterialUpdate {
	mu.mutation.SetBookID(id)
	return mu
}

// SetNillableBookID sets the Book edge to Book by id if the given value is not nil.
func (mu *MaterialUpdate) SetNillableBookID(id *int) *MaterialUpdate {
	if id != nil {
		mu = mu.SetBookID(*id)
	}
	return mu
}

// SetBook sets the Book edge to Book.
func (mu *MaterialUpdate) SetBook(b *Book) *MaterialUpdate {
	return mu.SetBookID(b.ID)
}

// SetNewspaperID sets the Newspaper edge to Newspaper by id.
func (mu *MaterialUpdate) SetNewspaperID(id int) *MaterialUpdate {
	mu.mutation.SetNewspaperID(id)
	return mu
}

// SetNillableNewspaperID sets the Newspaper edge to Newspaper by id if the given value is not nil.
func (mu *MaterialUpdate) SetNillableNewspaperID(id *int) *MaterialUpdate {
	if id != nil {
		mu = mu.SetNewspaperID(*id)
	}
	return mu
}

// SetNewspaper sets the Newspaper edge to Newspaper.
func (mu *MaterialUpdate) SetNewspaper(n *Newspaper) *MaterialUpdate {
	return mu.SetNewspaperID(n.ID)
}

// SetMagazineID sets the Magazine edge to Magazine by id.
func (mu *MaterialUpdate) SetMagazineID(id int) *MaterialUpdate {
	mu.mutation.SetMagazineID(id)
	return mu
}

// SetNillableMagazineID sets the Magazine edge to Magazine by id if the given value is not nil.
func (mu *MaterialUpdate) SetNillableMagazineID(id *int) *MaterialUpdate {
	if id != nil {
		mu = mu.SetMagazineID(*id)
	}
	return mu
}

// SetMagazine sets the Magazine edge to Magazine.
func (mu *MaterialUpdate) SetMagazine(m *Magazine) *MaterialUpdate {
	return mu.SetMagazineID(m.ID)
}

// Mutation returns the MaterialMutation object of the builder.
func (mu *MaterialUpdate) Mutation() *MaterialMutation {
	return mu.mutation
}

// ClearBook clears the Book edge to Book.
func (mu *MaterialUpdate) ClearBook() *MaterialUpdate {
	mu.mutation.ClearBook()
	return mu
}

// ClearNewspaper clears the Newspaper edge to Newspaper.
func (mu *MaterialUpdate) ClearNewspaper() *MaterialUpdate {
	mu.mutation.ClearNewspaper()
	return mu
}

// ClearMagazine clears the Magazine edge to Magazine.
func (mu *MaterialUpdate) ClearMagazine() *MaterialUpdate {
	mu.mutation.ClearMagazine()
	return mu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MaterialUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.NumberOfPages(); ok {
		if err := material.NumberOfPagesValidator(v); err != nil {
			return 0, &ValidationError{Name: "numberOfPages", err: fmt.Errorf("ent: validator failed for field \"numberOfPages\": %w", err)}
		}
	}
	if v, ok := mu.mutation.MaterialType(); ok {
		if err := material.MaterialTypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "materialType", err: fmt.Errorf("ent: validator failed for field \"materialType\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaterialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MaterialUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MaterialUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MaterialUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MaterialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   material.Table,
			Columns: material.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: material.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: material.FieldName,
		})
	}
	if value, ok := mu.mutation.DateOfEmission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: material.FieldDateOfEmission,
		})
	}
	if value, ok := mu.mutation.NumberOfPages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldNumberOfPages,
		})
	}
	if value, ok := mu.mutation.AddedNumberOfPages(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldNumberOfPages,
		})
	}
	if value, ok := mu.mutation.MaterialType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldMaterialType,
		})
	}
	if value, ok := mu.mutation.AddedMaterialType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldMaterialType,
		})
	}
	if mu.mutation.BookCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.BookIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.NewspaperCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.NewspaperIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MagazineCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MagazineIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{material.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MaterialUpdateOne is the builder for updating a single Material entity.
type MaterialUpdateOne struct {
	config
	hooks    []Hook
	mutation *MaterialMutation
}

// SetName sets the name field.
func (muo *MaterialUpdateOne) SetName(s string) *MaterialUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetDateOfEmission sets the dateOfEmission field.
func (muo *MaterialUpdateOne) SetDateOfEmission(t time.Time) *MaterialUpdateOne {
	muo.mutation.SetDateOfEmission(t)
	return muo
}

// SetNumberOfPages sets the numberOfPages field.
func (muo *MaterialUpdateOne) SetNumberOfPages(i int) *MaterialUpdateOne {
	muo.mutation.ResetNumberOfPages()
	muo.mutation.SetNumberOfPages(i)
	return muo
}

// AddNumberOfPages adds i to numberOfPages.
func (muo *MaterialUpdateOne) AddNumberOfPages(i int) *MaterialUpdateOne {
	muo.mutation.AddNumberOfPages(i)
	return muo
}

// SetMaterialType sets the materialType field.
func (muo *MaterialUpdateOne) SetMaterialType(i int) *MaterialUpdateOne {
	muo.mutation.ResetMaterialType()
	muo.mutation.SetMaterialType(i)
	return muo
}

// AddMaterialType adds i to materialType.
func (muo *MaterialUpdateOne) AddMaterialType(i int) *MaterialUpdateOne {
	muo.mutation.AddMaterialType(i)
	return muo
}

// SetBookID sets the Book edge to Book by id.
func (muo *MaterialUpdateOne) SetBookID(id int) *MaterialUpdateOne {
	muo.mutation.SetBookID(id)
	return muo
}

// SetNillableBookID sets the Book edge to Book by id if the given value is not nil.
func (muo *MaterialUpdateOne) SetNillableBookID(id *int) *MaterialUpdateOne {
	if id != nil {
		muo = muo.SetBookID(*id)
	}
	return muo
}

// SetBook sets the Book edge to Book.
func (muo *MaterialUpdateOne) SetBook(b *Book) *MaterialUpdateOne {
	return muo.SetBookID(b.ID)
}

// SetNewspaperID sets the Newspaper edge to Newspaper by id.
func (muo *MaterialUpdateOne) SetNewspaperID(id int) *MaterialUpdateOne {
	muo.mutation.SetNewspaperID(id)
	return muo
}

// SetNillableNewspaperID sets the Newspaper edge to Newspaper by id if the given value is not nil.
func (muo *MaterialUpdateOne) SetNillableNewspaperID(id *int) *MaterialUpdateOne {
	if id != nil {
		muo = muo.SetNewspaperID(*id)
	}
	return muo
}

// SetNewspaper sets the Newspaper edge to Newspaper.
func (muo *MaterialUpdateOne) SetNewspaper(n *Newspaper) *MaterialUpdateOne {
	return muo.SetNewspaperID(n.ID)
}

// SetMagazineID sets the Magazine edge to Magazine by id.
func (muo *MaterialUpdateOne) SetMagazineID(id int) *MaterialUpdateOne {
	muo.mutation.SetMagazineID(id)
	return muo
}

// SetNillableMagazineID sets the Magazine edge to Magazine by id if the given value is not nil.
func (muo *MaterialUpdateOne) SetNillableMagazineID(id *int) *MaterialUpdateOne {
	if id != nil {
		muo = muo.SetMagazineID(*id)
	}
	return muo
}

// SetMagazine sets the Magazine edge to Magazine.
func (muo *MaterialUpdateOne) SetMagazine(m *Magazine) *MaterialUpdateOne {
	return muo.SetMagazineID(m.ID)
}

// Mutation returns the MaterialMutation object of the builder.
func (muo *MaterialUpdateOne) Mutation() *MaterialMutation {
	return muo.mutation
}

// ClearBook clears the Book edge to Book.
func (muo *MaterialUpdateOne) ClearBook() *MaterialUpdateOne {
	muo.mutation.ClearBook()
	return muo
}

// ClearNewspaper clears the Newspaper edge to Newspaper.
func (muo *MaterialUpdateOne) ClearNewspaper() *MaterialUpdateOne {
	muo.mutation.ClearNewspaper()
	return muo
}

// ClearMagazine clears the Magazine edge to Magazine.
func (muo *MaterialUpdateOne) ClearMagazine() *MaterialUpdateOne {
	muo.mutation.ClearMagazine()
	return muo
}

// Save executes the query and returns the updated entity.
func (muo *MaterialUpdateOne) Save(ctx context.Context) (*Material, error) {
	if v, ok := muo.mutation.NumberOfPages(); ok {
		if err := material.NumberOfPagesValidator(v); err != nil {
			return nil, &ValidationError{Name: "numberOfPages", err: fmt.Errorf("ent: validator failed for field \"numberOfPages\": %w", err)}
		}
	}
	if v, ok := muo.mutation.MaterialType(); ok {
		if err := material.MaterialTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "materialType", err: fmt.Errorf("ent: validator failed for field \"materialType\": %w", err)}
		}
	}

	var (
		err  error
		node *Material
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MaterialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MaterialUpdateOne) SaveX(ctx context.Context) *Material {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MaterialUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MaterialUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MaterialUpdateOne) sqlSave(ctx context.Context) (m *Material, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   material.Table,
			Columns: material.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: material.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Material.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: material.FieldName,
		})
	}
	if value, ok := muo.mutation.DateOfEmission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: material.FieldDateOfEmission,
		})
	}
	if value, ok := muo.mutation.NumberOfPages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldNumberOfPages,
		})
	}
	if value, ok := muo.mutation.AddedNumberOfPages(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldNumberOfPages,
		})
	}
	if value, ok := muo.mutation.MaterialType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldMaterialType,
		})
	}
	if value, ok := muo.mutation.AddedMaterialType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: material.FieldMaterialType,
		})
	}
	if muo.mutation.BookCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.BookIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.NewspaperCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.NewspaperIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MagazineCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MagazineIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Material{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{material.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
