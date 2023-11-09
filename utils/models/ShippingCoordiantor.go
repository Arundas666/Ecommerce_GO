package models

type ShippingCoordinatorLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type ShippingCoordinatorDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"  `
	Lastname  string `json:"lastname" `
	Email     string `json:"email" `
}
type Shipment_status struct {
	Shipment_status string `json:"shipment_status" validate:"required" `
}
