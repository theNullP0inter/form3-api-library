package common

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cli := NewClient(MockConfigTest)
	assert.IsType(t, &Client{}, cli)
}

func TestCreateHttpRequest(t *testing.T) {

	reqBody := "request body"

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, CONTENT_TYPE_JSON, req.Header.Get("Content-Type"))

		reqBuf := new(strings.Builder)
		io.Copy(reqBuf, req.Body)
		assert.Equal(t, reqBody, reqBuf.String())

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))
	}))

	defer testServer.Close()

	cli := &Client{
		Config:             MockConfigTest,
		HttpClient:         &http.Client{},
		HttpRequestHandler: CreateHTTPRequest,
	}

	res, err := CreateHTTPRequest(cli, http.MethodPost, []byte(reqBody), testServer.URL)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}

func TestClientPost(t *testing.T) {

	reqBody := "request body"
	reqUrl := "http://example.com/"

	cli := &Client{
		Config:     MockConfigTest,
		HttpClient: &http.Client{},
		HttpRequestHandler: func(c *Client, httpMethod string, data []byte, url string) (*http.Response, error) {
			assert.Equal(t, http.MethodPost, httpMethod)
			assert.Equal(t, reqUrl, url)
			assert.Equal(t, []byte(reqBody), data)

			return nil, nil
		},
	}

	cli.MakeHTTPPostRequest([]byte(reqBody), reqUrl)
}

func TestClientGet(t *testing.T) {

	reqUrl := "http://example.com/"

	cli := &Client{
		Config:     MockConfigTest,
		HttpClient: &http.Client{},
		HttpRequestHandler: func(c *Client, httpMethod string, data []byte, url string) (*http.Response, error) {
			assert.Equal(t, http.MethodGet, httpMethod)
			assert.Equal(t, reqUrl, url)
			assert.Equal(t, []byte(nil), data)

			return nil, nil
		},
	}

	cli.MakeHTTPGetRequest(reqUrl)
}

func TestClientDelete(t *testing.T) {

	reqUrl := "http://example.com/"

	cli := &Client{
		Config:     MockConfigTest,
		HttpClient: &http.Client{},
		HttpRequestHandler: func(c *Client, httpMethod string, data []byte, url string) (*http.Response, error) {
			assert.Equal(t, http.MethodDelete, httpMethod)
			assert.Equal(t, reqUrl, url)
			assert.Equal(t, []byte(nil), data)

			return nil, nil
		},
	}

	cli.MakeHTTPDeleteRequest(reqUrl)
}
