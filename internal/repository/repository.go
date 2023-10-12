package repository

import "github.com/kenesparta/golang-solid/internal/contract/domain"

// Repository Persists Contracts Aggregates.Aggregates are clusters of data persistence.
type Repository interface {
	List() ([]domain.Contract, error)
}
