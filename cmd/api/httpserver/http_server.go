package httpserver

type HttpServerAdapter[T any] interface {
	Router() T
	Listen(port int)
}
