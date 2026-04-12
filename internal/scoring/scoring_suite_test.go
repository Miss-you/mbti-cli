package scoring

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScoring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scoring Suite")
}
