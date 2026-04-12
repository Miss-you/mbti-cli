package answers

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnswers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Answers Suite")
}
