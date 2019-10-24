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
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err != nil {
		result = gin.H{
			"status":  500,
			"message": "Internal Server Error",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	err = idb.DB.Create(&user).Error

	if err != nil {
		result = gin.H{
			"status":  501,
			"message": "Internal Server Error",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"status":  200,
			"message": "Success Creating User",
			"data":    user,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = c.BindJSON(&user)

	if err != nil {
		result = gin.H{
			"status":  502,
			"message": "Error While Binding JSON",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		err = idb.DB.Save(&user).Error

		if err != nil {
			result = gin.H{
				"status":  500,
				"message": "Failed to Save new Record",
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		} else {

			result = gin.H{
				"status":   200,
				"message":  "Success Update User Data",
				"new_data": user,
			}
			c.JSON(http.StatusOK, result)
			return
		}
	}
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Delete(&user).Error

	if err != nil {
		result = gin.H{
			"status":  304,
			"message": "Error Deleting User",
		}
		c.JSON(http.StatusNotModified, result)
		return
	} else {
		result = gin.H{
			"status":  200,
			"message": "Success Delete User",
		}
		c.JSON(http.StatusOK, result)
		return
	}
}

func (idb *InDB) Vote(c *gin.Context) {
	var (
		user   structs.UserVote
		result gin.H
	)

	err := c.BindJSON(&user)

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Response",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Create(&user).Error

	if err != nil {
		result = gin.H{
			"status":  501,
			"message": "Internal Server Error",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"status":  200,
			"message": "Voting Success",
			"data":    user,
		}
	}
	c.JSON(http.StatusOK, result)
}
