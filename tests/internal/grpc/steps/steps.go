package steps

import (
	"context"
	"testing"

	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"

	"google.golang.org/grpc/status"
)

func DescribeDevice(ctx context.Context, t *testing.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, id uint64) (*act_device_api.DescribeDeviceV1Response, error) {
	t.Helper()
	getRequest := &act_device_api.DescribeDeviceV1Request{
		DeviceId: id,
	}
	getResponse, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return getResponse, err
}

func CreateDevice(ctx context.Context, t *testing.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, platform string, userId uint64) (*act_device_api.CreateDeviceV1Response, error) {
	t.Helper()
	createRequest := &act_device_api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userId,
	}
	createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
	t.Logf("status.Code: %v", status.Code(err).String())

	return createResponse, err
}

func ListDevices(ctx context.Context, t *testing.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, page uint64, perPage uint64) (*act_device_api.ListDevicesV1Response, error) {
	t.Helper()
	listDevicesRequest := &act_device_api.ListDevicesV1Request{
		Page:    page,
		PerPage: perPage,
	}
	listDevicesResponse, err := deviceApiClient.ListDevicesV1(ctx, listDevicesRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return listDevicesResponse, err
}

func RemoveDevice(ctx context.Context, t *testing.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, id uint64) (*act_device_api.RemoveDeviceV1Response, error) {
	t.Helper()
	removeRequest := &act_device_api.RemoveDeviceV1Request{
		DeviceId: id,
	}
	removeResponse, err := deviceApiClient.RemoveDeviceV1(ctx, removeRequest)
	t.Logf("status.Code: %v", status.Code(err).String())
	return removeResponse, err
}
