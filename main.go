package main

import (
	"net/http"

	"github.com/gorilla/mux"

	handler "nakhun/web-server/handler"
	middleware "nakhun/web-server/middleware"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.Logging(handler.GetHome))

	fs := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", middleware.Logging(handler.CreateProduct)).Methods("POST")
	productRouter.HandleFunc("", middleware.Logging(handler.ReadProducts)).Methods("GET")
	productRouter.HandleFunc("/{productId}", middleware.Logging(handler.ReadProductDetail)).Methods("GET")
	productRouter.HandleFunc("/{productId}", middleware.Logging(handler.UpdateProduct)).Methods("PUT")
	productRouter.HandleFunc("/{productId}", middleware.Logging(handler.DeleteProduct)).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
