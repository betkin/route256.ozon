package atests

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
	"gitlab.ozon.dev/betkin/device-api/tests/allure-tests/config"
	"gitlab.ozon.dev/betkin/device-api/tests/allure-tests/internal/grpc/expects"
	"gitlab.ozon.dev/betkin/device-api/tests/allure-tests/internal/grpc/steps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDescribeDevice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	conn, err := grpc.Dial(config.GetGrpcURL(cfg), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	runner.Run(t, "Describe device returns correct ID", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "Windows", 4563)
		t.Require().Equal(codes.OK.String(), status.Code(err).String())
		t.Require().NoError(err)
		//act
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		//assert
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		t.Assert().Equal(createResponse.DeviceId, getResponse.Value.Id)
	})

	runner.Run(t, "Nonexistent ID return error", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		listItems, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, math.MaxUint32-1)
		t.Require().Equal(codes.OK.String(), status.Code(err).String())
		t.Require().NotNil(listItems.Items)
		//act
		_, err = steps.DescribeDevice(ctx, t, deviceAPIClient, listItems.Items[0].Id+1)
		//assert
		t.Assert().Equal(codes.NotFound.String(), status.Code(err).String())
	})

	runner.Run(t, "Zero ID value returns error", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.DescribeDevice(ctx, t, deviceAPIClient, 0)
		//assert
		t.Assert().Equal(codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("Device ID datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name string
			id   uint64
		}{
			{"test int8", 99},
			{"test uint8", 127},
			{"test int16", 32767},
			{"test uint16", 65535},
			{"test int32", 2147483647},
			{"test uint32", 4294967295},
			{"test int64", 9223372036854775807},
			{"test uint64", 18446744073709551611},
		}
		for _, data := range tests {
			runner.Run(t, data.name, func(t provider.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.DescribeDevice(ctx, t, deviceAPIClient, data.id)
				//assert
				t.Assert().Equal(codes.NotFound.String(), status.Code(err).String())
			})
		}
	})
}

func TestListDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	conn, err := grpc.Dial(config.GetGrpcURL(cfg), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	t.Run("View pages tests", func(t *testing.T) {
		tests := []struct {
			name   string
			page   uint64
			expect string
		}{
			{"Items count equal PerPage", 1, codes.OK.String()},
			{"Items count not greater PerPage", 2, codes.OK.String()},
			{"No items on page return error", 3, codes.NotFound.String()},
		}
		for _, data := range tests {
			runner.Run(t, data.name, func(t provider.T) {
				//arrange
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				allDevices, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, math.MaxUint32-1)
				t.Require().Equal(codes.OK.String(), status.Code(err).String())
				t.Require().NotNil(allDevices.Items)
				testCount := uint64(len(allDevices.Items)/2) + data.page
				//act
				listResponse, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, testCount)
				//arrange
				t.Assert().Equal(data.expect, status.Code(err).String())
				t.Assert().Equal(uint64(len(listResponse.Items)), testCount)
			})
		}
	})

	runner.Run(t, "Zero PerPage returns error", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.ListDevices(ctx, t, deviceAPIClient, 1, 0)
		//assert
		t.Assert().Equal(codes.Internal.String(), status.Code(err).String())
	})

	runner.Run(t, "Zero Page returns OK", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.ListDevices(ctx, t, deviceAPIClient, 0, 1)
		//assert
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
	})

	t.Run("Page and PerPage datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value uint64
		}{
			{"test int8", 99},
			{"test uint8", 127},
			{"test int16", 32767},
			{"test uint16", 65535},
			{"test int32", 2147483647},
			{"test uint32", 4294967295},
			{"test int64", 9223372036854775807},
			{"test uint64", 18446744073709551611},
		}
		for _, data := range tests {
			runner.Run(t, data.name, func(t provider.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.ListDevices(ctx, t, deviceAPIClient, data.value, data.value)
				//assert
				t.Assert().Equal(codes.OK.String(), status.Code(err).String())
			})
		}
	})
}

func TestCreateDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cfg, _err := config.GetConfig()
	if _err != nil {
		t.Fatalf("Config err:%v", _err)
	}
	conn, _err := grpc.Dial(config.GetGrpcURL(cfg), grpc.WithInsecure())
	if _err != nil {
		t.Fatalf("grpc.Dial err:%v", _err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	runner.Run(t, "Create Device returns ID", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "debian", 1304)
		//assert
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		t.Assert().Greater(createResponse.DeviceId, uint64(0))
	})

	runner.Run(t, "Create Device data equal request data", func(t provider.T) {
		//arrange
		testData := &act_device_api.CreateDeviceV1Request{
			Platform: "ChromeOS",
			UserId:   uint64(1304),
		}
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, testData.Platform, testData.UserId)
		//assert
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		expects.ExpectDeviceFields(t, createResponse.DeviceId, testData, getResponse)
	})

	runner.Run(t, "Creation date/time is correct", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createTime := timestamppb.Now().AsTime().UnixMilli()
		//act
		createResponse, err := steps.CreateDevice(ctx, t, deviceAPIClient, "Vista", 666)
		//assert
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		getResponse, err := steps.DescribeDevice(ctx, t, deviceAPIClient, createResponse.DeviceId)
		t.Assert().Equal(codes.OK.String(), status.Code(err).String())
		t.Assert().Less(getResponse.Value.EnteredAt.AsTime().UnixMilli()-createTime, int64(20))
	})

	runner.Run(t, "Zero UserID returns error", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "ZeroOS", 0)
		//assert
		t.Assert().Equal(codes.InvalidArgument.String(), status.Code(err).String())
	})

	runner.Run(t, "Empty Platform returns error", func(t provider.T) {
		//arrange
		deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
		//act
		_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "", 12345)
		//assert
		t.Assert().Equal(codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("UserID datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value uint64
		}{
			{"test int8", 99},
			{"test uint8", 127},
			{"test int16", 32767},
			{"test uint16", 65535},
			{"test int32", 2147483647},
			{"test uint32", 4294967295},
			{"test int64", 9223372036854775807},
			{"test uint64", 18446744073709551611},
		}
		for _, data := range tests {
			runner.Run(t, data.name, func(t provider.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.CreateDevice(ctx, t, deviceAPIClient, "TestOS", data.value)
				//assert
				t.Assert().Equal(codes.OK.String(), status.Code(err).String())
			})
		}
	})

	t.Run("Platform datatype testing", func(t *testing.T) {
		//arrange
		tests := []struct {
			name  string
			value string
		}{
			{"test char[1]", "I"},
			{"test char[2]", "OS"},
			{"test char[16]", "Loooooooooong OS"},
			{"test char[32]", "VeryLoooooooooooooooooooooong OS"},
			{"test char[64]", "SoLoooooooooooooooooooooooooooooooooooooooooooooooooooooooong OS"},
		}
		for _, data := range tests {
			runner.Run(t, data.name, func(t provider.T) {
				data := data
				t.Parallel()
				deviceAPIClient := act_device_api.NewActDeviceApiServiceClient(conn)
				//act
				_, err := steps.CreateDevice(ctx, t, deviceAPIClient, data.value, 54337)
				//assert
				t.Assert().Equal(codes.OK.String(), status.Code(err).String())
			})
		}
	})

}
