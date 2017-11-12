package dropbox

import (
	"encoding/json"
	"errors"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training/dropbox/dropboxfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("DomaincomuaHistoryTrainingRepo", func() {

	var repo DomainComAuHistoryDataRepo
	var fakeClient *dropboxfakes.FakeClient

	BeforeEach(func() {
		fakeClient = &dropboxfakes.FakeClient{}

		repo = DomainComAuHistoryDataRepo{
			dropboxClient: fakeClient,
		}
	})

	Describe("Add", func() {
		It("should add to dropbox", func() {
			propertyHistoryData := data.DomainComAuPropertyWrapper{
				PropertyObject: data.DomainComAuPropertyHistory{
					Suburb:        "North Sydney",
					State:         "NSW",
					Address:       "1 123-124 fake street, North Sydney NSW 2000",
					StreetAddress: "1 123-124 fake street",
				},
			}

			propertyHistoryDataJson, err := json.Marshal(propertyHistoryData)
			Expect(err).ToNot(HaveOccurred())

			err = repo.Add(&propertyHistoryData)
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeClient.UploadCallCount()).To(Equal(1))

			commitInfo, content := fakeClient.UploadArgsForCall(0)
			Expect(commitInfo.Path).To(Equal("/domaincomau/nsw/north_sydney/1_123-124_fake_street"))

			contents, err := ioutil.ReadAll(content)
			Expect(err).NotTo(HaveOccurred())

			Expect(contents).To(MatchJSON(propertyHistoryDataJson))
		})

		Context("when dropbox fails to upload", func() {
			It("should return an error", func() {
				propertyHistoryData := &data.DomainComAuPropertyWrapper{
					PropertyObject: data.DomainComAuPropertyHistory{
						Suburb:        "North Sydney",
						State:         "NSW",
						Address:       "1 123-124 fake street, North Sydney NSW 2000",
						StreetAddress: "1 123-124 fake street",
					},
				}

				fakeClient.UploadReturns(nil, errors.New("some dropbox error"))
				err := repo.Add(propertyHistoryData)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("some dropbox error"))
			})
		})
	})

})
