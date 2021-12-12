package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func isEnableUI() bool {
	cmd := exec.Command(`cmd`, `/C`, `tasklist`)
	out, _ := cmd.Output()
	buff := bytes.NewBuffer(out)
	c := strings.Index(buff.String(), `vpnui.exe`)
	if c > -1 {
		return true
	} else {
		return false
	}
}

func killVpnUI() bool {
	cmd := exec.Command(`cmd`, `/C`, `taskkill /F /T /IM vpnui.exe`)
	out, _ := cmd.Output()
	buff := bytes.NewBuffer(out)

	c1 := strings.Index(buff.String(), `SUCCESS`)
	c2 := strings.Index(buff.String(), `terminated`)

	if c1 > -1 && c2 > -1 {
		return true
	}

	return false
}

func runVpnUI() bool {
	cmd := exec.Command(`cmd`, `/C`, `START`, `Cisco AnyConnect Secure Mobility Client`, PathToExe())
	err := cmd.Start()
	if err == nil {
		return true
	} else {
		return false
	}
}
