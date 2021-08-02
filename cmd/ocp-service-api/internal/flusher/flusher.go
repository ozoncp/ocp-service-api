package flusher

import (
	"errors"
	"github.com/ozoncp/ocp-service-api/internal/models"
	"github.com/ozoncp/ocp-service-api/internal/repo"
	"github.com/ozoncp/ocp-service-api/internal/utils"
)

type flusher struct {
	chunkSize   int
	serviceRepo repo.Repo
}

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(services []models.Service) ([]models.Service, error)
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	serviceRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize:   chunkSize,
		serviceRepo: serviceRepo,
	}
}

// Flush Возвращает несохраненные записи и ошибку
func (f *flusher) Flush(services []models.Service) ([]models.Service, error) {
	if f.chunkSize <= 0 {
		return services, errors.New("batch size have to be greater than 0")
	}
	chunks := utils.SplitToBatches(services, f.chunkSize)
	for index := range chunks {
		if err := f.serviceRepo.AddServices(chunks[index]); err != nil {
			return services[f.chunkSize*index:], err
		}
	}
	return nil, nil
}
