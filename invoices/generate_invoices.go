package invoices

import (
	"github.com/kenesparta/golang-solid/contract"
	"time"

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
	contractRepo contract.Repository
}

func NewGenerateInvoices(contractRepo contract.Repository) *GenerateInvoices {
	return &GenerateInvoices{contractRepo: contractRepo}
}

func (gi *GenerateInvoices) Execute(input Input) ([]Output, error) {
	// dbRepo := gi.contractRepo.List()
	contracts, _ := gi.contractRepo.List()
	var output []Output
	for _, c := range contracts {
		if input.TypeInput == "cash" {
			for _, p := range c.Payments {
				paymentDate, errParseDate := time.Parse(time.RFC3339, p.Date)
				if errParseDate != nil {
					return nil, errParseDate
				}

				if int(paymentDate.Month()) == input.Month &&
					uint64(paymentDate.Year()) == input.Year {
					output = append(output, Output{
						Date:   p.Date,
						Amount: p.Amount,
					})
				}
			}
		}

		if input.TypeInput == "accrual" {
			var period uint64
			for period <= c.Periods {
				contractDate, errParseDate := time.Parse(time.RFC3339, c.Date)
				if errParseDate != nil {
					return nil, errParseDate
				}

				contractDate = contractDate.AddDate(0, int(period), 0)
				period++
				if int(contractDate.Month()) != input.Month ||
					uint64(contractDate.Year()) != input.Year {
					continue
				}

				output = append(output, Output{
					Date:   contractDate.Format(time.RFC3339),
					Amount: c.Amount / float64(c.Periods),
				})
			}
		}
	}

	return output, nil
}
