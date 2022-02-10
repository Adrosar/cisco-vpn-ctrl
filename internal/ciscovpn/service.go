package ciscovpn

import (
	"bytes"
	"os/exec"
	"strings"
)

func StopService() bool {
	cmd := exec.Command(`cmd`, `/C`, `sc config vpnagent start=disabled && sc stop vpnagent`)
	out, _ := cmd.CombinedOutput()
	str := bytes.NewBuffer(out).String()
	return strings.Contains(str, `ChangeServiceConfig SUCCESS`)
}

func StartService() bool {
	cmd := exec.Command(`cmd`, `/C`, `sc config vpnagent start=auto && sc start vpnagent`)
	out, _ := cmd.CombinedOutput()
	str := bytes.NewBuffer(out).String()

	c1 := strings.Contains(str, `ChangeServiceConfig SUCCESS`)
	c2 := strings.Contains(str, `START_PENDING`)
	c3 := strings.Contains(str, `StartService FAILED 1056`)

	return c1 && (c2 || c3)
}

func StateOfService() int {
	cmd := exec.Command(`sc`, `query`, `vpnagent`)
	out, _ := cmd.CombinedOutput()
	str := bytes.NewBuffer(out).String()

	c1 := strings.Contains(str, `STATE`)
	c2 := strings.Contains(str, `RUNNING`)
	c3 := strings.Contains(str, `STOPPED`)

	if c1 {
		if c2 {
			return 1 // Usługa działa :)
		}

		if c3 {
			return -1 // Usługa jest zatrzymana :/
		}
	}

	return 0 // Stan nieokreślony
}

func WaitForServiceStatus() int {
	for {
		s := StateOfService()
		if s != 0 {
			return s
		} else {
			Sleep(300)
		}
	}
}
