package main

import (
	"firstpro/cmd/api/docs"
	"firstpro/pkg/api/routes"
	"firstpro/pkg/config"
	database "firstpro/pkg/db"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title         Zog_festiv eCommerce API
// @version       1.0
// @description   API for ecommerce website

// @SecurityDefinition BearerAuth
// @TokenUrl /auth/token

// @securityDefinitions.Bearer		type apiKey
// @securityDefinitions.Bearer		name Authorization
// @securityDefinitions.Bearer		in header
// @securityDefinitions.BasicAuth	type basic
// @in header
// @name token
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host          www.arundas.cloud
// @BasePath      /
// @schemes       https
func main() {
	docs.SwaggerInfo.Title = "Ecommerce_site"
	docs.SwaggerInfo.Description = "Ecommerce shirt selling application suing Golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
	}
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(corsConfig))
	routes.UserRoutes(router.Group("/"), db)
	routes.AdminRoutes(router.Group("/admin"))
	routes.ShippingCoordinatorroutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	listenAddr := fmt.Sprintf("%s:%s", cfg.DBPort, cfg.DBHost)
	fmt.Printf("Starting server on %s...\n", cfg.BASE_URL)
	if err := router.Run(cfg.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s: %v", listenAddr, err)
	}
}
