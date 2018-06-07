# Kasa Go

Go library to remotely control TP-Link devices using their cloud web service so they can be
controlled from outside the network.

## Installation

```
go get github.com/ivandelabeldad/kasa-go
```

## Devices Supported

| Device | Type |
| :--- | :--- |
| HS100 | Smart Plug |
| HS110 | Smart Plug |

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

```
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
the [MIT license](https://github.com/ivandelabeldad/kasa-go/blob/master/LICENSE).
