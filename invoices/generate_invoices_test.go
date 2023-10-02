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
		generateInvoices.Execute()
	})
})
