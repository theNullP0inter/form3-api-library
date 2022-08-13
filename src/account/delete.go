package account

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (a *Account) Delete(id uuid.UUID, version int64) error {
	res, err := a.Client.MakeHTTPDeleteRequest(a.BasePath() + "/" + id.String() + "?version=" + strconv.Itoa(int(version)))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New(fmt.Sprintf("Failed to delete with http code %d", res.StatusCode))
	}

	return nil
}
