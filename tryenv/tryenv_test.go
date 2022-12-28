// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package tryenv

import (
	"testing"

	. "git.sr.ht/~smlavine/.xinitrc/assert"
)

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
			v := Tryenv(test.key, FALLBACK)
			Assert(t, v == test.value, "Expected %v, got %v", test.value, v)
		})
	}
	t.Run("Unset variable", func(t *testing.T) {
		v := Tryenv("_UNASSIGNED_", FALLBACK)
		Assert(t, v == FALLBACK, "Expected %v, got %v", FALLBACK, v)
	})
}
