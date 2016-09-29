package goping_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goping Suite")
}
