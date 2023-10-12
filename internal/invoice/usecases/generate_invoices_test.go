package usecases

import (
	"encoding/json"
	"github.com/kenesparta/golang-solid/internal/database"
	"github.com/kenesparta/golang-solid/internal/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAddition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Addition Suite")
}

var _ = Describe("generate Invoices", func() {
	It("should generate notas fiscais", func() {
		generateInvoices := NewGenerateInvoices(
			repository.NewDatabaseRepository(
				database.NewContractPgAdapter(),
			),
			NewJsonPresenter(),
		)
		output, _ := generateInvoices.Execute(Input{
			Month:     1,
			Year:      2022,
			TypeInput: "cash",
		})
		var outputSer []Output
		json.Unmarshal([]byte(output), &outputSer)
		Expect(outputSer[0].Date).To(Equal("2022-01-05T10:00:00Z"))
		Expect(outputSer[0].Amount).To(Equal(6000.0))
	})
})

var _ = Describe("generate Invoices", func() {
	It("should generate notas fiscais", func() {
		generateInvoices := NewGenerateInvoices(
			repository.NewDatabaseRepository(
				database.NewContractPgAdapter(),
			),
			NewJsonPresenter(),
		)
		output, _ := generateInvoices.Execute(Input{
			Month:     2,
			Year:      2022,
			TypeInput: "accrual",
		})
		var outputSer []Output
		json.Unmarshal([]byte(output), &outputSer)
		Expect(outputSer[0].Date).To(Equal("2022-02-01T10:00:00Z"))
		Expect(outputSer[0].Amount).To(Equal(500.0))
	})
})

var _ = Describe("generate Invoices", func() {
	It("should generate notas fiscais in CSV format", func() {
		generateInvoices := NewGenerateInvoices(
			repository.NewDatabaseRepository(
				database.NewContractPgAdapter(),
			),
			NewCsvPresenter(),
		)
		output, _ := generateInvoices.Execute(Input{
			Month:     1,
			Year:      2022,
			TypeInput: "cash",
		})
		resultCsv := `2022-01-05T10:00:00Z,6000.000000`
		Expect(output).To(Equal(resultCsv))
	})
})
