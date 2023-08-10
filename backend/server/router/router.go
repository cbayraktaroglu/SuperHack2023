package router

import (
	"SuperHack2023/server/api/encode"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Endpoints
	router.POST("/encode", encode.Handler())

	return router
}
