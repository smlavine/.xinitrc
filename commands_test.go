// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"testing"
)

func assert(t *testing.T, cond bool, format string, args ...interface{}) {
	if !cond {
		t.Errorf(format, args...)
	}
}

func TestWindowManager(t *testing.T) {
	t.Run("Default WM", func(t *testing.T) {
		// First make sure that WM actually isn't set.
		originalValue, existed := os.LookupEnv("WM")
		if existed {
			os.Unsetenv("WM")
			t.Cleanup(func() {
				os.Setenv("WM", originalValue)
			})
		}

		wm := windowManager()
		assert(t, wm.Args[0] == DEFAULT_WM,
			"Expected WM name to be %v, got %v",
			DEFAULT_WM, wm.Args[0])
	})

	t.Run("WM set", func(t *testing.T) {
		const WM = "examplewm"
		t.Setenv("WM", WM)
		wm := windowManager()
		assert(t, wm.Args[0] == WM,
			"Expected WM name to be %v, got %v", WM, wm.Args[0])
	})

	t.Run("Stdout and Stderr", func(t *testing.T) {
		wm := windowManager()
		assert(t, wm.Stdout == os.Stdout, "wm.Stdout != os.Stdout")
		assert(t, wm.Stderr == os.Stderr, "wm.Stderr != os.Stderr")
	})
}
