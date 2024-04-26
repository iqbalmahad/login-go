package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
}
