package main

import (
	"bufio"
	"cisco-vpn-ctrl/internal/ciscovpn"
	"flag"
	"fmt"
	"os"
)

func main() {
	var flagStart bool
	flag.BoolVar(&flagStart, "start", false, "Start the service and UI")

	var flagStop bool
	flag.BoolVar(&flagStop, "stop", false, "Turn off the UI and stop the service")

	flag.Parse()

	if flagStart {
		ciscovpn.Start()
		return
	}

	if flagStop {
		ciscovpn.Stop()
		return
	}

	fmt.Println("Cisco VPN Controller (v0.3.0)")
	fmt.Println("----------------------------------------------")
	fmt.Println("What do you want to do?")
	fmt.Println("- start : Start the service and UI")
	fmt.Println("- stop  : Turn off the UI and stop the service")
	fmt.Println("- exit  : Exit the application")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "start" {
			ciscovpn.Start()
			waitAndExit()
		}

		if text == "stop" {
			ciscovpn.Stop()
			waitAndExit()
		}

		if text == "exit" {
			os.Exit(0)
		}
	}
}

func waitAndExit() {
	fmt.Print(`Terminal closing in:`)
	for i := 5; i > 0; i-- {
		fmt.Print(` `, i)
		ciscovpn.Sleep(1000)
	}
	os.Exit(0)
}
