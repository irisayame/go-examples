package main

import (
	"fmt"
	"github.com/avct/user-agent-surfer"
)

func GetDeviceType(userAgent string) string {
	ua := uasurfer.Parse(userAgent)

	switch ua.DeviceType {
	case uasurfer.DeviceComputer:
		return "desktop"
	case uasurfer.DevicePhone:
		return "mobile"
	default:
		return "unknown"
	}
}

func main() {
	user_agent := "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1C28 Safari/419.3"
	device := GetDeviceType(user_agent)
	fmt.Println("Device:", device)
}
