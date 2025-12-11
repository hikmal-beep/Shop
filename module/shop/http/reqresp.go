package http

// Request structs
type CreateShopRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type UpdateShopRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

// Response structs
type ShopResponse struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}