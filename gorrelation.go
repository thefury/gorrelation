// Package gorrelation provides middleware for adding a correlation id
// into the HTTP headers.
package gorrelation

import (
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

// Gorrelation data structure
type Gorrelation struct {
	HeaderField string
}

// New constructs a new Gorrelation structure
func New() *Gorrelation {
	gorrelation := &Gorrelation{
		HeaderField: "Correlation-Id",
	}

	return gorrelation
}

// Handler is a MiddlewareFunc that makes Gorrelation implement the Middleware interface
func (gr *Gorrelation) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get(gr.HeaderField)) == 0 {
			correlationId, err := uuid.NewV4()
			if err != nil {
				log.Printf("could not assign correlation id: %+v", err)
			}

			r.Header.Add(gr.HeaderField, correlationId.String())
		}

		h.ServeHTTP(w, r)
	})
}
