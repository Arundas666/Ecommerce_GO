package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	// r.GET("/", handlers.ShowAllProducts)

	r.GET("/products/page/:page", handlers.ShowAllProducts)
	r.POST("/products/filter", handlers.FilterCategory)

	r.POST("/crop", handlers.CropImage)

	//PAYMENT

	r.GET("/payment", handlers.MakePaymentRazorPay)
	r.GET("/payment-success", handlers.VerifyPayment)

	//admin

	r.Group("/users")

	r.Use(middleware.AuthMiddleware())
	{

		r.GET("/address", handlers.GetAllAddress)
		r.POST("/address", handlers.AddAddress)
		r.GET("/show-user-details", handlers.UserDetails)
		r.PATCH("/edit-user-profile", handlers.UpdateUserDetails)
		r.POST("/update-password", handlers.UpdatePassword)

		//CART
		r.POST("/addtocart/:product_id", handlers.AddToCart)
		r.DELETE("/removefromcart/:product_id", handlers.RemoveFromCart)
		r.GET("/cart", handlers.DisplayCart)
		r.DELETE("/emptycart", handlers.EmptyCart)

		//ORDERS
		r.GET("/orders/:page", handlers.GetOrderDetails)
		r.PUT("/cancel-orders/:order_id", handlers.CancelOrder)
		r.GET("/checkout", handlers.CheckOut)
		r.GET("/place-order/:order_id/:payment", handlers.PlaceOrder)
		r.POST("/apply-coupon", handlers.ApplyCoupon)
		r.GET("/referral/apply", handlers.ApplyReferral)

	}

}
