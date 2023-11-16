package middleware

import (
	"firstpro/helper"
	"firstpro/utils/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Authorization")
		fmt.Println(tokenHeader, "this is the token header")
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		// All good, set the Claims in the context
		c.Set("tokenClaims", tokenClaims)

		c.Next()

	}

}

func AuthorizationMiddlewareForShipmentCoordinator() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		fmt.Println(tokenHeader, "this is the token header")
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateTokenShippingCoordinator(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		// All good, set the Claims in the context
		c.Set("tokenClaims", tokenClaims)

		c.Next()

	}

}
