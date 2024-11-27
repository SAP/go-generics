/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package sets

import "github.com/sap/go-generics/maps"

// Set.
// Always create sets with the New() function, do not use unininizialized sets (i.e. sets having the zero value).
type Set[T comparable] struct {
	m map[T]struct{}
}

// Create new set.
func New[T comparable](x ...T) Set[T] {
	s := Set[T]{m: make(map[T]struct{})}
	for _, y := range x {
		s.m[y] = struct{}{}
	}
	return s
}

// Clone set.
func Clone[T comparable](s Set[T]) Set[T] {
	t := Set[T]{m: make(map[T]struct{})}
	for x := range s.m {
		t.m[x] = struct{}{}
	}
	return t
}

// Get number of elements in the set.
func Len[T comparable](s Set[T]) int {
	return len(s.m)
}

// Get values of set as slice; order is not predictable.
// Will return an empty non-nil slice in case the set is empty.
func Values[T comparable](s Set[T]) []T {
	return maps.Keys(s.m)
}

// Check if set contains specified element.
func Contains[T comparable](s Set[T], x T) bool {
	_, ok := s.m[x]
	return ok
}

// Add specified element to set.
func Add[T comparable](s Set[T], x T) {
	s.m[x] = struct{}{}
}

// Delete specified element from set.
func Delete[T comparable](s Set[T], x T) {
	delete(s.m, x)
}

// Compare two sets.
func Equal[T comparable](s Set[T], t Set[T]) bool {
	if len(s.m) != len(t.m) {
		return false
	}
	for x := range s.m {
		_, ok := t.m[x]
		if !ok {
			return false
		}
	}
	return true
}
