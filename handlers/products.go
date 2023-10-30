package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get Products Details to users
// @Description Retrieve all product Details with pagination to users
// @Tags User Product
// @Accept json
// @Produce json
// @Param page path string true "Page number"
// @Param count query string true "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/page/{page} [get]
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

// @Summary Show Products of specified category
// @Description Show all the Products belonging to a specified category
// @Tags User Product
// @Accept json
// @Produce json
// @Param data body map[string]int true "Category IDs and quantities"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/filter [post]
func FilterCategory(c *gin.Context) {
	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	productCategory, err := usecase.FilterCategory(data)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve products by category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productCategory, nil)
	c.JSON(http.StatusOK, successRes)
}
