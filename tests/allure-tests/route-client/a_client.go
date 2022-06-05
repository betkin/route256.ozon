package route_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.ozon.dev/betkin/device-api/internal/pkg/logger"
)

// Client describes HTTP client methods
type Client interface {
	Do(req *http.Request) (*http.Response, error)
	CreateDevice(ctx context.Context, body CreateDeviceRequest) (CreateDeviceResponse, *http.Response, error)
	ListDevices(ctx context.Context, opts url.Values) (ListDevicesResponse, *http.Response, error)
	DescribeDevice(ctx context.Context, deviceID string) (DescribeDeviceResponse, *http.Response, error)
	RemoveDevice(ctx context.Context, deviceID string) (RemoveDeviceResponse, *http.Response, error)
	UpdateDevice(ctx context.Context, deviceID string, body UpdateDeviceRequest) (UpdateDeviceResponse, *http.Response, error)
}

// Body structure for client object
type client struct {
	client   *retryablehttp.Client
	BasePath string
}

// CreateDeviceRequest - HTTP request body structure
type CreateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// CreateDeviceResponse - HTTP response body structure
type CreateDeviceResponse struct {
	DeviceID int `json:"deviceId,string"`
}

// ListDevicesResponse - HTTP response body structure
type ListDevicesResponse struct {
	Items []struct {
		Item
	} `json:"items"`
}

// Item describes structure for DescribeDeviceResponse
type Item struct {
	ID        string     `json:"id"`
	Platform  string     `json:"platform"`
	UserID    string     `json:"userId"`
	EnteredAt *time.Time `json:"enteredAt"`
}

// DescribeDeviceResponse HTTP response body structure
type DescribeDeviceResponse struct {
	Value Item `json:"value"`
}

// RemoveDeviceResponse HTTP response body structure
type RemoveDeviceResponse struct {
	Found bool `json:"found"`
}

// UpdateDeviceRequest HTTP request body structure
type UpdateDeviceRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

// UpdateDeviceResponse HTTP response body structure
type UpdateDeviceResponse struct {
	Success bool `json:"success"`
}

// NewHTTPClient creates client object for HTTP connection
func NewHTTPClient(basePath string, retryMax int, timeout time.Duration) Client {
	c := &retryablehttp.Client{
		HTTPClient:      &http.Client{Timeout: timeout},
		RetryMax:        retryMax,
		RetryWaitMin:    1 * time.Second,
		RetryWaitMax:    10 * time.Second,
		CheckRetry:      retryablehttp.DefaultRetryPolicy,
		Backoff:         retryablehttp.DefaultBackoff,
		RequestLogHook:  requestHook,
		ResponseLogHook: responseHook,
	}

	client := &client{client: c, BasePath: basePath}

	return client
}

func requestHook(_ retryablehttp.Logger, req *http.Request, retry int) {
	logger.InfoKV(
		req.Context(),
		fmt.Sprintf("Retry request %d", retry),
		"request", req,
		"url", req.URL.String(),
	)
}

func responseHook(_ retryablehttp.Logger, res *http.Response) {
	logger.InfoKV(
		res.Request.Context(),
		"Responded",
		"response", res,
		"url", res.Request.URL.String(),
		"status_code", res.StatusCode,
	)
}

func (c *client) Do(request *http.Request) (*http.Response, error) {
	req, err := retryablehttp.FromRequest(request)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *client) makeRequest(ctx context.Context, deviceResponse interface{}, method string, urlString string, body interface{}) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)
	if body != nil {
		b := new(bytes.Buffer)
		err = json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, urlString, b)
	} else {
		req, err = http.NewRequest(method, urlString, nil)
	}
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
		return res, err
	}
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(data, &deviceResponse)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *client) CreateDevice(ctx context.Context, body CreateDeviceRequest) (CreateDeviceResponse, *http.Response, error) {
	createResponse := new(CreateDeviceResponse)
	urlString := c.BasePath + "/api/v1/devices"
	res, err := c.makeRequest(ctx, createResponse, http.MethodPost, urlString, body)
	if err != nil {
		return CreateDeviceResponse{}, res, err
	}
	return *createResponse, res, nil
}

func (c *client) ListDevices(ctx context.Context, opts url.Values) (ListDevicesResponse, *http.Response, error) {
	listDevicesResponse := new(ListDevicesResponse)
	apiURL, err := url.Parse(c.BasePath + "/api/v1/devices")
	if err != nil {
		return ListDevicesResponse{}, nil, err
	}
	query := apiURL.Query()
	for k, v := range opts {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	apiURL.RawQuery = query.Encode()
	res, err := c.makeRequest(ctx, listDevicesResponse, http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return ListDevicesResponse{}, res, err
	}
	return *listDevicesResponse, res, nil
}

func (c *client) DescribeDevice(ctx context.Context, deviceID string) (DescribeDeviceResponse, *http.Response, error) {
	describeResponse := new(DescribeDeviceResponse)
	urlString := c.BasePath + "/api/v1/devices/{deviceId}"
	urlString = strings.Replace(urlString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	res, err := c.makeRequest(ctx, describeResponse, http.MethodGet, urlString, nil)
	if err != nil {
		return DescribeDeviceResponse{}, res, err
	}
	return *describeResponse, res, nil
}

func (c *client) RemoveDevice(ctx context.Context, deviceID string) (RemoveDeviceResponse, *http.Response, error) {
	removeResponse := new(RemoveDeviceResponse)
	urlString := c.BasePath + "/api/v1/devices/{deviceId}"
	urlString = strings.Replace(urlString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	res, err := c.makeRequest(ctx, removeResponse, http.MethodDelete, urlString, nil)
	if err != nil {
		return RemoveDeviceResponse{}, res, err
	}
	return *removeResponse, res, nil
}

func (c *client) UpdateDevice(ctx context.Context, deviceID string, body UpdateDeviceRequest) (UpdateDeviceResponse, *http.Response, error) {
	updateResponse := new(UpdateDeviceResponse)
	urlString := c.BasePath + "/api/v1/devices/{deviceId}"
	urlString = strings.Replace(urlString, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	res, err := c.makeRequest(ctx, updateResponse, http.MethodPut, urlString, body)
	if err != nil {
		return UpdateDeviceResponse{}, res, err
	}
	return *updateResponse, res, nil
}
