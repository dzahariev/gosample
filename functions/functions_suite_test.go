package functions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFunctions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Functions Suite")
}
