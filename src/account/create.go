package account

import (
	"encoding/json"
	"net/http"

	"github.com/theNullP0inter/form3-api-library/src/common"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func (a *Account) Create(acc models.Account) (models.Account, error) {

	req := struct {
		Data models.Account `json:"data"`
	}{
		Data: acc,
	}
	reqData, _ := json.Marshal(req)

	res, err := a.Client.MakeHTTPPostRequest(reqData, a.BasePath())

	if err != nil {
		return models.Account{}, err
	}

	if res.StatusCode != http.StatusCreated {
		return models.Account{}, common.ErrCreationFailed
	}

	return AccountFromResponse(res)

}
