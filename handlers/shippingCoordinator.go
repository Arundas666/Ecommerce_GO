package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShippingCoordinatorLogin(c *gin.Context) { // login handler for the admin

	var shippingCoordinatorDetails models.ShippingCoordinatorLogin
	if err := c.ShouldBindJSON(&shippingCoordinatorDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := usecase.ShippingCoordinatorLogin(shippingCoordinatorDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "ShippingCoordinator authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

func UpdateShipmentStatus(c *gin.Context) {
	fmt.Println("Hey")
	orderID := c.Query("id")
	// shipment_status := c.Query("shipment-status")
	var shipmentStatus models.Shipment_status
	if err := c.ShouldBindJSON(&shipmentStatus); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := usecase.UpdateShipmentStatus(shipmentStatus.Shipment_status, orderID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Cannot update shipment status", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Shipment Status updated successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
