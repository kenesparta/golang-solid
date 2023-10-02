package invoices

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Contract struct {
	ID          string
	Description string
	Amount      float64
	Periods     int64
	Date        string
}

type Output struct {
}

type GenerateInvoices struct{}

func NewGenerateInvoices() *GenerateInvoices {
	return &GenerateInvoices{}
}

func (gi *GenerateInvoices) Execute() []Contract {
	connStr := "user=user dbname=user sslmode=disable password=user host=127.0.0.1 port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT id, description, amount, periods, date FROM ken.contract")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var contracts []Contract
	for rows.Next() {
		var c Contract
		if errRow := rows.Scan(
			&c.ID,
			&c.Description,
			&c.Amount,
			&c.Periods,
			&c.Date); errRow != nil {
			log.Fatal(errRow)
		}
		contracts = append(contracts, c)
	}

	return contracts
}
