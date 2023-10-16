package usecase

import (
	"errors"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"
)

func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil

}
func CancelOrders(orderId string, userId int) error {
	userTest, err := repository.UserOrderRelationship(orderId, userId)
	if err != nil {
		return err
	}
	if userTest != userId {
		return errors.New("the order is not done by this user")
	}
	orderProductDetails, err := repository.GetProductDetailsFromOrders(orderId)
	if err != nil {
		return err
	}
	shipmentStatus, err := repository.GetShipmentStatus(orderId)
	if err != nil {
		return err
	}
	if shipmentStatus == "delivered" {
		return errors.New("item already delivered, cannot cancel")
	}

	if shipmentStatus == "pending" || shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)
		return errors.New("the order is in" + message + ", so no point in cancelling")
	}

	if shipmentStatus == "cancelled" {
		return errors.New("the order is already cancelled, so no point in cancelling")
	}
	err = repository.CancelOrders(orderId)
	if err != nil {
		return err
	}
	err = repository.UpdateQuantityOfProduct(orderProductDetails)
	if err != nil {
		return err
	}
	return nil

}
func ExecutePurchaseCOD(userID int, orderID string) (models.Invoice, error) {
	ok, err := repository.CartExist(userID)
	if err != nil {
		return models.Invoice{}, err
	}
	if !ok {
		return models.Invoice{}, errors.New("cart doesnt exist")
	}
	
	err=repository.EmptyCart(userID)
	if err!=nil{
		return models.Invoice{},err
	}
	
	address, err := repository.GetAddressFromOrderId(orderID)
	if err != nil {
		return models.Invoice{}, err
	}
	orderDetails, err := repository.GetOrderDetailOfAproduct(orderID)
	if err != nil {
		return models.Invoice{}, err
	}

	Invoice := models.Invoice{
		OrderDetails: orderDetails,
		AddressInfo:  address,
		
	}
	return Invoice, nil

}
