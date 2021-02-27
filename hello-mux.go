package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, GDG-ES!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
}
