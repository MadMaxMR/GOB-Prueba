package models

type Product struct {
	SKU            string        `json:"sku" gorm:"primary_key"`
	Name           string        `json:"name" gorm:"type:varchar;size:3,50;not null;required" `
	Brand          string        `json:"brand" gorm:"type:varchar;size:3,50;not null;required"`
	Size           string        `json:"size" gorm:"type:varchar;not null"`
	Price          float32       `json:"price" gorm:"type:float;not null;required"`
	PrincipalImage string        `json:"principal_image" gorm:"type:varchar;required;not null"`
	OtherImages    []OtherImages `json:"other_images"`
}
