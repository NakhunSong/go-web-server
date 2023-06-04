package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "createProduct")
}

func ReadProducts(w http.ResponseWriter, r *http.Request) {
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
