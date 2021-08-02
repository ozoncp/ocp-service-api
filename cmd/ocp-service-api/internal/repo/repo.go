package repo

import "github.com/ozoncp/ocp-service-api/internal/models"

// Repo - интерфейс хранилища для сущности Entity
type Repo interface {
	AddServices(services []models.Service) error
	ListServices(limit, offset uint64) ([]models.Service, error)
	DescribeService(serviceId uint64) (*models.Service, error)
}
