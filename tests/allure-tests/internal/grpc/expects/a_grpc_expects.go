package expects

import (
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/assert"
)

func ExpectDeviceFields(t provider.T, expectId uint64, expectData *act_device_api.CreateDeviceV1Request, actualResponse *act_device_api.DescribeDeviceV1Response) {
	assert.Equal(t, actualResponse.Value.Id, expectId)
	assert.Equal(t, actualResponse.Value.UserId, expectData.UserId)
	assert.Equal(t, actualResponse.Value.Platform, expectData.Platform)
}
