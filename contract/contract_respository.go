package contract

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// Repository Persists Contracts Aggregates.Aggregates are clusters of data persistence.
type Repository interface {
	List() ([]Contract, error)
}

type DatabaseRepository struct{}

func NewDatabaseRepository() *DatabaseRepository {
	return &DatabaseRepository{}
}

func (dbRepo *DatabaseRepository) List() ([]Contract, error) {
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

	for i, c := range contracts {
		paymentsRow, paymentErr := db.Query("SELECT amount, date FROM ken.payment WHERE id_contract = $1", c.ID)
		if paymentErr != nil {
			return nil, paymentErr
		}

		var paymentsArray []Payment
		for paymentsRow.Next() {
			var p Payment
			if errRow := paymentsRow.Scan(&p.Amount, &p.Date); errRow != nil {
				return nil, errRow
			}
			paymentsArray = append(paymentsArray, p)
		}
		contracts[i].Payments = paymentsArray
	}
	return contracts, nil
}
