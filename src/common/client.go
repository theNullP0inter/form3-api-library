package common

import (
	"bytes"
	"net/http"
)

const (
	CONTENT_TYPE_JSON = "application/vnd.api+json"
)

func CreateHTTPRequest(c *Client, httpMethod string, data []byte, url string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)
	return c.HttpClient.Do(req)

}

type Client struct {
	Config     *Config
	HttpClient *http.Client
}

func (c *Client) MakeHTTPPostRequest(req []byte, url string) (*http.Response, error) {
	return CreateHTTPRequest(c, http.MethodPost, req, url)
}

func (c *Client) MakeHTTPGetRequest(url string) (*http.Response, error) {
	return CreateHTTPRequest(c, http.MethodGet, nil, url)
}

func (c *Client) MakeHTTPDeleteRequest(url string) (*http.Response, error) {
	return CreateHTTPRequest(c, http.MethodDelete, nil, url)
}
