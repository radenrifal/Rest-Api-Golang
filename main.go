package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radenrifal/Rest-Api-Golang/controllers/productcontroller"
	"github.com/radenrifal/Rest-Api-Golang/models"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
