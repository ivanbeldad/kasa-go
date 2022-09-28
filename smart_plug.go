package kasa

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

// get timers
// change wifi network
// factory reset 			{"system":{"reset":{"delay":1}}}

// SmartPlug Allow to interact with either the SmartPlug HS100 and HS110
type SmartPlug interface {
	GetAlias() string
	GetInfo() (DeviceInfo, error)
	TurnOn() error
	TurnOff() error
	SwitchOnOff() error
	Reboot() error
	ScanAPs() ([]AP, error)
}

// HS100 Allow to interact with the SmartPlug HS100
type HS100 SmartPlug

// HS103 Allow to interact with the SmartPlug HS103
type HS103 SmartPlug

// HS110 Allow to interact with the SmartPlug HS110
type HS110 SmartPlug

// HS200 Allow to interact with the SmartPlug HS200
type HS200 SmartPlug

type smartPlug struct {
	DeviceID string
	Alias    string
	Auth     auth
}

func (s smartPlug) GetAlias() string {
	return s.Alias
}

func (s smartPlug) GetEnergyUsage() (DeviceInfo, error) {
	var deviceInfo DeviceInfo
	res, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"system\": {\"get_sysinfo\": {}}}",
		},
	}).execute()
	if err != nil {
		return deviceInfo, err
	}
	res.ResponseData = strings.Replace(res.ResponseData, "{\"system\":{\"get_sysinfo\":", "", 1)
	res.ResponseData = strings.Replace(res.ResponseData, "}}", "", 1)
	err = json.Unmarshal([]byte(res.ResponseData), &deviceInfo)
	if err != nil {
		log.Fatal(err)
	}
	return deviceInfo, nil
}

func (s smartPlug) GetInfo() (DeviceInfo, error) {
	var deviceInfo DeviceInfo
	res, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"system\": {\"get_sysinfo\": {}}}",
		},
	}).execute()
	if err != nil {
		return deviceInfo, err
	}
	res.ResponseData = strings.Replace(res.ResponseData, "{\"system\":{\"get_sysinfo\":", "", 1)
	res.ResponseData = strings.Replace(res.ResponseData, "}}", "", 1)
	err = json.Unmarshal([]byte(res.ResponseData), &deviceInfo)
	if err != nil {
		log.Fatal(err)
	}
	return deviceInfo, nil
}

func (s smartPlug) TurnOn() error {
	_, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"system\": {\"set_relay_state\": {\"state\": 1}}}",
		},
	}).execute()
	return err
}

func (s smartPlug) TurnOff() error {
	_, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"system\": {\"set_relay_state\": {\"state\": 0}}}",
		},
	}).execute()
	return err
}

func (s smartPlug) SwitchOnOff() error {
	info, err := s.GetInfo()
	if err != nil {
		return err
	}
	if info.RelayState == 1 {
		return s.TurnOff()
	}
	return s.TurnOn()
}

func (s smartPlug) Reboot() error {
	_, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"system\":{\"reboot\":{\"delay\":1}}}",
		},
	}).execute()
	if err.Error() == "Request timeout" {
		err = nil
	}
	return err
}

func (s smartPlug) ScanAPs() ([]AP, error) {
	aps := make([]AP, 0)
	res, err := s.getAuthRequest(requestBody{
		Method: methodPassthrough,
		Params: params{
			DeviceID:    s.DeviceID,
			RequestData: "{\"netif\":{\"get_scaninfo\":{\"refresh\":1}}}",
		},
	}).execute()
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`.*ap_list":(\[.*\])`)
	data := re.FindStringSubmatch(res.ResponseData)
	if len(data) < 2 {
		return aps, nil
	}
	err = json.Unmarshal([]byte(data[1]), &aps)
	if err != nil {
		log.Fatal(err)
	}
	return aps, nil
}

func (s smartPlug) getAuthRequest(reqBody requestBody) authRequest {
	return authRequest{
		Auth: s.Auth,
		Request: request{
			URL:         cloudURL,
			RequestBody: reqBody,
		},
	}
}
