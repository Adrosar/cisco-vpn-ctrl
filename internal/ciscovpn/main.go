package ciscovpn

import (
	"fmt"
)

func Start() {
	if WaitForServiceStatus() == -1 {
		if !StartService() {
			fmt.Println("[ERR01] Failed to start VPN service!")
			fmt.Println("[.....] Maybe you don't have admin permissions?")
			return
		}
	}

	if IsEnableUI() {
		fmt.Println("[WAR01] UI is already up and running. Check next to the clock (tray)")
		return
	}

	if !RunUI() {
		fmt.Println(`[ERR02] Failed to launch UI!`)
		fmt.Println(`[.....] Maybe the file "vpnui.exe" doesn't exist!`)
		return
	}

	fmt.Println("VPN is running")
}

func Stop() {
	if IsEnableUI() {
		if !KillUI() {
			fmt.Println(`[ERR06] Cannot kill the UI process!`)
			return
		}
	}

	if WaitForServiceStatus() == 1 {
		if !DisconnectVPN() {
			fmt.Println(`[ERR03] Unable to disconnect VPN!`)
			fmt.Println(`[.....] Maybe the file "vpncli.exe" does not exist!`)
			return
		}

		if !StopService() {
			fmt.Println("[ERR04] Failed to stop VPN service!")
			fmt.Println("[.....] Maybe you don't have admin permissions?")
			return
		}
	}

	fmt.Println("VPN is stopped")
}
