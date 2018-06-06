package kasa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	codeNoError = 0
)

type reponseBody struct {
	ErrorCode    int    `json:"error_code"`
	Result       result `json:"result"`
	ErrorMessage string `json:"msg"`
}

type result struct {
	ResponseData   string       `json:"responseData,omitempty"`
	DeviceInfoList []DeviceInfo `json:"deviceList,omitempty"`
	AccountID      string       `json:"accountID,omitempty"`
	RegTime        string       `json:"regTime,omitempty"`
	Email          string       `json:"email,omitempty"`
	Token          string       `json:"token,omitempty"`
}

// DeviceInfo ...
type DeviceInfo struct {
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

type request struct {
	URL         string
	Token       string
	RequestBody requestBody
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

// Exec ...
func (r request) Exec() (res result, err error) {
	if r.Token != "" {
		r.URL = r.URL + "?token=" + r.Token
	}
	reqBodyJSON, err := json.MarshalIndent(&r.RequestBody, "", "  ")
	if err != nil {
		return res, err
	}
	response, err := http.Post(r.URL, "application/json", bytes.NewReader(reqBodyJSON))
	if err != nil {
		return res, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	rb := reponseBody{}
	err = json.Unmarshal(bodyBytes, &rb)
	if rb.ErrorCode != codeNoError {
		if rb.ErrorMessage == "" {
			return res, fmt.Errorf("unknow API error")
		}
		return res, fmt.Errorf("%s", rb.ErrorMessage)
	}
	return rb.Result, err
}
