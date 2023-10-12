package repository

import (
	"github.com/kenesparta/golang-solid/internal/contract/domain"
	"github.com/kenesparta/golang-solid/internal/database"
	_ "github.com/lib/pq"
)

type DatabaseRepository struct {
	adapter database.Adapter[[]domain.Contract, *domain.Contract]
}

func NewDatabaseRepository(adapter database.Adapter[[]domain.Contract, *domain.Contract]) *DatabaseRepository {
	return &DatabaseRepository{adapter}
}

func (dbRepo *DatabaseRepository) List() ([]domain.Contract, error) {
	return dbRepo.adapter.QueryAll()
}
