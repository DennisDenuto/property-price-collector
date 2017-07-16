package integration_test

import (
	. "github.com/DennisDenuto/property-price-collector/data/training"

	"github.com/DennisDenuto/property-price-collector/data"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pachyderm/pachyderm/src/client"
)

var _ = Describe("Integration tests", func() {

	var trainingDataRepo Repo

	BeforeEach(func() {
		client, err := client.NewFromAddress("0.0.0.0:30650")
		Expect(err).ToNot(HaveOccurred())

		trainingDataRepo = NewTrainingDataRepo(client)
	})

	It("should allow writing a file in a transaction", func() {
		By("should create a repo correctly named", func() {
			Expect(trainingDataRepo.Create()).To(Succeed())
			Expect(ListRepos()).To(ContainElement("training-data-properties"))
		})

		var startedTxn string
		var err error
		By("starting a txn", func() {
			startedTxn, err = trainingDataRepo.StartTxn()
			Expect(err).To(Succeed())
			Expect(ListCommits("training-data-properties")).To(ContainElement("training-data-properties"))
		})

		By("writing a file in txn", func() {
			propertyHistoryData := data.PropertyHistoryData{
				Property: data.Property{
					Address: data.Address{
						AddressLine1: "123 fake street",
						State:        "NSW",
						Suburb:       "north sydney",
					},
				},
			}
			Expect(trainingDataRepo.Add(propertyHistoryData)).To(Succeed())
		})

		By("commiting transaction", func() {
			Expect(trainingDataRepo.Commit(startedTxn)).To(Succeed())
		})

		Expect(GetFile("training-data-properties", "master", "/nsw/north_sydney/123_fake_street")).ToNot(BeEmpty())
	})
})
