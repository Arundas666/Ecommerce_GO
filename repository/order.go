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

func UserOrderRelationship(orderID string, userID int) (int, error) {

	var testUserID int
	err := database.DB.Raw("select user_id from orders where order_id = ?", orderID).Scan(&testUserID).Error
	if err != nil {
		return -1, err
	}
	return testUserID, nil
}

func CancelOrders(orderID string) error {
	shipmentStatus := "cancelled"
	err := database.DB.Exec("update orders set shipment_status = ? where order_id = ?", shipmentStatus, orderID).Error
	if err != nil {
		return err
	}
	var paymentMethod int
	err = database.DB.Raw("select payment_method_id from orders where order_id = ?", orderID).Scan(&paymentMethod).Error
	if err != nil {
		return err
	}
	if paymentMethod == 3 || paymentMethod == 2 {
		err = database.DB.Exec("update orders set payment_status = 'refunded'  where order_id = ?", orderID).Error
		if err != nil {
			return err
		}
	}
	return nil

}
func GetProductDetailsFromOrders(orderID string) ([]models.OrderProducts, error) {

	var orderProductDetails []models.OrderProducts
	if err := database.DB.Raw("select product_id,quantity from order_items where order_id = ?", orderID).Scan(&orderProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}

	return orderProductDetails, nil
}
func GetShipmentStatus(orderID string) (string, error) {

	var shipmentStatus string
	err := database.DB.Raw("select shipment_status from orders where order_id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}

	return shipmentStatus, nil

}
func UpdateQuantityOfProduct(orderProducts []models.OrderProducts) error {

	for _, od := range orderProducts {

		var quantity int
		if err := database.DB.Raw("select quantity from products where id = ?", od.ProductId).Scan(&quantity).Error; err != nil {
			return err
		}

		od.Quantity += quantity
		if err := database.DB.Exec("update products set quantity = ? where id = ?", od.Quantity, od.ProductId).Error; err != nil {
			return err
		}
	}
	return nil

}

func CheckOrderID(orderID string) (bool, error) {

	var count int
	err := database.DB.Raw("select count(*) from orders where order_id = ?", orderID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}
func  ApproveOrder(orderID string) error {

	err := database.DB.Exec("update orders set shipment_status = 'order placed',approval = true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}

