package main

import (
	"./api"
	"github.com/gin-gonic/gin"
)

func main() {
	db := api.ConnectDB()
	defer db.Close()

	router := gin.Default()

	// Routes
	api.SetupRoutes(router)

	router.Run(":8080")
}
