package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func CancelOrder(c *gin.Context) {

	orderID := c.Param("id")
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
