package router

import (
	"nasi-golang/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ProductRouter struct {
}

func (h ProductRouter) Initialize(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/product/list", handler.GetAllProducts)
}

func NewProductPouter() *ProductRouter {
	return &ProductRouter{}
}
