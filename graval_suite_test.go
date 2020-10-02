package graval_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGraval(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graval Suite")
}
