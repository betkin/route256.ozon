package steps

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
	"google.golang.org/grpc/status"
)

//These functions contain the actions for the gRPS tests

// DescribeDevice gives describe device response and runtime error
func DescribeDevice(ctx context.Context, t provider.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, id uint64) (*act_device_api.DescribeDeviceV1Response, error) {
	t.Helper()
	getRequest := &act_device_api.DescribeDeviceV1Request{
		DeviceId: id,
	}
	getResponse, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return getResponse, err
}

// CreateDevice gives create device response and runtime error
func CreateDevice(ctx context.Context, t provider.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, platform string, userId uint64) (*act_device_api.CreateDeviceV1Response, error) {
	t.Helper()
	createRequest := &act_device_api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userId,
	}
	createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
	t.Logf("status.Code: %v", status.Code(err).String())

	return createResponse, err
}

// ListDevices gives response for list of devices and runtime error
func ListDevices(ctx context.Context, t provider.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, page uint64, perPage uint64) (*act_device_api.ListDevicesV1Response, error) {
	t.Helper()
	listDevicesRequest := &act_device_api.ListDevicesV1Request{
		Page:    page,
		PerPage: perPage,
	}
	listDevicesResponse, err := deviceApiClient.ListDevicesV1(ctx, listDevicesRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return listDevicesResponse, err
}

// RemoveDevice gives remove device response and runtime error
func RemoveDevice(ctx context.Context, t provider.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, id uint64) (*act_device_api.RemoveDeviceV1Response, error) {
	t.Helper()
	removeRequest := &act_device_api.RemoveDeviceV1Request{
		DeviceId: id,
	}
	removeResponse, err := deviceApiClient.RemoveDeviceV1(ctx, removeRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return removeResponse, err
}
