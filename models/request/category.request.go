package request

type CategoryRequest struct {
	Name string `json:"Name" validate:"required"`
	
}
