package models

import "time"

type Product struct {
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UID         string    `json:"uid"     binding:"required"`
	ImageUrl    string    `json:"image_url"`
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"`
	Price       int       `json:"price" binding:"required"`
	IsDelete    bool      `json:"-" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProductRequest struct {
	ImageUrl    string `json:"image"`
	Name        string `json:"name"  binding:"required"`
	Description string `json:"des"`
	Price       int    `json:"price" binding:"required"`
}

type NoteUpdateRequest struct {
	ID          string `json:"id" binding:"required"`
	ImageUrl    string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"des"`
	Price       int    `json:"price"`
}
