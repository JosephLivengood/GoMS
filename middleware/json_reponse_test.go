package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestJsonResponse(t *testing.T) {

	t.Run("all requests have json type", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Use(JsonResponse)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

		req, _ := http.NewRequest("GET", "/", nil)

		r.ServeHTTP(recorder, req)
		res := recorder.Result()
		contentType := res.Header.Get("Content-Type")

		if contentType != "application/json" {
			t.Errorf("response is incorrect, got %s", contentType)
		}
	})

	t.Run("can be overridden", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Use(JsonResponse)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "gibberish")
		})

		req, _ := http.NewRequest("GET", "/", nil)

		r.ServeHTTP(recorder, req)
		res := recorder.Result()
		contentType := res.Header.Get("Content-Type")

		if contentType != "gibberish" {
			t.Errorf("response is incorrect, got %s", contentType)
		}
	})

	t.Run("errored requests have json types", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Use(JsonResponse)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		})

		req, _ := http.NewRequest("GET", "/", nil)

		r.ServeHTTP(recorder, req)
		res := recorder.Result()
		contentType := res.Header.Get("Content-Type")

		if contentType != "application/json" {
			t.Errorf("response is incorrect, got %s", contentType)
		}
	})

}
