package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-or/vivo-synq/gateway/internal/repository"
)

type ProductHandler struct {
	productRepository repository.ProductRepository
}

func NewProductHandler(productRepository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepository: productRepository}
}

func (h *ProductHandler) FindUserProducts(ctx *gin.Context) {
	// Implementação da busca de produtos de um usuário
	// 1 - FindInCache -> SaveInCache
	// 2 - Find -> SaveInCache
	ctx.Status(http.StatusNotImplemented)
}
