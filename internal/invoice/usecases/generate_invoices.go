package usecases

import (
	"github.com/kenesparta/golang-solid/internal/repository"
	_ "github.com/lib/pq"
)

type Output struct {
	Date   string
	Amount float64
}

type Input struct {
	Month     int
	Year      uint64
	TypeInput string
}

type GenerateInvoices struct {
	contractRepo repository.Repository
}

func NewGenerateInvoices(contractRepo repository.Repository) *GenerateInvoices {
	return &GenerateInvoices{contractRepo: contractRepo}
}

func (gi *GenerateInvoices) Execute(input Input) ([]Output, error) {
	contracts, listErr := gi.contractRepo.List()
	if listErr != nil {
		return nil, listErr
	}

	var output []Output
	for _, c := range contracts {
		invoices, invErr := c.GenerateInvoices(input.Month, input.Year, input.TypeInput)
		if invErr != nil {
			return nil, invErr
		}

		for _, inv := range invoices {
			output = append(output, Output{
				Date:   inv.Date,
				Amount: inv.Amount,
			})
		}
	}

	return output, nil
}
