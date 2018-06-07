package kasa

import "fmt"

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
	ActiveMode string  `json:"active_mode,omitempty"`
	Feature    string  `json:"feature,omitempty"`
	RelayState int     `json:"relay_state,omitempty"`
	OnTime     int     `json:"on_time,omitempty"`
	Updating   int     `json:"updating,omitempty"`
	Rssi       int     `json:"rssi,omitempty"`
	LedOff     int     `json:"led_off,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

func (d *DeviceInfo) fromListedDeviceInfo(l listedDeviceInfo) {
	*d = DeviceInfo{
		FwID:       l.FwID,
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

func (d DeviceInfo) String() string {
	var s string
	if d.SwVer != "" {
		s += fmt.Sprintf("SwVer:       %40s\n", d.SwVer)
	}
	if d.HwVer != "" {
		s += fmt.Sprintf("HwVer:       %40s\n", d.HwVer)
	}
	if d.Type != "" {
		s += fmt.Sprintf("Type:        %40s\n", d.Type)
	}
	if d.Model != "" {
		s += fmt.Sprintf("Model:       %40s\n", d.Model)
	}
	if d.Mac != "" {
		s += fmt.Sprintf("Mac:         %40s\n", d.Mac)
	}
	if d.DeviceID != "" {
		s += fmt.Sprintf("DeviceID:    %40s\n", d.DeviceID)
	}
	if d.HwID != "" {
		s += fmt.Sprintf("HwID:        %40s\n", d.HwID)
	}
	if d.FwID != "" {
		s += fmt.Sprintf("FwID:        %40s\n", d.FwID)
	}
	if d.Alias != "" {
		s += fmt.Sprintf("Alias:       %40s\n", d.Alias)
	}
	if d.DeviceName != "" {
		s += fmt.Sprintf("DeviceName:  %40s\n", d.DeviceName)
	}
	if d.IconHash != "" {
		s += fmt.Sprintf("IconHash:    %40s\n", d.IconHash)
	}
	if d.ActiveMode != "" {
		s += fmt.Sprintf("ActiveMode:  %40s\n", d.ActiveMode)
	}
	if d.Feature != "" {
		s += fmt.Sprintf("Feature:     %40s\n", d.Feature)
	}
	s += fmt.Sprintf("RelayState:  %40d\n", d.RelayState)
	s += fmt.Sprintf("OnTime:      %40d\n", d.OnTime)
	s += fmt.Sprintf("Updating:    %40d\n", d.Updating)
	s += fmt.Sprintf("Rssi:        %40d\n", d.Rssi)
	s += fmt.Sprintf("LedOff:      %40d\n", d.LedOff)
	s += fmt.Sprintf("Latitude:    %40.2f\n", d.Latitude)
	s += fmt.Sprintf("Longitude:   %40.2f\n", d.Longitude)

	return s
}
