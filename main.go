package main

import (
	"html/template"
	"net/http"
)

// Product is a struct representation to a product in E-Commerce.
type Product struct {
	Name   string
	Desc   string
	Price  float64
	Amount int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Product{
		{Name: "Camiseta", Desc: "Azul escura com bordas", Price: 39.99, Amount: 10},
		{"Bermuda", "Jeans", 99.99, 20},
		{"Tenis", "Confortável", 89.99, 3},
		{"Fone", "Beats", 199.99, 30},
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
