package main

import (
	"net/http"

	"github.com/guilherme-or/vivo-synq/api/internal/handler"
)

func main() {
	h := handler.NewUserProductsHandler(nil)
	http.HandleFunc("/users/{id}/products", h.GetUserProductsByID)
}
