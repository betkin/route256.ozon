package tests

import (
	"context"
	"fmt"
	"github.com/ozonmp/act-device-api/tests/allure-tests/config"
	test_config "github.com/ozonmp/act-device-api/tests/config"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/ozonmp/act-device-api/tests/internal/http/steps"
	route_client "github.com/ozonmp/act-device-api/tests/route-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpServerRemove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cfg, err := test_config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	client := route_client.NewHTTPClient(test_config.GetApiURL(cfg), 5, 1*time.Second)
	ctx := context.TODO()

	t.Run("Device removing returns true", func(t *testing.T) {
		//arrange
		testId, err := steps.CreateDevice(ctx, t, client, "Windows", "12")
		require.NoError(t, err, "CreateDevices error")
		//act
		responseBody, _, err := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		//assert
		require.NoError(t, err, "RemoveDevices error")
		assert.Equal(t, true, responseBody.Found, "RemoveDevice error")
	})

	t.Run("Nonexistent device removing returns false", func(t *testing.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		list, _, err := client.ListDevices(ctx, opts)
		require.NoError(t, err, "ListDevices error")
		nonId, err := strconv.Atoi(list.Items[0].ID)
		require.NoError(t, err, "ID to string error")
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, err := client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		//assert
		require.NoError(t, err, "RemoveDevices error")
		assert.Equal(t, false, responseBody.Found, "RemoveDevice error")
	})

	t.Run("Number of devices hasn't changed after a failed removal", func(t *testing.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		list, _, err := client.ListDevices(ctx, opts)
		require.NoError(t, err, "ListDevices error")
		beforeRemoval := len(list.Items)
		nonId, err := strconv.Atoi(list.Items[0].ID)
		require.NoError(t, err, "ID to string error")
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		//assert
		require.NoError(t, err, "RemoveDevices error")
		list, _, err = client.ListDevices(ctx, opts)
		require.NoError(t, err, "ListDevices error")
		afterRemoval := len(list.Items)
		assert.Equal(t, beforeRemoval, afterRemoval)
	})

	t.Run("Modified device can be removal", func(t *testing.T) {
		//arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "6",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Xubuntu", "7")
		require.NoError(t, err, "CreateDevices error")
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		require.NoError(t, err, "UpdateDevices error")
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		//assert
		require.NoError(t, err, "RemoveDevices error")
		assert.Equal(t, true, responseBody.Found)
	})

	t.Run("Double removal", func(t *testing.T) {
		//arrange
		testId, err := steps.CreateDevice(ctx, t, client, "Windows", "4567")
		require.NoError(t, err, "CreateDevices error")
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		require.NoError(t, err, "RemoveDevices error")
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		//assert
		require.NoError(t, err, "RemoveDevices error")
		assert.Equal(t, false, responseBody.Found)
	})
}

func TestHttpServerUpdate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}
	client := route_client.NewHTTPClient(config.GetApiURL(cfg), 5, 1*time.Second)
	ctx := context.Background()

	t.Run("Device updating returns true", func(t *testing.T) {
		//arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "MsDos",
			UserId:   "66",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Dos", "99")
		require.NoError(t, err, "CreateDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		//assert
		require.NoError(t, err, "UpdateDevices error")
		assert.Equal(t, true, responseBody.Success)
	})

	t.Run("Nonexistent device updating returns false", func(t *testing.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Windows Me",
			UserId:   "666",
		}
		list, _, err := client.ListDevices(ctx, opts)
		require.NoError(t, err, "ListDevices error")
		nonId, err := strconv.Atoi(list.Items[0].ID)
		if err != nil {
			t.Fatalf("Read ID error!")
		}
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", nonId), deviceUpdate)
		//assert
		require.NoError(t, err, "UpdateDevices error")
		assert.Equal(t, false, responseBody.Success)
	})

	t.Run("Date/time field hasn't changed after update", func(t *testing.T) {
		//arrange

		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "9600",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "RedHat", "6900")
		require.NoError(t, err, "CreateDevices error")
		testedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", testId))
		require.NoError(t, err, "DescribeDevices error")
		//action
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		require.NoError(t, err, "UpdateDevices error")
		//assert
		updatedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", testId))
		require.NoError(t, err, "DescribeDevices error")
		assert.Equal(t, testedBody.Value.EnteredAt, updatedBody.Value.EnteredAt)
	})

	t.Run("Double change", func(t *testing.T) {
		// arrange

		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "AltLinux",
			UserId:   "999",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "MacOS", "10")
		require.NoError(t, err, "CreateDevices error")
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		require.NoError(t, err, "UpdateDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		//assert
		require.NoError(t, err, "UpdateDevices error")
		assert.Equal(t, true, responseBody.Success)
	})

	t.Run("Removed device can't be modified", func(t *testing.T) {
		// arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Lubuntu",
			UserId:   "7707",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Mint", "101010")
		require.NoError(t, err, "CreateDevices error")
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		require.NoError(t, err, "RemoveDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		//assert
		require.NoError(t, err, "UpdateDevices error")
		assert.Equal(t, false, responseBody.Success)
	})
}
