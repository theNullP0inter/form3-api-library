package common

type Service struct {
	Client   *Client
	BasePath func() string
}
