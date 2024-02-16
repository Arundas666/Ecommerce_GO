package routes

import (
	"firstpro/pkg/api/handlers"
	"firstpro/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func ShippingCoordinatorroutes(r *gin.Engine) *gin.Engine {
	//shippingcoordinator
	r.POST("/shipping-coordinator-login", handlers.ShippingCoordinatorLogin)
	r.POST("/update-shipment-status", middleware.AuthorizationMiddlewareForShipmentCoordinator(), handlers.UpdateShipmentStatus)

	return r

}
