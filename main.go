// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"os"
	"os/exec"
)

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

func main() {
	wm := exec.Command(Tryenv("WM", DEFAULT_WM))
	wm.Stdout = os.Stdout
	wm.Stderr = os.Stderr
	if err := wm.Run(); err != nil {
		log.Println(err)
	}
}
