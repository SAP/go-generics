package maps_test

import (
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sap/go-generics/maps"
)

func TestMaps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maps Suite")
}

var _ = Describe("maps", func() {
	var nilMap map[int]string
	var emptyMap map[int]string
	var mapA map[int]string
	var mapB map[int]string
	var mapC map[int]float64
	var mapD map[int]string

	BeforeEach(func() {
		nilMap = nil
		emptyMap = map[int]string{}
		mapA = map[int]string{1: "a", 2: "b", 3: "b"}
		mapB = map[int]string{1: "u", 2: "v", 3: "w", 4: "w"}
		mapC = map[int]float64{1: 2.1, 2: 3.14}
		mapD = map[int]string{1: "2.1", 2: "3.14"}
	})

	AfterEach(func() {
		Expect(nilMap).To(BeNil())
		Expect(emptyMap).To(Equal(map[int]string{}))
		Expect(mapA).To(Equal(map[int]string{1: "a", 2: "b", 3: "b"}))
		Expect(mapB).To(Equal(map[int]string{1: "u", 2: "v", 3: "w", 4: "w"}))
		Expect(mapC).To(Equal(map[int]float64{1: 2.1, 2: 3.14}))
		Expect(mapD).To(Equal(map[int]string{1: "2.1", 2: "3.14"}))
	})

	Describe("tests for Keys()", func() {
		Context("with a nil map", func() {
			It("should return nil", func() {
				Expect(maps.Keys(nilMap)).To(Equal([]int(nil)))
			})
		})
		Context("with an empty map", func() {
			It("should return an empty slice", func() {
				Expect(maps.Keys(emptyMap)).To(Equal([]int{}))
			})
		})
		Context("with a more cmoplex map", func() {
			It("should return a slice containing the map keys", func() {
				Expect(maps.Keys(mapA)).To(ConsistOf(1, 2, 3))
			})
		})
	})

	Describe("tests for Values()", func() {
		Context("with a nil map", func() {
			It("should return nil", func() {
				Expect(maps.Values(nilMap)).To(Equal([]string(nil)))
			})
		})
		Context("with an empty map", func() {
			It("should return an empty slice", func() {
				Expect(maps.Values(emptyMap)).To(Equal([]string{}))
			})
		})
		Context("with a more cmoplex map", func() {
			It("should return a slice containing the map values", func() {
				Expect(maps.Values(mapA)).To(ConsistOf("a", "b", "b"))
			})
		})
	})

	Describe("tests for EqualBy()", func() {
		Context("with different lengths", func() {
			It("should return false", func() {
				Expect(maps.EqualBy(mapA, mapB, func(string, string) bool { return true })).To(BeFalse())
			})
		})
		Context("with nil/nil maps", func() {
			It("should return true", func() {
				Expect(maps.EqualBy(nilMap, nilMap, func(string, string) bool { return false })).To(BeTrue())
			})
		})
		Context("with empty/empty maps", func() {
			It("should return true", func() {
				Expect(maps.EqualBy(emptyMap, emptyMap, func(string, string) bool { return false })).To(BeTrue())
			})
		})
		Context("with nil/empty maps", func() {
			It("should return true", func() {
				Expect(maps.EqualBy(nilMap, emptyMap, func(string, string) bool { return false })).To(BeTrue())
			})
		})
		Context("with more complex maps", func() {
			It("should return true", func() {
				f := func(x float64, s string) bool {
					y, err := strconv.ParseFloat(s, 64)
					Expect(err).NotTo(HaveOccurred())
					return x == y
				}
				Expect(maps.EqualBy(mapC, mapD, f)).To(BeTrue())
			})
		})
	})

	Describe("tests for Equal()", func() {
		Context("with different lengths", func() {
			It("should return false", func() {
				Expect(maps.Equal(mapA, mapB)).To(BeFalse())
			})
		})
		Context("with nil/nil maps", func() {
			It("should return true", func() {
				Expect(maps.Equal(nilMap, nilMap)).To(BeTrue())
			})
		})
		Context("with empty/empty maps", func() {
			It("should return true", func() {
				Expect(maps.Equal(emptyMap, emptyMap)).To(BeTrue())
			})
		})
		Context("with nil/empty maps", func() {
			It("should return true", func() {
				Expect(maps.Equal(nilMap, emptyMap)).To(BeTrue())
			})
		})
		Context("with identical maps", func() {
			It("should return true", func() {
				Expect(maps.Equal(mapB, map[int]string{2: "v", 1: "u", 4: "w", 3: "w"})).To(BeTrue())
			})
		})
		Context("with different maps (different keys)", func() {
			It("should return false", func() {
				Expect(maps.Equal(mapB, map[int]string{2: "v", 1: "u", 4: "w", 5: "w"})).To(BeFalse())
			})
		})
		Context("with different maps (different values)", func() {
			It("should return false", func() {
				Expect(maps.Equal(mapB, map[int]string{2: "v", 1: "u", 4: "w", 3: "x"})).To(BeFalse())
			})
		})
	})

	Describe("tests for Collect()", func() {
		Context("with a nil map", func() {
			It("should return nil", func() {
				Expect(maps.Collect(nilMap, func(string) string { return "" })).To(Equal(nilMap))
			})
		})
		Context("with an empty map", func() {
			It("should return en empty map", func() {
				Expect(maps.Collect(emptyMap, func(string) string { return "" })).To(Equal(emptyMap))
			})
		})
		Context("with a more complex map", func() {
			It("should return a map containing the mapped valuese", func() {
				f := func(x float64) int {
					return int(x)
				}
				Expect(maps.Collect(mapC, f)).To(Equal(map[int]int{1: 2, 2: 3}))
			})
		})
	})
})
