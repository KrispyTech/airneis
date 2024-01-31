package httpclient

import (
	"errors"
)

type MockResponse struct {
	Data   []byte
	Status int
}

type HttpClientMock struct {
	ShouldFail bool
	Resp       []MockResponse
}

func FailedResponse() MockResponse {
	return MockResponse{
		Status: 400,
		Data:   []byte(""),
	}
}

func OKResp(payload []byte) MockResponse {
	return MockResponse{
		Status: 400,
		Data:   payload,
	}
}

func (c *HttpClientMock) Get(url string, headers map[string]string) ([]byte, int, error) {
	if c.ShouldFail {
		return FailedResponse().Data, FailedResponse().Status, errors.New("Failed to get resource")
	}

	return OKResp([]byte("")).Data, OKResp([]byte("")).Status, nil
}

func (c *HttpClientMock) Post(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error) {
	if c.ShouldFail {
		return FailedResponse().Data, FailedResponse().Status, errors.New("Failed to get resource")
	}

	return OKResp([]byte("")).Data, OKResp([]byte("")).Status, nil
}

func (c *HttpClientMock) PrepareBody(reqBody interface{}) (body []byte, err error) {
	if c.ShouldFail {
		return FailedResponse().Data, errors.New("Failed to prepare body")
	}

	return OKResp([]byte("")).Data, nil
}
