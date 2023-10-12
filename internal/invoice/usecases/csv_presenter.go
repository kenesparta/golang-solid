package usecases

import (
	"fmt"
	"strings"
)

type CsvPresenter struct{}

func NewCsvPresenter() *CsvPresenter {
	return &CsvPresenter{}
}

func (*CsvPresenter) Present(output []Output) (string, error) {
	var lines []string
	for _, out := range output {
		var row []string
		row = append(row, out.Date)
		row = append(row, fmt.Sprintf("%f", out.Amount))
		lines = append(lines, strings.Join(row, ","))
	}

	return strings.Join(lines, "\n"), nil
}
