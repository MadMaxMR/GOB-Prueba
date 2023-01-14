package models

type OtherImages struct {
	Id         string `json:"id" gorm:"primary_key;auto_increment"`
	ProductSKU string `json:"product_sku" gorm:"type:varchar REFERENCES products(sku)"`
	Image      string `json:"image" gorm:"type:varchar;not null"`
}
