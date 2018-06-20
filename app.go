package main

import (
	"log"
	"net/http"

	. "./repository"
	"github.com/gorilla/mux"
)

// Routes definition
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/goods", AllGoodsEndPoint).Methods("GET")
	r.HandleFunc("/goods", CreateGoodEndPoint).Methods("POST")
	r.HandleFunc("/goods", UpdateGoodEndPoint).Methods("PUT")
	r.HandleFunc("/goods", DeleteGoodEndPoint).Methods("DELETE")
	r.HandleFunc("/goods/{id}", FindGoodEndpoint).Methods("GET")
	r.HandleFunc("/catalog/{id}", FindGoodHtmlEndpoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
