package routes

import (
	"firstpro/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.POST("/signup", handlers.Signup)
	r.POST("/login-with-password", handlers.UserLoginWithPassword)

	r.POST("/send-otp", handlers.SendOTP)
	r.POST("/verify-otp", handlers.VerifyOTP)



		r.GET("", handlers.ShowAllProducts)
		r.GET("/page/:page", handlers.ShowAllProducts)

	return r

}
