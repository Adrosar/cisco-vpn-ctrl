package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func vpnStop() bool {
	cmd := exec.Command(`cmd`, `/C`, `sc config vpnagent start=disabled && sc stop vpnagent`)
	out, _ := cmd.Output()
	buff := bytes.NewBuffer(out)
	n := strings.Index(buff.String(), `ChangeServiceConfig SUCCESS`)
	if n > -1 {
		return true
	} else {
		return false
	}
}

func vpnStart() bool {
	cmd := exec.Command(`cmd`, `/C`, `sc config vpnagent start=auto && sc start vpnagent`)
	out, _ := cmd.Output()
	buff := bytes.NewBuffer(out)

	c1 := strings.Index(buff.String(), `ChangeServiceConfig SUCCESS`)
	c2 := strings.Index(buff.String(), `START_PENDING`)
	c3 := strings.Index(buff.String(), `1056`)

	if c1 > -1 {
		if c2 > -1 {
			return true
		}

		if c3 > -1 {
			return true
		}
	}

	return false
}
