// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"os/exec"

	"git.sr.ht/~smlavine/.xinitrc/tryenv"
)

// The notification daemon to use.
const NOTIFICATION_DAEMON string = "dunst"

// The window manager to use if WM is not set in the environment.
const DEFAULT_WM string = "dwm"

// Returns a Cmd that sets the desktop background.
func backgroundSetter() *exec.Cmd {
	const file = "/usr/include/X11/bitmaps/xsnow" // For the Holidays
	//const file = "/usr/include/X11/bitmaps/xlogo32"
	const bg = "rgb:14/13/7E"
	const fg = "rgb:2A/89/F5"
	cmd := exec.Command("xsetroot", "-bitmap", file, "-bg", bg, "-fg", fg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
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
	cmd := exec.Command(tryenv.Tryenv("WM", DEFAULT_WM))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
