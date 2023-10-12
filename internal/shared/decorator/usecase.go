package decorator

type UseCase[T any] interface {
	Execute(input T) ([]byte, error)
}
