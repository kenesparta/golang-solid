package contract

import (
	"github.com/kenesparta/golang-solid/internal/invoice"
	"time"
)

type Contract struct {
	ID          string
	Description string
	Amount      float64
	Periods     uint64
	Date        string
	payments    []Payment
}

func (c *Contract) AddPayments(payment Payment) {
	c.payments = append(c.payments, payment)
}

func (c *Contract) GetPayments() []Payment {
	return c.payments
}

func (c *Contract) GetInvoices(
	month int,
	year uint64,
	typeInput string,
) ([]invoice.Invoice, error) {
	var invoices []invoice.Invoice

	if typeInput == "cash" {
		for _, p := range c.GetPayments() {
			paymentDate, errParseDate := time.Parse(time.RFC3339, p.Date)
			if errParseDate != nil {
				return nil, errParseDate
			}

			if int(paymentDate.Month()) == month &&
				uint64(paymentDate.Year()) == year {
				invoices = append(invoices, invoice.Invoice{
					Date:   p.Date,
					Amount: p.Amount,
				})
			}
		}
	}

	if typeInput == "accrual" {
		var period uint64
		for period <= c.Periods {
			contractDate, errParseDate := time.Parse(time.RFC3339, c.Date)
			if errParseDate != nil {
				return nil, errParseDate
			}

			contractDate = contractDate.AddDate(0, int(period), 0)
			period++
			if int(contractDate.Month()) != month ||
				uint64(contractDate.Year()) != year {
				continue
			}

			invoices = append(invoices, invoice.Invoice{
				Date:   contractDate.Format(time.RFC3339),
				Amount: c.Amount / float64(c.Periods),
			})
		}
	}
	return invoices, nil
}

type Payment struct {
	Date   string
	Amount float64
}
