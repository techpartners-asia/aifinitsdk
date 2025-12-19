package aifinitsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"
)

type MockClient struct {
	RestyClient *resty.Client
}

func (m *MockClient) GetSignature(timestamp int64) (string, error) {
	return "mock_signature", nil
}

func (m *MockClient) IsDebug() bool {
	return true
}

func (m *MockClient) RestyDebug() bool {
	return true
}

func (m *MockClient) SetConfig(config Config) {}

func (m *MockClient) GetRestyClient() *resty.Client {
	return m.RestyClient
}

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestControlAdStatus(t *testing.T) {
	// Test case 1: Success
	restyClient := resty.New()
	restyClient.SetTransport(RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PUT", req.Method)
		assert.Contains(t, req.URL.String(), "/facade/open/materials/updatePromotionStatus/123/2")
		assert.Equal(t, "mock_signature", req.Header.Get("Authorization"))

		response := AdControlStatusResponse{
			Status:  200,
			Message: "OK",
			Ok:      true,
		}
		respBytes, _ := json.Marshal(response)

		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
			Header:     header,
		}, nil
	}))

	mockClient := &MockClient{
		RestyClient: restyClient,
	}

	adsClient := &advertisementManageClientImpl{
		Client: mockClient,
		Resty:  restyClient,
	}

	resp, err := adsClient.ControlAdStatus(123, AdStatusApproved)
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, 200, resp.Status)
	}

	// Test case 2: Error 4446
	restyClient4446 := resty.New()
	restyClient4446.SetTransport(RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		response := AdControlStatusResponse{
			Status:  4446,
			Message: "Advertisements are not allowed to be removed from shelves",
			Ok:      false,
		}
		respBytes, _ := json.Marshal(response)

		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		return &http.Response{
			StatusCode: 200, // HTTP status 200 but result status 4446
			Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
			Header:     header,
		}, nil
	}))

	mockClient4446 := &MockClient{
		RestyClient: restyClient4446,
	}

	adsClient4446 := &advertisementManageClientImpl{
		Client: mockClient4446,
		Resty:  restyClient4446,
	}

	resp4446, err4446 := adsClient4446.ControlAdStatus(123, AdStatusApproved)
	assert.Error(t, err4446)
	assert.Nil(t, resp4446)
	assert.Contains(t, err4446.Error(), "AdRemoveNotAllowed: 4446")
}

func TestAdAssociatedToVm(t *testing.T) {
	// Test case: Success with ScanCodeList
	restyClient := resty.New()
	restyClient.SetTransport(RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PUT", req.Method)
		assert.Contains(t, req.URL.String(), "/facade/open/materials/bind/123")
		assert.Equal(t, "mock_signature", req.Header.Get("Authorization"))

		// Check request body
		var body AdAssociatedToVmRequest
		err := json.NewDecoder(req.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Contains(t, body.VmList, "vm1")
		assert.Contains(t, body.ScanCodeList, "scan1")

		response := AdAssociatedToVmResponse{
			Status:  200,
			Message: "OK",
		}
		respBytes, _ := json.Marshal(response)

		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
			Header:     header,
		}, nil
	}))

	mockClient := &MockClient{
		RestyClient: restyClient,
	}

	adsClient := &advertisementManageClientImpl{
		Client: mockClient,
		Resty:  restyClient,
	}

	req := &AdAssociatedToVmRequest{
		VmList:       []string{"vm1"},
		ScanCodeList: []string{"scan1"},
	}
	resp, err := adsClient.AdAssociatedToVm(123, req)
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, 200, resp.Status)
		assert.Equal(t, "OK", resp.Message)
	}
}
