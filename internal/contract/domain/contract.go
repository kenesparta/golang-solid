package domain

import (
	"github.com/kenesparta/golang-solid/internal/invoice/domain"
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

func (c *Contract) GetBalance() float64 {
	balance := c.Amount
	for _, p := range c.payments {
		balance -= p.Amount
	}

	return balance
}

func (c *Contract) GenerateInvoices(
	month int,
	year uint64,
	typeInput string,
) ([]domain.Invoice, error) {
	generationFactory := InvoiceGenerationFactory{}
	invoiceGenStrategy, createErr := generationFactory.Create(typeInput)
	if createErr != nil {
		return nil, createErr
	}

	return invoiceGenStrategy.Generate(c, month, year)
}

type Payment struct {
	Date   string
	Amount float64
}
