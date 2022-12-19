// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package main

import "testing"

func TestTryenv(t *testing.T) {
	const FALLBACK = "FALLBACK"
	setVarSubtests := []struct {
		key, value string
	}{
		{"SOMEVAR", "12345678"},
		{"SET_BUT_EMPTY", ""},
		{"TO_BE_REASSIGNED", "fakeuser"},
	}
	t.Setenv("TO_BE_REASSIGNED", "exampleuser")
	for _, test := range setVarSubtests {
		t.Setenv(test.key, test.value)
		t.Run(test.key, func(t *testing.T) {
			if v := Tryenv(test.key, FALLBACK); v != test.value {
				t.Errorf("Expected %v, got %v", test.value, v)
			}
		})
	}
	t.Run("Unset variable", func(t *testing.T) {
		if v := Tryenv("_UNASSIGNED_", FALLBACK); v != FALLBACK {
			t.Errorf("Expected %v, got %v", FALLBACK, v)
		}
	})
}
