package aifinitsdk

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"
)

func TestCustomBaseURL(t *testing.T) {
	customBaseURL := "https://custom.api.endpoint"
	credentials := Crendetials{
		MerchantCode: "test_merchant",
		SecretKey:    "test_secret",
	}

	restyClient := resty.New()
	restyClient.SetTransport(RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		assert.Contains(t, req.URL.String(), customBaseURL)
		assert.Contains(t, req.URL.String(), Get_VendingMachineList)

		// Return a dummy success response
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       http.NoBody,
		}, nil
	}))

	client := New(credentials, restyClient, customBaseURL)
	deviceClient := NewDeviceClient(client)

	_, _ = deviceClient.List(&ListMachineRequest{Page: 1, Limit: 10})
}

func TestDefaultBaseURL(t *testing.T) {
	credentials := Crendetials{
		MerchantCode: "test_merchant",
		SecretKey:    "test_secret",
	}

	restyClient := resty.New()
	restyClient.SetTransport(RoundTripFunc(func(req *http.Request) (*http.Response, error) {
		assert.Contains(t, req.URL.String(), DefaultBaseURL)
		assert.Contains(t, req.URL.String(), Get_VendingMachineList)

		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       http.NoBody,
		}, nil
	}))

	// Pass empty string for baseUrl to use default
	client := New(credentials, restyClient, "")
	deviceClient := NewDeviceClient(client)

	_, _ = deviceClient.List(&ListMachineRequest{Page: 1, Limit: 10})
}
