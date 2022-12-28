// Copyright (C) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package assert

import "testing"

// Calls t.Errorf(format, args...) if cond is false.
func Assert(t *testing.T, cond bool, format string, args ...interface{}) {
	if !cond {
		t.Errorf(format, args...)
	}
}
