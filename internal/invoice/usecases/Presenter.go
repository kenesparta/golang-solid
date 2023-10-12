package usecases

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Presenter interface {
	Present(output []Output) (string, error)
}

type JsonPresenter struct{}

func NewJsonPresenter() *JsonPresenter {
	return &JsonPresenter{}
}

func (jp *JsonPresenter) Present(output []Output) (string, error) {
	bytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type CsvPresenter struct{}

func NewCsvPresenter() *CsvPresenter {
	return &CsvPresenter{}
}

func (jp *CsvPresenter) Present(output []Output) (string, error) {
	var lines []string
	for _, out := range output {
		var row []string
		row = append(row, out.Date)
		row = append(row, fmt.Sprintf("%f", out.Amount))
		lines = append(lines, strings.Join(row, ","))
	}

	return strings.Join(lines, "\n"), nil
}
