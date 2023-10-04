package domain
type Products struct {
	ID         uint   `json:"id" gorm:"unique;not null"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	CategoryID uint   `json:"category_id"`
	// Category           Category  `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	Color             string  `json:"color"`
	Size              string  `json:"size"`
	Material          string  `json:"material"`
	DesignDescription string  `json:"design_description"`
	BrandID           uint    `json:"brand_id"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	ProductStatus string	`json:"product_status"`
	IsDeleted         bool    `json:"is_deleted" gorm:"default:false"`
}