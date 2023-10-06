package repository

import (
	"github.com/kenesparta/golang-solid/internal/contract"
	"github.com/kenesparta/golang-solid/internal/database"
	_ "github.com/lib/pq"
)

// Repository Persists Contracts Aggregates.Aggregates are clusters of data persistence.
type Repository interface {
	List() ([]contract.Contract, error)
}

type DatabaseRepository struct {
	adapter database.Adapter[[]contract.Contract, *contract.Contract]
}

func NewDatabaseRepository(adapter database.Adapter[[]contract.Contract, *contract.Contract]) *DatabaseRepository {
	return &DatabaseRepository{adapter}
}

func (dbRepo *DatabaseRepository) List() ([]contract.Contract, error) {
	return dbRepo.adapter.QueryAll()
}
