package tests

import (
	"context"
	"fmt"
	route_client "github.com/ozonmp/act-device-api/tests/route-client"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"
)

/*
type ListOfItemsResponse struct {
	Items []struct {
		ID        string     `json:"id"`
		Platform  string     `json:"platform"`
		UserID    string     `json:"userId"`
		EnteredAt *time.Time `json:"enteredAt"`
	} `json:"items"`
}

type ItemRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

func Test_HttpServer_List(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Run("GET on list return 200", func(t *testing.T) {
		response, err := http.Get("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=1")
		if err != nil {
			panic(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Got %v, but want %v", response.StatusCode, http.StatusOK)
		}
	})

	t.Run("GET on list return devices list", func(t *testing.T) {
		countOfItems := 10
		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=%d", countOfItems))
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		list := new(ListOfItemsResponse)
		err = json.Unmarshal(data, &list)
		if err != nil {
			panic(err)
		}

		if len(list.Items) != countOfItems {
			t.Errorf("Want %d, get %d items", countOfItems, len(list.Items))
		}
	})

	t.Run("GET on list return devices list if zeroed", func(t *testing.T) {
		countOfItems := 0
		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=%d", countOfItems))
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		list := new(ListOfItemsResponse)
		err = json.Unmarshal(data, &list)
		if err != nil {
			panic(err)
		}

		if len(list.Items) != countOfItems {
			t.Errorf("Want %d, get %d items", countOfItems, len(list.Items))
		}
	})

	t.Run("POST on creating device", func(t *testing.T) {
		data := []byte(`{"platform": "Android", "userId": "123456"}`)
		r := bytes.NewReader(data)
		contentType := "application/json"

		_, err := http.Post("http://127.0.0.1:8080/api/v1/devices", contentType, r)
		if err != nil {
			panic(err)
		}

		payload := ItemRequest{Platform: "Android", UserID: "123456"}
		payloadJSON, _ := json.Marshal(payload)

		_, err = http.Post("http://127.0.0.1:8080/api/v1/devices", contentType, bytes.NewBuffer(payloadJSON))
		if err != nil {
			panic(err)
		}
	})

	t.Run("Why do we need a client?", func(t *testing.T) {
		t.Skip()
		// nc -lp 9090
		_, err := http.Get("http://127.0.0.1:9090")
		if err != nil {
			panic(err)
		}
	})

	t.Run("POST with client", func(t *testing.T) {
		// arrange
		payload := ItemRequest{Platform: "Android", UserID: "666"}
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(payload)
		if err != nil {
			panic(err)
		}
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		// action
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/v1/devices", b)
		if err != nil {
			panic(err)
		}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				t.Log(err)
			}
		}(res.Body)
		//assert

		if res.StatusCode != http.StatusOK {
			t.Errorf("Got %v, but want %v", res.StatusCode, http.StatusOK)
		}
		data, _ := ioutil.ReadAll(res.Body)
		if len(data) != 0 {
			t.Log(string(data))
		}
	})

	t.Run("Create device via client API", func(t *testing.T) {
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		device := route_client.CreateDeviceRequest{
			Platform: "Ubuntu",
			UserId:   "701",
		}
		ctx := context.TODO()
		id, _, _ := client.CreateDevice(ctx, device)
		t.Logf("New device is %d", id.DeviceId)
		assert.GreaterOrEqual(t, id.DeviceId, 0)
	})

	t.Run("List devices via client API", func(t *testing.T) {
		client := route_client.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "100")
		ctx := context.TODO()
		items, _, _ := client.ListDevices(ctx, opts)
		assert.GreaterOrEqual(t, len(items.Items), 1)
	})

}*/

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
