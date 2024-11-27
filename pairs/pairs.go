/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package pairs

// Pair.
type Pair[S any, T any] struct {
	X S
	Y T
}

// Create new pair.
func New[S any, T any](x S, y T) *Pair[S, T] {
	return &Pair[S, T]{X: x, Y: y}
}
