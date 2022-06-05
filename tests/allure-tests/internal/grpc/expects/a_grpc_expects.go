package expects

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/assert"
	act_device_api "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
)

// The function compares the response fields with the expectation

func ExpectDeviceFields(t provider.T, expectId uint64, expectData *act_device_api.CreateDeviceV1Request, actualResponse *act_device_api.DescribeDeviceV1Response) {
	assert.Equal(t, actualResponse.Value.Id, expectId)
	assert.Equal(t, actualResponse.Value.UserId, expectData.UserId)
	assert.Equal(t, actualResponse.Value.Platform, expectData.Platform)
}
