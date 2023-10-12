package usecases

import "encoding/json"

type JsonPresenter struct{}

func NewJsonPresenter() *JsonPresenter {
	return &JsonPresenter{}
}

func (*JsonPresenter) Present(output []Output) (string, error) {
	bytes, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
