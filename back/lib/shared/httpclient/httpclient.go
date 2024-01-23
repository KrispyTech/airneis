package httpclient

import (
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type HttpApi interface {
	Get(url string, headers map[string]string) ([]byte, int, error)
	// Post(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error)
	// Patch(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error)
	// Put(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error)
	// Delete(url string, headers map[string]string) ([]byte, int, error)
}

type HttpClient struct {
	api *http.Client
}

func CreateHTTPClient() *http.Client {
	return &http.Client{Timeout: time.Duration(60) * time.Second}
}

func InitializeHttpApi() HttpApi {
	return &HttpClient{api: CreateHTTPClient()}
}

func addHeaders(req *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req
}

func (c *HttpClient) Get(url string, headers map[string]string) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	if headers != nil {
		req = addHeaders(req, headers)
	}

	resp, err := c.api.Do(req)
	if err != nil {
		return nil, 0, err
	}

	statusCode = resp.StatusCode
	if statusCode != 200 {
		logrus.Errorf("Error %v: %s", statusCode, resp.Request.Host)
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return
}
