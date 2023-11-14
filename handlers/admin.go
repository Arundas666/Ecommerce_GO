package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Admin Login
// @Description Login handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param  admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/adminlogin [post]
func LoginHandler(c *gin.Context) { // login handler for the admin

	var adminDetails models.AdminLogin

	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := usecase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Admin Dashboard
// @Description Get Amin Home Page with Complete Details
// @Tags Admin Dash Board
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/dashboard [GET]
func DashBoard(c *gin.Context) {

	adminDashBoard, err := usecase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed fine", adminDashBoard, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Approve Order
// @Description Approve Order from admin side which is in processing state
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/approve-order/{id} [get]
func ApproveOrder(c *gin.Context) {

	orderId := c.Param("order_id")

	err := usecase.ApproveOrder(orderId)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not approve the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order approved successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Cancel Order Admin
// @Description Cancel Order from admin side
// @Tags Admin Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/cancel-order/{id} [get]
func CancelOrderFromAdminSide(c *gin.Context) {

	orderID := c.Param("order_id")

	err := usecase.CancelOrderFromAdminSide(orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Filtered Sales Report
// @Description Get Filtered sales report by week, month and year
// @Tags Admin Dash Board
// @Accept json
// @Produce json
// @Security Bearer
// @Param period path string true "sales report"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/sales-report/{period} [GET]
func FilteredSalesReport(c *gin.Context) {

	timePeriod := c.Param("period")
	salesReport, err := usecase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, successRes)

}
