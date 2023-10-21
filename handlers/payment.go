package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func MakePaymentRazorPay(c *gin.Context) {
	orderID := c.Query("id")
	userID := c.Query("user_id")
	user_Id, _ := strconv.Atoi(userID)
	orderDetail, razorID, err := usecase.MakePaymentRazorPay(orderID, user_Id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"final_price": orderDetail.FinalPrice * 100,
		"razor_id":    razorID,
		"user_id":     userID,
		"order_id":    orderDetail.OrderId,
		"user_name":   orderDetail.Name,
		"total":       int(orderDetail.FinalPrice),
	})
}
func VerifyPayment(c *gin.Context) {
	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")

	err := usecase.SavePaymentDetails(paymentID, orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
