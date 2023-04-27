/*
Copyright 2022 Cloud Services.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sets

import "github.com/sap/go-generics/maps"

// Set.
type Set[T comparable] struct {
	m map[T]struct{}
}

// Create new set.
func New[T comparable](x ...T) *Set[T] {
	s := Set[T]{m: make(map[T]struct{})}
	for _, y := range x {
		s.m[y] = struct{}{}
	}
	return &s
}

// Get number of elements in the set.
func Len[T comparable](s Set[T]) int {
	return len(s.m)
}

// Get values of set as slice; order is not predictable.
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
