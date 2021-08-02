package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-service-api/internal/flusher"
	"github.com/ozoncp/ocp-service-api/internal/mocks"
	"github.com/ozoncp/ocp-service-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl  *gomock.Controller
		mockRepo  *mocks.MockRepo
		f         flusher.Flusher
		services  []models.Service
		chunkSize int
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)

		services = []models.Service{
			{Id: 1, UserId: 1, Name: "serv1", Link: "https://serv.ru"},
			{Id: 2, UserId: 1, Name: "serv2", Link: "https://serv2.ru"},
			{Id: 3, UserId: 2, Name: "serv3", Link: "https://serv3.ru"},
			{Id: 4, UserId: 3, Name: "serv 4", Link: "https://serv4.ru"},
			{Id: 5, UserId: 2, Name: "serv 5", Link: "https://serv5.ru"},
		}
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	// TODO: empty services
	Context("Flush", func() {
		When("negative chunk size", func() {
			BeforeEach(func() {
				chunkSize = -1
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("returns original slice and error", func() {
				result, err := f.Flush(services)
				Expect(result).Should(Equal(services))
				Expect(err).Should(HaveOccurred())
			})
			It("doesn't call AddServices from Repo", func() {
				mockRepo.EXPECT().AddServices(gomock.Any()).Times(0)
				f.Flush(services)
			})
		})
		When("empty services", func() {
			BeforeEach(func() {
				chunkSize = 2
				services = []models.Service{}
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("no errors and no unsaved elements", func() {
				result, err := f.Flush(services)
				Expect(result).Should(BeNil())
				Expect(err).Should(BeNil())
			})
			It("doesn't call AddServices from Repo", func() {
				mockRepo.EXPECT().AddServices(gomock.Any()).Times(0)
				f.Flush(services)
			})
		})
		When("Chunk size greater than elements length", func() {
			BeforeEach(func() {
				chunkSize = 10
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("no errors and no unsaved elements;", func() {
				mockRepo.EXPECT().AddServices(gomock.Any()).Times(1)
				result, err := f.Flush(services)
				Expect(result).Should(BeNil())
				Expect(err).Should(BeNil())
			})
		})
		When("Chunk size equal", func() {
			BeforeEach(func() {
				chunkSize = 5
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("no errors and no unsaved elements", func() {
				mockRepo.EXPECT().AddServices(gomock.Any()).Times(1)
				result, err := f.Flush(services)
				Expect(result).Should(BeNil())
				Expect(err).Should(BeNil())
			})
		})
		When("Chunk size less", func() {
			BeforeEach(func() {
				chunkSize = 2
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("no errors and no unsaved elements", func() {
				mockRepo.EXPECT().AddServices(gomock.Any()).Times(3)
				result, err := f.Flush(services)
				Expect(result).Should(BeNil())
				Expect(err).Should(BeNil())
			})
		})
		When("Add service returns error immediately", func() {
			BeforeEach(func() {
				chunkSize = 2
				mockRepo.EXPECT().AddServices(gomock.Any()).Return(errors.New("error"))
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("error raised and all elements unsaved", func() {
				result, err := f.Flush(services)
				Expect(result).Should(Equal(services))
				Expect(err).Should(HaveOccurred())
			})
		})
		When("Add service returns error after iterations", func() {
			BeforeEach(func() {
				chunkSize = 2
				mockRepo.EXPECT().AddServices(gomock.Any()).Return(nil).Times(1)
				mockRepo.EXPECT().AddServices(gomock.Any()).Return(errors.New("error")).Times(1)
				f = flusher.NewFlusher(chunkSize, mockRepo)
			})
			It("error raised saved only two first elements", func() {
				result, err := f.Flush(services)
				Expect(result).Should(Equal(services[2:]))
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
