package repository

import (
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
)

func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := database.DB.Raw("select * from users where email = ? AND isadmin=true ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}
func DashboardUserDetails() (models.DashboardUser, error) {

	var userDetails models.DashboardUser
	err := database.DB.Raw("select count(*) from users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	err = database.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	return userDetails, nil
}

func DashBoardProductDetails() (models.DashBoardProduct, error) {

	var productDetails models.DashBoardProduct
	err := database.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	err = database.DB.Raw("select count(*) from products where quantity = 0").Scan(&productDetails.OutOfStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	return productDetails, nil
}

func DashBoardOrder() (models.DashboardOrder, error) {

	var orderDetails models.DashboardOrder
	err := database.DB.Raw("select count(*) from orders where payment_status = 'paid' and approval = true ").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders where shipment_status = 'pending' or shipment_status = 'processing'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders where shipment_status = 'cancelled'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select count(*) from orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = database.DB.Raw("select sum(quantity) from order_items").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}
	return orderDetails, nil
}
