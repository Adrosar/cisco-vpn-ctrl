package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func DisconnectVPN() bool {
	cmd := exec.Command(ExecVPNCLI(), `disconnect`)
	out, _ := cmd.CombinedOutput()
	str := bytes.NewBuffer(out).String()

	c1 := strings.Contains(str, `state: Disconnecting`)
	c2 := strings.Contains(str, `state: Disconnected`)
	c3 := strings.Contains(str, `not connected`)

	return (c1 && c2) || (c2 && c3)
}
