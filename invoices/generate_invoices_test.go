package invoices

import (
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
		generateInvoices := NewGenerateInvoices()
		output, _ := generateInvoices.Execute(Input{
			Month:     1,
			Year:      2022,
			TypeInput: "cash",
		})
		Expect(output[0].Date).To(Equal("2022-01-05T10:00:00Z"))
		Expect(output[0].Amount).To(Equal(6000.0))
	})
})

var _ = Describe("generate Invoices", func() {
	It("should generate notas fiscais", func() {
		generateInvoices := NewGenerateInvoices()
		output, _ := generateInvoices.Execute(Input{
			Month:     2,
			Year:      2022,
			TypeInput: "accrual",
		})
		Expect(output[0].Date).To(Equal("2022-02-01T10:00:00Z"))
		Expect(output[0].Amount).To(Equal(500.0))
	})
})
