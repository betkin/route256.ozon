package tests

import (
	"context"
	"fmt"

	route_client "gitlab.ozon.dev/betkin/device-api/homework/10-go-http-tests/tests/route-client"

	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHttpServerRemove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Run("Device removing returns true", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "Windows",
			UserId:   "11",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device creation error!")
		}
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		//assert
		t.Logf("Unexpected result - %t", responseBody.Found)
		assert.Equal(t, true, responseBody.Found)
	})

	t.Run("Nonexistent device removing returns false", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		ctx := context.TODO()
		list, _, err := client.ListDevices(ctx, opts)
		if err != nil {
			t.Errorf("Device list error!")
		}
		nonId, err := strconv.Atoi(list.Items[0].ID)
		if err != nil {
			t.Errorf("Read ID error!")
		}
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		//assert
		t.Logf("Unexpected result - %t", responseBody.Found)
		assert.Equal(t, false, responseBody.Found)
	})

	t.Run("Number of devices hasn't changed after a failed removal", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		ctx := context.TODO()
		list, _, err := client.ListDevices(ctx, opts)
		if err != nil {
			t.Errorf("Device list error!")
		}
		beforeRemoval := len(list.Items)
		nonId, err := strconv.Atoi(list.Items[0].ID)
		if err != nil {
			t.Errorf("Read ID error!")
		}
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		if err != nil {
			t.Errorf("Device removing error!")
		}
		//assert
		list, _, err = client.ListDevices(ctx, opts)
		if err != nil {
			t.Errorf("Device list error!")
		}
		afterRemoval := len(list.Items)
		t.Log("Unexpected action")
		assert.Equal(t, beforeRemoval, afterRemoval)
	})

	t.Run("Modified device can be removal", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "Xubuntu",
			UserId:   "7",
		}
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "6",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		if err != nil {
			t.Errorf("Device updete error!")
		}
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		//assert
		t.Logf("Unexpected result - %t", responseBody.Found)
		assert.Equal(t, true, responseBody.Found)
	})

	t.Run("Double removal", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "Android",
			UserId:   "222",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		if err != nil {
			t.Errorf("Device remove error!")
		}
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		//assert
		t.Logf("Unexpected result - %t", responseBody.Found)
		assert.Equal(t, false, responseBody.Found)
	})
}

func TestHttpServerUpdate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Run("Device updating returns true", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "Dos",
			UserId:   "99",
		}
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "MsDos",
			UserId:   "66",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		//assert
		t.Logf("Unexpected result - %t", responseBody.Success)
		assert.Equal(t, true, responseBody.Success)
	})

	t.Run("Nonexistent device updating returns false", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		ctx := context.TODO()
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Windows Me",
			UserId:   "666",
		}
		list, _, err := client.ListDevices(ctx, opts)
		if err != nil {
			t.Errorf("Device list error!")
		}
		nonId, err := strconv.Atoi(list.Items[0].ID)
		if err != nil {
			t.Errorf("Read ID error!")
		}
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", nonId), deviceUpdate)
		//assert
		t.Logf("Unexpected result - %t", responseBody.Success)
		assert.Equal(t, false, responseBody.Success)
	})

	t.Run("Date/time field hasn't changed after update", func(t *testing.T) {
		//arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "RedHat",
			UserId:   "6900",
		}
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "9600",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		testedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		if err != nil {
			t.Errorf("Get device error!")
		}
		//action
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		if err != nil {
			t.Errorf("Device update error!")
		}
		//assert
		updatedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		if err != nil {
			t.Errorf("Get device error!")
		}
		t.Logf("Unexpected result - %s", updatedBody.Value.EnteredAt)
		assert.Equal(t, testedBody.Value.EnteredAt, updatedBody.Value.EnteredAt)
	})

	t.Run("Double change", func(t *testing.T) {
		// arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "MacOS",
			UserId:   "10",
		}
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "AltLinux",
			UserId:   "999",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		if err != nil {
			t.Errorf("Device update error!")
		}
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		//assert
		t.Logf("Unexpected result - %t", responseBody.Success)
		assert.Equal(t, true, responseBody.Success)
	})

	t.Run("Removed device can't be modified", func(t *testing.T) {
		// arrange
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		ctx := context.TODO()
		device := route_client.CreateDeviceRequest{
			Platform: "Mint",
			UserId:   "101110",
		}
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Lubuntu",
			UserId:   "7707",
		}
		tested, _, err := client.CreateDevice(ctx, device)
		if err != nil {
			t.Errorf("Device create error!")
		}
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", tested.DeviceId))
		if err != nil {
			t.Errorf("Device remove error!")
		}
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", tested.DeviceId), deviceUpdate)
		//assert
		t.Logf("Unexpected result - %t", responseBody.Success)
		assert.Equal(t, false, responseBody.Success)
	})
}
