package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/guilherme-or/vivo-synq/api/internal/repository"
)

// Controlador de requisições dos produtos de um usuário
type ProductHandler struct {
	productRepository repository.ProductRepository
}

func NewProductHandler(productRepository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepository: productRepository}
}

func (h *ProductHandler) FindUserProducts(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")

	// 1 - FindInCache -> SaveInCache
	products, err := h.productRepository.FindInCache(userID)
	if err != nil {
		// 2 - FindInCache -> Error -> Find -> SaveInCache
		log.Print("FindInCache Error: ", err)

		products, err = h.productRepository.Find(userID)
		if err != nil {
			// 3 - FindInCache -> Error -> Find -> Error
			log.Print("Find Error: ", err)
			return ctx.Status(fiber.StatusNotFound).JSON(ErrNotFound)
		}

		if err := h.productRepository.SaveInCache(userID, products); err != nil {
			log.Print("SaveInCache Error: ", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}
