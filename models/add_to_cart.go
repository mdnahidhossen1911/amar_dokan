package models

type AddToCart struct {
	ID        string   `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID string   `json:"product_id" binding:"required"`
	UserID    string   `json:"user_id" binding:"required"`
	Quantity  int      `json:"quantity" binding:"required"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type AddToCartRequest struct {
	ProductID string `json:"product_id" binding:"required"`
}

type AddToCartUpdateRequest struct {
	ID        string `json:"id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}
