package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kenesparta/golang-solid/cmd/api/httpserver"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"github.com/kenesparta/golang-solid/internal/shared/decorator"
)

type ChiServer struct{}

func (ChiServer) NewServer(
	adapter httpserver.HttpServerAdapter[*chi.Mux],
	usecase decorator.UseCase[usecases.Input],
) {
	mc := &MainController{Usecase: usecase, Adapter: adapter}
	r := mc.Adapter.Router()
	r.Get("/invoices", mc.ReadInvoices)
}
