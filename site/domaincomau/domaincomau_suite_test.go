package domaincomau_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"testing"
)

var DomainComAuPropertyProfile string
var DomainComAuPropertyProfileWithMultipleSoldAndRented string

func TestDomaincomau(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domaincomau Suite")
}

var _ = BeforeEach(func() {
	contents, err := ioutil.ReadFile("test_assets/property-profile.html")
	Expect(err).ToNot(HaveOccurred())
	DomainComAuPropertyProfile = string(contents)

	contents, err = ioutil.ReadFile("test_assets/property-profile-multiple-sold-and-rent.html")
	Expect(err).ToNot(HaveOccurred())
	DomainComAuPropertyProfileWithMultipleSoldAndRented = string(contents)
})
