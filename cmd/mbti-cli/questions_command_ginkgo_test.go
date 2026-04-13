package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Questions command", func() {
	When("the canonical bank is exported as JSON", func() {
		It("writes parseable prompt data without scoring internals", func() {
			stdout, stderr, err := runRoot(
				"questions",
				"--questions", canonicalBankPath(),
				"--format", "json",
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(stderr).To(BeEmpty())
			Expect(stdout).To(HaveSuffix("\n"))

			var got questionSetOutput
			Expect(json.Unmarshal([]byte(stdout), &got)).To(Succeed())
			Expect(got.Meta.Title).To(Equal("AI Behavioral Style Assessment v3"))
			Expect(got.Meta.Version).To(Equal("0.3.0"))
			Expect(got.Meta.Language).To(Equal("zh"))
			Expect(got.Meta.Count).To(Equal(70))
			Expect(got.Meta.Total).To(Equal(70))
			Expect(got.Meta.Source).To(Equal(canonicalBankPath()))
			Expect(got.Questions).To(HaveLen(70))
			Expect(got.Questions[0].ID).To(Equal("q01"))
			Expect(got.Questions[0].Scenario).To(ContainSubstring("用户让你帮忙"))
			Expect(got.Questions[0].Options).To(HaveLen(4))
			Expect(got.Questions[0].Options[0].Code).To(Equal("A"))
			Expect(got.Questions[0].Options[0].Label).To(ContainSubstring("主动标注文化梗"))

			lower := strings.ToLower(stdout)
			Expect(lower).NotTo(ContainSubstring("\"score\""))
			Expect(lower).NotTo(ContainSubstring("\"reverse\""))
			Expect(lower).NotTo(ContainSubstring("\"thresholds\""))
			Expect(lower).NotTo(ContainSubstring("\"dimensions\""))
		})
	})

	When("a bounded English text export is requested", func() {
		It("writes one localized question with option labels", func() {
			stdout, stderr, err := runRoot(
				"questions",
				"--questions", canonicalBankPath(),
				"--format", "text",
				"--count", "1",
				"--lang", "en",
			)

			Expect(err).NotTo(HaveOccurred())
			Expect(stderr).To(BeEmpty())
			Expect(stdout).To(Equal(readGoldenFile("questions-one-en.txt.golden")))
		})
	})

	When("the same seed and count are used twice", func() {
		It("selects the same question IDs in the same order", func() {
			firstStdout, firstStderr, err := runRoot(
				"questions",
				"--questions", canonicalBankPath(),
				"--format", "json",
				"--count", "3",
				"--seed", "7",
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(firstStderr).To(BeEmpty())

			secondStdout, secondStderr, err := runRoot(
				"questions",
				"--questions", canonicalBankPath(),
				"--format", "json",
				"--count", "3",
				"--seed", "7",
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(secondStderr).To(BeEmpty())

			Expect(extractQuestionIDs(firstStdout)).To(Equal(extractQuestionIDs(secondStdout)))
			Expect(extractQuestionIDs(firstStdout)).To(HaveLen(3))
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
		Entry("unsupported format", func() []string { return []string{"questions", "--questions", canonicalBankPath(), "--format", "yaml"} }, "unsupported format"),
		Entry("unsupported language", func() []string { return []string{"questions", "--questions", canonicalBankPath(), "--lang", "fr"} }, "unsupported language"),
		Entry("negative count", func() []string { return []string{"questions", "--questions", canonicalBankPath(), "--count", "-1"} }, "count must be non-negative"),
		Entry("count too large", func() []string { return []string{"questions", "--questions", canonicalBankPath(), "--count", "71"} }, "count 71 exceeds available questions 70"),
		Entry("missing path", func() []string { return []string{"questions"} }, "question bank path is required"),
		Entry("missing file", func() []string {
			return []string{"questions", "--questions", filepath.Join(GinkgoT().TempDir(), "missing.json")}
		}, "read question bank"),
		Entry("invalid bank content", func() []string {
			return []string{"questions", "--questions", invalidQuestionBankPath("invalid-meta-total.json"), "--format", "json"}
		}, "validate question bank"),
	)
})

func invalidQuestionBankPath(name string) string {
	return filepath.Join("..", "..", "internal", "questionbank", "testdata", name)
}

func runRoot(args ...string) (string, string, error) {
	cmd := NewRootCmd()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return stdout.String(), stderr.String(), err
}

func readGoldenFile(name string) string {
	data, err := os.ReadFile(cliTestdataPath(name))
	Expect(err).NotTo(HaveOccurred())
	return string(data)
}

func extractQuestionIDs(stdout string) []string {
	var output questionSetOutput
	Expect(json.Unmarshal([]byte(stdout), &output)).To(Succeed())

	ids := make([]string, 0, len(output.Questions))
	for _, question := range output.Questions {
		ids = append(ids, question.ID)
	}
	return ids
}
