// handler/handler.go

package handler

import (
	"net/http"
)

// CustomHandler is your custom handler interface that includes the repository as a dependency.
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// Decorator is the interface for decorators.
type Decorator interface {
	Decorate(handler Handler) Handler
}

type HTTPRequest interface {
	Parse(r *http.Request) error
}

type HTTPResponse interface {
	Write(w http.ResponseWriter) error
}
