package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My handler"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	m := myHandler{}

	r.Handle("/handler", m)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		println("endpoint")
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

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"Message": "success"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product

		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}
