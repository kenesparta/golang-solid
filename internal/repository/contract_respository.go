package repository

import (
	"github.com/kenesparta/golang-solid/internal/contract/domain"
	"github.com/kenesparta/golang-solid/internal/database"
	_ "github.com/lib/pq"
)

// Repository Persists Contracts Aggregates.Aggregates are clusters of data persistence.
type Repository interface {
	List() ([]domain.Contract, error)
}

type DatabaseRepository struct {
	adapter database.Adapter[[]domain.Contract, *domain.Contract]
}

func NewDatabaseRepository(adapter database.Adapter[[]domain.Contract, *domain.Contract]) *DatabaseRepository {
	return &DatabaseRepository{adapter}
}

func (dbRepo *DatabaseRepository) List() ([]domain.Contract, error) {
	return dbRepo.adapter.QueryAll()
}
