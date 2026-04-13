package scoring

import (
	"github.com/Miss-you/mbti-cli/internal/questionbank"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scoring classifier", func() {
	When("dimension scores land on threshold boundaries", func() {
		DescribeTable("labels the boundary score with the matching strength while preserving the raw score",
			func(score int, strength questionbank.Strength, pole string, letter string) {
				classification, err := Classify(classifierSpecBank(), Result{
					DimensionScores: map[questionbank.Dimension]int{
						questionbank.DimensionEI: score,
					},
				})

				Expect(err).NotTo(HaveOccurred())
				got := classification.Dimensions[questionbank.DimensionEI]
				Expect(got.Score).To(Equal(score))
				Expect(got.Strength).To(Equal(strength))
				Expect(got.Pole).To(Equal(pole))
				Expect(got.Letter).To(Equal(letter))
				Expect(got.Balanced).To(BeFalse())
			},
			Entry("strong A lower bound", 13, questionbank.StrengthStrongA, "E (Expansive)", "E"),
			Entry("strong A upper bound", 999, questionbank.StrengthStrongA, "E (Expansive)", "E"),
			Entry("moderate A lower bound", 5, questionbank.StrengthModerateA, "E (Expansive)", "E"),
			Entry("moderate A upper bound", 12, questionbank.StrengthModerateA, "E (Expansive)", "E"),
			Entry("slight A lower bound", 1, questionbank.StrengthSlightA, "E (Expansive)", "E"),
			Entry("slight A upper bound", 4, questionbank.StrengthSlightA, "E (Expansive)", "E"),
			Entry("slight B lower bound", -4, questionbank.StrengthSlightB, "I (Focused)", "I"),
			Entry("slight B upper bound", -1, questionbank.StrengthSlightB, "I (Focused)", "I"),
			Entry("moderate B lower bound", -12, questionbank.StrengthModerateB, "I (Focused)", "I"),
			Entry("moderate B upper bound", -5, questionbank.StrengthModerateB, "I (Focused)", "I"),
			Entry("strong B lower bound", -999, questionbank.StrengthStrongB, "I (Focused)", "I"),
			Entry("strong B upper bound", -13, questionbank.StrengthStrongB, "I (Focused)", "I"),
		)
	})

	When("raw dimension scores mix positive, negative, and zero values", func() {
		It("generates the type in EI/SN/TF/JP order and marks zero as balanced", func() {
			classification, err := Classify(classifierSpecBank(), Result{
				DimensionScores: map[questionbank.Dimension]int{
					questionbank.DimensionJP: -2,
					questionbank.DimensionEI: 2,
					questionbank.DimensionSN: 0,
					questionbank.DimensionTF: 1,
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(classification.Type).To(Equal("EXTP"))

			ei := classification.Dimensions[questionbank.DimensionEI]
			Expect(ei.Score).To(Equal(2))
			Expect(ei.Pole).To(Equal("E (Expansive)"))
			Expect(ei.Letter).To(Equal("E"))
			Expect(ei.Balanced).To(BeFalse())

			sn := classification.Dimensions[questionbank.DimensionSN]
			Expect(sn.Score).To(Equal(0))
			Expect(sn.Pole).To(Equal("balanced"))
			Expect(sn.Letter).To(Equal("X"))
			Expect(sn.Balanced).To(BeTrue())

			tf := classification.Dimensions[questionbank.DimensionTF]
			Expect(tf.Score).To(Equal(1))
			Expect(tf.Pole).To(Equal("T (Analytical)"))
			Expect(tf.Letter).To(Equal("T"))
			Expect(tf.Balanced).To(BeFalse())

			jp := classification.Dimensions[questionbank.DimensionJP]
			Expect(jp.Score).To(Equal(-2))
			Expect(jp.Pole).To(Equal("P (Flexible)"))
			Expect(jp.Letter).To(Equal("P"))
			Expect(jp.Balanced).To(BeFalse())
		})
	})

	When("required classifier metadata is unavailable", func() {
		It("returns a contextual error for a missing threshold range", func() {
			bank := classifierSpecBank()
			delete(bank.Meta.Scoring.Thresholds, questionbank.StrengthSlightA)

			_, err := Classify(bank, Result{
				DimensionScores: map[questionbank.Dimension]int{
					questionbank.DimensionEI: 1,
				},
			})

			Expect(err).To(MatchError(ContainSubstring("classify EI score 1")))
			Expect(err).To(MatchError(ContainSubstring("no threshold range")))
		})

		It("returns a contextual error for missing dimension metadata", func() {
			bank := classifierSpecBank()
			delete(bank.Meta.Dimensions, questionbank.DimensionEI)

			_, err := Classify(bank, Result{
				DimensionScores: map[questionbank.Dimension]int{
					questionbank.DimensionEI: 1,
				},
			})

			Expect(err).To(MatchError(ContainSubstring("missing dimension metadata for EI")))
		})
	})
})

func classifierSpecBank() questionbank.Bank {
	return questionbank.Bank{
		Meta: questionbank.Meta{
			Dimensions: map[questionbank.Dimension]questionbank.DimensionMeta{
				questionbank.DimensionEI: {PoleA: "E (Expansive)", PoleB: "I (Focused)"},
				questionbank.DimensionSN: {PoleA: "S (Concrete)", PoleB: "N (Abstract)"},
				questionbank.DimensionTF: {PoleA: "T (Analytical)", PoleB: "F (Empathetic)"},
				questionbank.DimensionJP: {PoleA: "J (Structured)", PoleB: "P (Flexible)"},
			},
			Scoring: questionbank.ScoringMeta{
				Thresholds: map[questionbank.Strength]questionbank.ThresholdRange{
					questionbank.StrengthStrongA:   {13, 999},
					questionbank.StrengthModerateA: {5, 12},
					questionbank.StrengthSlightA:   {1, 4},
					questionbank.StrengthSlightB:   {-4, -1},
					questionbank.StrengthModerateB: {-12, -5},
					questionbank.StrengthStrongB:   {-999, -13},
				},
			},
		},
	}
}
