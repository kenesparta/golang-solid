package decorator

import (
	"github.com/kenesparta/golang-solid/internal/invoice/usecases"
	"log"
)

type LoggerDecorator struct {
	usecase UseCase[usecases.Input]
}

func NewLoggerDecorator(usecase UseCase[usecases.Input]) *LoggerDecorator {
	return &LoggerDecorator{usecase}
}

func (ld *LoggerDecorator) Execute(input usecases.Input) ([]byte, error) {
	log.Println(input)
	return ld.usecase.Execute(input)
}
