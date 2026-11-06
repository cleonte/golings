package exercises_test

import (
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("The hints of an exercise", func() {
	It("goes from the nudge to the full hint", func() {
		ex := exercises.Exercise{Nudge: "a", Stronger: "b", Hint: "c"}

		Expect(ex.HintLevels()).To(Equal([]string{"a", "b", "c"}))
	})
	It("leaves out the levels that are not there", func() {
		Expect(exercises.Exercise{Hint: "c"}.HintLevels()).To(Equal([]string{"c"}))
		Expect(exercises.Exercise{Nudge: "a", Hint: "c"}.HintLevels()).To(Equal([]string{"a", "c"}))
		Expect(exercises.Exercise{}.HintLevels()).To(BeEmpty())
	})
	It("trims the text", func() {
		Expect(exercises.Exercise{Hint: "  c\n\n"}.HintLevels()).To(Equal([]string{"c"}))
	})
})
