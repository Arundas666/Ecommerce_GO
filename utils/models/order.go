package models

type OrderDetails struct {
	OrderId        string
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}

type OrderProductDetails struct {
	ProductID  uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
type FullOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}