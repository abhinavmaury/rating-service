package passanger

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rating-service/models/commonModel"
)

func PassengerRating(c *gin.Context) {
	ctx := context.Background()
	var passenger commonModel.User
	if err := c.BindJSON(&passenger); err != nil {
		errors.New("invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   false,
			"response": "invalid request",
		})
		return
	}
	avgRating, err := commonModel.GetAggregateRating(&ctx, passenger)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":   false,
			"response": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"rating": avgRating,
	})
	return
}

func PassengerRide(c *gin.Context)  {
	ctx := context.Background()
	var passenger commonModel.User
	if err := c.BindJSON(&passenger); err != nil {
		errors.New("invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   false,
			"response": "invalid request",
		})
		return
	}
	passenger.UserType = commonModel.Passenger
	if err := commonModel.SetRating(&ctx, passenger); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":   false,
			"response": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"response": "Rating Accepted Successfully",
	})
	return
}