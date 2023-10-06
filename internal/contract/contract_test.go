package contract

import (
	"github.com/kenesparta/golang-solid/internal/invoice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract_GetPayments(t *testing.T) {
	tests := []struct {
		name          string
		contract      Contract
		invoices      invoice.Invoice
		invoicesIndex int
	}{
		{
			name: "",
			contract: Contract{
				ID:          "",
				Description: "",
				Amount:      6000.0,
				Periods:     12,
				Date:        "2022-01-05T10:00:00Z",
				payments:    nil,
			},
			invoices: invoice.Invoice{
				Date:   "2022-01-05T10:00:00Z",
				Amount: 500,
			},
			invoicesIndex: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoices, _ := tt.contract.GetInvoices(1, 2022, "accrual")
			assert.Equal(t, tt.invoices.Amount, invoices[tt.invoicesIndex].Amount)
			assert.Equal(t, tt.invoices.Date, invoices[tt.invoicesIndex].Date)
		})
	}
}
