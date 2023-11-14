package repository

import (
	database "firstpro/db"
)

func AddRazorPayDetails(orderID string, razorPayOrderID string) error {

	err := database.DB.Exec("insert into razer_pays (order_id,razor_id) values (?,?)", orderID, razorPayOrderID).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckPaymentStatus(orderID string) (string, error) {

	var paymentStatus string
	err := database.DB.Raw("select payment_status from orders where order_id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}
	return paymentStatus, nil
}

func UpdatePaymentDetails(orderID string, paymentID string) error {

	err := database.DB.Exec("update razer_pays set payment_id = ? where order_id = ?", paymentID, orderID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateShipmentAndPaymentByOrderID(shipmentStatus string, paymentStatus string, orderID string) error {

	err := database.DB.Exec("update orders set payment_status = ?, shipment_status = ? where order_id = ?", paymentStatus, shipmentStatus, orderID).Error
	if err != nil {
		return err
	}

	return nil

}
