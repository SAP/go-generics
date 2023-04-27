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
