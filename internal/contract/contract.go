package contract

type Contract struct {
	ID          string
	Description string
	Amount      float64
	Periods     uint64
	Date        string
	Payments    []Payment
}

type Payment struct {
	Date   string
	Amount float64
}
