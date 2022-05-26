package atests

import (
	"context"
	"fmt"
	"github.com/ozonmp/act-device-api/tests/allure-tests/config"
	"github.com/ozonmp/act-device-api/tests/allure-tests/internal/http/steps"
	route_client "github.com/ozonmp/act-device-api/tests/allure-tests/route-client"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"net/url"
	"strconv"

	"math/rand"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type CustomSuite struct {
	suite.Suite
}

func TestHttp(t *testing.T) {
	suite.RunSuite(t, new(CustomSuite))
}

func TestHttpServerRemove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("Config err:%v", err)
	}

	client := route_client.NewHTTPClient(config.GetApiURL(cfg), 5, 1*time.Second)
	ctx := context.TODO()

	runner.Run(t, "Device removing returns true", func(t provider.T) {
		testId, err := steps.CreateDevice(ctx, t, client, "Windows", "12")
		t.Require().NoError(err, "CreateDevices error")
		//act
		responseBody, _, err := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "RemoveDevices error")
			ctx.Assert().Equal(true, responseBody.Found, "RemoveDevice error")
		})
	})

	runner.Run(t, "Nonexistent device removing returns false", func(t provider.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		list, _, err := client.ListDevices(ctx, opts)
		t.Require().NoError(err, "ListDevices error")
		nonId, err := strconv.Atoi(list.Items[0].ID)
		t.Require().NoError(err, "ID to string error")
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, err := client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "RemoveDevices error")
			ctx.Assert().Equal(false, responseBody.Found, "RemoveDevice error")
		})
	})

	runner.Run(t, "Number of devices hasn't changed after a failed removal", func(t provider.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		list, _, err := client.ListDevices(ctx, opts)
		t.Require().NoError(err, "ListDevices error")
		beforeRemoval := len(list.Items)
		nonId, err := strconv.Atoi(list.Items[0].ID)
		t.Require().NoError(err, "ID to string error")
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", nonId))
		//assert
		t.Require().NoError(err, "RemoveDevices error")
		list, _, err = client.ListDevices(ctx, opts)
		t.Require().NoError(err, "ListDevices error")
		afterRemoval := len(list.Items)
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			ctx.Assert().Equal(beforeRemoval, afterRemoval)
		})
	})

	runner.Run(t, "Modified device can be removal", func(t provider.T) {
		//arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "6",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Xubuntu", "7")
		t.Require().NoError(err, "CreateDevices error")
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		t.Require().NoError(err, "UpdateDevices error")
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		//assert
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "RemoveDevices error")
			ctx.Assert().Equal(true, responseBody.Found, "RemoveDevice error")
		})
	})

	runner.Run(t, "Double removal", func(t provider.T) {
		//arrange
		testId, err := steps.CreateDevice(ctx, t, client, "Windows", "4567")
		t.Require().NoError(err, "CreateDevices error")
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		t.Require().NoError(err, "RemoveDevices error")
		//act
		responseBody, _, _ := client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		//assert
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "RemoveDevices error")
			ctx.Assert().Equal(false, responseBody.Found, "RemoveDevice error")
		})
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

	runner.Run(t, "Device updating returns true", func(t provider.T) {
		//arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "MsDos",
			UserId:   "66",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Dos", "99")
		t.Require().NoError(err, "CreateDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "UpdateDevices error")
			ctx.Assert().Equal(true, responseBody.Success)
		})
	})

	runner.Run(t, "Nonexistent device updating returns false", func(t provider.T) {
		//arrange
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "9223372036854775807")
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Windows Me",
			UserId:   "666",
		}
		list, _, err := client.ListDevices(ctx, opts)
		t.Require().NoError(err, "ListDevices error")
		nonId, err := strconv.Atoi(list.Items[0].ID)
		if err != nil {
			t.Fatalf("Read ID error!")
		}
		nonId += 1 // last device ID + 1 = nonexistent device ID
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", nonId), deviceUpdate)
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "UpdateDevices error")
			ctx.Assert().Equal(false, responseBody.Success)
		})
	})

	runner.Run(t, "Date/time field hasn't changed after update", func(t provider.T) {
		//arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "9600",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "RedHat", "6900")
		t.Require().NoError(err, "CreateDevices error")
		testedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", testId))
		t.Require().NoError(err, "DescribeDevices error")
		//action
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		t.Require().NoError(err, "UpdateDevices error")
		//assert
		updatedBody, _, err := client.DescribeDevice(ctx, fmt.Sprintf("%d", testId))
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "DescribeDevices error")
			ctx.Assert().Equal(testedBody.Value.EnteredAt, updatedBody.Value.EnteredAt)
		})
	})

	runner.Run(t, "Double change", func(t provider.T) {
		// arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "AltLinux",
			UserId:   "999",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "MacOS", "10")
		t.Require().NoError(err, "CreateDevices error")
		_, _, err = client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		t.Require().NoError(err, "UpdateDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		//assert
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "UpdateDevices error")
			ctx.Assert().Equal(true, responseBody.Success)
		})
	})

	runner.Run(t, "Removed device can't be modified", func(t provider.T) {
		// arrange
		deviceUpdate := route_client.UpdateDeviceRequest{
			Platform: "Lubuntu",
			UserId:   "7707",
		}
		testId, err := steps.CreateDevice(ctx, t, client, "Mint", "101010")
		t.Require().NoError(err, "CreateDevices error")
		_, _, err = client.RemoveDevice(ctx, fmt.Sprintf("%d", testId))
		t.Require().NoError(err, "RemoveDevices error")
		//act
		responseBody, _, _ := client.UpdateDevice(ctx, fmt.Sprintf("%d", testId), deviceUpdate)
		//assert
		t.WithNewStep("Test checks", func(ctx provider.StepCtx) {
			//assert
			ctx.Require().NoError(err, "UpdateDevices error")
			ctx.Assert().Equal(false, responseBody.Success)
		})
	})
}