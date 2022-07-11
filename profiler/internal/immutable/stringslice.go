// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

// Package immutable provides read-only types
package immutable

// StringSlice is a wrapped slice which cannot be modified and must be copied to
// access. The zero value for a StringSlice is ready to use
type StringSlice struct {
	s []string
}

// NewStringSlice wraps a copy of the input slice
func NewStringSlice(s []string) StringSlice {
	return StringSlice{s: append([]string{}, s...)}
}

// Get returns a copy of the wrapped slice
func (s StringSlice) Get() []string {
	return append([]string{}, s.s...)
}

// Append creates a new StringSlice by concatenating the given strings to a copy
// of the slice wrapped by f.
func (s StringSlice) Append(strings ...string) StringSlice {
	dup := make([]string, len(s.s)+len(strings))
	copy(dup, s.s)
	copy(dup[len(s.s):], strings)
	return StringSlice{s: dup}
}
