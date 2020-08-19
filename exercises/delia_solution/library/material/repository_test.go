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
