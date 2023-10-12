package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	rt := GetRouter()
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
