package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
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
		Client: common.NewClient(&common.Config{Live: false}),
		BasePath: func() string {
			return mockServer.URL
		},
		Validator: validator.New(),
	}

	accID, _ := uuid.Parse(MockAccountJsonID)
	_, err := service.Fetch(accID)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to fetch with http code 400", err.Error())
}

func TestFetchAccountSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(MockAccountJson))
	}))

	service := &Account{
		Client: common.NewClient(&common.Config{Live: false}),
		BasePath: func() string {
			return mockServer.URL
		},
		Validator: validator.New(),
	}

	accID, _ := uuid.Parse(MockAccountJsonID)
	acc, err := service.Fetch(accID)
	assert.Nil(t, err)
	assert.Equal(t, accID, acc.ID)

}
