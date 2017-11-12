package dropbox

import (
	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training/dropbox/dropboxfakes"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
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

	Describe("Add", func() {
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

	Describe("List", func() {
		var returnedFile *files.FileMetadata

		BeforeEach(func() {
			returnedFile = &files.FileMetadata{
				Id: "some-id",
			}
			returnedFile.Name = "filename"
			returnedFile.PathLower = "/some/path"
		})

		It("should list files in directory", func() {
			fakeClient.ListFolderReturns(&files.ListFolderResult{
				Entries: []files.IsMetadata{
					returnedFile,
				},
				HasMore: true,
				Cursor:  "some-cursor",
			}, nil)

			fakeClient.ListFolderContinueReturns(&files.ListFolderResult{
				Entries: []files.IsMetadata{
					returnedFile,
				},
				HasMore: false,
			}, nil)

			fakeClient.DownloadStub = func(arg *files.DownloadArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
				return nil, getTestDropboxAddressFile(), nil
			}

			propertiesChan, err := repo.List("NSW", "north sydney")
			Eventually(err).ShouldNot(Receive())

			var expectedProperty1 *data.PropertyHistoryData
			var expectedProperty2 *data.PropertyHistoryData
			Eventually(propertiesChan).Should(Receive(&expectedProperty1))
			Eventually(propertiesChan).Should(Receive(&expectedProperty2))
			Eventually(propertiesChan).Should(BeClosed())

			Expect(expectedProperty1.Price).To(Equal("N/A"))
			Expect(expectedProperty2.Price).To(Equal("N/A"))

			Expect(fakeClient.ListFolderCallCount()).To(Equal(1))
			Expect(fakeClient.ListFolderContinueCallCount()).To(Equal(1))
			Expect(fakeClient.DownloadCallCount()).To(Equal(2))

			listFolderArg := fakeClient.ListFolderArgsForCall(0)
			Expect(listFolderArg.Path).To(Equal("/nsw/north_sydney"))

			listFolderContinueArg := fakeClient.ListFolderContinueArgsForCall(0)
			Expect(listFolderContinueArg.Cursor).To(Equal("some-cursor"))
		})

		Context("when unable to list files in folder", func() {
			It("should generate a useful error", func() {
				fakeClient.ListFolderReturns(nil, errors.New("some error"))

				_, err := repo.List("NSW", "north sydney")
				Eventually(err).Should(Receive(MatchError("Unable to list files in directory: some error")))

			})
		})

		Context("when listing more files from a directory fails", func() {
			It("should generate a useful error", func() {
				fakeClient.ListFolderReturns(&files.ListFolderResult{
					Entries: []files.IsMetadata{
						returnedFile,
					},
					HasMore: true,
					Cursor:  "some-cursor",
				}, nil)
				dropboxPropertyFile := getTestDropboxAddressFile()
				defer dropboxPropertyFile.Close()
				fakeClient.DownloadReturns(nil, dropboxPropertyFile, nil)

				fakeClient.ListFolderContinueReturns(nil, errors.New("some error"))

				propertiesChan, err := repo.List("NSW", "north sydney")
				Eventually(propertiesChan).Should(Receive())
				Eventually(err).Should(Receive(MatchError("Unable to list files in directory: some error")))
			})

		})

		Context("when unable to download a file", func() {
			It("should generate a useful error", func() {
				fakeClient.ListFolderReturns(&files.ListFolderResult{
					Entries: []files.IsMetadata{
						returnedFile,
					},
					HasMore: true,
					Cursor:  "some-cursor",
				}, nil)

				fakeClient.ListFolderContinueReturns(&files.ListFolderResult{
					Entries: []files.IsMetadata{
						returnedFile,
					},
					HasMore: false,
				}, nil)

				fakeClient.DownloadReturns(nil, nil, errors.New("some-error"))

				_, err := repo.List("NSW", "north sydney")
				Eventually(err).Should(Receive(MatchError("Unable to download file: some-error")))
			})

		})
	})
})

func getTestDropboxAddressFile() *os.File {
	file, err := os.Open("./test_assets/dropbox_address_test1.json")
	Expect(err).NotTo(HaveOccurred())
	return file
}
