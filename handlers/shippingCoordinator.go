package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Shipment Coordinator Login
// @Description Login handler for ShipmnetCoordinator
// @Tags Shipment coordinator
// @Accept json
// @Produce json
// @Param  shipmentCoordinator body models.ShippingCoordinatorLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /shipping-coordinator-login [post]
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

// @Summary Update Shipment Status
// @Description Update shipment status from shipment coordinator's side
// @Tags Shipment Order Management
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param shipmentStatus  body models.Shipment_status true "Shipment status"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/approve-order/{id} [put]
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
