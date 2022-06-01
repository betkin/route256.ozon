package steps

import (
	"context"
	"fmt"
	"testing"

	"github.com/ozonmp/act-device-api/internal/config"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	route_sql "github.com/ozonmp/act-device-api/tests/route-sql"
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

func UpdateDevice(ctx context.Context, t *testing.T, deviceApiClient act_device_api.ActDeviceApiServiceClient, id uint64, platform string, userId uint64) (*act_device_api.UpdateDeviceV1Response, error) {
	t.Helper()
	updateRequest := &act_device_api.UpdateDeviceV1Request{
		DeviceId: id,
		Platform: platform,
		UserId:   userId,
	}
	updateResponse, err := deviceApiClient.UpdateDeviceV1(ctx, updateRequest)
	t.Logf("status.Code: %v", status.Code(err).String())

	return updateResponse, err
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

func ConnectDB(t *testing.T) *route_sql.Storage {
	if err := config.ReadConfigYML("../config.yml"); err != nil {
		t.Fatalf("Configuration fail! err:%v", err)
	}
	cfgApi := config.GetConfigInstance()
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", cfgApi.Database.Host, cfgApi.Database.Port, cfgApi.Database.User, cfgApi.Database.Password, cfgApi.Database.Name, cfgApi.Database.SslMode)
	apiDB, err := route_sql.NewPostgres(dsn, cfgApi.Database.Driver)
	if err != nil {
		t.Fatalf("Postgres init err:%v", err)
	}
	return apiDB
}
