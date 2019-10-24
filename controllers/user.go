package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (idb *InDB) GetUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		result = gin.H{
			"status":  403,
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving User Data",
			"data_count": 1,
			"data":       user,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		users  []structs.User
		result gin.H
	)

	idb.DB.Find(&users)

	if len(users) <= 0 {
		result = gin.H{
			"status":  404,
			"message": "User is null",
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving User Data",
			"data_count": len(users),
			"data":       users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	err := c.BindJSON(&user)

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Response",
		}
		c.JSON(http.StatusBadRequest, result)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err != nil {
		result = gin.H{
			"status":  500,
			"message": "Internal Server Error",
		}
		c.JSON(http.StatusInternalServerError, result)
	}

	err = idb.DB.Create(&user).Error

	if err != nil {
		result = gin.H{
			"status":  501,
			"message": "Internal Server Error",
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"status":  200,
			"message": "Success Creating User",
			"data":    user,
		}
	}

	c.JSON(http.StatusOK, result)
}
