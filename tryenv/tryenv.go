// Copyright (c) 2022 Sebastian LaVine <mail@smlavine.com>
// SPDX-License-Identifier: MIT

package tryenv

import "os"

// Tryenv retrieves the value of the environment variable named by the key, or
// returns the fallback value if the variable is not present.
func Tryenv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return fallback
	}
}
