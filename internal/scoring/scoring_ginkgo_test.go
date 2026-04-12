package scoring

import (
	"errors"

	"github.com/Miss-you/mbti-cli/internal/answers"
	"github.com/Miss-you/mbti-cli/internal/questionbank"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scoring engine", func() {
	bank := func() questionbank.Bank {
		return questionbank.Bank{
			Questions: []questionbank.Question{
				scoringSpecQuestion("q01", questionbank.DimensionEI, false),
				scoringSpecQuestion("q02", questionbank.DimensionSN, true),
				scoringSpecQuestion("q03", questionbank.DimensionTF, false),
				scoringSpecQuestion("q04", questionbank.DimensionJP, false),
			},
		}
	}

	When("a complete strict answer set is scored", func() {
		It("adds selected option scores by dimension and reports answer counts", func() {
			result, err := Score(bank(), answers.Set{
				"q01": "A",
				"q02": "D",
				"q03": "B",
				"q04": "C",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Answered).To(Equal(4))
			Expect(result.Total).To(Equal(4))
			Expect(result.DimensionScores).To(Equal(map[questionbank.Dimension]int{
				questionbank.DimensionEI: 2,
				questionbank.DimensionSN: -2,
				questionbank.DimensionTF: -1,
				questionbank.DimensionJP: 1,
			}))
		})
	})

	When("a selected option belongs to a reversed question", func() {
		It("adds the stored signed score without applying reverse again", func() {
			result, err := Score(bank(), answers.Set{
				"q01": "A",
				"q02": "D",
				"q03": "A",
				"q04": "A",
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Answered).To(Equal(4))
			Expect(result.Total).To(Equal(4))
			Expect(result.DimensionScores[questionbank.DimensionSN]).To(Equal(-2))
			Expect(result.DimensionScores[questionbank.DimensionEI]).To(Equal(2))
		})
	})

	When("answers are incomplete", func() {
		It("returns a validation error without partial dimension totals", func() {
			result, err := Score(bank(), answers.Set{
				"q01": "A",
				"q02": "D",
				"q03": "B",
			})

			Expect(err).To(HaveOccurred())
			var validationErr *answers.ValidationError
			Expect(errors.As(err, &validationErr)).To(BeTrue())
			Expect(validationErr.Issues).To(Equal([]string{"missing answer for question q04"}))
			Expect(result).To(BeZero())
		})
	})
})

func scoringSpecQuestion(id string, dimension questionbank.Dimension, reverse bool) questionbank.Question {
	return questionbank.Question{
		ID:        id,
		Dimension: dimension,
		Reverse:   reverse,
		Options: []questionbank.Option{
			{Code: "A", Score: 2},
			{Code: "B", Score: -1},
			{Code: "C", Score: 1},
			{Code: "D", Score: -2},
		},
	}
}
