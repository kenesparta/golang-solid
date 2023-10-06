package database

type Adapter[T, P any] interface {
	query(statement string, params T) P
}
