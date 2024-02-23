/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package slices

// Orderable constraint.
type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Check if slice contains given element.
func Contains[T comparable](s []T, x T) bool {
	for _, y := range s {
		if y == x {
			return true
		}
	}
	return false
}

// Remove all occurrences of given element from slice (and return new slice; old slice remains unchanged).
// If the input is nil, it will return nil; otherwise, if the result is empty, it will return an empty slice.
func Remove[T comparable](s []T, x T) (r []T) {
	if s == nil {
		return
	}
	r = make([]T, 0)
	for _, y := range s {
		if y == x {
			continue
		}
		r = append(r, y)
	}
	return
}

// Get first n elements of a slice.
// If n is greater than the length of the slice, the input slice will be returned.
// If the input is nil, it will return nil; otherwise, if the result is empty, it will return an empty slice.
func First[T any](s []T, n uint) []T {
	i := int(n)
	if i > len(s) {
		i = len(s)
	}
	return s[:i]
}

// Get last n elements of a slice.
// If n is greater than the length of the slice, the input slice will be returned.
// If the input is nil, it will return nil; otherwise, if the result is empty, it will return an empty slice.
func Last[T any](s []T, n uint) []T {
	i := len(s) - int(n)
	if i < 0 {
		i = 0
	}
	return s[i:]
}

// Reverse slice.
// If the input is nil, it will return nil; otherwise, if the input is empty, it will return an empty slice.
func Reverse[T any](s []T) (r []T) {
	if s == nil {
		return
	}
	l := len(s)
	r = make([]T, l)
	for i, x := range s {
		r[l-i-1] = x
	}
	return
}

// Sort slice by given comparator function.
// If the input is nil, it will return nil; otherwise, if the input is empty, it will return an empty slice.
// The comparator function f(x,y) must return true if x is larger than y, and false if x is smaller than y;
// the return value in case of equality does not matter (may be true or false).
func SortBy[T any](s []T, f func(x, y T) bool) (r []T) {
	l := len(s)
	if l <= 1 {
		return s
	}
	s1 := SortBy(s[0:l/2], f)
	s2 := SortBy(s[l/2:l], f)
	r = make([]T, l)
	i1, i2 := 0, 0
	for j := 0; j < l; j++ {
		if i1 >= len(s1) {
			r[j] = s2[i2]
			i2++
		} else if i2 >= len(s2) {
			r[j] = s1[i1]
			i1++
		} else {
			if f(s1[i1], s2[i2]) {
				r[j] = s2[i2]
				i2++
			} else {
				r[j] = s1[i1]
				i1++
			}
		}
	}
	return
}

// Sort slice of orderable elements.
// If the input is nil, it will return nil; otherwise, if the input is empty, it will return an empty slice.
func Sort[T Orderable](s []T) []T {
	f := func(x, y T) bool {
		return x > y
	}
	return SortBy(s, f)
}

// Compare two slices by a given equality function.
// Slices with different length are never equal.
// Empty and nil slices are always equal (in particular, comparing an empty with a nil slice yields true).
// Otherwise, the two slices are equal if all their elements are equal by means of the provided equality function.
func EqualBy[S any, T any](s []S, t []T, f func(S, T) bool) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !f(s[i], t[i]) {
			return false
		}
	}
	return true
}

// Compare two slices of comparable elements.
// Slices with different length are never equal.
// Empty and nil slices are always equal (in particular, comparing an empty with a nil slice yields true).
// Otherwise, the two slices are equal, if all their elements are equal by means of the == operator.
func Equal[T comparable](s []T, t []T) bool {
	f := func(x T, y T) bool {
		return x == y
	}
	return EqualBy(s, t, f)
}

// Remove duplicates from slice by given mapper function.
// Two elements are considered equal if the mapper function returns the same value for them.
// Preserves order. The first occurrence of an element will be kept, the other occurrences will be dropped.
// If the input is nil, it will return nil; otherwise, if the input is empty, it will return an empty slice.
func UniqBy[S any, T comparable](s []S, f func(S) T) (r []S) {
	if s == nil {
		return
	}
	r = make([]S, 0)
	m := make(map[T]struct{})
	for _, x := range s {
		y := f(x)
		if _, ok := m[y]; !ok {
			m[y] = struct{}{}
			r = append(r, x)
		}
	}
	return
}

// Remove duplicates from slice of comparable elements.
// Preserves order. The first occurrence of an element will be kept, the other occurrences will be dropped.
// If the input is nil, it will return nil; otherwise, if the input is empty, it will return an empty slice.
func Uniq[T comparable](s []T) (r []T) {
	f := func(x T) T {
		return x
	}
	return UniqBy(s, f)
}

// Collect (map) slice through given function.
// If the input is nil, it will return nil; if the input is empty, it will return an empty slice.
// Otherweise, it will return return a slice of the same length as the input slice, containing the
// elements mapped through the provided function f.
func Collect[S any, T any](s []S, f func(S) T) (r []T) {
	if s == nil {
		return nil
	}
	r = make([]T, len(s))
	for i, x := range s {
		r[i] = f(x)
	}
	return
}

// Report whether the given boolean function evaluates to true for at least one element of the given slice.
// Returns false for nil or empty slices.
// Any(s, f) is equivalent to !All(s, !f) and !None(s, f).
func Any[T any](s []T, f func(T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}
	return false
}

// Report whether the given boolean function evaluates to true for all elements of the given slice.
// Returns true for nil or empty slices.
// All(s, f) is equivalent to !Any(s, !f) and None(s, !f).
func All[T any](s []T, f func(T) bool) bool {
	for _, x := range s {
		if !f(x) {
			return false
		}
	}
	return true
}

// Report whether the given boolean function evaluates to true for none of the elements of the given slice.
// Equivalently: evaluates to false for all of the elements of the given slice.
// Returns true for nil or empty slices.
// None(s, f) is equivalent to !Any(s, f) and All(s, !f).
func None[T any](s []T, f func(T) bool) bool {
	for _, x := range s {
		if f(x) {
			return false
		}
	}
	return true
}
