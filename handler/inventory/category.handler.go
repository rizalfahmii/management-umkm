package inventory

import (
	"backend-umkm/database"
	"backend-umkm/models/entity"
	"backend-umkm/models/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(ctx *fiber.Ctx) error {
	var category []entity.Category
	result := database.DB.Find(&category)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(fiber.Map{
		"data": category,
	})

}

func PostDataCategory(ctx *fiber.Ctx) error {
	category := new(request.CategoryRequest)

	err := ctx.BodyParser(category)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	newCategory := entity.Category{
		Name: category.Name,
	}

	errCreate := database.DB.Create(&newCategory).Error

	if errCreate != nil {
		return ctx.JSON(fiber.Map{
			"message": "failde store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newCategory,
	})
}

func DeleteDataCategory(ctx *fiber.Ctx) error {
	cId := ctx.Params("id")

	if cId == "" {
		return ctx.JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	err := database.DB.Delete(cId)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed deleted data!",
		})
	}

	return ctx.JSON(fiber.Map{
		"messsage": "success deleted data !",
	})
}
