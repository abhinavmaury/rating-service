package driverRouter

import (
	"github.com/gin-gonic/gin"
	driver2 "rating-service/handler/driver"
)

func DriverRoute(router *gin.RouterGroup, path string) {
	driver := router.Group(path)
	driver.POST("", driver2.DriverRating)
	driver.POST("/ride", driver2.DriverRide)
}
