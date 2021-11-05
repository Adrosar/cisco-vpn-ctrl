package ciscovpn

import "time"

func Sleep(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}
