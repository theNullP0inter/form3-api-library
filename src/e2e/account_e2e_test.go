package e2e

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
	"github.com/theNullP0inter/form3-api-library/src/form3"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

var MockAccount = models.Account{
	Type: "accounts",
	Attributes: &models.AccountAttributes{
		Country:       "GB",
		BankIDCode:    "GBDSC",
		Bic:           "NWBKGB22",
		BankID:        "123456",
		AccountNumber: "12345678",
		Name:          []string{"Jon", "Snow"},
	},
}

func TestAccountE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Skipping E2E tests")
	}

	form3_cli := form3.NewForm3Client(&common.Config{
		Live: false,
	})

	acc, err := form3_cli.Account.Create(MockAccount)

	assert.Nil(t, err)
	assert.Equal(t, acc.Attributes.Bic, "NWBKGB22")

	accs, err := form3_cli.Account.List(common.NewPagination(0, 10), map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(accs))
	assert.Equal(t, accs[0].Attributes.Bic, "NWBKGB22")

	accID := accs[0].ID

	acc, err = form3_cli.Account.Fetch(accID)
	assert.Nil(t, err)
	assert.Equal(t, accID, acc.ID)

	err = form3_cli.Account.Delete(accID, acc.Version)
	assert.Nil(t, err)

}
