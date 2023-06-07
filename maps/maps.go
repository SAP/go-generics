/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package maps

// Get keys of map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Get values of map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Compare two maps by a given equality function (for values, keys are still compared by ==).
func EqualBy[K comparable, V any, W any](m map[K]V, n map[K]W, f func(V, W) bool) bool {
	if len(m) != len(n) {
		return false
	}
	for k, v := range m {
		if w, ok := n[k]; !ok || !f(v, w) {
			return false
		}
	}
	return true
}

// Compare two maps of comparable values.
func Equal[K comparable, V comparable](m map[K]V, n map[K]V) bool {
	f := func(x V, y V) bool {
		return x == y
	}
	return EqualBy(m, n, f)
}
