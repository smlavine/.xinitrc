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

// Returns a Cmd for the notification daemon.
func notifDaemon() *exec.Cmd {
	cmd := exec.Command(NOTIFICATION_DAEMON)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

// Returns a Cmd for the window manager.
func windowManager() *exec.Cmd {
	cmd := exec.Command(Tryenv("WM", DEFAULT_WM))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func main() {
	notifications := notifDaemon()
	wm := windowManager()

	if err := notifications.Start(); err != nil {
		log.Println(err)
	}

	if err := wm.Run(); err != nil {
		log.Fatalln(err)
	}

}
