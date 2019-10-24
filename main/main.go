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

	user := router.Group("/user")
	{
		user.GET("/:id", inDB.GetUser)
		user.GET("/", inDB.GetUsers)
		user.POST("/", inDB.CreateUser)
		user.PUT("/:id", inDB.UpdateUser)
		user.DELETE("/:id", inDB.DeleteUser)
		user.POST("/vote", inDB.Vote)
	}

	votingList := router.Group("/vl")
	{
		votingList.GET("/:id", inDB.GetVotingList)
		votingList.GET("/", inDB.GetVotingLists)
		votingList.POST("/", inDB.CreateVotingList)
		votingList.PUT("/:id", inDB.UpdateVotingList)
		votingList.DELETE("/:id", inDB.DeleteVotingList)
	}

	votingListItem := router.Group("/vli")
	{
		votingListItem.GET("/:id", inDB.GetVotingListItem)
		votingListItem.GET("/", inDB.GetVotingListItems)
		votingListItem.POST("/", inDB.CreateVotingListItem)
		votingListItem.PUT("/:id", inDB.UpdateVotingListItem)
		votingListItem.DELETE("/:id", inDB.DeleteVotingListItem)
	}

	err := router.Run(":7080")

	if err != nil {
		panic("Error when running router")
	}
}
