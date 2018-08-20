package gorrelation

import (
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

type Gorrelation struct {
	header string
}

func New() *Gorrelation {
	gorrelation := &Gorrelation{
		header: "Correlation-Id",
	}

	return gorrelation
}

func (gr *Gorrelation) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get(gr.header)) == 0 {
			correlationId, err := uuid.NewV4()
			if err != nil {
				log.Printf("could not assign correlation id: %+v", err)
			}

			r.Header.Add(gr.header, correlationId.String())
		}

		h.ServeHTTP(w, r)
	})
}
