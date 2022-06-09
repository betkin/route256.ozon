//go:build grpctest
// +build grpctest

package tests

import (
	"context"
	"math/rand"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
	test_config "gitlab.ozon.dev/betkin/device-api/test/grpctest/internal/config"
	"gitlab.ozon.dev/betkin/device-api/test/grpctest/internal/expects"
	"gitlab.ozon.dev/betkin/device-api/test/grpctest/internal/models"
	"gitlab.ozon.dev/betkin/device-api/test/grpctest/internal/steps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDescribeDevice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfg, err := test_config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	conn, err := grpc.Dial(test_config.GetGrpcURL(cfg), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	apiDB := steps.ConnectDB(t)
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			t.Logf("Postgres init err:%v", err)
		}
	}(apiDB.DB)

	t.Run("Describe device returns correct ID", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "Windows", 4563)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		require.NoError(t, err)
		//act
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		//assert
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Equal(t, createResponse.DeviceId, getResponse.Value.Id)
	})

	// I FOUND BUG !!!
	t.Run("Nonexistent ID return error", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		devicesCount, err := apiDB.GetCountDevices(ctx, false)
		require.NoError(t, err, "GetCountDevices error!")
		//act
		t.Logf("%v", devicesCount.Count)
		_, err = steps.DescribeDevice(ctx, t, deviceAPIClient, devicesCount.Count+1)
		//assert
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})

	t.Run("Zero ID value returns error", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.DescribeDevice(ctx, t, deviceAPIClient, 0)
		//assert
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("Device ID datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name string
			id   uint64
		}{
			{"tests int8", 99},
			{"tests uint8", 127},
			{"tests int16", 32767},
			{"tests uint16", 65535},
			{"tests int32", 2147483647},
			//{"tests uint32", 4294967295},
			//{"tests int64", 9223372036854775807},
			//{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.DescribeDevice(ctx, t, deviceAPIClient, data.id)
				//assert
				assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
			})
		}
	})
}

func TestListDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfg, err := test_config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	conn, err := grpc.Dial(test_config.GetGrpcURL(cfg), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	apiDB := steps.ConnectDB(t)
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			t.Logf("Postgres init err:%v", err)
		}
	}(apiDB.DB)

	t.Run("View pages tests", func(t *testing.T) {
		tests := []struct {
			name   string
			page   uint64
			expect string
		}{
			{"Items count equal PerPage", 1, codes.OK.String()},
			{"Items count not greater PerPage", 2, codes.OK.String()},
			{"No items on page return error", 3, codes.NotFound.String()}, // I FOUND BUG!!!
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				//arrange
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				allDevices, err := apiDB.GetCountDevices(ctx, true)
				require.NoError(t, err)
				testCount := uint64(allDevices.Count/2) + data.page
				//act
				listResponse, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, testCount)
				//arrange
				assert.Equal(t, data.expect, status.Code(err).String())
				assert.Equal(t, uint64(len(listResponse.Items)), testCount)
			})
		}
	})

	// I FOUND BUG !!!
	t.Run("Zero PerPage returns error", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, 0)
		//assert
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
	})

	// I FOUND BUG !!!
	t.Run("Zero Page returns OK", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.ListDevices(ctx, t, deviceAPIClient, 0, 1)
		//assert
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
	})

	t.Run("Page and PerPage datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value uint64
		}{
			{"tests int8", 99},
			{"tests uint8", 127},
			{"tests int16", 32767},
			{"tests uint16", 65535},
			{"tests int32", 2147483647},
			//{"tests uint32", 4294967295},
			//{"tests int64", 9223372036854775807},
			//{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.ListDevices(ctx, t, deviceAPIClient, data.value, data.value)
				//assert
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
			})
		}
	})
}

func TestCreateDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfgTest, _err := test_config.GetConfig()
	if _err != nil {
		t.Fatalf("Config err:%v", _err)
	}
	conn, _err := grpc.Dial(test_config.GetGrpcURL(cfgTest), grpc.WithInsecure())
	if _err != nil {
		t.Fatalf("grpc.Dial err:%v", _err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	apiDB := steps.ConnectDB(t)
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			t.Logf("Postgres init err:%v", err)
		}
	}(apiDB.DB)

	t.Run("Create Device returns ID", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "debian", 1304)
		//assert
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Greater(t, createResponse.DeviceId, uint64(0))
	})

	t.Run("Create Device data equal request data", func(t *testing.T) {
		//arrange
		testData := &act_device_api.CreateDeviceV1Request{
			Platform: "ChromeOS",
			UserId:   uint64(1304),
		}
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, testData.Platform, testData.UserId)
		//assert
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		expects.ExpectDeviceFields(t, createResponse.DeviceId, testData, getResponse)
	})

	t.Run("Creation date/time is correct", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createTime, err := apiDB.GetDBTime(ctx)
		require.NoError(t, err, "GetDB time error!")
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "Vista", 666)
		//assert
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Less(t, getResponse.Value.EnteredAt.AsTime().UnixMilli()-createTime.Time.UnixMilli(), int64(20))
	})

	t.Run("Zero UserID returns error", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "ZeroOS", 0)
		//assert
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("Empty Platform returns error", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "", 12345)
		//assert
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("UserID datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value uint64
		}{
			{"tests int8", 99},
			{"tests uint8", 127},
			{"tests int16", 32767},
			{"tests uint16", 65535},
			{"tests int32", 2147483647},
			{"tests uint32", 4294967295},
			{"tests int64", 9223372036854775807},
			//{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "TestOS", data.value)
				//assert
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
			})
		}
	})

	t.Run("Platform datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value string
		}{
			{"tests char[1]", "I"},
			{"tests char[2]", "OS"},
			{"tests char[16]", "Loooooooooong OS"},
			{"tests char[32]", "VeryLoooooooooooooooooooooong OS"},
			//{"tests char[64]", "SoLoooooooooooooooooooooooooooooooooooooooooooooooooooooooong OS"},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.CreateDevice(ctx, t, deviceAPIClient, data.value, 54337)
				//assert
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
			})
		}
	})
}

func TestLogDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfgTest, err := test_config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	conn, err := grpc.Dial(test_config.GetGrpcURL(cfgTest), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	apiDB := steps.ConnectDB(t)
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			t.Logf("Postgres init err:%v", err)
		}
	}(apiDB.DB)

	t.Run("CreateDevice was logged", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		timeNow, err := apiDB.GetDBTime(ctx)
		if err != nil {
			t.Fatalf("GetDBTime err:%v", err)
		}
		expectEvent := models.DeviceEvent{
			Type:   models.Created,
			Status: models.Processed,
			Device: &models.Device{
				Platform:  "MacOS",
				UserID:    8877,
				Removed:   false,
				EnteredAt: &timeNow.Time,
			},
			CreatedAt: timeNow.Time,
			UpdatedAt: timeNow.Time,
		}
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, expectEvent.Device.Platform, expectEvent.Device.UserID)
		//assert
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		require.Greater(t, createResponse.DeviceId, uint64(0))
		expectEvent.DeviceID = createResponse.DeviceId
		expectEvent.Device.ID = createResponse.DeviceId
		actualEvent, err := apiDB.ByDeviceID(ctx, expectEvent.DeviceID)
		require.NoError(t, err) // error = empty
		expects.ExpectEventFields(t, &expectEvent, actualEvent)
	})

	t.Run("RemoveDevice was logged", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		timeNow, err := apiDB.GetDBTime(ctx)
		if err != nil {
			t.Fatalf("GetDBTime err:%v", err)
		}
		expectEvent := models.DeviceEvent{
			Type:      models.Removed,
			Status:    models.Deferred,
			Device:    nil,
			CreatedAt: timeNow.Time,
			UpdatedAt: timeNow.Time,
		}
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "debian", 1304)
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		//act
		_, err = steps.RemoveDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		//assert
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		expectEvent.DeviceID = createResponse.DeviceId
		actualEvent, err := apiDB.ByDeviceID(ctx, expectEvent.DeviceID)
		require.NoError(t, err) // error = empty
		expects.ExpectEventFields(t, &expectEvent, actualEvent)
	})

	t.Run("UpdateDevice was logged", func(t *testing.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		timeNow, err := apiDB.GetDBTime(ctx)
		if err != nil {
			t.Fatalf("GetDBTime err:%v", err)
		}
		expectEvent := models.DeviceEvent{
			Type:   models.Updated,
			Status: models.Processed,
			Device: &models.Device{
				Platform:  "ubuntu",
				UserID:    3576,
				Removed:   false,
				EnteredAt: nil,
			},
		}
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "RedHat", 1234)
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		timeNow, err = apiDB.GetDBTime(ctx)
		if err != nil {
			t.Fatalf("GetDBTime err:%v", err)
		}
		expectEvent.CreatedAt = timeNow.Time
		expectEvent.UpdatedAt = timeNow.Time
		//act
		_, err = steps.UpdateDevice(ctx, t, deviceAPIClient, createResponse.DeviceId, expectEvent.Device.Platform, expectEvent.Device.UserID)
		//assert
		require.Equal(t, codes.OK.String(), status.Code(err).String())
		expectEvent.DeviceID = createResponse.DeviceId
		expectEvent.Device.ID = createResponse.DeviceId
		actualEvent, err := apiDB.ByDeviceID(ctx, expectEvent.DeviceID)
		require.NoError(t, err) // error = empty
		expects.ExpectEventFields(t, &expectEvent, actualEvent)
	})

}
