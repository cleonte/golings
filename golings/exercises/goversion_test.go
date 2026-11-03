package exercises_test

import (
	"github.com/cleonte/golings/golings/exercises"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// The exercises that need Go 1.27 have to be skipped, cleanly, on Go 1.26 (the
// project floor) and run on 1.27. These specs pretend to be each Go in turn,
// against the real manifest, so they hold whichever Go runs the tests.
var _ = Describe("The Go 1.27 exercises", func() {
	const infoFile = "../../info.toml"

	var all, newer []exercises.Exercise

	pretendToBeGo := func(local string) {
		previous := exercises.LocalGoVersion
		exercises.LocalGoVersion = func() (string, error) { return local, nil }
		DeferCleanup(func() { exercises.LocalGoVersion = previous })
	}

	BeforeEach(func() {
		var err error
		all, err = exercises.List(infoFile)
		Expect(err).NotTo(HaveOccurred())

		newer = nil
		for _, ex := range all {
			if ex.GoVersion == "1.27" {
				newer = append(newer, ex)
			}
		}
		Expect(newer).NotTo(BeEmpty())
	})

	Context("on Go 1.26", func() {
		BeforeEach(func() { pretendToBeGo("go1.26.10") })

		It("skips every one of them and says why", func() {
			for _, ex := range newer {
				Expect(ex.SkipReason()).To(Equal("needs Go 1.27"), "exercise %s", ex.Name)
			}
		})
		It("does not build or run them", func() {
			for _, ex := range newer {
				result, err := ex.Run()

				Expect(err).To(MatchError(exercises.ErrGoVersionTooOld), "exercise %s", ex.Name)
				Expect(exercises.IsSkipped(err)).To(BeTrue())
				Expect(result.BuildFailed).To(BeFalse(), "exercise %s was built", ex.Name)
			}
		})
		It("lists them as skipped and leaves them out of the progress", func() {
			skipped, err := exercises.Skipped(infoFile)
			Expect(err).NotTo(HaveOccurred())
			var names []string
			for _, ex := range skipped {
				names = append(names, ex.Name)
			}
			for _, ex := range newer {
				Expect(names).To(ContainElement(ex.Name))
			}

			_, _, total, err := exercises.Progress(infoFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(total + len(skipped)).To(Equal(len(all)))
		})
	})

	Context("on Go 1.27", func() {
		BeforeEach(func() { pretendToBeGo("go1.27.3") })

		It("runs them", func() {
			for _, ex := range newer {
				Expect(ex.SkipReason()).To(BeEmpty(), "exercise %s", ex.Name)
			}
		})
	})
})
