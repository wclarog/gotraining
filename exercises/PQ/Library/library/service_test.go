package library

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	repMock  = NewRepositoryMock()
	logger   = log.NewLogfmtLogger(os.Stderr)
	servMock = NewService(repMock, logger)
)

func TestService_GetMaterialsSuccess(t *testing.T) {
	materials, err := servMock.GetMaterials(context.Background())
	expected, _ := servMock.dtoToMaterials(MockMaterials)
	assert.Equal(t, expected, materials)
	assert.Nil(t, err)
}

func TestService_GetMaterialsError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetMaterialsError", true)
	materials, err := servMock.GetMaterials(ctx)
	assert.EqualError(t, err, ErrDefault.Error())
	assert.Nil(t, materials)
}

func TestService_GetMaterialByCodeSuccess(t *testing.T) {
	materials, err := servMock.GetMaterialByCode(context.Background(), "abc")
	expected, _ := servMock.dtoToMaterial(MockMaterials[0])
	assert.Equal(t, expected, materials)
	assert.Nil(t, err)
}

func TestService_GetMaterialByCodeError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetMaterialByCodeError", true)
	_, err := servMock.GetMaterialByCode(ctx, "abc")
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_DeleteMaterialSuccess(t *testing.T) {
	err := servMock.DeleteMaterial(context.Background(), "abc")
	assert.Nil(t, err)
}

func TestService_DeleteMaterialError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "DeleteMaterialError", true)
	err := servMock.DeleteMaterial(ctx, "abc")
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_GetBooksSuccess(t *testing.T) {
	books, err := servMock.GetBooks(context.Background())
	expected, _ := servMock.dtoToBooks(MockBooks)
	assert.Equal(t, expected, books)
	assert.Nil(t, err)
}

func TestService_GetBooksError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetBooksError", true)
	expected, err := servMock.GetBooks(ctx)
	assert.EqualError(t, err, ErrDefault.Error())
	assert.Nil(t, expected)
}

func TestService_GetBookByCodeSuccess(t *testing.T) {
	book, err := servMock.GetBookByCode(context.Background(), "abc")
	expected, _ := servMock.dtoToBook(MockBooks[0])
	assert.Equal(t, expected, book)
	assert.Nil(t, err)
}

func TestService_GetBookByCodeError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetBookByCodeError", true)
	_, err := servMock.GetBookByCode(ctx, "abc")
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_AddBookSuccess(t *testing.T) {
	book, err := servMock.AddBook(context.Background(), Book{})
	expected, _ := servMock.dtoToBook(MockBooks[0])
	assert.Equal(t, expected, book)
	assert.Nil(t, err)
}

func TestService_AddBookError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "AddBookError", true)
	_, err := servMock.AddBook(ctx, Book{})
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_UpdateBookSuccess(t *testing.T) {
	book, err := servMock.UpdateBook(context.Background(), "abc", Book{})
	expected, _ := servMock.dtoToBook(MockBooks[0])
	assert.Equal(t, expected, book)
	assert.Nil(t, err)
}

func TestService_UpdateBookError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "UpdateBookError", true)
	_, err := servMock.UpdateBook(ctx, "abc", Book{})
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_GetNewspapersSuccess(t *testing.T) {
	newspapers, err := servMock.GetNewspapers(context.Background())
	expected, _ := servMock.dtoToNewspapers(MockNewspapers)
	assert.Equal(t, expected, newspapers)
	assert.Nil(t, err)
}

func TestService_GetNewspapersError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetNewspapersError", true)
	expected, err := servMock.GetNewspapers(ctx)
	assert.EqualError(t, err, ErrDefault.Error())
	assert.Nil(t, expected)
}

func TestService_GetNewspaperByCodeSuccess(t *testing.T) {
	newspaper, err := servMock.GetNewspaperByCode(context.Background(), "abc")
	expected, _ := servMock.dtoToNewspaper(MockNewspapers[0])
	assert.Equal(t, expected, newspaper)
	assert.Nil(t, err)
}

func TestService_GetNewspaperByCodeError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetNewspaperByCodeError", true)
	_, err := servMock.GetNewspaperByCode(ctx, "abc")
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_AddNewspaperSuccess(t *testing.T) {
	newspaper, err := servMock.AddNewspaper(context.Background(), Newspaper{})
	expected, _ := servMock.dtoToNewspaper(MockNewspapers[0])
	assert.Equal(t, expected, newspaper)
	assert.Nil(t, err)
}

func TestService_AddNewspaperError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "AddNewspaperError", true)
	_, err := servMock.AddNewspaper(ctx, Newspaper{})
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_UpdateNewspaperSuccess(t *testing.T) {
	newspaper, err := servMock.UpdateNewspaper(context.Background(), "abc", Newspaper{})
	expected, _ := servMock.dtoToNewspaper(MockNewspapers[0])
	assert.Equal(t, expected, newspaper)
	assert.Nil(t, err)
}

func TestService_UpdateNewspaperError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "UpdateNewspaperError", true)
	_, err := servMock.UpdateNewspaper(ctx, "abc", Newspaper{})
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_GetMagazinesSuccess(t *testing.T) {
	magazines, err := servMock.GetMagazines(context.Background())
	expected, _ := servMock.dtoToMagazines(MockMagazines)
	assert.Equal(t, expected, magazines)
	assert.Nil(t, err)
}

func TestService_GetMagazinesError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetMagazinesError", true)
	expected, err := servMock.GetMagazines(ctx)
	assert.EqualError(t, err, ErrDefault.Error())
	assert.Nil(t, expected)
}

func TestService_GetMagazineByCodeSuccess(t *testing.T) {
	magazine, err := servMock.GetMagazineByCode(context.Background(), "abc")
	expected, _ := servMock.dtoToMagazine(MockMagazines[0])
	assert.Equal(t, expected, magazine)
	assert.Nil(t, err)
}

func TestService_GetMagazineByCodeError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetMagazineByCodeError", true)
	_, err := servMock.GetMagazineByCode(ctx, "abc")
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_AddMagazineSuccess(t *testing.T) {
	magazine, err := servMock.AddMagazine(context.Background(), Magazine{})
	expected, _ := servMock.dtoToMagazine(MockMagazines[0])
	assert.Equal(t, expected, magazine)
	assert.Nil(t, err)
}

func TestService_AddMagazineError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "AddMagazineError", true)
	_, err := servMock.AddMagazine(ctx, Magazine{})
	assert.EqualError(t, err, ErrDefault.Error())
}

func TestService_UpdateMagazineSuccess(t *testing.T) {
	magazine, err := servMock.UpdateMagazine(context.Background(), "abc", Magazine{})
	expected, _ := servMock.dtoToMagazine(MockMagazines[0])
	assert.Equal(t, expected, magazine)
	assert.Nil(t, err)
}

func TestService_UpdateMagazineError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "UpdateMagazineError", true)
	_, err := servMock.UpdateMagazine(ctx, "abc", Magazine{})
	assert.EqualError(t, err, ErrDefault.Error())
}
