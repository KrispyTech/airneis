package httpclient

import (
	"errors"

	"github.com/KrispyTech/airneis/lib/shared/constants"
)

type MockResponse struct {
	Data   []byte
	Status int
}

type HTTPClientMock struct {
	ShouldFail bool
	Resp       []MockResponse
}

func FailedResponse() MockResponse {
	return MockResponse{
		Status: constants.BadRequestStatus,
		Data:   []byte(""),
	}
}

func OKResp(payload []byte) MockResponse {
	return MockResponse{
		Status: constants.BadRequestStatus,
		Data:   payload,
	}
}

func (c *HTTPClientMock) Get(url string, headers map[string]string) ([]byte, int, error) {
	if c.ShouldFail || url == "" || headers == nil {
		return FailedResponse().Data, FailedResponse().Status, errors.New("Failed to get resource")
	}

	return OKResp([]byte("")).Data, OKResp([]byte("")).Status, nil
}

func (c *HTTPClientMock) Post(url string, headers map[string]string, jsonBody []byte) ([]byte, int, error) {
	if c.ShouldFail || url == "" || headers == nil || jsonBody == nil {
		return FailedResponse().Data, FailedResponse().Status, errors.New("Failed to get resource")
	}

	return OKResp([]byte("")).Data, OKResp([]byte("")).Status, nil
}

func (c *HTTPClientMock) PrepareBody(reqBody interface{}) (body []byte, err error) {
	if c.ShouldFail || reqBody == nil {
		return FailedResponse().Data, errors.New("Failed to prepare body")
	}

	return OKResp([]byte("")).Data, nil
}
