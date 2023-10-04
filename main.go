package main

import (
	"firstpro/config"
	database "firstpro/db"
	"firstpro/routes"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
	}
	fmt.Println(cfg, "😊")
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println(db, "😢")

	
	router := gin.Default()
	routes.UserRoutes(router, db)

	listenAddr := fmt.Sprintf("%s:%s", cfg.DBPort, cfg.DBHost)
	fmt.Printf("Starting server on %s...\n", cfg.BASE_URL)
	if err := router.Run(cfg.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s: %v", listenAddr, err)
	}
	// fmt.Println("Starting server on port 8080...")
	// err = http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
