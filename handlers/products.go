package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllProducts(c *gin.Context) {
	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	count, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := usecase.ShowAllProducts(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)
}
