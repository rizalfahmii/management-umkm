package request

type StockLogRequest struct {
	Qty       int    `json:"qty" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Note      string `json:"note"`
	ProductId int    `json:"productId"`
}
