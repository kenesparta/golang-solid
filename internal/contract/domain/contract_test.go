package domain

import (
	"testing"

	"github.com/kenesparta/golang-solid/internal/invoice/domain"
	"github.com/stretchr/testify/assert"
)

func TestContract_GetPayments(t *testing.T) {
	tests := []struct {
		name          string
		contract      Contract
		invoices      domain.Invoice
		invoicesIndex int
	}{
		{
			name: "Should generate invoices from a contract",
			contract: Contract{
				ID:          "",
				Description: "",
				Amount:      6000.0,
				Periods:     12,
				Date:        "2022-01-05T10:00:00Z",
				payments:    nil,
			},
			invoices: domain.Invoice{
				Date:   "2022-01-05T10:00:00Z",
				Amount: 500,
			},
			invoicesIndex: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoices, _ := tt.contract.GenerateInvoices(1, 2022, "accrual")
			assert.Equal(t, tt.invoices.Amount, invoices[tt.invoicesIndex].Amount)
			assert.Equal(t, tt.invoices.Date, invoices[tt.invoicesIndex].Date)
		})
	}
}

func TestContract_Getbalance(t *testing.T) {
	tests := []struct {
		name     string
		contract Contract
		invoices domain.Invoice
	}{
		{
			name: "Should generate the total from a contract",
			contract: Contract{
				ID:          "",
				Description: "",
				Amount:      6000.0,
				Periods:     12,
				Date:        "2022-01-01T10:00:00Z",
				payments: []Payment{
					{
						Date:   "2022-01-01T10:00:00Z",
						Amount: 2000,
					},
					{
						Date:   "2022-02-01T10:00:00Z",
						Amount: 500,
					},
				},
			},
			invoices: domain.Invoice{
				Date:   "2022-01-05T10:00:00Z",
				Amount: 3500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance := tt.contract.GetBalance()
			assert.Equal(t, tt.invoices.Amount, balance)
		})
	}
}
