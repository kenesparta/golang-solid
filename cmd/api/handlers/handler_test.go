package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kenesparta/golang-solid/cmd/api/httpserver"
	"github.com/kenesparta/golang-solid/internal/database"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"github.com/kenesparta/golang-solid/internal/repository"
	"github.com/kenesparta/golang-solid/internal/shared/decorator"
	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	generateInvoices := decorator.NewLoggerDecorator(
		usecases.NewGenerateInvoices(
			repository.NewDatabaseRepository(database.NewContractPgAdapter()),
			usecases.NewJsonPresenter(),
		),
	)
	chiAdapter := httpserver.NewChiHttpAdapter()
	chiServer := ChiServer{}
	chiServer.NewServer(chiAdapter, generateInvoices)
	rt := chiAdapter.Router()

	in := usecases.Input{
		Month:     1,
		Year:      2022,
		TypeInput: "cash",
	}
	marshRes, _ := json.Marshal(in)
	req, err := http.NewRequest(http.MethodGet, "/invoices", bytes.NewReader(marshRes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	rt.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	respBody, _ := io.ReadAll(rr.Body)
	var outputSer []usecases.Output
	json.Unmarshal(respBody, &outputSer)
	assert.Equal(t, "2022-01-05T10:00:00Z", outputSer[0].Date, "Expected response body to match")
	assert.Equal(t, 6000.0, outputSer[0].Amount, "Expected response body to match")
}
