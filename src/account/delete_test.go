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

func TestDeleteAccountWithWrongHttpStatus(t *testing.T) {
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
	err := service.Delete(accID)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to delete with http code 400", err.Error())
}

func TestDeleteAccountSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.WriteHeader(http.StatusNoContent)
		res.Write([]byte(nil))
	}))

	service := &Account{
		Client: common.NewClient(&common.Config{Live: false}),
		BasePath: func() string {
			return mockServer.URL
		},
		Validator: validator.New(),
	}

	accID, _ := uuid.Parse(MockAccountJsonID)
	err := service.Delete(accID)
	assert.Nil(t, err)

}
