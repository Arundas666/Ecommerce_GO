package repository

import (
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
)

func ShippingCoordinatorLogin(shippingCoordinatorDetails models.ShippingCoordinatorLogin) (domain.ShippingCoordinator, error) {
	var shippingCoordinatorCompareDetails domain.ShippingCoordinator
	if err := database.DB.Raw("select * from users where email = ? AND is_shipping_coordinator=true ", shippingCoordinatorDetails.Email).Scan(&shippingCoordinatorCompareDetails).Error; err != nil {
		return domain.ShippingCoordinator{}, err
	}
	return shippingCoordinatorCompareDetails, nil
}

func UpdateShipmentStatus(orderID string, shipmentStatus string) error {
	var shipmentStatusUpdater = make(map[string]string)
	shipmentStatusUpdater["1"] = "Processing"
	shipmentStatusUpdater["2"] = "Ready to Ship"
	shipmentStatusUpdater["3"] = "Shipped/Dispatched"
	shipmentStatusUpdater["4"] = "In Transit"
	shipmentStatusUpdater["5"] = "Out for Delivery"
	shipmentStatusUpdater["6"] = "Delivered"
	err := database.DB.Exec("update orders set shipment_status = ? where order_id = ?", shipmentStatusUpdater[shipmentStatus], orderID).Error
	if err != nil {
		return err
	}
	return nil
}
