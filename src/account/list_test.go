package account

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

func TestAccountListWithWrongHttpStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("mock error"))
	}))

	service := &Account{
		Client: common.NewClient(&common.Config{Live: false}),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	_, err := service.List(common.NewPagination(1, 10), map[string]string{})
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrBadRequest, err)
}

func TestAccountListSuccess(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		assert.Equal(t, "1", req.URL.Query().Get("page[number]"))
		assert.Equal(t, "10", req.URL.Query().Get("page[size]"))
		assert.Equal(t, "bar", req.URL.Query().Get("filter[foo]"))

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(MockAccountListJson))
	}))

	service := &Account{
		Client: common.NewClient(&common.Config{Live: false}),
		BasePath: func() string {
			return mockServer.URL
		},
	}

	accs, err := service.List(common.NewPagination(1, 10), map[string]string{"foo": "bar"})
	assert.Nil(t, err)
	expectedId, _ := uuid.Parse(MockAccountJsonID)
	assert.Equal(t, expectedId, accs[0].ID)

}
