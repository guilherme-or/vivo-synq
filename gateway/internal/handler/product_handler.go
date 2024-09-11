package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/guilherme-or/vivo-synq/gateway/internal/repository"
)

type ProductHandler struct {
	productRepository repository.ProductRepository
}

func NewProductHandler(productRepository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepository: productRepository}
}

func (h *ProductHandler) FindUserProducts(ctx *fiber.Ctx) error {
	// Implementação da busca de produtos de um usuário
	// 1 - FindInCache -> SaveInCache
	// 2 - Find -> SaveInCache
	return ctx.Status(http.StatusNotImplemented).JSON(fiber.Map{
		"message": "not implemented",
	})
}
