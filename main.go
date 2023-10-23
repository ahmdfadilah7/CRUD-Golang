package main

import (
	"go-lang-web/config"
	"go-lang-web/controllers/categories"
	"go-lang-web/controllers/home"
	"go-lang-web/controllers/products"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// HomePage
	http.HandleFunc("/", home.Welcome)

	// Categories
	http.HandleFunc("/categories", categories.Index)
	http.HandleFunc("/categories/add", categories.Add)
	http.HandleFunc("/categories/edit", categories.Edit)
	http.HandleFunc("/categories/delete", categories.Delete)

	// Products
	http.HandleFunc("/products", products.Index)
	http.HandleFunc("/products/detail", products.Detail)
	http.HandleFunc("/products/add", products.Add)
	http.HandleFunc("/products/edit", products.Edit)
	http.HandleFunc("/products/delete", products.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
