package request

type ProductStockRequest struct {
	Stock int `json:"stock" validate:"required"`
}
