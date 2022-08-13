package account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/theNullP0inter/form3-api-library/src/common"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func AccountFromResponse(res *http.Response) (models.Account, error) {

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return models.Account{}, err
	}

	var accountData struct {
		Data  models.Account `json:"data"`
		Links common.Links   `json:"links"`
	}

	err = json.Unmarshal(body, &accountData)

	if err != nil {
		return models.Account{}, err
	}

	return accountData.Data, nil

}

func AccountListFromResponse(res *http.Response) ([]models.Account, error) {

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []models.Account{}, err
	}

	var accountListData struct {
		Data  []models.Account `json:"data"`
		Links common.Links     `json:"links"`
	}

	err = json.Unmarshal(body, &accountListData)

	if err != nil {
		return []models.Account{}, err
	}

	return accountListData.Data, nil

}
