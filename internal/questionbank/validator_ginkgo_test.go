package questionbank

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type invalidValidationCase struct {
	mutate    func(*Bank)
	wantIssue func(Bank) string
}

var _ = Describe("Question bank validator", func() {
	When("the canonical v3 question bank JSON is unmarshaled directly into the typed model", func() {
		It("accepts the canonical bank without applying scoring behavior", func() {
			bank := loadCanonicalV3BankBDD()

			Expect(Validate(bank)).To(Succeed())
		})
	})

	DescribeTable("rejects representative schema violations with validation issues",
		func(tc invalidValidationCase) {
			bank := loadCanonicalV3BankBDD()
			tc.mutate(&bank)

			err := Validate(bank)

			Expect(err).To(HaveOccurred())
			validationErr, ok := err.(*ValidationError)
			Expect(ok).To(BeTrue())

			wantIssue := tc.wantIssue(bank)
			Expect(err.Error()).To(ContainSubstring(wantIssue))
			Expect(validationErr.Issues).To(ContainElement(wantIssue))
		},
		Entry("meta.total does not match the number of questions", invalidValidationCase{
			mutate: func(bank *Bank) {
				bank.Meta.Total++
			},
			wantIssue: func(bank Bank) string {
				return fmt.Sprintf("meta.total=%d does not match questions length %d", bank.Meta.Total, len(bank.Questions))
			},
		}),
		Entry("dimension metadata is missing for EI", invalidValidationCase{
			mutate: func(bank *Bank) {
				delete(bank.Meta.Dimensions, DimensionEI)
			},
			wantIssue: func(Bank) string {
				return "missing dimension metadata for EI"
			},
		}),
		Entry("question identity is missing", invalidValidationCase{
			mutate: func(bank *Bank) {
				bank.Questions[0].ID = ""
			},
			wantIssue: func(Bank) string {
				return "questions[0].id is required"
			},
		}),
		Entry("question localized scenario text is missing", invalidValidationCase{
			mutate: func(bank *Bank) {
				bank.Questions[0].Scenario.ZH = ""
			},
			wantIssue: func(Bank) string {
				return "questions[0].scenario.zh is required"
			},
		}),
		Entry("option score is not supported", invalidValidationCase{
			mutate: func(bank *Bank) {
				bank.Questions[0].Options[0].Score = 0
			},
			wantIssue: func(Bank) string {
				return "questions[0].options[0].score 0 is not supported"
			},
		}),
		Entry("a supported threshold bucket is missing", invalidValidationCase{
			mutate: func(bank *Bank) {
				delete(bank.Meta.Scoring.Thresholds, StrengthSlightA)
			},
			wantIssue: func(Bank) string {
				return "missing threshold slight_a"
			},
		}),
		Entry("threshold ranges overlap", invalidValidationCase{
			mutate: func(bank *Bank) {
				bank.Meta.Scoring.Thresholds[StrengthModerateA] = ThresholdRange{4, 12}
			},
			wantIssue: func(Bank) string {
				return "threshold moderate_a overlaps threshold slight_a"
			},
		}),
	)
})

func loadCanonicalV3BankBDD() Bank {
	data := mustReadCanonicalV3JSON()

	var bank Bank
	Expect(json.Unmarshal(data, &bank)).To(Succeed())
	return bank
}

func mustReadCanonicalV3JSON() []byte {
	data, err := os.ReadFile(filepath.Join("..", "..", "questions", "mbti-questions-v3.json"))
	Expect(err).NotTo(HaveOccurred())
	return data
}
