package handlers

import (
	"encoding/json"
	"github.com/kenesparta/golang-solid/internal/database"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"github.com/kenesparta/golang-solid/internal/repository"
	"net/http"
)

func ReadInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input usecases.Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	generateInvoices := usecases.NewGenerateInvoices(
		repository.NewDatabaseRepository(
			database.NewContractPgAdapter(),
		),
		usecases.NewJsonPresenter(),
	)

	output, executeErr := generateInvoices.Execute(input)

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
