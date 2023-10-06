package domain

import (
	"errors"
)

type InvoiceGenerationFactory struct{}

func (igf *InvoiceGenerationFactory) Create(typeInput string) (InvoiceGenerationStrategy, error) {
	switch typeInput {
	case "cash":
		return &CashBasis{}, nil
	case "accrual":
		return &AccrualBasis{}, nil
	default:
		return nil, errors.New("type does not exist")
	}
}
