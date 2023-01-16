package models

type Product struct {
	SKU            string        `json:"sku" gorm:"primary_key" binding:"required"`
	Name           string        `json:"name" gorm:"type:varchar;size:3,50;not null" binding:"required"`
	Brand          string        `json:"brand" gorm:"type:varchar;size:3,50;not null" binding:"required"`
	Size           string        `json:"size" gorm:"type:varchar;not null"`
	Price          float32       `json:"price" gorm:"type:float;not null" binding:"required"`
	PrincipalImage string        `json:"image_urls" gorm:"type:varchar;not null" binding:"required"`
	OtherImages    []OtherImages `json:"other_images"`
}
