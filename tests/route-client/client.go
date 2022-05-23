package route_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/ozonmp/act-device-api/internal/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
	CreateDevice(ctx context.Context, body CreateDeviceRequest) (CreateDeviceResponse, *http.Response, error)
	ListDevices(ctx context.Context, opts url.Values) (ListDevicesResponse, *http.Response, error)
	DescribeDevice(ctx context.Context, deviceId string) (DescribeDeviceResponse, *http.Response, error)
	RemoveDevice(ctx context.Context, deviceId string) (RemoveDeviceResponse, *http.Response, error)
	UpdateDevice(ctx context.Context, deviceId string, body UpdateDeviceRequest) (UpdateDeviceResponse, *http.Response, error)
}

type client struct {
	client   *retryablehttp.Client
	BasePath string
}

type CreateDeviceRequest struct {
	Platform string `json:"platform"`
	UserId   string `json:"userId"`
}

type CreateDeviceResponse struct {
	DeviceId int `json:"deviceId,string"`
}

type ListDevicesResponse struct {
	Items []struct {
		Item
	} `json:"items"`
}

type Item struct {
	ID        string     `json:"id"`
	Platform  string     `json:"platform"`
	UserID    string     `json:"userId"`
	EnteredAt *time.Time `json:"enteredAt"`
}

type DescribeDeviceResponse struct {
	Value Item `json:"value"`
}

type RemoveDeviceResponse struct {
	Found bool `json:"found"`
}

type UpdateDeviceRequest struct {
	Platform string `json:"platform"`
	UserId   string `json:"userId"`
}

type UpdateDeviceResponse struct {
	Success bool `json:"success"`
}

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

func (c *client) CreateDevice(ctx context.Context, body CreateDeviceRequest) (CreateDeviceResponse, *http.Response, error) {
	var (
		localResponse CreateDeviceResponse
	)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BasePath+"/api/v1/devices", b)
	if err != nil {
		return localResponse, nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
		return localResponse, res, err
	}
	data, _ := ioutil.ReadAll(res.Body)
	device := new(CreateDeviceResponse)
	err = json.Unmarshal(data, &device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) ListDevices(ctx context.Context, opts url.Values) (ListDevicesResponse, *http.Response, error) {
	var (
		localResponse ListDevicesResponse
	)

	apiUrl, err := url.Parse(c.BasePath + "/api/v1/devices")
	if err != nil {
		return localResponse, nil, err
	}

	query := apiUrl.Query()
	for k, v := range opts {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}
	apiUrl.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	devices := new(ListDevicesResponse)

	err = json.Unmarshal(data, &devices)
	if err != nil {
		return localResponse, res, err
	}
	return *devices, res, nil
}

func (c *client) DescribeDevice(ctx context.Context, deviceId string) (DescribeDeviceResponse, *http.Response, error) {
	var (
		localResponse DescribeDeviceResponse
	)

	apiUrlString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiUrlString = strings.Replace(apiUrlString, "{deviceId}", fmt.Sprintf("%v", deviceId), -1)
	apiUrl, err := url.Parse(apiUrlString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	device := new(DescribeDeviceResponse)

	err = json.Unmarshal(data, &device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) RemoveDevice(ctx context.Context, deviceId string) (RemoveDeviceResponse, *http.Response, error) {
	var (
		localResponse RemoveDeviceResponse
	)

	apiUrlString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiUrlString = strings.Replace(apiUrlString, "{deviceId}", fmt.Sprintf("%v", deviceId), -1)
	apiUrl, err := url.Parse(apiUrlString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, apiUrl.String(), nil)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	device := new(RemoveDeviceResponse)

	err = json.Unmarshal(data, &device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}

func (c *client) UpdateDevice(ctx context.Context, deviceId string, body UpdateDeviceRequest) (UpdateDeviceResponse, *http.Response, error) {
	var (
		localResponse UpdateDeviceResponse
	)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return localResponse, nil, err
	}

	apiUrlString := c.BasePath + "/api/v1/devices/{deviceId}"
	apiUrlString = strings.Replace(apiUrlString, "{deviceId}", fmt.Sprintf("%v", deviceId), -1)
	apiUrl, err := url.Parse(apiUrlString)
	if err != nil {
		return localResponse, nil, err
	}

	req, err := http.NewRequest(http.MethodPut, apiUrl.String(), b)
	if err != nil {
		return localResponse, nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return localResponse, res, err
	}

	if res.StatusCode != http.StatusOK {
		logger.ErrorKV(ctx, "Bad status code", res.StatusCode)
	}

	data, _ := ioutil.ReadAll(res.Body)
	device := new(UpdateDeviceResponse)

	err = json.Unmarshal(data, &device)
	if err != nil {
		return localResponse, res, err
	}
	return *device, res, nil
}
