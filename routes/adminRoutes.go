package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {

	r.POST("/adminlogin", handlers.LoginHandler)
	r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBoard)
	r.GET("/approve-order/:order_id", middleware.AuthorizationMiddleware(), handlers.ApproveOrder)
	r.GET("/cancel-order/:order_id", middleware.AuthorizationMiddleware(), handlers.CancelOrderFromAdminSide)
	r.GET("/sales-report/:period", handlers.FilteredSalesReport)
	r.POST("/add-coupon", middleware.AuthorizationMiddleware(), handlers.AddCoupon)
	r.POST("/add-product-offer", middleware.AuthorizationMiddleware(), handlers.AddProdcutOffer)
	r.POST("/add-category-offer", middleware.AuthorizationMiddleware(), handlers.AddCategoryOffer)
	r.GET("/coupons", middleware.AuthorizationMiddleware(), handlers.GetCoupon)
	r.PATCH("coupons/expire/:id", middleware.AuthorizationMiddleware(), handlers.ExpireCoupon)
}
