package form3

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/theNullP0inter/form3-api-library/src/account"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

const (
	AccountsAPIVersion = "v1"
)

type Form3Client struct {
	client *common.Client

	// Services
	Account *account.Account
}

func NewForm3Client(cfg *common.Config) *Form3Client {

	c := &Form3Client{}
	c.client = &common.Client{
		Config:     cfg,
		HttpClient: &http.Client{},
	}

	c.Account = &account.Account{
		Client: c.client,
		BasePath: func() string {
			return fmt.Sprintf("%s/%s/organisation/accounts/", cfg.GetBaseUrl(), AccountsAPIVersion)
		},
		Validator: validator.New(),
	}

	return c
}
