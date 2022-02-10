package ciscovpn

import (
	"os"
	"time"
)

func Sleep(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FirstExistFile(list []string) string {
	for _, path := range list {
		if IsFileExists(path) {
			return path
		}
	}

	return ""
}

func ExecVPNUI() string {
	return FirstExistFile([]string{
		`vpnui.exe`,
		`C:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`D:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`E:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
		`Cisco AnyConnect Secure Mobility Client\vpnui.exe`,
	})
}

func ExecVPNCLI() string {
	return FirstExistFile([]string{
		`vpncli.exe`,
		`C:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpncli.exe`,
		`D:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpncli.exe`,
		`E:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpncli.exe`,
		`Cisco AnyConnect Secure Mobility Client\vpncli.exe`,
	})
}
