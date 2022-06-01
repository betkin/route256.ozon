package expects

import (
	"github.com/ozonmp/act-device-api/tests/internal/models"
	"testing"

	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/assert"
)

func ExpectDeviceFields(t *testing.T, expectId uint64, expectData *act_device_api.CreateDeviceV1Request, actualResponse *act_device_api.DescribeDeviceV1Response) {
	assert.Equal(t, actualResponse.Value.Id, expectId)
	assert.Equal(t, actualResponse.Value.UserId, expectData.UserId)
	assert.Equal(t, actualResponse.Value.Platform, expectData.Platform)
}

func ExpectEventFields(t *testing.T, expect *models.DeviceEvent, actual *models.DeviceEvent) {
	assert.Equal(t, expect.DeviceId, actual.DeviceId)
	assert.Equal(t, expect.Type, actual.Type)
	if expect.Device != nil {
		assert.Equal(t, expect.Device.UserID, actual.Device.UserID)
		assert.Equal(t, expect.Device.Platform, actual.Device.Platform)
		if expect.Device.EnteredAt != nil {
			assert.Less(t, expect.Device.EnteredAt.UnixMilli()-actual.Device.EnteredAt.UnixMilli(), int64(20))
		}
	}
	assert.Less(t, expect.CreatedAt.UnixMilli()-actual.CreatedAt.UnixMilli(), int64(20))
	assert.Less(t, expect.UpdatedAt.UnixMilli()-actual.UpdatedAt.UnixMilli(), int64(20))
}
