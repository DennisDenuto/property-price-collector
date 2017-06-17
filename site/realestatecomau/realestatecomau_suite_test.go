package realestatecomau_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"testing"
)

func TestRealestatecomau(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Realestatecomau Suite")
}

var ReadRealEstateComAu_Buy_list_1 string

var _ = BeforeSuite(func() {
	contents, err := ioutil.ReadFile("test_assets/realestate_com_au_buy_list_1.html")
	Expect(err).ToNot(HaveOccurred())
	ReadRealEstateComAu_Buy_list_1 = string(contents)
})
