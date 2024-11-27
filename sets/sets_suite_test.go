/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package sets_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sap/go-generics/sets"
)

func TestSets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sets Suite")
}

var _ = Describe("sets", func() {
	var emptySet sets.Set[int]
	var setA sets.Set[int]

	BeforeEach(func() {
		emptySet = sets.New[int]()
		setA = sets.New(1, 2, 3, 3)
	})

	AfterEach(func() {
		Expect(sets.Equal(emptySet, sets.New[int]())).To(BeTrue())
		Expect(sets.Equal(setA, sets.New(1, 2, 3))).To(BeTrue())
	})

	Describe("tests for Len()", func() {
		Context("with an empty set", func() {
			It("should return 0", func() {
				Expect(sets.Len(emptySet)).To(Equal(0))
			})
		})
		Context("with a set with three elements", func() {
			It("should return 3", func() {
				Expect(sets.Len(setA)).To(Equal(3))
			})
		})
	})

	Describe("tests for Values()", func() {
		Context("with an empty set", func() {
			It("should return an empty slice", func() {
				Expect(sets.Values(emptySet)).To(Equal([]int{}))
			})
		})
		Context("with a non-empty set", func() {
			It("should return the slice of values", func() {
				Expect(sets.Values(setA)).To(ConsistOf(1, 2, 3))
			})
		})
	})

	Describe("tests for Contains()", func() {
		Context("with an empty set", func() {
			It("should return false", func() {
				Expect(sets.Contains(emptySet, 1)).To(BeFalse())
			})
		})
		Context("with a non-empty set that contains the value", func() {
			It("should return true", func() {
				Expect(sets.Contains(setA, 2)).To(BeTrue())
			})
		})
		Context("with a non-empty set that does not contain the value", func() {
			It("should return false", func() {
				Expect(sets.Contains(setA, 4)).To(BeFalse())
			})
		})
	})

	Describe("tests for Add()", func() {
		Context("with a set that contains the value", func() {
			It("should return the set unchanged", func() {
				set := sets.Clone(setA)
				sets.Add(set, 3)
				Expect(sets.Equal(set, sets.New(1, 2, 3))).To(BeTrue())
			})
		})
		Context("with a set that does not contain the value", func() {
			It("should return the augmented set", func() {
				set := sets.Clone(setA)
				sets.Add(set, 4)
				Expect(sets.Equal(set, sets.New(1, 2, 3, 4))).To(BeTrue())
			})
		})
	})

	Describe("tests for Delete()", func() {
		Context("with a set that contains the value", func() {
			It("should return the diminished set", func() {
				set := sets.Clone(setA)
				sets.Delete(set, 3)
				Expect(sets.Equal(set, sets.New(1, 2))).To(BeTrue())
			})
		})
		Context("with a set that does not contain the value", func() {
			It("should return the set unchanged", func() {
				set := sets.Clone(setA)
				sets.Delete(set, 4)
				Expect(sets.Equal(set, sets.New(1, 2, 3))).To(BeTrue())
			})
		})
	})
})
