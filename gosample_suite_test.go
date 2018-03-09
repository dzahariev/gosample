package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGosample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gosample Suite")
}
