package steps

import (
	"context"
	"github.com/ozonmp/act-device-api/tests/route-client"
	"testing"
)

func CreateDevice(ctx context.Context, t *testing.T, client route_client.Client, platform string, userId string) (uint64, error) {
	t.Helper()
	device := route_client.CreateDeviceRequest{
		Platform: platform,
		UserId:   userId,
	}
	tested, _, err := client.CreateDevice(ctx, device)

	return uint64(tested.DeviceId), err
}
