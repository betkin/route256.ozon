package mocks_learning

import (
	"bou.ke/monkey"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/betkin/device-api/internal/api"
	"gitlab.ozon.dev/betkin/device-api/internal/model"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
)

func TestAPIDevice(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 06, 10, 11, 59, 59, 651387237, time.UTC)
	})
	t.Run("Unit testing: CreateDevice", func(t *testing.T) {
		ctx := context.TODO()
		span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateDeviceV1")
		defer span.Finish()
		ctrlDevice := gomock.NewController(t)
		mockDevice := NewMockRepo(ctrlDevice)
		now := time.Now()
		testDevice := &model.Device{
			Platform:  "BSD",
			UserID:    666,
			EnteredAt: &now,
		}
		callMockDevice := mockDevice.EXPECT().CreateDevice(ctx, testDevice).Return(uint64(13), nil).AnyTimes()
		ctrlEvent := gomock.NewController(t)
		mockEvent := NewMockEventRepo(ctrlEvent)
		testEvent := &model.DeviceEvent{
			DeviceID: 13,
			Type:     model.Created,
			Status:   model.Deferred,
			Device:   testDevice,
		}
		mockEvent.EXPECT().Add(ctx, testEvent).Return(nil).After(callMockDevice).AnyTimes()

		testApi := api.NewDeviceAPI(mockDevice, mockEvent)
		createRequest := &act_device_api.CreateDeviceV1Request{
			Platform: "BSD",
			UserId:   666,
		}
		createResponse, err := testApi.CreateDeviceV1(context.TODO(), createRequest)
		require.NoError(t, err)
		expectResponse := &act_device_api.CreateDeviceV1Response{DeviceId: 13}
		assert.Equal(t, expectResponse, createResponse)
	})

	t.Run("Unit testing: DescribeDevice", func(t *testing.T) {
		ctx := context.TODO()
		span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeDeviceV1")
		defer span.Finish()

		ctrlDevice := gomock.NewController(t)
		mockDevice := NewMockRepo(ctrlDevice)
		now := time.Now()
		testDevice := &model.Device{
			ID:        13,
			Platform:  "BSD",
			UserID:    666,
			EnteredAt: &now,
		}
		mockDevice.EXPECT().DescribeDevice(ctx, uint64(13)).Return(testDevice, nil).AnyTimes()

		ctrlEvent := gomock.NewController(t)
		mockEvent := NewMockEventRepo(ctrlEvent)
		testApi := api.NewDeviceAPI(mockDevice, mockEvent)

		describeRequest := &act_device_api.DescribeDeviceV1Request{DeviceId: 13}
		describeResponse, err := testApi.DescribeDeviceV1(context.TODO(), describeRequest)
		require.NoError(t, err)
		expectResponse := &act_device_api.DescribeDeviceV1Response{
			Value: &act_device_api.Device{
				Id:        13,
				Platform:  "BSD",
				UserId:    666,
				EnteredAt: timestamppb.Now(),
			},
		}
		assert.Equal(t, expectResponse, describeResponse)
	})

	t.Run("Unit testing: RemoveDevice", func(t *testing.T) {
		ctx := context.TODO()
		span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveDeviceV1")
		defer span.Finish()

		ctrlDevice := gomock.NewController(t)
		mockDevice := NewMockRepo(ctrlDevice)

		callMockDevice := mockDevice.EXPECT().RemoveDevice(ctx, uint64(13)).Return(true, nil).AnyTimes()

		ctrlEvent := gomock.NewController(t)
		mockEvent := NewMockEventRepo(ctrlEvent)
		testEvent := &model.DeviceEvent{
			DeviceID: 13,
			Type:     model.Removed,
			Status:   model.Deferred,
		}
		mockEvent.EXPECT().Add(ctx, testEvent).Return(nil).After(callMockDevice).AnyTimes()

		testApi := api.NewDeviceAPI(mockDevice, mockEvent)
		removeRequest := &act_device_api.RemoveDeviceV1Request{
			DeviceId: 13,
		}
		removeResponse, err := testApi.RemoveDeviceV1(context.TODO(), removeRequest)
		require.NoError(t, err)
		expectResponse := &act_device_api.RemoveDeviceV1Response{Found: true}
		assert.Equal(t, expectResponse, removeResponse)
	})
}
