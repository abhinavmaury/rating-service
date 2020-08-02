package driver

import (
"context"
"errors"
"github.com/gin-gonic/gin"
"net/http"
"rating-service/models/commonModel"
)

func DriverRating(c *gin.Context) {
	ctx := context.Background()
	var driver commonModel.User
	if err := c.BindJSON(&driver); err != nil {
		errors.New("invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   false,
			"response": "invalid request",
		})
		return
	}
	avgRating, err := commonModel.GetAggregateRating(&ctx, driver)
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

func DriverRide(c *gin.Context)  {
	ctx := context.Background()
	var driver commonModel.User
	if err := c.BindJSON(&driver); err != nil {
		errors.New("invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   false,
			"response": "invalid request",
		})
		return
	}
	driver.UserType = commonModel.Driver
	if err := commonModel.SetRating(&ctx, driver); err != nil {
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