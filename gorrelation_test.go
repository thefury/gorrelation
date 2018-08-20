package gorrelation

import (
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCorrelationId(t *testing.T) {
	g := New()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/test", nil)

	g.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 36, len(r.Header.Get(g.HeaderField)), "should return UUID length of 36")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})).ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
}

func TestCorrelationIdExists(t *testing.T) {
	g := New()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/test", nil)
	req.Header.Add(g.HeaderField, "example")

	g.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "example", r.Header.Get(g.HeaderField))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})).ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
}

func TestCorrelationIdHeaderDifferent(t *testing.T) {
	g := New()
	g.HeaderField = "Some-Other-Header"

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://example.com/test", nil)

	g.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, 36, len(r.Header.Get(g.HeaderField)), "should return UUID length of 36")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})).ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusOK)
}
