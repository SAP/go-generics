/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package maps

// Get keys of map.
// If the input is nil, it will return a nil slice; otherwise, if the input is empty, it will return an empty slice.
// Otherwise it will return a slice containing the keys of the given map.
// Note that there is no guarantee about the order of the returned keys.
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Get values of map.
// If the input is nil, it will return a nil slice; otherwise, if the input is empty, it will return an empty slice.
// Otherwise it will return a slice containing the values of the given map (repeating identical values, if any).
// Note that there is no guarantee about the order of the returned values.
func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Compare two maps by a given equality function (for values, keys are still compared by ==).
// Maps with different length are never equal.
// Empty and nil maps are always equal (in particular, comparing an empty with a nil map yields true).
// Otherwise, the two maps are equal if corresponding values are equal by means of the provided equality function.
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
// Maps with different length are never equal.
// Empty and nil maps are always equal (in particular, comparing an empty with a nil map yields true).
// Otherwise, the two maps are equal if corresponding values are equal by means of the == operator.
func Equal[K comparable, V comparable](m map[K]V, n map[K]V) bool {
	f := func(x V, y V) bool {
		return x == y
	}
	return EqualBy(m, n, f)
}

// Collect map through given function.
// If the input is nil, it will return nil; if the input is empty, it will return an empty map.
// Otherweise, it will return return a map of the same length as the input map, containing the
// same keys, and the values mapped through the provided function f.
func Collect[K comparable, V any, W any](m map[K]V, f func(V) W) map[K]W {
	if m == nil {
		return nil
	}
	n := make(map[K]W)
	for k, v := range m {
		n[k] = f(v)
	}
	return n
}

// Select sub-map from map by given function.
// If the input is nil, it will return nil; if the input is empty, it will return an empty map.
// Otherwise, it will return a map containing the keys and according values, for which the provided
// function f evaluates to true.
func Select[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	if m == nil {
		return nil
	}
	n := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			n[k] = v
		}
	}
	return n
}

// Select sub-map from map by the given keys.
// If the input is nil, it will return nil; if the input is empty, it will return an empty map.
// Otherwise, it will return a map containing those keys and according values, if existing in the input map;
// Keys which do not exist in the input map will be ignored.
func SelectByKeys[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if m == nil {
		return nil
	}
	n := make(map[K]V)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			n[k] = v
		}
	}
	return n
}

// Count elements for which the given boolean function evaluates to true.
func Count[K comparable, V any](m map[K]V, f func(K, V) bool) (c int) {
	for k, v := range m {
		if f(k, v) {
			c++
		}
	}
	return
}
