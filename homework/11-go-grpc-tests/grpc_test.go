package tests

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDescribeDevice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	host := "localhost:8082"
	conn, _err := grpc.Dial(host, grpc.WithInsecure())
	if _err != nil {
		t.Fatalf("grpc.Dial err:%v", _err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	t.Run("Describe device returns correct ID", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "Windows",
			UserId:   9966,
		}
		createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		require.NoError(t, err)
		getRequest := &act_device_api.DescribeDeviceV1Request{
			DeviceId: createResponse.DeviceId,
		}
		//act
		getResponse, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Equal(t, createResponse.DeviceId, getResponse.Value.Id)
	})

	t.Run("Nonexistent ID return error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		listDevicesRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: math.MaxUint32 - 1,
		}
		listDevicesResponse, err := deviceApiClient.ListDevicesV1(ctx, listDevicesRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		require.NoError(t, err)
		require.NotNil(t, listDevicesResponse.Items)
		getRequest := &act_device_api.DescribeDeviceV1Request{
			DeviceId: listDevicesResponse.Items[0].Id + 1,
		}
		//act
		_, err = deviceApiClient.DescribeDeviceV1(ctx, getRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
	})

	t.Run("Zero ID value returns error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		getRequest := &act_device_api.DescribeDeviceV1Request{
			DeviceId: 0,
		}
		//act
		_, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
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
			{"tests uint32", 4294967295},
			{"tests int64", 9223372036854775807},
			{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
				getRequest := &act_device_api.DescribeDeviceV1Request{
					DeviceId: data.id,
				}
				//act
				_, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
				//assert
				t.Logf("status.Code: %v", status.Code(err).String())
				assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
			})
		}
	})
}

func TestListDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	host := "localhost:8082"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	t.Run("Items count equal PerPage", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		testRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: math.MaxUint32,
		}
		testResponse, err := deviceApiClient.ListDevicesV1(ctx, testRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		testCount := uint64(len(testResponse.Items)/2) + 1
		listRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: testCount,
		}
		//act
		listResponse, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Equal(t, uint64(len(listResponse.Items)), testCount)
	})

	t.Run("Items count not greater PerPage", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		testRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: math.MaxUint32,
		}
		testResponse, err := deviceApiClient.ListDevicesV1(ctx, testRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		testCount := uint64(len(testResponse.Items)/2) + 2
		listRequest := &act_device_api.ListDevicesV1Request{
			Page:    2,
			PerPage: testCount,
		}
		//act
		listResponse, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Less(t, uint64(len(listResponse.Items)), testCount)
	})

	t.Run("No items on page return error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		testRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: math.MaxUint32,
		}
		testResponse, err := deviceApiClient.ListDevicesV1(ctx, testRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		testCount := uint64(len(testResponse.Items)/2) + 1
		listRequest := &act_device_api.ListDevicesV1Request{
			Page:    3,
			PerPage: testCount,
		}
		//act
		listResponse, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
		assert.Nil(t, listResponse.Items)
	})

	t.Run("Zero PerPage returns error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		listRequest := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: 0,
		}
		//act
		_, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
	})

	t.Run("Zero Page returns OK", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		listRequest := &act_device_api.ListDevicesV1Request{
			Page:    0,
			PerPage: 1,
		}
		//act
		_, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
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
			{"tests uint32", 4294967295},
			{"tests int64", 9223372036854775807},
			{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
				listRequest := &act_device_api.ListDevicesV1Request{
					Page:    data.value,
					PerPage: data.value,
				}
				//act
				_, err := deviceApiClient.ListDevicesV1(ctx, listRequest)
				//assert
				t.Logf("status.Code: %v", status.Code(err).String())
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
			})
		}
	})
}

func TestCreateDevices(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	host := "localhost:8082"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)

	t.Run("Create Device returns ID", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "Debian",
			UserId:   1304,
		}
		//act
		createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		assert.Greater(t, createResponse.DeviceId, uint64(0))
	})

	t.Run("Create Device data equal request data", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "ChromeOS",
			UserId:   7896,
		}
		//act
		createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		getRequest := &act_device_api.DescribeDeviceV1Request{
			DeviceId: createResponse.DeviceId,
		}
		getResponse, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, createRequest.Platform, getResponse.Value.Platform)
		assert.Equal(t, createRequest.UserId, getResponse.Value.UserId)
	})

	t.Run("Creation date/time is correct", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "Vista",
			UserId:   666,
		}
		createTime := timestamppb.Now().AsTime().UnixMilli()
		//act
		createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		getRequest := &act_device_api.DescribeDeviceV1Request{
			DeviceId: createResponse.DeviceId,
		}
		getResponse, err := deviceApiClient.DescribeDeviceV1(ctx, getRequest)
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Less(t, getResponse.Value.EnteredAt.AsTime().UnixMilli()-createTime, int64(20))
	})

	t.Run("Zero UserID returns error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "ZeroOS",
			UserId:   0,
		}
		//act
		_, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
	})

	t.Run("Empty Platform returns error", func(t *testing.T) {
		//arrange
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "",
			UserId:   12345,
		}
		//act
		_, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
		//assert
		t.Logf("status.Code: %v", status.Code(err).String())
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
			{"tests uint64", 18446744073709551611},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
				createRequest := &act_device_api.CreateDeviceV1Request{
					Platform: "TestOS",
					UserId:   data.value,
				}
				//act
				createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
				//assert
				t.Logf("status.Code: %v", status.Code(err).String())
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
				assert.Greater(t, createResponse.DeviceId, uint64(0))
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
			{"tests char[64]", "SoLoooooooooooooooooooooooooooooooooooooooooooooooooooooooong OS"},
		}
		for _, data := range tests {
			t.Run(data.name, func(t *testing.T) {
				deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
				createRequest := &act_device_api.CreateDeviceV1Request{
					Platform: data.value,
					UserId:   54337,
				}
				//act
				createResponse, err := deviceApiClient.CreateDeviceV1(ctx, createRequest)
				//assert
				t.Logf("status.Code: %v", status.Code(err).String())
				assert.Equal(t, codes.OK.String(), status.Code(err).String())
				assert.Greater(t, createResponse.DeviceId, uint64(0))
			})
		}
	})

}
