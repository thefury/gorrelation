// Package gorrelation provides middleware for adding a correlation id
// into the HTTP headers.
package gorrelation

import (
	"github.com/satori/go.uuid"
	"net/http"
)

// Gorrelation data structure
type Gorrelation struct {
	HeaderField string
	Value       string
}

// New constructs a new Gorrelation structure
func New() *Gorrelation {
	return &Gorrelation{HeaderField: "Correlation-Id"}
}

// Handler is a MiddlewareFunc that makes Gorrelation implement the Middleware interface
func (gr *Gorrelation) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gr.EnsureContextId(r)
		h.ServeHTTP(w, r)
	})
}

// EnsureContextId checks the incoming headers and adds a Context id if
// one does not already exist.
func (gr *Gorrelation) EnsureContextId(r *http.Request) {
	if len(r.Header.Get(gr.HeaderField)) == 0 {
		correlationId := uuid.NewV4()
		r.Header.Add(gr.HeaderField, correlationId.String())
	}
}
