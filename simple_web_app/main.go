package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my basic web app </h1><a href=\"http://localhost:3000/articles/\"><h4 color=\"red\">ARticles</h4></a><a href=\"http://localhost:3000/products/\"><h4 color=\"red\">Products</h4></a>")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Product !!!</h1>")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Articles!!!</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/", ProductsHandler)
	r.HandleFunc("/articles/", ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3000", r)
}
