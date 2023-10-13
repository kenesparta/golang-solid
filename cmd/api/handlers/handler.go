package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/kenesparta/golang-solid/cmd/api/httpserver"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"github.com/kenesparta/golang-solid/internal/shared/decorator"
)

type MainController struct {
	Usecase decorator.UseCase[usecases.Input]
	Adapter httpserver.HttpServerAdapter[*chi.Mux]
}

func (mc *MainController) ReadInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input usecases.Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	output, executeErr := mc.Usecase.Execute(input)
	if executeErr != nil {
		http.Error(w, "Failed execution", http.StatusInternalServerError)
		return
	}

	var outputSer []usecases.Output
	unmErr := json.Unmarshal(output, &outputSer)
	if unmErr != nil {
		http.Error(w, "Failed unmarshal", http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(outputSer)
	if err != nil {
		http.Error(w, "Failed encoding", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
