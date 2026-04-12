package questionbank

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQuestionbank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Questionbank Suite")
}
