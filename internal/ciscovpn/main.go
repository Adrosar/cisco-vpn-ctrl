package ciscovpn

import "fmt"

func Start() {
	var result bool = false

	result = vpnStart()
	if !result {
		fmt.Println("[ERR01] Failed to start VPN service!")
		fmt.Println("[.....] Maybe you don't have admin permissions?")
		return
	}

	result = isEnableUI()
	if result {
		fmt.Println("UI is already up and running. Check next to the clock (tray)")
		return
	}

	Sleep(300)
	result = runVpnUI()
	if !result {
		fmt.Println("[ERR02] Failed to launch UI!")
		return
	}

	fmt.Println("VPN has been launched :)")
}

func Stop() {
	var result bool = false

	result = vpnDisconnect()
	if !result {
		fmt.Println("[WAR01] A problem occurred while disconnecting from VPN.")
		fmt.Println("[.....] The service was probably stopped earlier.")
	}

	result = isEnableUI()
	if result {
		result = killVpnUI()
		if !result {
			fmt.Println("[ERR03] UI not disabled!")
			return
		}

		Sleep(300)
	}

	result = vpnStop()
	if !result {
		fmt.Println("[ERR04] Failed to stop VPN service!")
		fmt.Println("[.....] Maybe you don't have admin permissions?")
		return
	}

	fmt.Println("VPN has been stopped :)")
}
