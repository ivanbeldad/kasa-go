package kasa

import (
	"encoding/json"
	"log"
)

// system info 				{"system":{"get_sysinfo":{}}}
// get token
// get devices
// switch on / off			{“system":{"set_relay_state":{"state":1}}}
// reboot					{"system":{"reboot":{"delay":1}}}
// get timers
// get ssids
// change wifi network
// factory reset 			{"system":{"reset":{"delay":1}}}

// HS100 Allow to interact with the SmartPlug HS100
type HS100 interface {
	GetAlias() string
	GetInfo() (DeviceInfo, error)
	TurnOn()
	TurnOff()
	SwitchOnOff()
	Reboot()
}

// HS110 Allow to interact with the SmartPlug HS110
type HS110 interface {
	HS100
}

type smartPlug struct {
	DeviceID string
	Alias    string
	Auth     auth
}

func (s smartPlug) GetAlias() string {
	return s.Alias
}

func (s smartPlug) GetInfo() (DeviceInfo, error) {
	var deviceInfo DeviceInfo
	res, err := authRequest{
		Auth: s.Auth,
		Request: request{
			URL: cloudURL,
			RequestBody: requestBody{
				Method: methodPassthrough,
				Params: params{
					DeviceID:    s.DeviceID,
					RequestData: "{\"system\": {\"get_sysinfo\": {}}}",
				},
			},
		},
	}.execute()
	if err != nil {
		return deviceInfo, err
	}
	var sys system
	err = json.Unmarshal([]byte(res.ResponseData), &sys)
	if err != nil {
		log.Fatal(err)
	}
	return deviceInfo, err
}

// TurnOn ...
func (s smartPlug) TurnOn() {

}

// TurnOff ...
func (s smartPlug) TurnOff() {

}

// SwitchOnOff ...
func (s smartPlug) SwitchOnOff() {

}

// Reboot ...
func (s smartPlug) Reboot() {

}
