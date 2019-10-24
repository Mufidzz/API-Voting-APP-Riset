package main

import (
	"../config"
	"../controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:id", inDB.GetUser)
	router.GET("/user", inDB.GetUsers)
	router.POST("/user", inDB.CreateUser)

	err := router.Run(":7080")

	if err != nil {
		panic("Error when running router")
	}
}
