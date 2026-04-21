package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá lá"))
	})

	r.Get("/{product}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "product")
		w.Write([]byte(param))
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")

		if name != "" {
			w.Write([]byte(name + " " + age))
		} else {
			w.Write([]byte("users"))
		}
	})

	http.ListenAndServe(":3000", r)
}
