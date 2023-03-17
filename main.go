package main

import (
	"net/http"
	"path/mod/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/SecondPage", routes.SecondPage)

	http.ListenAndServe(":8080", r)
}
