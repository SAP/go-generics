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
func Remove[T comparable](s []T, x T) (r []T) {
	for _, y := range s {
		if y == x {
			continue
		}
		r = append(r, y)
	}
	return
}

// Reverse slice.
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
func Sort[T Orderable](s []T) []T {
	f := func(x, y T) bool {
		return x > y
	}
	return SortBy(s, f)
}

// Compare two slices by a given equality function.
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
func Equal[T comparable](s []T, t []T) bool {
	f := func(x T, y T) bool {
		return x == y
	}
	return EqualBy(s, t, f)
}

// Remove duplicates from slice of comparable elements.
func Uniq[T comparable](s []T) (r []T) {
	for _, x := range s {
		if !Contains(r, x) {
			r = append(r, x)
		}
	}
	return
}
