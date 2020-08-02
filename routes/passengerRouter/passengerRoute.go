package passengerRouter

import (
	"github.com/gin-gonic/gin"
	"rating-service/handler/passanger"
)

func PassengerRoute(router *gin.RouterGroup, path string) {

	passenger := router.Group(path)
	passenger.POST("", passanger.PassengerRating)
	passenger.POST("/ride", passanger.PassengerRide)
}