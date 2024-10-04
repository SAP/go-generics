/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and go-generics contributors
SPDX-License-Identifier: Apache-2.0
*/

package slices_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sap/go-generics/slices"
)

func TestSlices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slices Suite")
}

var _ = Describe("slices", func() {
	var nilSlice []int
	var emptySlice []int
	var sliceA []int
	var sliceB []int
	var sliceC []int
	var sliceD []int
	var sliceE []float64
	var sliceF []float64

	BeforeEach(func() {
		nilSlice = nil
		emptySlice = []int{}
		sliceA = []int{1}
		sliceB = []int{1, 2, 3}
		sliceC = []int{1, 2, 3, 1, 2, 3}
		sliceD = []int{9, 6, 5, 6, 3, 7, 7, 1, 2, 8}
		sliceE = []float64{1.5, 2.99, 3.14}
		sliceF = []float64{2.1, 1.5, 2.6, 3.4, 1.9}
	})

	AfterEach(func() {
		Expect(nilSlice).To(BeNil())
		Expect(emptySlice).To(Equal([]int{}))
		Expect(sliceA).To(Equal([]int{1}))
		Expect(sliceB).To(Equal([]int{1, 2, 3}))
		Expect(sliceC).To(Equal([]int{1, 2, 3, 1, 2, 3}))
		Expect(sliceD).To(Equal([]int{9, 6, 5, 6, 3, 7, 7, 1, 2, 8}))
		Expect(sliceE).To(Equal([]float64{1.5, 2.99, 3.14}))
		Expect(sliceF).To(Equal([]float64{2.1, 1.5, 2.6, 3.4, 1.9}))
	})

	Describe("tests for Contains()", func() {
		Context("with a nil slice", func() {
			It("should return false", func() {
				Expect(slices.Contains(nilSlice, 1)).To(BeFalse())
			})
		})
		Context("with an empty slice", func() {
			It("should return false", func() {
				Expect(slices.Contains(emptySlice, 1)).To(BeFalse())
			})
		})
		Context("with a slice not containing the element", func() {
			It("should return false", func() {
				Expect(slices.Contains(sliceB, 4)).To(BeFalse())
			})
		})
		Context("with a slice containing the element", func() {
			It("should return true", func() {
				Expect(slices.Contains(sliceC, 2)).To(BeTrue())
			})
		})
	})

	Describe("tests for Remove()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Remove(nilSlice, 1)).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Remove(emptySlice, 1)).To(Equal(emptySlice))
			})
		})
		Context("with an empty result", func() {
			It("should return an empty slice", func() {
				Expect(slices.Remove(sliceA, 1)).To(Equal(emptySlice))
			})
		})
		Context("with a slice not containing the element", func() {
			It("should return the slice unchanged", func() {
				Expect(slices.Remove(sliceB, 4)).To(Equal(sliceB))
			})
		})
		Context("with a slice containing the element", func() {
			It("should return the slice without the element", func() {
				Expect(slices.Remove(sliceC, 2)).To(Equal([]int{1, 3, 1, 3}))
			})
		})
	})

	Describe("tests for First()", func() {
		Context("with a nil slice (first 0)", func() {
			It("should return nil", func() {
				Expect(slices.First(nilSlice, 0)).To(Equal(nilSlice))
			})
		})
		Context("with a nil slice (first 1)", func() {
			It("should return nil", func() {
				Expect(slices.First(nilSlice, 1)).To(Equal(nilSlice))
			})
		})
		Context("with a nil slice (first 2)", func() {
			It("should return nil", func() {
				Expect(slices.First(nilSlice, 2)).To(Equal(nilSlice))
			})
		})

		Context("with an empty slice (first 0)", func() {
			It("should return an empty slice", func() {
				Expect(slices.First(emptySlice, 0)).To(Equal(emptySlice))
			})
		})
		Context("with an empty slice (first 1)", func() {
			It("should return an empty slice", func() {
				Expect(slices.First(emptySlice, 1)).To(Equal(emptySlice))
			})
		})
		Context("with an empty slice (first 2)", func() {
			It("should return an empty slice", func() {
				Expect(slices.First(emptySlice, 2)).To(Equal(emptySlice))
			})
		})

		Context("with a slice of length 3 (first 0)", func() {
			It("should return an empty slice", func() {
				Expect(slices.First(sliceB, 0)).To(Equal(emptySlice))
			})
		})
		Context("with a slice of length 3 (first 2)", func() {
			It("should return the first two elements", func() {
				Expect(slices.First(sliceB, 2)).To(Equal([]int{1, 2}))
			})
		})
		Context("with a slice of length 3 (first 3)", func() {
			It("should return the slice unchanged", func() {
				Expect(slices.First(sliceB, 3)).To(Equal(sliceB))
			})
		})
		Context("with a slice of length 3 (first 4)", func() {
			It("should return the slice unchanged", func() {
				Expect(slices.First(sliceB, 4)).To(Equal(sliceB))
			})
		})
	})

	Describe("tests for Last()", func() {
		Context("with a nil slice (last 0)", func() {
			It("should return nil", func() {
				Expect(slices.Last(nilSlice, 0)).To(Equal(nilSlice))
			})
		})
		Context("with a nil slice (last 1)", func() {
			It("should return nil", func() {
				Expect(slices.Last(nilSlice, 1)).To(Equal(nilSlice))
			})
		})
		Context("with a nil slice (last 2)", func() {
			It("should return nil", func() {
				Expect(slices.Last(nilSlice, 2)).To(Equal(nilSlice))
			})
		})

		Context("with an empty slice (last 0)", func() {
			It("should return an empty slice", func() {
				Expect(slices.Last(emptySlice, 0)).To(Equal(emptySlice))
			})
		})
		Context("with an empty slice (last 1)", func() {
			It("should return an empty slice", func() {
				Expect(slices.Last(emptySlice, 1)).To(Equal(emptySlice))
			})
		})
		Context("with an empty slice (last 2)", func() {
			It("should return an empty slice", func() {
				Expect(slices.Last(emptySlice, 2)).To(Equal(emptySlice))
			})
		})

		Context("with a slice of length 3 (last 0)", func() {
			It("should return an empty slice", func() {
				Expect(slices.Last(sliceB, 0)).To(Equal(emptySlice))
			})
		})
		Context("with a slice of length 3 (last 2)", func() {
			It("should return the last two elements", func() {
				Expect(slices.Last(sliceB, 2)).To(Equal([]int{2, 3}))
			})
		})
		Context("with a slice of length 3 (last 3)", func() {
			It("should return the slice unchanged", func() {
				Expect(slices.Last(sliceB, 3)).To(Equal(sliceB))
			})
		})
		Context("with a slice of length 3 (last 4)", func() {
			It("should return the slice unchanged", func() {
				Expect(slices.Last(sliceB, 4)).To(Equal(sliceB))
			})
		})
	})

	Describe("tests for Reverse()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Reverse(nilSlice)).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Reverse(emptySlice)).To(Equal(emptySlice))
			})
		})
		Context("with a slice of length three", func() {
			It("should return slice in reverse order", func() {
				Expect(slices.Reverse(sliceB)).To(Equal([]int{3, 2, 1}))
			})
		})
	})

	Describe("tests for SortBy()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.SortBy(nilSlice, func(int, int) bool { return false })).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.SortBy(emptySlice, func(int, int) bool { return false })).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a sorted slice, ordered as defined by the provided order function", func() {
				// sort as follows: even before odd; then all even numbers descending, all odd numbers ascending
				f := func(x int, y int) bool {
					if x%2 == 0 && y%2 != 0 {
						return false
					}
					if x%2 != 0 && y%2 == 0 {
						return true
					}
					if x%2 == 0 && y%2 == 0 {
						return y >= x
					}
					if x%2 != 0 && y%2 != 0 {
						return x >= y
					}
					panic("cannot happen")
				}
				Expect(slices.SortBy(sliceD, f)).To(Equal([]int{8, 6, 6, 2, 1, 3, 5, 7, 7, 9}))
			})
		})
	})

	Describe("tests for Sort()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Sort(nilSlice)).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Sort(emptySlice)).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a sorted slice", func() {
				Expect(slices.Sort(sliceD)).To(Equal([]int{1, 2, 3, 5, 6, 6, 7, 7, 8, 9}))
			})
		})
	})

	Describe("tests for EqualBy()", func() {
		Context("with different lengths", func() {
			It("should return false", func() {
				Expect(slices.EqualBy(sliceA, sliceB, func(int, int) bool { return true })).To(BeFalse())
			})
		})
		Context("with nil/nil slices", func() {
			It("should return true", func() {
				Expect(slices.EqualBy(nilSlice, nilSlice, func(int, int) bool { return false })).To(BeTrue())
			})
		})
		Context("with empty/empty slices", func() {
			It("should return true", func() {
				Expect(slices.EqualBy(emptySlice, emptySlice, func(int, int) bool { return false })).To(BeTrue())
			})
		})
		Context("with nil/empty slices", func() {
			It("should return true", func() {
				Expect(slices.EqualBy(nilSlice, emptySlice, func(int, int) bool { return false })).To(BeTrue())
			})
		})
		Context("with more complex slices", func() {
			It("should return true", func() {
				f := func(x int, y float64) bool {
					return x == int(y)
				}
				Expect(slices.EqualBy(sliceB, sliceE, f)).To(BeTrue())
			})
		})
	})

	Describe("tests for Equal()", func() {
		Context("with different lengths", func() {
			It("should return false", func() {
				Expect(slices.Equal(sliceA, sliceB)).To(BeFalse())
			})
		})
		Context("with nil/nil slices", func() {
			It("should return true", func() {
				Expect(slices.Equal(nilSlice, nilSlice)).To(BeTrue())
			})
		})
		Context("with empty/empty slices", func() {
			It("should return true", func() {
				Expect(slices.Equal(emptySlice, emptySlice)).To(BeTrue())
			})
		})
		Context("with nil/empty slices", func() {
			It("should return true", func() {
				Expect(slices.Equal(nilSlice, emptySlice)).To(BeTrue())
			})
		})
		Context("with identical slices", func() {
			It("should return true", func() {
				Expect(slices.Equal(sliceB, []int{1, 2, 3})).To(BeTrue())
			})
		})
		Context("with different slices", func() {
			It("should return false", func() {
				Expect(slices.Equal(sliceB, []int{1, 2, 4})).To(BeFalse())
			})
		})
	})

	Describe("tests for UniqBy()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.UniqBy(nilSlice, func(int) bool { return false })).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.UniqBy(emptySlice, func(int) bool { return false })).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a slice preserving order, but without duplicates", func() {
				f := func(x float64) int {
					return int(x)
				}
				Expect(slices.UniqBy(sliceF, f)).To(Equal([]float64{2.1, 1.5, 3.4}))
			})
		})
	})

	Describe("tests for Uniq()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Uniq(nilSlice)).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Uniq(emptySlice)).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a slice preserving order, but without duplicates", func() {
				Expect(slices.Uniq(sliceD)).To(Equal([]int{9, 6, 5, 3, 7, 1, 2, 8}))
			})
		})
	})

	Describe("tests for Collect()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Collect(nilSlice, func(int) int { return -1 })).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Collect(emptySlice, func(int) int { return -1 })).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a slice containing the mapped elements", func() {
				f := func(x float64) int {
					return int(x)
				}
				Expect(slices.Collect(sliceF, f)).To(Equal([]int{2, 1, 2, 3, 1}))
			})
		})
	})

	Describe("tests for Select()", func() {
		Context("with a nil slice", func() {
			It("should return nil", func() {
				Expect(slices.Select(nilSlice, func(int) bool { return true })).To(Equal(nilSlice))
			})
		})
		Context("with an empty slice", func() {
			It("should return an empty slice", func() {
				Expect(slices.Select(emptySlice, func(int) bool { return true })).To(Equal(emptySlice))
			})
		})
		Context("with a more complex slice", func() {
			It("should return a slice containing the selected elements", func() {
				f := func(x int) bool {
					return x%2 != 0
				}
				Expect(slices.Select(sliceD, f)).To(Equal([]int{9, 5, 3, 7, 7, 1}))
			})
		})
	})

	Describe("tests for Any()", func() {
		Context("with a nil slice", func() {
			It("should return false", func() {
				Expect(slices.Any(nilSlice, func(int) bool { return true })).To(BeFalse())
			})
		})
		Context("with an empty slice", func() {
			It("should return false", func() {
				Expect(slices.Any(emptySlice, func(int) bool { return true })).To(BeFalse())
			})
		})
		Context("with a more complex slice", func() {
			It("should detect that slice contains a number greater than 2", func() {
				f := func(x int) bool {
					return x > 2
				}
				Expect(slices.Any(sliceB, f)).To(BeTrue())
			})
			It("should detect that slice contains no number greater than 3", func() {
				f := func(x int) bool {
					return x > 3
				}
				Expect(slices.Any(sliceB, f)).To(BeFalse())
			})
		})
	})

	Describe("tests for All()", func() {
		Context("with a nil slice", func() {
			It("should return true", func() {
				Expect(slices.All(nilSlice, func(int) bool { return false })).To(BeTrue())
			})
		})
		Context("with an empty slice", func() {
			It("should return true", func() {
				Expect(slices.All(emptySlice, func(int) bool { return false })).To(BeTrue())
			})
		})
		Context("with a more complex slice", func() {
			It("should detect that slice contains only numbers smaller than 4", func() {
				f := func(x int) bool {
					return x < 4
				}
				Expect(slices.All(sliceB, f)).To(BeTrue())
			})
			It("should detect that slice does not only contain numbers smaller than 3", func() {
				f := func(x int) bool {
					return x < 3
				}
				Expect(slices.All(sliceB, f)).To(BeFalse())
			})
		})
	})

	Describe("tests for None()", func() {
		Context("with a nil slice", func() {
			It("should return true", func() {
				Expect(slices.None(nilSlice, func(int) bool { return false })).To(BeTrue())
			})
		})
		Context("with an empty slice", func() {
			It("should return true", func() {
				Expect(slices.None(emptySlice, func(int) bool { return false })).To(BeTrue())
			})
		})
		Context("with a more complex slice", func() {
			It("should detect that slice contains no numbers greater than 3", func() {
				f := func(x int) bool {
					return x > 3
				}
				Expect(slices.None(sliceB, f)).To(BeTrue())
			})
			It("should detect that slice does not contain no numbers greater than 2", func() {
				f := func(x int) bool {
					return x > 2
				}
				Expect(slices.None(sliceB, f)).To(BeFalse())
			})
		})
	})

	Describe("tests for Count()", func() {
		Context("with a nil slice", func() {
			It("should return zero", func() {
				Expect(slices.Count(nilSlice, func(int) bool { return true })).To(Equal(0))
			})
		})
		Context("with an empty slice", func() {
			It("should return zero", func() {
				Expect(slices.Count(emptySlice, func(int) bool { return true })).To(Equal(0))
			})
		})
		Context("with a more complex slice", func() {
			It("should detect that slice contains 4 numbers greater than 6", func() {
				f := func(x int) bool {
					return x > 6
				}
				Expect(slices.Count(sliceD, f)).To(Equal(4))
			})
		})
	})
})
