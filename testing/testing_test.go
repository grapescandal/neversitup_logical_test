package testing_test

import (
	"neversitup_logical_test/testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing", func() {
	Describe("Test_1", func() {
		It("should pass", func() {
			input := "17283"
			isPass := testing.LengthCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "172839"
			isPass = testing.LengthCheck(input)

			Expect(isPass).Should(BeTrue())
		})
	})

	Describe("Test_2", func() {
		It("should pass", func() {
			input := "118822"
			isPass := testing.AdjacentDuplicateCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "111762"
			isPass = testing.AdjacentDuplicateCheck(input)

			Expect(isPass).Should(BeTrue())
		})
	})

	Describe("Test_3", func() {
		It("should pass", func() {
			input := "123743"
			isPass := testing.SortingCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "321895"
			isPass = testing.SortingCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "124578"
			isPass = testing.SortingCheck(input)

			Expect(isPass).Should(BeTrue())
		})
	})

	Describe("Test_4", func() {
		It("should pass", func() {
			input := "112233"
			isPass := testing.DuplicateCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "882211"
			isPass = testing.DuplicateCheck(input)

			Expect(isPass).Should(BeFalse())

			input = "887712"
			isPass = testing.DuplicateCheck(input)

			Expect(isPass).Should(BeTrue())
		})
	})
})
