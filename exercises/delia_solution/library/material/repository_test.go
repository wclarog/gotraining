package material

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRepository_GetMaterials(t *testing.T) {

	name := "GetMaterials should work fine"

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectationsWereMet()

	ctx := context.TODO() // context.WithValue(context.Background(), "test", "abc123") //
	repository := NewRepository(db)

	materials, err := repository.GetMaterials(ctx)

	t.Logf("Running test case: %s", name)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(materials), 100)

}

func TestRepository_GetMaterialByCode(t *testing.T) {
	wrongCode := "abc"
	sampleErr := errors.New(fmt.Sprintf("Material %s not found.", wrongCode))
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectationsWereMet()

	ctx := context.TODO() // context.WithValue(context.Background(), "test", "abc123") //
	repository := NewRepository(db)

	materials, _ := repository.GetMaterials(ctx)
	material := materials[0]

	tests := map[string]struct {
		response   MaterialDTO
		uniqueCode string
		err        error
	}{
		"successful": {
			response:   material,
			err:        nil,
			uniqueCode: material.UniqueCode,
		},
		"not found material error": {
			response:   MaterialDTO{},
			err:        sampleErr,
			uniqueCode: wrongCode,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response, err := repository.GetMaterialByCode(ctx, test.uniqueCode)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.response, response)

	}
}

func TestRepository_GetBooks(t *testing.T) {
	name := "GetBooks should work fine"

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectationsWereMet()

	ctx := context.TODO() // context.WithValue(context.Background(), "test", "abc123") //
	repository := NewRepository(db)

	books, err := repository.GetBooks(ctx)

	t.Logf("Running test case: %s", name)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(books), 40)

	for _, book := range books {
		assert.Equal(t, book.Type, BookType)
	}
}

func TestRepository_GetBookByCode(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectationsWereMet()

	ctx := context.TODO() // context.WithValue(context.Background(), "test", "abc123") //
	repository := NewRepository(db)

	mats, _ := repository.GetMaterials(ctx)
	wrongMatCode := mats[41].UniqueCode
	materialTypeErr := errors.New(fmt.Sprintf("Material %s is not a book.", wrongMatCode))

	books, _ := repository.GetBooks(ctx)
	book := books[0]
	wrongCode := "abc"
	bookNotFoundErr := errors.New(fmt.Sprintf("Book %s not found.", wrongCode))

	tests := map[string]struct {
		response   BookDTO
		uniqueCode string
		err        error
	}{
		"successful": {
			response:   book,
			err:        nil,
			uniqueCode: book.UniqueCode,
		},
		"not found book error": {
			response:   BookDTO{},
			err:        bookNotFoundErr,
			uniqueCode: wrongCode,
		},
		"invalid book type error": {
			response:   BookDTO{},
			err:        materialTypeErr,
			uniqueCode: wrongMatCode,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response, err := repository.GetBookByCode(ctx, test.uniqueCode)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.response, response)
	}
}
