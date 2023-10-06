package domain

import (
	"github.com/kenesparta/golang-solid/internal/invoice/domain"
)

type InvoiceGenerationStrategy interface {
	Generate(
		contract *Contract,
		month int,
		year uint64,
	) ([]domain.Invoice, error)
}
