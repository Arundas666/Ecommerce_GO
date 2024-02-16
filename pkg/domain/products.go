package domain

import "time"

type Products struct {
	ID                uint     `json:"id" gorm:"unique;not null"`
	Name              string   `json:"name"`
	SKU               string   `json:"sku"`
	CategoryID        uint     `json:"category_id"`
	Category          Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	Color             string   `json:"color"`
	Size              string   `json:"size"`
	Material          string   `json:"material"`
	DesignDescription string   `json:"design_description"`
	BrandID           uint     `json:"brand_id"`
	Quantity          int      `json:"quantity"`
	Price             float64  `json:"price"`
	ProductStatus     string   `json:"product_status"`
	IsDeleted         bool     `json:"is_deleted" gorm:"default:false"`
}

type Category struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"`
}

type ProductImages struct {
	ID              uint   `json:"id" gorm:"unique; not null"`
	ProductImageUrl string `json:"product_image_url"`
}

type ProductOffer struct {
	ID                 uint      `json:"id" gorm:"unique; not null"`
	ProductID          uint      `json:"product_id"`
	Products           Products  `json:"-" gorm:"foreignkey:ProductID"`
	OfferName          string    `json:"offer_name"`
	DiscountPercentage int       `json:"discount_percentage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	OfferLimit         int       `json:"offer_limit"`
	OfferUsed          int       `json:"offer_used"`
}

type CategoryOffer struct {
	ID                 uint      `json:"id" gorm:"unique; not null"`
	CategoryID         uint      `json:"category_id"`
	Category           Category  `json:"-" gorm:"foreignkey:CategoryID"`
	OfferName          string    `json:"offer_name"`
	DiscountPercentage int       `json:"discount_percentage"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	OfferLimit         int       `json:"offer_limit"`
	OfferUsed          int       `json:"offer_used"`
}
