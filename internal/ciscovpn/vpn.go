package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func vpnDisconnect() bool {
	cmd := exec.Command(`C:\Program Files (x86)\Cisco\Cisco AnyConnect Secure Mobility Client\vpncli.exe`, `disconnect`)
	go waitAndKill(cmd)

	out, _ := cmd.Output()
	buff := bytes.NewBuffer(out)

	c1 := strings.Index(buff.String(), `Disconnecting`)
	c2 := strings.Index(buff.String(), `Disconnected`)
	c3 := strings.Index(buff.String(), `not connected`)

	if c1 > -1 && c2 > -1 {
		return true
	}

	if c2 > -1 && c3 > -1 {
		return true
	}

	return false
}

func waitAndKill(cmd *exec.Cmd) {
	Sleep(1500)
	if cmd != nil {
		cmd.Process.Kill()
	}
}
