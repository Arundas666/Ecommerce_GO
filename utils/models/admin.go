package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

// type AdminDetails struct {
// 	ID        uint   `json:"id" gorm:"uniquekey; not null"`
// 	Firstname string `json:"firstname"  gorm:"validate:required"`
// 	Lastname  string `json:"lastname"  gorm:"validate:required"`
// 	Email     string `json:"email"  gorm:"validate:required"`
// }

// type AdminSignUp struct {
// 	Name            string `json:"name" binding:"required" gorm:"validate:required"`
// 	Email           string `json:"email" binding:"required" gorm:"validate:required"`
// 	Password        string `json:"password" binding:"required" gorm:"validate:required"`
// 	ConfirmPassword string `json:"confirmpassword" binding:"required"`
// }

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"  `
	Lastname  string `json:"lastname" `
	Email     string `json:"email" `
}

// ADMIN DASHBOARD COMPLETE DETAILS

type DashboardUser struct {
	TotalUsers  int
	BlockedUser int
}

type DashBoardProduct struct {
	TotalProducts     int
	OutOfStockProduct int
}

type CompleteAdminDashboard struct {
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
}
