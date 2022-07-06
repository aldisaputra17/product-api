package main

import (
	"fmt"
	"net/http"

	"github.com/aldisaputra17/product-API/product"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/product", product.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products", product.GetAllProduct).Methods("GET")
	r.HandleFunc("/product/{id}", product.GetByIdProduct).Methods("GET")
	r.HandleFunc("/product/{id}", product.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", product.DeleteProduct).Methods("DELETE")

	fmt.Println("Start Listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}
