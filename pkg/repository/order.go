package repository

import (
	database "firstpro/pkg/db"
	"firstpro/pkg/domain"
	"firstpro/pkg/utils/models"
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
func GetOrderDetailOfAproduct(orderId string) (models.OrderDetails, error) {
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
func ApproveOrder(orderID string) error {

	err := database.DB.Exec("update orders set shipment_status = 'order placed',approval = true where order_id = ?", orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrderDetailsByOrderId(orderID string) (models.CombinedOrderDetails, error) {
	var orderDetails models.CombinedOrderDetails
	err := database.DB.Raw("select orders.order_id,orders.final_price,orders.shipment_status,orders.payment_status,users.firstname,users.email,users.phone,addresses.house_name,addresses.state,addresses.pin,addresses.street,addresses.city from orders inner join users on orders.user_id = users.id inner join addresses on users.id = addresses.user_id where order_id = ?", orderID).Scan(&orderDetails).Error

	if err != nil {
		return models.CombinedOrderDetails{}, nil
	}
	return orderDetails, nil
}

func CreateOrder(orderDetails domain.Order) error {
	err := database.DB.Create(&orderDetails).Error
	if err != nil {
		return err
	}
	return nil
}

func AddOrderItems(orderItemDetails domain.OrderItem, UserID int, ProductID uint, Quantity float64) error {

	// after creating the order delete all cart items and also update the quantity of the product
	err := database.DB.Omit("id").Create(&orderItemDetails).Error
	if err != nil {
		return err
	}

	err = database.DB.Exec("delete from carts where user_id = ? and product_id = ?", UserID, ProductID).Error
	if err != nil {
		return err
	}

	err = database.DB.Exec("update products set quantity = quantity - ? where id = ?", Quantity, ProductID).Error
	if err != nil {
		return err
	}

	return nil

}

func GetBriefOrderDetails(orderID string) (domain.OrderSuccessResponse, error) {

	var orderSuccessResponse domain.OrderSuccessResponse
	database.DB.Raw("select order_id,shipment_status from orders where order_id = ?", orderID).Scan(&orderSuccessResponse)
	return orderSuccessResponse, nil

}
