package domain

import "firstpro/pkg/utils/models"

type ShippingCoordinator struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenShippingCoordinator struct {
	ShippingCoordinator models.ShippingCoordinatorDetailsResponse
	Token               string
}
