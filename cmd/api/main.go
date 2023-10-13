package main

import (
	"github.com/kenesparta/golang-solid/cmd/api/handlers"
	"github.com/kenesparta/golang-solid/cmd/api/httpserver"
	"github.com/kenesparta/golang-solid/internal/database"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"github.com/kenesparta/golang-solid/internal/repository"
	"github.com/kenesparta/golang-solid/internal/shared/decorator"
)

func main() {
	generateInvoices := decorator.NewLoggerDecorator(
		usecases.NewGenerateInvoices(
			repository.NewDatabaseRepository(database.NewContractPgAdapter()),
			usecases.NewJsonPresenter(),
		),
	)
	chiAdapter := httpserver.NewChiHttpAdapter()
	chiServer := handlers.ChiServer{}
	chiServer.NewServer(chiAdapter, generateInvoices)
	chiAdapter.Listen(8080)
}
