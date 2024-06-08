package engine

import "net/http"

// Route represents a web handler with optional middlewares
type Route struct {
	Logger  bool
	Handler http.Handler
}
