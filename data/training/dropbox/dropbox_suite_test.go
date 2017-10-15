package dropbox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDropbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dropbox Suite")
}
