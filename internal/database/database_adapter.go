package database

type Adapter[T, U any] interface {
	QueryAll() (T, error)
	Query(U) (U, error)
	Close()
}
