package models

import "time"

type Category struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UID       string    `json:"uid"     binding:"required"`
	ImageUrl  string    `json:"image_url"`
	Name      string    `json:"name"  binding:"required"`
	IsDelete  bool      `json:"-" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CategoryRequest struct {
	ID       string `json:"id"`
	UID      string `json:"uid"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	IsDelete bool   `json:"is_delete"`
}
