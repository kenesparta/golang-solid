package usecases

type Presenter interface {
	Present(output []Output) ([]byte, error)
}
