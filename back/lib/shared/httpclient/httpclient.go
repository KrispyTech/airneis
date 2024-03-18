package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const methodGet string = "GET"
const methodPost string = "POST"

type HttpApi interface {
	Get(url string, headers map[string]string) ([]byte, int, error)
	Post(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error)
}

type HttpClient struct {
	api *http.Client
}

type Request struct {
	Body       []byte
	Headers    map[string]string
	Method     string
	StatusCode int
	URL        string
}

func InitializeHTTPClient() HttpApi {
	return &HttpClient{api: &http.Client{}}
}

func PrepareBody(reqBody interface{}) (body []byte, err error) {
	body, err = json.Marshal(reqBody)
	if err != nil {
		return nil, errors.Errorf("PrepareBody, Unable to marshall body:  %s", err)
	}

	return body, nil
}

func (c *HttpClient) Get(url string, headers map[string]string) (body []byte, status int, err error) {
	return c.makeRequest(Request{URL: url, Method: methodGet, Headers: headers})
}

func (c *HttpClient) Post(url string, headers map[string]string, jsonBody []byte) (body []byte, status int, err error) {
	return c.makeRequest(Request{URL: url, Method: methodPost, Headers: headers, Body: jsonBody})
}

func addHeaders(req *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req
}

func (c *HttpClient) makeRequest(r Request) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.Body))
	if err != nil {
		return nil, 0, errors.Errorf("makeRequest,unable to make new request:  %s", err)
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
		return nil, 0, errors.Errorf("Error %v: %s", resp.StatusCode, resp.Request.URL)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Errorf("showResponse,unable to ReadAll body:  %s", err)
	}

	return body, resp.StatusCode, nil
}
