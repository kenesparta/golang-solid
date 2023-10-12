package handlers

import (
	"github.com/go-chi/chi/v5"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/invoices", ReadInvoices)
	return r
}
