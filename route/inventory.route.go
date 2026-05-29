package route

import (
	"backend-umkm/handler/inventory"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/product", inventory.GetAllProduct)
	r.Post("/product", inventory.PostDataProduct)
	r.Put("/product/:id", inventory.UpdateDataProduct)
	r.Delete("/product/:id", inventory.DeleteDataProduct)
	r.Put("/product/stcok/:id", inventory.UpdateStockProduct)

	r.Get("/category", inventory.GetAllCategory)
	r.Delete("/category/:id", inventory.DeleteDataCategory)
	r.Post("/category", inventory.PostDataCategory)

}
