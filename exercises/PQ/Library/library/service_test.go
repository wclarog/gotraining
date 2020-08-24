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

func TestService_GetMaterials(t *testing.T) {
	materials, err := servMock.GetMaterials(context.Background())
	mat, _ := servMock.dtoToMaterials(MockMaterials)
	assert.Equal(t, mat, materials)
	assert.Nil(t, err)
}

func TestService_GetMaterialsError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "GetMaterialsError", true)
	materials, err := servMock.GetMaterials(ctx)
	assert.EqualError(t, err, ErrDefault.Error())
	assert.Nil(t, materials)
}
