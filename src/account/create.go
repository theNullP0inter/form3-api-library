package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/theNullP0inter/form3-api-library/src/models"
)

func (a *Account) Create(acc models.Account) (models.Account, error) {

	a.Validator.RegisterStructValidation(AccountValidation, models.Account{})
	err := a.Validator.Struct(acc)

	if err != nil {
		return models.Account{}, err
	}

	req := struct {
		Data models.Account `json:"data"`
	}{
		Data: acc,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return models.Account{}, err
	}

	res, err := a.Client.MakeHTTPPostRequest(reqData, a.BasePath())

	if err != nil {
		return models.Account{}, err
	}

	if res.StatusCode != http.StatusCreated {
		return models.Account{}, errors.New(fmt.Sprintf("Failed to create with http code %d", res.StatusCode))
	}

	return AccountFromResponse(res)

}
