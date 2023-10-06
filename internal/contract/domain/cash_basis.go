package domain

import (
	"time"

	"github.com/kenesparta/golang-solid/internal/invoice/domain"
)

type CashBasis struct{}

func (cb *CashBasis) Generate(
	contract *Contract,
	month int,
	year uint64,
) ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	for _, p := range contract.GetPayments() {
		paymentDate, errParseDate := time.Parse(time.RFC3339, p.Date)
		if errParseDate != nil {
			return nil, errParseDate
		}

		if int(paymentDate.Month()) == month &&
			uint64(paymentDate.Year()) == year {
			invoices = append(invoices, domain.Invoice{
				Date:   p.Date,
				Amount: p.Amount,
			})
		}
	}

	return invoices, nil
}
