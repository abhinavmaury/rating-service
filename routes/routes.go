package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rating-service/handler"
	"rating-service/routes/driverRouter"
	"rating-service/routes/passengerRouter"
)

func InitRoutes(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(cors.Default())
	api := router.Group("/api")
	apiRatingGroup := api.Group("/rating")

	passengerRouter.PassengerRoute(apiRatingGroup, "/passenger")
	driverRouter.DriverRoute(apiRatingGroup, "/driver")
	router.GET("/health", handler.HealthCheck)
}
