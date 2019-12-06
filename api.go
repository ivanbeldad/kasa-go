package kasa

import (
	"fmt"
)

const (
	cloudURL          = "https://wap.tplinkcloud.com"
	methodLogin       = "login"
	methodGetDevices  = "getDeviceList"
	methodPassthrough = "passthrough"
	appType           = "Kasa_Android"
)

// API Public interface to get information, devices and interact with them
type API interface {
	GetDevicesInfo() ([]DeviceInfo, error)
	GetHS100(alias string) (HS100, error)
	GetHS110(alias string) (HS110, error)
	GetHS200(alias string) (HS200, error)
}

type api struct {
	Auth        auth
	DevicesInfo []listedDeviceInfo
}

// Connect Create an authenticated API
func Connect(username, password string) (API, error) {
	a := api{
		Auth: auth{
			Username: username,
			Password: password,
			URL:      cloudURL,
		},
	}
	err := a.Auth.generateToken()
	if err != nil {
		return a, err
	}
	return a, nil
}

func (a api) GetHS100(alias string) (HS100, error) {
	var hs100 HS100
	devices, err := a.GetDevicesInfo()
	if err != nil {
		return hs100, err
	}
	for _, device := range devices {
		if device.Alias == alias {
			hs100 = smartPlug{Alias: alias, Auth: a.Auth, DeviceID: device.DeviceID}
			return hs100, nil
		}
	}
	return smartPlug{}, fmt.Errorf("there is no device with alias %s", alias)
}

func (a api) GetHS110(alias string) (HS110, error) {
	return a.GetHS100(alias)
}

func (a api) GetHS200(alias string) (HS200, error) {
	return a.GetHS100(alias)
}

func (a api) GetDevicesInfo() ([]DeviceInfo, error) {
	res, err := a.getAuthRequest(requestBody{Method: methodGetDevices}).execute()
	if err != nil {
		return nil, err
	}
	deviceInfoList := make([]DeviceInfo, 0)
	for _, device := range res.DeviceInfoList {
		var deviceInfo DeviceInfo
		deviceInfo.fromListedDeviceInfo(device)
		deviceInfoList = append(deviceInfoList, deviceInfo)
	}
	return deviceInfoList, nil
}

func (a api) getAuthRequest(reqBody requestBody) authRequest {
	return authRequest{
		Auth: a.Auth,
		Request: request{
			URL:         cloudURL,
			RequestBody: reqBody,
		},
	}
}
