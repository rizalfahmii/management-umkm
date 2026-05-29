package request

type ProductRequest struct {
	Name       string `json:"name" validate:"required"`
	PriceBuy   int    `json:"priceBuy" validate:"required"`
	PriceSell  int    `json:"priceSell" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
}
