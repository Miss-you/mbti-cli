package questionbank

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Question bank loader", func() {
	canonicalV3Path := filepath.Join("..", "..", "questions", "mbti-questions-v3.json")

	When("the canonical v3 question bank path is loaded", func() {
		It("returns the typed bank with source metadata", func() {
			loaded, err := LoadFile(canonicalV3Path)

			Expect(err).NotTo(HaveOccurred())
			Expect(loaded.Bank.Meta.Title).To(Equal("AI Behavioral Style Assessment v3"))
			Expect(loaded.Bank.Meta.Version).To(Equal("0.3.0"))
			Expect(loaded.Bank.Questions[0].ID).To(Equal("q01"))
			Expect(loaded.Source.Path).To(Equal(canonicalV3Path))
			Expect(loaded.Source.FileName).To(Equal("mbti-questions-v3.json"))
			Expect(loaded.Source.SizeBytes).To(BeNumerically(">", 0))
		})
	})

	When("the path does not exist", func() {
		It("returns a read error with path context and preserves the missing-file cause", func() {
			missingPath := filepath.Join(GinkgoT().TempDir(), "missing-bank.json")

			_, err := LoadFile(missingPath)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("read question bank")))
			Expect(err).To(MatchError(ContainSubstring(missingPath)))
			Expect(errors.Is(err, os.ErrNotExist)).To(BeTrue())
		})
	})

	When("the file is not valid question bank JSON", func() {
		It("returns a parse error with source path context", func() {
			dir := GinkgoT().TempDir()
			malformedPath := filepath.Join(dir, "malformed-bank.json")
			Expect(os.WriteFile(malformedPath, []byte(`{"meta":`), 0o600)).To(Succeed())

			_, err := LoadFile(malformedPath)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("parse question bank")))
			Expect(err).To(MatchError(ContainSubstring(malformedPath)))
		})
	})

	When("the path is empty", func() {
		It("returns a path-required error", func() {
			_, err := LoadFile("")

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("question bank path is required"))
		})
	})
})
