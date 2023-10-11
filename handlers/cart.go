package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	id := c.Param("id")

	product_id, err := strconv.Atoi(id)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadGateway, "Prodcut id is given in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errResponse)
		return
	}
	// user_ID, _ := c.Get("user_id")
	user_ID := c.Request.Header.Get("User_id")

	user_id, _ := strconv.Atoi(user_ID)
	cartResponse, err := usecase.AddToCart(product_id, user_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := response.ClientResponse(200, "Added porduct Successfully to the cart", cartResponse, nil)
	c.JSON(200, successRes)

}
