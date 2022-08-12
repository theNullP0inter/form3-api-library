package account

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (a *Account) Delete(id uuid.UUID) error {
	res, err := a.Client.MakeHTTPDeleteRequest(a.BasePath() + "/" + id.String())

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New(fmt.Sprintf("Failed to delete with http code %d", res.StatusCode))
	}

	return nil
}
