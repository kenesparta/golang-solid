package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type ChiHttpAdapter struct {
	routes *chi.Mux
}

func NewChiHttpAdapter() *ChiHttpAdapter {
	return &ChiHttpAdapter{chi.NewRouter()}
}

func (cha *ChiHttpAdapter) Router() *chi.Mux {
	return cha.routes
}

func (cha *ChiHttpAdapter) Listen(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), cha.routes)
	if err != nil {
		log.Fatal(err)
		return
	}
}
