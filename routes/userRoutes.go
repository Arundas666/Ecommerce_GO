package routes

import (
	"firstpro/handlers"
	"firstpro/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.POST("/signup", handlers.Signup)
	r.POST("/login-with-password", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)

	// r.GET("/", handlers.ShowAllProducts)
	r.GET("/page/:page", handlers.ShowAllProducts)
	r.POST("/filter", handlers.FilterCategory)
	r.GET("/showaddress", middleware.AuthMiddleware(), handlers.GetAllAddress)
	r.POST("/add-address", middleware.AuthMiddleware(), handlers.AddAddress)

	//CART
	r.POST("/addtocart/:id", middleware.AuthMiddleware(), handlers.AddToCart)
	r.DELETE("/removefromcart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)
	r.GET("/displaycart", middleware.AuthMiddleware(), handlers.DisplayCart)
	r.DELETE("/emptycart", middleware.AuthMiddleware(), handlers.EmptyCart)

	//ORDERS
	r.GET("/orders/:page", middleware.AuthMiddleware(), handlers.GetOrderDetails)

	r.POST("/adminlogin", handlers.LoginHandler)

	r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBoard)

	return r

}
