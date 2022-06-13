package mocks_learning

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/betkin/device-api/internal/api"
	"gitlab.ozon.dev/betkin/device-api/internal/model"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
)

type stubRandom struct{}

func (r *stubRandom) Random(n int) int64 {
	return int64(15)
}

func (r *stubRandom) Doer() {}

type spyRandom struct{ target int }

func (r *spyRandom) Random(n int) int64 {
	r.target = n
	return int64(15)
}

func (r *spyRandom) Doer() {}

func TestGen(t *testing.T) {
	// Stub
	a1 := myRandom(5, 20, new(stubRandom))
	assert.Equal(t, int64(20), a1)
	// Spy
	spy := new(spyRandom)
	a2 := myRandom(5, 20, spy)
	assert.Equal(t, int64(20), a2)
	assert.Equal(t, 21, spy.target)
	// Mock
	ctrl := gomock.NewController(t)
	mockRandom := NewMockgenRand(ctrl)
	call := mockRandom.EXPECT().Doer().AnyTimes()
	mockRandom.EXPECT().Random(21).Return(int64(15)).After(call).AnyTimes()
	a3 := myRandom(5, 20, mockRandom)
	assert.Equal(t, int64(20), a3)
	// Real Mock - CreateDeviceV1 from api.go
	ctx := context.TODO()
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateDeviceV1")
	defer span.Finish()
	ctrlDevice := gomock.NewController(t)
	mockDevice := NewMockRepo(ctrlDevice)
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2022, 06, 10, 11, 59, 59, 651387237, time.UTC)
	})
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
}
