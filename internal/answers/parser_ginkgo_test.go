package answers

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	When("the answers object contains option codes with extra whitespace or lowercase letters", func() {
		It("normalizes option codes before returning the answer set", func() {
			got, err := Parse([]byte(`{
				"answers": {
					"q01": "a",
					"q02": " C "
				}
			}`))

			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(Set{
				"q01": "A",
				"q02": "C",
			}))
		})
	})

	When("the answers object is missing or null", func() {
		DescribeTable("returns an answers-required error",
			func(input string) {
				_, err := Parse([]byte(input))

				Expect(err).To(MatchError(And(
					ContainSubstring("parse answer file"),
					ContainSubstring("answers object is required"),
				)))
			},
			Entry("missing answers object", `{}`),
			Entry("null answers object", `{"answers": null}`),
		)
	})

	When("the JSON cannot be decoded into an answer file", func() {
		It("wraps malformed JSON errors with parser context", func() {
			_, err := Parse([]byte(`{"answers":`))

			Expect(err).To(MatchError(ContainSubstring("parse answer file")))
		})

		It("wraps structurally invalid answer values with parser context", func() {
			_, err := Parse([]byte(`{"answers":{"q01":1}}`))

			Expect(err).To(MatchError(ContainSubstring("parse answer file")))
		})
	})

	When("answers reference unknown questions or option codes", func() {
		It("normalizes the data and leaves bank-aware checks to validation", func() {
			got, err := Parse([]byte(`{"answers":{"unknown":"z"}}`))

			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(Set{"unknown": "Z"}))
		})
	})
})
