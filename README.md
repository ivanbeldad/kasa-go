# Kasa Go

[![Version](https://img.shields.io/github/tag/ivanbeldad/kasa-go.svg)](https://github.com/ivanbeldad/kasa-go)
[![License](https://img.shields.io/badge/license-MIT-orange.svg)](https://github.com/ivanbeldad/kasa-go/blob/master/LICENSE)

Go library to remotely control TP-Link devices using their cloud web service so they can be
controlled from outside the network.

## Installation

```
go get github.com/ivanbeldad/kasa-go
```

## Devices Supported

| Device | Type |
| :--- | :--- |
| HS100 | Smart Plug |
| HS103 | Smart Plug |
| HS110 | Smart Plug |
| HS200 | Smart Light Switch |

### Smart Plugs

| Method | Description |
| :--- | :--- |
| GetInfo | Get information about the smart plug |
| TurnOn | Turn on the device |
| TurnOff | Turn off the device |
| SwitchOnOff | Change the state of the device |
| Reboot | Power off the device and power on after a few seconds |
| GetAlias | Get the name of the device |
| ScanAPs | Scan near access points and return info about them |

## Example

```go
api, err := kasa.Connect("tplink@mail.com", "myStrongPassword")
if err != nil {
  log.Fatal(err)
}
hs100, err := api.GetHS100("livingroom")
if err != nil {
  log.Fatal(err)
}
err = hs100.TurnOn()
if err != nil {
  log.Fatal(err)
}
fmt.Print("Living Room turned on successfully!")
```

## License

Kasa Go is open-sourced software licensed under
the [MIT license](https://github.com/ivanbeldad/kasa-go/blob/master/LICENSE).
