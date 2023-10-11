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

	r.POST("/addtocart/:id", middleware.AuthMiddleware(), handlers.AddToCart)
	r.DELETE("/removefromcart/:id", middleware.AuthMiddleware(), handlers.RemoveFromCart)

	r.POST("/adminlogin", handlers.LoginHandler)

	r.GET("/dashboard", middleware.AuthorizationMiddleware(), handlers.DashBoard)

	return r

}
