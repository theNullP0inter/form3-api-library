package account

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func (a *Account) Fetch(id uuid.UUID) (models.Account, error) {
	res, err := a.Client.MakeHTTPGetRequest(a.BasePath() + "/" + id.String())

	if err != nil {
		return models.Account{}, err
	}

	if res.StatusCode != http.StatusOK {
		return models.Account{}, errors.New(fmt.Sprintf("Failed to fetch with http code %d", res.StatusCode))
	}

	return AccountFromResponse(res)
}
