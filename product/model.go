package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Rating float32 `json:"rating"`
	Likes  int     `json:"likes"`
}

var products []Product

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}

	products = append(products, product)

	response, err := json.Marshal(&product)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func IndexByID(products []Product, id string) int {
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			return i
		}
	}
	return -1
}

func GetByIdProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := IndexByID(products, id)

	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products[index]); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := IndexByID(products, id)

	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	p := Product{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}

	products[index] = p

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&p); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := IndexByID(products, id)
	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	products = append(products[:index], products[index+1:]...)
	w.WriteHeader(http.StatusOK)
}
