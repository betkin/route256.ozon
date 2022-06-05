package steps

import (
	"context"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	route_client "gitlab.ozon.dev/betkin/device-api/tests/allure-tests/route-client"
)

// These function contain the actions for the HTTP test
// Get create device response

func CreateDevice(ctx context.Context, t provider.T, client route_client.Client, platform string, userId string) (uint64, error) {
	t.Helper()
	device := route_client.CreateDeviceRequest{
		Platform: platform,
		UserID:   userId,
	}
	tested, _, err := client.CreateDevice(ctx, device)

	return uint64(tested.DeviceID), err
}
