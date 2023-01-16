package models

type OtherImages struct {
	Id         int    `json:"id" gorm:"primary_key;auto_increment"`
	ProductSKU string `json:"product_sku" gorm:"type:varchar REFERENCES products(sku) on DELETE CASCADE on UPDATE CASCADE;not null"`
	ImageURL   string `json:"image" gorm:"type:varchar;not null"`
}
