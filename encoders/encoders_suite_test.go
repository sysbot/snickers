package encoders

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEncoders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Encoders Suite")
}
