package database

import (
	"database/sql"
	"github.com/kenesparta/golang-solid/internal/contract"
	"log"
)

type ContractPgAdapter struct {
	conn *sql.DB
}

func NewContractPgAdapter() *ContractPgAdapter {
	connStr := "user=user dbname=user sslmode=disable password=user host=127.0.0.1 port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	return &ContractPgAdapter{
		conn: db,
	}
}

func (pga *ContractPgAdapter) Query(c *contract.Contract) (*contract.Contract, error) {
	return nil, nil
}

func (pga *ContractPgAdapter) QueryAll() ([]contract.Contract, error) {
	rows, err := pga.conn.Query("SELECT id, description, amount, periods, date FROM ken.contract")
	if err != nil {
		log.Fatal(err)
	}

	var contracts []contract.Contract
	for rows.Next() {
		var c contract.Contract
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
		paymentsRow, paymentErr := pga.conn.Query("SELECT amount, date FROM ken.payment WHERE id_contract = $1", c.ID)
		if paymentErr != nil {
			return nil, paymentErr
		}

		// var paymentsArray []contract.Payment
		for paymentsRow.Next() {
			var p contract.Payment
			if errRow := paymentsRow.Scan(&p.Amount, &p.Date); errRow != nil {
				return nil, errRow
			}
			contracts[i].AddPayments(p)
		}
	}

	defer pga.Close()
	defer rows.Close()
	return contracts, nil
}

func (pga *ContractPgAdapter) Close() {
	err := pga.conn.Close()
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}
}
