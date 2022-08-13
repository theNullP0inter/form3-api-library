package account

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/theNullP0inter/form3-api-library/src/common"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func (a *Account) List(pagination common.Pagination, filters map[string]string) ([]models.Account, error) {

	params := url.Values{}
	params.Add("page[number]", strconv.Itoa(pagination.Number))
	params.Add("page[size]", strconv.Itoa(pagination.Size))

	for k, v := range filters {
		params.Add(fmt.Sprintf("filter[%s]", k), v)
	}

	uri, _ := url.Parse(a.BasePath())
	uri.RawQuery = params.Encode()

	res, err := a.Client.MakeHTTPGetRequest(uri.String())

	if err != nil {
		return []models.Account{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []models.Account{}, errors.New(fmt.Sprintf("Failed to fetch with http code %d", res.StatusCode))
	}

	return AccountListFromResponse(res)
}
