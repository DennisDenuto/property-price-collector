package dropbox

import (
	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training/dropbox/dropboxfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"io/ioutil"
)

var _ = Describe("TrainingRepo", func() {

	var repo PropertyHistoryDataRepo
	var fakeClient *dropboxfakes.FakeClient

	BeforeEach(func() {
		fakeClient = &dropboxfakes.FakeClient{}

		repo = PropertyHistoryDataRepo{
			dropboxClient: fakeClient,
		}
	})

	It("should add to dropbox", func() {
		propertyHistoryData := data.PropertyHistoryData{
			Address: data.Address{
				AddressLine1: "1/123-124 fake street",
				State:        "NSW",
				Suburb:       "north sydney",
			},
		}

		propertyHistoryDataJson, err := json.Marshal(propertyHistoryData)
		Expect(err).ToNot(HaveOccurred())

		err = repo.Add(propertyHistoryData)
		Expect(err).NotTo(HaveOccurred())
		Expect(fakeClient.UploadCallCount()).To(Equal(1))

		commitInfo, content := fakeClient.UploadArgsForCall(0)
		Expect(commitInfo.Path).To(Equal("/nsw/north_sydney/1_123-124_fake_street"))

		contents, err := ioutil.ReadAll(content)
		Expect(err).NotTo(HaveOccurred())

		Expect(contents).To(MatchJSON(propertyHistoryDataJson))
	})

	Context("when dropbox fails to upload", func() {
		It("should return an error", func() {
			propertyHistoryData := data.PropertyHistoryData{
				Address: data.Address{
					AddressLine1: "1/123-124 fake street",
					State:        "NSW",
					Suburb:       "north sydney",
				},
			}

			fakeClient.UploadReturns(nil, errors.New("some dropbox error"))
			err := repo.Add(propertyHistoryData)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("some dropbox error"))
		})
	})
})
