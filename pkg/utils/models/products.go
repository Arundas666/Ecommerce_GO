package models

type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category_name"`
}
type ProductBrief struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	SKU           string  `json:"sku"`
	BrandID       uint    `json:"brand_id"`
	Color         string  `json:"color"`
	Size          string  `json:"size"`
	Material      string  `json:"material"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	ProductStatus string  `json:"product_status"`
}

type ProductImage struct {
	ID      uint   `json:"id" gorm:"unique; not null"`
	ImageID string `json:"image_id"`
}
