package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Score command", func() {
	When("complete canonical answers are scored as JSON", func() {
		It("writes the stable result summary with all four dimensions", func() {
			stdout, stderr, err := runRoot(
				"score",
				"--questions", canonicalBankPath(),
				"--answers", cliTestdataPath("answers-all-a.json"),
				"--format", "json",
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(stderr).To(BeEmpty())
			Expect(stdout).To(HaveSuffix("\n"))
			Expect(stdout).To(Equal(readGoldenFile("score-all-a.json.golden")))

			var got scoreJSONOutput
			Expect(json.Unmarshal([]byte(stdout), &got)).To(Succeed())
			Expect(got.Meta.Title).To(Equal("AI Behavioral Style Assessment v3"))
			Expect(got.Meta.Version).To(Equal("0.3.0"))
			Expect(got.Meta.Answered).To(Equal(70))
			Expect(got.Meta.Total).To(Equal(70))
			Expect(got.Type).To(Equal("ESTJ"))
			Expect(got.Dimensions.EI).To(Equal(scoreJSONDimension{Letter: "E", Score: 36, Strength: "strong_a", Pole: "E (Expansive)"}))
			Expect(got.Dimensions.SN).To(Equal(scoreJSONDimension{Letter: "S", Score: 36, Strength: "strong_a", Pole: "S (Concrete)"}))
			Expect(got.Dimensions.TF).To(Equal(scoreJSONDimension{Letter: "T", Score: 34, Strength: "strong_a", Pole: "T (Analytical)"}))
			Expect(got.Dimensions.JP).To(Equal(scoreJSONDimension{Letter: "J", Score: 34, Strength: "strong_a", Pole: "J (Structured)"}))
		})
	})

	When("complete canonical answers are scored as text", func() {
		It("writes the factual renderer output without diagnostic claims", func() {
			stdout, stderr, err := runRoot(
				"score",
				"--questions", canonicalBankPath(),
				"--answers", cliTestdataPath("answers-all-d.json"),
				"--format", "text",
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(stderr).To(BeEmpty())
			Expect(stdout).To(Equal(readGoldenFile("score-all-d.txt.golden")))

			lower := strings.ToLower(stdout)
			Expect(lower).NotTo(ContainSubstring("diagnosis"))
			Expect(lower).NotTo(ContainSubstring("personality"))
			Expect(lower).NotTo(ContainSubstring("you are"))
		})
	})

	DescribeTable("returns an error and leaves stdout empty when inputs are invalid",
		func(buildArgs func() []string, want string) {
			args := buildArgs()
			stdout, stderr, err := runRoot(args...)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring(want)))
			Expect(stderr).To(ContainSubstring(want))
			Expect(stdout).To(BeEmpty())
		},
		Entry("unsupported format", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", cliTestdataPath("answers-all-a.json"), "--format", "yaml"}
		}, "unsupported format"),
		Entry("missing question bank path", func() []string { return []string{"score"} }, "question bank path is required"),
		Entry("missing answer file path", func() []string { return []string{"score", "--questions", canonicalBankPath()} }, "answer file path is required"),
		Entry("missing answer file", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", filepath.Join(GinkgoT().TempDir(), "missing.json")}
		}, "read answer file"),
		Entry("invalid bank content", func() []string {
			return []string{"score", "--questions", invalidQuestionBankPath("invalid-meta-total.json"), "--answers", cliTestdataPath("answers-all-a.json"), "--format", "json"}
		}, "validate question bank"),
		Entry("malformed answer json", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", cliWriteAnswerFile([]byte("{"), "answers.json")}
		}, "parse answer file"),
		Entry("missing answer", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", cliWriteAnswersFile(map[string]string{"q01": "A"})}
		}, "missing answer for question q02"),
		Entry("unknown question id", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", cliWriteAnswersFile(map[string]string{"q999": "A"})}
		}, "unknown question id q999"),
		Entry("invalid option code", func() []string {
			return []string{"score", "--questions", canonicalBankPath(), "--answers", cliWriteAnswersFile(map[string]string{"q01": "Z"})}
		}, "question q01 option Z is not valid"),
	)
})

func cliWriteAnswerFile(data []byte, name string) string {
	path := filepath.Join(GinkgoT().TempDir(), name)
	Expect(os.WriteFile(path, data, 0o600)).To(Succeed())
	return path
}

func cliWriteAnswersFile(answers map[string]string) string {
	data, err := json.Marshal(map[string]map[string]string{
		"answers": answers,
	})
	Expect(err).NotTo(HaveOccurred())
	return cliWriteAnswerFile(data, "answers.json")
}
