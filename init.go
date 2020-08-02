package main

import (
	"fmt"
	"os"
	"rating-service/db"
)

func initMain() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}