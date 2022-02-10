package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func IsEnableUI() bool {
	cmd := exec.Command(`tasklist`)
	out, _ := cmd.CombinedOutput()
	buff := bytes.NewBuffer(out)
	return strings.Contains(buff.String(), `vpnui.exe`)
}

func KillUI() bool {
	cmd := exec.Command(`taskkill`, `/F`, `/T`, `/IM`, `vpnui.exe`)
	out, _ := cmd.CombinedOutput()
	buff := bytes.NewBuffer(out)
	c1 := strings.Contains(buff.String(), `SUCCESS`)
	c2 := strings.Contains(buff.String(), `terminated`)
	return c1 && c2
}

func RunUI() bool {
	cmd := exec.Command(ExecVPNUI())
	err := cmd.Start()
	return err == nil
}
