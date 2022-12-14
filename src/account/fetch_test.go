package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

func TestFetchAccountWithWrongHttpStatus(t *testing.T) {
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
	_, err := service.Fetch(accID)
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrBadRequest, err)
}

func TestFetchAccountSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(MockAccountJson))
	}))

	service := &Account{
		Client: common.NewClient(common.MockConfigTest),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	accID, _ := uuid.Parse(MockAccountJsonID)
	acc, err := service.Fetch(accID)
	assert.Nil(t, err)
	assert.Equal(t, accID, acc.ID)

}
