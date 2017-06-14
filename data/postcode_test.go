package data_test

import (
	. "github.com/DennisDenuto/property-price-collector/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Postcode", func() {

	It("should generate postcode numbers for nsw", func() {
		postCodes := ListNswPostcodes()
		Eventually(postCodes, 1*time.Minute, 1*time.Millisecond).Should(Receive(Equal(1001)))
		Eventually(postCodes, 1*time.Minute, 1*time.Millisecond).Should(Receive(Equal(3707)))
	})
})
