package propertypricehistorycom_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"testing"
)

var PropertyPriceHistory_list_nsw_2155 string
var PropertyPriceHistory_list_nsw_2155_last_page string

func TestPropertypricehistorycom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Propertypricehistorycom Suite")
}

var _ = BeforeSuite(func() {
	contents, err := ioutil.ReadFile("test_assets/property_price_com_list_nsw_2155.html")
	Expect(err).ToNot(HaveOccurred())
	PropertyPriceHistory_list_nsw_2155 = string(contents)

	contents, err = ioutil.ReadFile("test_assets/property_price_com_list_nsw_2155_last_page.html")
	Expect(err).ToNot(HaveOccurred())
	PropertyPriceHistory_list_nsw_2155_last_page = string(contents)
})
