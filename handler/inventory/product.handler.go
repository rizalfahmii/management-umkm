package inventory

import (
	"backend-umkm/database"
	"backend-umkm/models/entity"
	"backend-umkm/models/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(ctx *fiber.Ctx) error {
	var Product []entity.Product
	result := database.DB.Find(&Product)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(fiber.Map{
		"data": Product,
	})

}

func PostDataProduct(ctx *fiber.Ctx) error {
	product := new(request.ProductRequest)

	err := ctx.BodyParser(product)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	newProduct := entity.Product{
		Name:       product.Name,
		PriceBuy:   product.PriceBuy,
		PriceSell:  product.PriceSell,
		Stock:      product.Stock,
		CategoryId: product.CategoryId,
	}
	log.Println(product)
	errCreate := database.DB.Create(&newProduct).Error

	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "succes create data",
		"data":    newProduct,
	})

}

func UpdateDataProduct(ctx *fiber.Ctx) error {
	request := new(request.ProductRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "invalid request!",
		})
	}

	userId := ctx.Params("id")
	if userId == "" {
		return ctx.JSON(fiber.Map{
			"message": "user id not found!",
		})
	}

	var product entity.Product
	result := database.DB.First(&product, userId)
	if result != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found!",
		})
	}

	product.Name = request.Name
	product.PriceBuy = request.PriceBuy
	product.PriceSell = request.PriceSell
	product.Stock = request.Stock

	errDb := database.DB.Save(product)
	if errDb != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed update",
		})
	}

	return ctx.JSON(fiber.Map{
		"data": product,
	})
}

func UpdateStockProduct(ctx *fiber.Ctx) error {
	reqStock := new(request.ProductStockRequest)
	errReq := ctx.BodyParser(reqStock)
	if errReq != nil {
		return ctx.JSON(fiber.Map{
			"message": "invalid request!",
		})
	}
	pId := ctx.Params("id")

	if pId == "" {
		return ctx.JSON(fiber.Map{
			"message": "invalid id product",
		})
	}

	var product entity.Product

	errDb := database.DB.First(&product, pId)
	if errDb != nil {
		return ctx.JSON(fiber.Map{
			"message": "product not found",
		})
	}

	product.Stock = reqStock.Stock

	err := database.DB.Save(product)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed update stock",
		})
	}
	return ctx.JSON(fiber.Map{
		"data": product,
	})
}

func DeleteDataProduct(ctx *fiber.Ctx) error {
	pId := ctx.Params("id")
	if pId == "" {
		return ctx.JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	errDb := database.DB.Delete(pId)
	if errDb != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed deleted data!",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success deleted data!",
	})
}

func PostDataStocklog(ctx *fiber.Ctx) error {
	stocklog := new(request.StockLogRequest)

	err := ctx.BodyParser(stocklog)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	newStockLog := entity.Stocklog{
		Type:      stocklog.Type,
		Qty:       stocklog.Qty,
		Note:      stocklog.Note,
		ProductId: stocklog.ProductId,
	}
	errDB := database.DB.Create(&newStockLog).Error
	if errDB != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed store data",
		})
	}
	database.DB.Preload("Product.Category").First(&newStockLog, newStockLog.ID)

	//ambil product dulu
	var product entity.Product
	errGet := database.DB.First(&product, stocklog.ProductId).Error
	if errGet != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	//hitung product
	if stocklog.Type == "IN" {
		product.Stock += stocklog.Qty
	} else if stocklog.Type == "OUT" {
		product.Stock -= stocklog.Qty
	}
	if product.Stock < 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "stok tidak valid",
		})
	}
	database.DB.Save(&product)
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newStockLog,
	})
}
