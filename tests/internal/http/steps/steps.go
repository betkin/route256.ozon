package steps

import (
	"context"
	"testing"

	route_client "github.com/ozonmp/act-device-api/tests/route-client"
)

// These function contain the actions for the HTTP test
// Get create device response

func CreateDevice(ctx context.Context, t *testing.T, client route_client.Client, platform string, userID string) (uint64, error) {
	t.Helper()
	device := route_client.CreateDeviceRequest{
		Platform: platform,
		UserID:   userID,
	}
	tested, _, err := client.CreateDevice(ctx, device)

	return uint64(tested.DeviceID), err
}
