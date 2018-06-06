package kasa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	codeNoError       = 0
	expiredTokenError = "Token expired"
)

type reponseBody struct {
	ErrorCode    int    `json:"error_code"`
	Result       result `json:"result"`
	ErrorMessage string `json:"msg"`
}

type result struct {
	ResponseData   string             `json:"responseData,omitempty"`
	DeviceInfoList []listedDeviceInfo `json:"deviceList,omitempty"`
	AccountID      string             `json:"accountID,omitempty"`
	RegTime        string             `json:"regTime,omitempty"`
	Email          string             `json:"email,omitempty"`
	Token          string             `json:"token,omitempty"`
}

type system struct {
	GetSysInfo getSysInfo `json:"system,omitempty"`
}

type getSysInfo struct {
	DeviceInfo listedDeviceInfo `json:"get_sysinfo,omitempty"`
}

type listedDeviceInfo struct {
	FwVer        string `json:"fwVer,omitempty"`
	DeviceName   string `json:"deviceName,omitempty"`
	Status       int    `json:"status,omitempty"`
	Alias        string `json:"alias,omitempty"`
	DeviceType   string `json:"deviceType,omitempty"`
	AppServerURL string `json:"appServerUrl,omitempty"`
	DeviceModel  string `json:"deviceModel,omitempty"`
	DeviceMac    string `json:"deviceMac,omitempty"`
	Role         int    `json:"role,omitempty"`
	IsSameRegion bool   `json:"isSameRegion,omitempty"`
	HwID         string `json:"hwId,omitempty"`
	FwID         string `json:"fwId,omitempty"`
	OemID        string `json:"oemId,omitempty"`
	DeviceID     string `json:"deviceId,omitempty"`
	DeviceHwVer  string `json:"deviceHwVer,omitempty"`
}

func (l listedDeviceInfo) toDeviceInfo() DeviceInfo {
	return DeviceInfo{
		FwID:       l.FwVer,
		DeviceName: l.DeviceName,
		Alias:      l.Alias,
		Type:       l.DeviceType,
		Model:      l.DeviceModel,
		Mac:        l.DeviceMac,
		HwID:       l.HwID,
		OemID:      l.OemID,
		DeviceID:   l.DeviceID,
		HwVer:      l.DeviceHwVer,
	}
}

// DeviceInfo Get all the information about a device
type DeviceInfo struct {
	SwVer      string  `json:"sw_ver,omitempty"`
	HwVer      string  `json:"hw_ver,omitempty"`
	Type       string  `json:"type,omitempty"`
	Model      string  `json:"model,omitempty"`
	Mac        string  `json:"mac,omitempty"`
	DeviceID   string  `json:"deviceId,omitempty"`
	HwID       string  `json:"hwId,omitempty"`
	FwID       string  `json:"fwId,omitempty"`
	OemID      string  `json:"oemId,omitempty"`
	Alias      string  `json:"alias,omitempty"`
	DeviceName string  `json:"dev_name,omitempty"`
	IconHash   string  `json:"icon_hash,omitempty"`
	RelayState int     `json:"relay_state,omitempty"`
	OnTime     int     `json:"on_time,omitempty"`
	ActiveMode string  `json:"active_mode,omitempty"`
	Feature    string  `json:"feature,omitempty"`
	Updating   int     `json:"updating,omitempty"`
	Rssi       int     `json:"rssi,omitempty"`
	LedOff     int     `json:"led_off,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

type request struct {
	URL         string
	RequestBody requestBody
}

type authRequest struct {
	Request request
	Auth    auth
}

type requestBody struct {
	Method string `json:"method,omitempty"`
	Params params `json:"params,omitempty"`
}

type params struct {
	AppType       string `json:"appType,omitempty"`
	CloudUserName string `json:"cloudUserName,omitempty"`
	CloudPassword string `json:"cloudPassword,omitempty"`
	TerminalUUID  string `json:"terminalUUID,omitempty"`
	DeviceID      string `json:"deviceId,omitempty"`
	RequestData   string `json:"requestData,omitempty"`
}

func (r request) execute() (result, error) {
	res := result{}
	reqBodyJSON, err := json.MarshalIndent(&r.RequestBody, "", "  ")
	if err != nil {
		return res, err
	}
	response, err := http.Post(r.URL, "application/json", bytes.NewReader(reqBodyJSON))
	if err != nil {
		return res, err
	}
	return parseResponse(response)
}

func (r authRequest) execute() (result, error) {
	var res result
	var err error
	if r.Auth.Token == "" {
		err = r.Auth.generateToken()
		if err != nil {
			return res, err
		}
	}
	r.Request.URL = r.Request.URL + "?token=" + r.Auth.Token
	res, err = r.Request.execute()
	if err != nil {
		if err.Error() == expiredTokenError {
			err = r.Auth.generateToken()
			if err != nil {
				return res, err
			}
			r.Request.URL = r.Request.URL + "?token=" + r.Auth.Token
			return r.Request.execute()
		}
	}
	return res, err
}

func parseResponse(res *http.Response) (result, error) {
	r := result{}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}
	rb := reponseBody{}
	err = json.Unmarshal(bodyBytes, &rb)
	if rb.ErrorCode != codeNoError {
		if rb.ErrorMessage == "" {
			return r, fmt.Errorf("unknow API error")
		}
		return r, fmt.Errorf("%s", rb.ErrorMessage)
	}
	return rb.Result, nil
}
