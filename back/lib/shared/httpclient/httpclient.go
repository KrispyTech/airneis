package httpclient

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const methodGet string = "GET"

type HttpApi interface {
	Get(url string, headers map[string]string) ([]byte, int, error)
}

type HttpClient struct {
	api *http.Client
}

type Request struct {
	Body       []byte
	Headers    map[string]string
	Method     string
	StatusCode int
	Url        string
}

func InitializeHttpApi() HttpApi {
	return &HttpClient{api: &http.Client{}}
}

func (c *HttpClient) Get(url string, headers map[string]string) (body []byte, statusCode int, err error) {
	return c.makeRequest(Request{Url: url, Method: methodGet, Headers: headers})
}

func addHeaders(req *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req
}

func (c *HttpClient) makeRequest(r Request) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest(r.Method, r.Url, bytes.NewReader(r.Body))
	if err != nil {
		return nil, 0, err
	}

	if r.Headers != nil {
		req = addHeaders(req, r.Headers)
	}

	resp, err := c.api.Do(req)
	if err != nil {
		return nil, 0, err
	}

	return showResponse(resp)
}

func showResponse(resp *http.Response) ([]byte, int, error) {
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Error %v: %s", resp.StatusCode, resp.Request.URL)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
