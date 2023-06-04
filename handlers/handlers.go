package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nakhun/web-server/services"

	"github.com/gorilla/mux"
)

type ProductDto struct {
	Title string `json:"title"`
	Price int64  `json:"price"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p ProductDto

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "createProduct <title: %s, price: %d>", p.Title, p.Price)
	services.CreateProduct(p.Title, p.Price)
}

func ReadProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetProducts()
	fmt.Fprintf(w, "%v", products)
	fmt.Fprintf(w, "getProducts")
}

func ReadProductDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	fmt.Fprintf(w, "You've requested the productId: %s", productId)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateProduct")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteProduct")
}
