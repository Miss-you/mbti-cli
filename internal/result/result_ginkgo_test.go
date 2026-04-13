package result

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Miss-you/mbti-cli/internal/questionbank"
)

var _ = Describe("Result rendering", func() {
	Describe("NewSummary", func() {
		It("preserves summary metadata and normalizes balanced dimensions", func() {
			summary, err := NewSummary(resultBank(), resultScore(), resultClassification())

			Expect(err).NotTo(HaveOccurred())
			Expect(summary.Meta.Title).To(Equal("AI Behavioral Style Assessment v3"))
			Expect(summary.Meta.Version).To(Equal("0.3.0"))
			Expect(summary.Meta.Answered).To(Equal(70))
			Expect(summary.Meta.Total).To(Equal(70))
			Expect(summary.Type).To(Equal("EXTP"))
			Expect(summary.Dimensions.EI).To(Equal(Dimension{
				Letter:   "E",
				Score:    8,
				Strength: string(questionbank.StrengthModerateA),
				Pole:     "E (Expansive)",
				Balanced: false,
			}))
			Expect(summary.Dimensions.SN).To(Equal(Dimension{
				Letter:   "X",
				Score:    0,
				Strength: "balanced",
				Pole:     "balanced",
				Balanced: true,
			}))
			Expect(summary.Dimensions.TF).To(Equal(Dimension{
				Letter:   "T",
				Score:    1,
				Strength: string(questionbank.StrengthSlightA),
				Pole:     "T (Analytical)",
				Balanced: false,
			}))
			Expect(summary.Dimensions.JP).To(Equal(Dimension{
				Letter:   "P",
				Score:    -2,
				Strength: string(questionbank.StrengthSlightB),
				Pole:     "P (Flexible)",
				Balanced: false,
			}))
		})

		It("returns a contextual error when a required dimension classification is missing", func() {
			classification := resultClassification()
			delete(classification.Dimensions, questionbank.DimensionSN)

			_, err := NewSummary(resultBank(), resultScore(), classification)

			Expect(err).To(MatchError(ContainSubstring("missing classification for SN")))
		})
	})

	Describe("RenderJSON", func() {
		It("renders the stable golden output with a trailing newline", func() {
			summary, err := NewSummary(resultBank(), resultScore(), resultClassification())
			Expect(err).NotTo(HaveOccurred())

			got, err := RenderJSON(summary)

			Expect(err).NotTo(HaveOccurred())
			Expect(string(got)).To(Equal(readGoldenFile("summary.json.golden")))
			Expect(string(got)).To(HaveSuffix("\n"))
		})
	})

	Describe("RenderText", func() {
		It("renders factual text without diagnostic or personality claims", func() {
			summary, err := NewSummary(resultBank(), resultScore(), resultClassification())
			Expect(err).NotTo(HaveOccurred())

			got := RenderText(summary)

			Expect(got).To(Equal(readGoldenFile("summary.txt.golden")))
			lower := strings.ToLower(got)
			Expect(lower).NotTo(ContainSubstring("diagnosis"))
			Expect(lower).NotTo(ContainSubstring("personality"))
			Expect(lower).NotTo(ContainSubstring("you are"))
		})
	})
})

func readGoldenFile(name string) string {
	data, err := os.ReadFile(filepath.Join("testdata", name))
	Expect(err).NotTo(HaveOccurred())
	return string(data)
}
