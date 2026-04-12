package answers

import (
	"errors"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Answer validation", func() {
	answerBank := func() questionbank.Bank {
		return questionbank.Bank{
			Questions: []questionbank.Question{
				{
					ID: "q01",
					Options: []questionbank.Option{
						{Code: "A"},
						{Code: "B"},
						{Code: "C"},
						{Code: "D"},
					},
				},
				{
					ID: "q02",
					Options: []questionbank.Option{
						{Code: "A"},
						{Code: "B"},
						{Code: "C"},
						{Code: "D"},
					},
				},
				{
					ID: "q03",
					Options: []questionbank.Option{
						{Code: "A"},
						{Code: "B"},
						{Code: "C"},
						{Code: "D"},
					},
				},
			},
		}
	}

	When("every bank question has a defined selected option", func() {
		It("accepts the complete strict answer set", func() {
			err := Validate(answerBank(), Set{
				"q01": "A",
				"q02": "D",
				"q03": "B",
			})

			Expect(err).NotTo(HaveOccurred())
		})
	})

	When("a strict answer set contains a single caller-visible validation issue", func() {
		DescribeTable("rejects the bad answer without normalizing it",
			func(answerSet Set, expectedIssue string) {
				err := Validate(answerBank(), answerSet)

				Expect(err).To(HaveOccurred())
				var validationErr *ValidationError
				Expect(errors.As(err, &validationErr)).To(BeTrue())
				Expect(validationErr.Issues).To(Equal([]string{expectedIssue}))
			},
			Entry("unknown question id", Set{
				"q01":     "A",
				"q02":     "A",
				"q03":     "A",
				"unknown": "A",
			}, "unknown question id unknown"),
			Entry("invalid option code", Set{
				"q01": "Z",
				"q02": "A",
				"q03": "A",
			}, "question q01 option Z is not valid"),
			Entry("missing answer", Set{
				"q01": "A",
				"q03": "A",
			}, "missing answer for question q02"),
			Entry("lowercase non-canonical input", Set{
				"q01": "a",
				"q02": "A",
				"q03": "A",
			}, "question q01 option a is not valid"),
		)
	})

	When("unknown ids, invalid option codes, and missing answers appear together", func() {
		It("returns all validation issues in stable order", func() {
			answerSet := Set{
				"q03":   "x",
				"later": "A",
				"alpha": "B",
			}
			wantIssues := []string{
				"unknown question id alpha",
				"unknown question id later",
				"question q03 option x is not valid",
				"missing answer for question q01",
				"missing answer for question q02",
			}

			for range 5 {
				err := Validate(answerBank(), answerSet)

				Expect(err).To(HaveOccurred())
				var validationErr *ValidationError
				Expect(errors.As(err, &validationErr)).To(BeTrue())
				Expect(validationErr.Issues).To(Equal(wantIssues))
			}
		})
	})
})
