package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"rating-service/routes"
)

func main() {
	initMain()
	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run(":7000")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
