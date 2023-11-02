package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/models"
	"firstpro/utils/response"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Add to Cart
// @Description Add product to the cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Param id path string true "product-id"
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/addtocart/{id} [post]
func AddToCart(c *gin.Context) {
	id := c.Param("product_id")

	product_id, err := strconv.Atoi(id)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadGateway, "Prodcut id is given in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errResponse)
		return
	}
	user_ID, _ := c.Get("user_id")
	// user_ID := c.Request.Header.Get("User_id")

	// user_id, _ := strconv.Atoi(user_ID)

	cartResponse, err := usecase.AddToCart(product_id, user_ID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}
	successRes := response.ClientResponse(200, "Added porduct Successfully to the cart", cartResponse, nil)
	c.JSON(200, successRes)

}

// @Summary Remove product from cart
// @Description Remove specified product of quantity 1 from cart using product id
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/removefromcart/{id} [delete]
func RemoveFromCart(c *gin.Context) {
	id := c.Param("product_id")
	product_id, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")
	updatedCart, err := usecase.RemoveFromCart(product_id, userID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "cannot remove product from the cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}
	succesRes := response.ClientResponse(200, "product removed successfully", updatedCart, nil)
	c.JSON(200, succesRes)

}

// @Summary Display Cart
// @Description Display all products of the cart along with price of the product and grand total
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart [get]
func DisplayCart(c *gin.Context) {

	userID, _ := c.Get("user_id")
	cart, err := usecase.DisplayCart(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot display cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Cart items displayed successfully", cart, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Delete all Items Present inside the Cart
// @Description Remove all product from cart
// @Tags User Cart
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart [delete]
func EmptyCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	cart, err := usecase.EmptyCart(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot empty the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Cart emptied successfully", cart, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Apply coupon on Checkout Section
// @Description Add coupon to get discount on Checkout section
// @Tags User Checkout
// @Accept json
// @Produce json
// @Security Bearer
// @Param couponDetails body models.CouponAddUser true "Add coupon to order"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /apply-coupon [post]
func ApplyCoupon(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var couponDetails models.CouponAddUser

	if err := c.ShouldBindJSON(&couponDetails); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not bind the coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := usecase.ApplyCoupon(couponDetails.CouponName, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "coupon could not be added", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Coupon added successfully", nil, nil)
	c.JSON(http.StatusCreated, successRes)
}
