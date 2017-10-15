package pachyderm_test

import (
	"github.com/DennisDenuto/property-price-collector/data/training"

	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training/pachyderm/pachydermfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pkg/errors"
	"io/ioutil"
	. "github.com/DennisDenuto/property-price-collector/data/training/pachyderm"
)

var _ = Describe("TrainingDataRepo", func() {

	var trainingDataRepo training.TxnRepo
	var fakeApiClient *pachydermfakes.FakeAPIClient

	BeforeEach(func() {
		fakeApiClient = &pachydermfakes.FakeAPIClient{}
		trainingDataRepo = NewTrainingDataRepo(fakeApiClient)
	})

	Describe("Create", func() {
		BeforeEach(func() {
			fakeApiClient.ListRepoReturns([]*pfs.RepoInfo{}, nil)
		})

		It("should create a repo correctly named", func() {
			Expect(trainingDataRepo.Create()).To(Succeed())

			Expect(fakeApiClient.CreateRepoCallCount()).To(Equal(1))
			Expect(fakeApiClient.CreateRepoArgsForCall(0)).To(Equal("training-data-properties"))
		})

		Context("when repo already exists", func() {
			BeforeEach(func() {
				fakeApiClient.ListRepoReturns([]*pfs.RepoInfo{{Repo: &pfs.Repo{"training-data-repo"}}}, nil)
			})

			It("should not error", func() {
				Expect(trainingDataRepo.Create()).To(Succeed())

				Expect(fakeApiClient.CreateRepoCallCount()).To(Equal(0))
			})

			Context("when listing a repo fails", func() {
				BeforeEach(func() {
					fakeApiClient.ListRepoReturns(nil, errors.New("some-error"))
				})

				It("should return error", func() {
					err := trainingDataRepo.Create()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("listing repos: some-error"))
				})
			})
		})

		Context("when api client returns an error", func() {
			BeforeEach(func() {
				fakeApiClient.CreateRepoReturns(errors.New("some-error"))
			})

			It("should handle error", func() {
				err := trainingDataRepo.Create()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("creating repo: some-error"))
			})
		})
	})

	Describe("StartTxn", func() {
		BeforeEach(func() {
			fakeApiClient.ListCommitByRepoReturns([]*pfs.CommitInfo{}, nil)
			fakeApiClient.StartCommitReturns(&pfs.Commit{ID: "123"}, nil)
		})

		It("should create a commit for correct repo", func() {
			commitId, err := trainingDataRepo.StartTxn()
			Expect(err).ToNot(HaveOccurred())
			Expect(commitId).To(Equal("123"))

			Expect(fakeApiClient.ListCommitByRepoCallCount()).To(Equal(1))
			Expect(fakeApiClient.ListCommitByRepoArgsForCall(0)).To(Equal("training-data-properties"))
			Expect(fakeApiClient.StartCommitCallCount()).To(Equal(1))
			repo, branch := fakeApiClient.StartCommitArgsForCall(0)
			Expect(repo).To(Equal("training-data-properties"))
			Expect(branch).To(Equal("master"))
		})

		Context("when starting a txn fails", func() {
			BeforeEach(func() {
				fakeApiClient.StartCommitReturns(nil, errors.New("some-error"))
			})

			It("should return an error", func() {
				_, err := trainingDataRepo.StartTxn()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("starting txn: some-error"))
			})
		})

		Context("when transaction was previously started but not finished", func() {
			BeforeEach(func() {
				fakeApiClient.ListCommitByRepoReturns([]*pfs.CommitInfo{
					{Finished: nil, Commit: &pfs.Commit{ID: "existing-commit-id"}},
				}, nil)
			})

			It("should not try to start it again", func() {
				commitId, err := trainingDataRepo.StartTxn()
				Expect(err).NotTo(HaveOccurred())
				Expect(commitId).To(Equal("existing-commit-id"))
				Expect(fakeApiClient.StartCommitCallCount()).To(Equal(0))
			})

			Context("when listing commits for a repo fail", func() {
				BeforeEach(func() {
					fakeApiClient.ListCommitByRepoReturns([]*pfs.CommitInfo{{}}, errors.New("some-error"))
				})

				It("should return error", func() {
					_, err := trainingDataRepo.StartTxn()
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("listing txn: some-error"))
				})
			})
		})
	})

	Describe("Add", func() {
		var propertyHistoryData data.PropertyHistoryData
		var propertyJson string

		BeforeEach(func() {
			propertyHistoryData = data.PropertyHistoryData{
				Property: data.Property{
					Address: data.Address{
						AddressLine1: "1/123-124 fake street",
						State:        "NSW",
						Suburb:       "north sydney",
					},
				},
			}

			json, err := json.Marshal(propertyHistoryData)
			Expect(err).ToNot(HaveOccurred())

			propertyJson = string(json)
		})

		It("should add a file to repo", func() {

			Expect(trainingDataRepo.Add(propertyHistoryData)).To(Succeed())

			Expect(fakeApiClient.PutFileCallCount()).To(Equal(1))

			repo, commitID, path, reader := fakeApiClient.PutFileArgsForCall(0)
			Expect(repo).To(Equal("training-data-properties"))
			Expect(commitID).To(Equal("master"))
			Expect(path).To(Equal("/nsw/north_sydney/1_123_124_fake_street"))

			propertyArg, err := ioutil.ReadAll(reader)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(propertyArg)).To(Equal(propertyJson))
		})
	})

	Describe("Commit", func() {
		It("should finish transaction", func() {
			Expect(trainingDataRepo.Commit("master-123")).To(Succeed())

			Expect(fakeApiClient.FinishCommitCallCount()).To(Equal(1))
			repo, commit := fakeApiClient.FinishCommitArgsForCall(0)
			Expect(repo).To(Equal("training-data-properties"))
			Expect(commit).To(Equal("master-123"))
		})

		Context("when finishing a commit fails", func() {

			BeforeEach(func() {
				fakeApiClient.FinishCommitReturns(errors.New("some-error"))
			})

			It("should return an error", func() {
				err := trainingDataRepo.Commit("master")
				Expect(err).To(HaveOccurred())

				Expect(err).To(MatchError("committing: some-error"))
			})
		})
	})
})
