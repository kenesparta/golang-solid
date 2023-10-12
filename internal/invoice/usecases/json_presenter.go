package usecases

import "encoding/json"

type JsonPresenter struct{}

func NewJsonPresenter() *JsonPresenter {
	return &JsonPresenter{}
}

func (*JsonPresenter) Present(output []Output) ([]byte, error) {
	bytes, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
