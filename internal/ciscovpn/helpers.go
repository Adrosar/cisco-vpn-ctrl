package ciscovpn

import (
	"os"
	"time"
)

func Sleep(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var pathToExe string = ""

func PathToExe() string {
	if len(pathToExe) > 0 {
		return pathToExe
	}

	list := []string{
		"vpnui.exe",
		`C:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`D:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`E:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
	}

	for _, v := range list {
		if FileExists(v) {
			pathToExe = v
			return v
		}
	}

	return ""
}
