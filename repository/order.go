package repository

import (
	database "firstpro/db"
	"firstpro/utils/models"
	"fmt"
)

func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var orderDetails []models.OrderDetails
	fmt.Println("userid is", userId, "page is ", page, "count is ", count, "offset is", offset)
	database.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where user_id = ? limit ? offset ? ", userId, count, offset).Scan(&orderDetails)
	fmt.Println("order details is ", orderDetails)

	var fullOrderDetails []models.FullOrderDetails
	// for each order select all the associated products and their details
	for _, od := range orderDetails {

		var orderProductDetails []models.OrderProductDetails
		database.DB.Raw("select order_items.product_id,products.name as product_name,order_items.quantity,order_items.total_price from order_items inner join products on order_items.product_id = products.id where order_items.order_id = ?", od.OrderId).Scan(&orderProductDetails)
		fullOrderDetails = append(fullOrderDetails, models.FullOrderDetails{OrderDetails: od, OrderProductDetails: orderProductDetails})

	}

	return fullOrderDetails, nil

}

func GetOrderDetail(orderId int) (models.OrderDetails, error) {
	var OrderDetails models.OrderDetails

	if err := database.DB.Raw("select order_id,final_price,shipment_status,payment_status from orders where order_id = ?", orderId).Scan(&OrderDetails).Error; err != nil {
		return models.OrderDetails{}, err
	}
	return OrderDetails, nil
}
