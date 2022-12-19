// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"os"
	"os/exec"
)

// The notification daemon to use.
const NOTIFICATION_DAEMON string = "dunst"

// The window manager to use if WM is not set in the environment.
const DEFAULT_WM string = "dwm"

// Tryenv retrieves the value of the environment variable named by the key, or
// returns the fallback value if the variable is not present.
func Tryenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return fallback
	}
}

// Starts the notification daemon.
func startNotificationDaemon() error {
	nd := exec.Command(NOTIFICATION_DAEMON)
	nd.Stdout = os.Stdout
	nd.Stderr = os.Stderr
	return nd.Start()
}

// Runs the window manager. Uses WM if it exists; otherwise uses DEFAULT_WM.
func runWM() error {
	wm := exec.Command(Tryenv("WM", DEFAULT_WM))
	wm.Stdout = os.Stdout
	wm.Stderr = os.Stderr
	return wm.Run()
}

func main() {
	if err := startNotificationDaemon(); err != nil {
		log.Println(err)
	}
	if err := runWM(); err != nil {
		log.Fatalln(err)
	}
}
