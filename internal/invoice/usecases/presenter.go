package usecases

type Presenter interface {
	Present(output []Output) (string, error)
}
