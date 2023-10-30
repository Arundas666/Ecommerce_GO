package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get Order Details to user side
// @Description Get all order details done by user to user side
// @Tags User Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/orders/{id} [get]
func GetOrderDetails(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, _ := c.Get("user_id")
	userID := id.(int)
	fullOrderDetails, err := usecase.GetOrderDetails(userID, page, pageSize)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	fmt.Println("full order details is ", fullOrderDetails)

	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", fullOrderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Cancel order
// @Description Cancel order by the user using order ID
// @Tags User Order
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /users/cancel-order/{id} [put]
func CancelOrder(c *gin.Context) {

	orderID := c.Param("order_id")
	fmt.Println("ordr id ", orderID)

	id, _ := c.Get("user_id")
	userID := id.(int)

	err := usecase.CancelOrders(orderID, userID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Place Order
// @Description Place order from the user side
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Order ID"
// @Param id path string true "Payment"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/approve-order/{id} [get]
func PlaceOrder(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userId := userID.(int)
	orderId := c.Param("order_id")
	paymentMethod := c.Param("payment")

	fmt.Println("payment is ", paymentMethod, "order id is is ", orderId)

	if paymentMethod == "cash_on_delivery" {

		Invoice, err := usecase.ExecutePurchaseCOD(userId, orderId)
		if err != nil {
			errorRes := response.ClientResponse(http.StatusInternalServerError, "error in making cod ", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errorRes)
			return
		}
		successRes := response.ClientResponse(http.StatusOK, "Placed Order with cash on delivery", Invoice, nil)
		c.JSON(http.StatusOK, successRes)
	}
}
