package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

func TestCreateAccountWithWrongHttpStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("mock error"))
	}))

	service := &Account{
		Client: common.NewClient(common.MockConfigTest),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	_, err := service.Create(MockAccount)
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrBadRequest, err)
}

func TestCreateAccountSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(MockAccountJson))
	}))

	service := &Account{
		Client: common.NewClient(common.MockConfigTest),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	acc, err := service.Create(MockAccount)
	assert.Nil(t, err)
	expectedId, _ := uuid.Parse(MockAccountJsonID)
	assert.Equal(t, expectedId, acc.ID)

}
