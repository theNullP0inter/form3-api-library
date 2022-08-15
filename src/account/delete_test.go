package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

func TestDeleteAccountWithWrongHttpStatus(t *testing.T) {
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

	accID, _ := uuid.Parse(MockAccountJsonID)
	err := service.Delete(accID, 1)
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrBadRequest, err)
}

func TestDeleteAccountSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.WriteHeader(http.StatusNoContent)
		res.Write([]byte(nil))
	}))

	service := &Account{
		Client: common.NewClient(common.MockConfigTest),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	accID, _ := uuid.Parse(MockAccountJsonID)
	err := service.Delete(accID, 1)
	assert.Nil(t, err)

}
