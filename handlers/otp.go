package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {
	var phone models.OTPData

	if err := c.ShouldBindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := usecase.SendOTP(phone.PhoneNumber)

	if err != nil {

		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
