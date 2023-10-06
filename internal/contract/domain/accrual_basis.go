package domain

import (
	"time"

	"github.com/kenesparta/golang-solid/internal/invoice/domain"
)

type AccrualBasis struct{}

func (cb *AccrualBasis) Generate(
	contract *Contract,
	month int,
	year uint64,
) ([]domain.Invoice, error) {
	var (
		invoices []domain.Invoice
		period   uint64
	)

	for period <= contract.Periods {
		contractDate, errParseDate := time.Parse(time.RFC3339, contract.Date)
		if errParseDate != nil {
			return nil, errParseDate
		}

		contractDate = contractDate.AddDate(0, int(period), 0)
		period++
		if int(contractDate.Month()) != month ||
			uint64(contractDate.Year()) != year {
			continue
		}

		invoices = append(invoices, domain.Invoice{
			Date:   contractDate.Format(time.RFC3339),
			Amount: contract.Amount / float64(contract.Periods),
		})
	}

	return invoices, nil
}
