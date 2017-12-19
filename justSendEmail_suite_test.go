package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJustSendEmail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JustSendEmail Suite")
}
