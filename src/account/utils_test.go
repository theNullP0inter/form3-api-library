package account

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func TestAccountFromResponse(t *testing.T) {

	resData := map[string]interface{}{}

	json.Unmarshal([]byte(MockAccountJson), &resData)

	res, err := httpmock.NewJsonResponse(200, resData)
	assert.NoError(t, err)

	acc, err := AccountFromResponse(res)
	assert.NoError(t, err)
	assert.IsType(t, models.Account{}, acc)

	expectedId, _ := uuid.Parse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Equal(t, expectedId, acc.ID)

}

func TestAccountListFromResponse(t *testing.T) {
	resData := map[string]interface{}{}

	json.Unmarshal([]byte(MockAccountListJson), &resData)

	res, err := httpmock.NewJsonResponse(200, resData)
	assert.NoError(t, err)

	accs, err := AccountListFromResponse(res)
	assert.NoError(t, err)
	assert.IsType(t, []models.Account{}, accs)

	assert.Equal(t, 2, len(accs))

	expectedId, _ := uuid.Parse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Equal(t, expectedId, accs[0].ID)

}
