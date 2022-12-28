// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"testing"

	. "git.sr.ht/~smlavine/.xinitrc/assert"
)

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
		Assert(t, wm.Args[0] == DEFAULT_WM,
			"Expected WM name to be %v, got %v",
			DEFAULT_WM, wm.Args[0])
	})

	t.Run("WM set", func(t *testing.T) {
		const WM = "examplewm"
		t.Setenv("WM", WM)
		wm := windowManager()
		Assert(t, wm.Args[0] == WM,
			"Expected WM name to be %v, got %v", WM, wm.Args[0])
	})

	t.Run("Stdout and Stderr", func(t *testing.T) {
		wm := windowManager()
		Assert(t, wm.Stdout == os.Stdout, "wm.Stdout != os.Stdout")
		Assert(t, wm.Stderr == os.Stderr, "wm.Stderr != os.Stderr")
	})
}
