package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetVotingListItem(c *gin.Context) {
	var (
		votingListItem structs.VotingListItem
		result         gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingListItem).Error

	if err != nil {
		result = gin.H{
			"status":  403,
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving VotingListItem Data",
			"data_count": 1,
			"data":       votingListItem,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetVotingListItems(c *gin.Context) {
	var (
		votingListItems []structs.VotingListItem
		result          gin.H
	)

	idb.DB.Find(&votingListItems)

	if len(votingListItems) <= 0 {
		result = gin.H{
			"status":  404,
			"message": "VotingListItem is null",
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving VotingListItem Data",
			"data_count": len(votingListItems),
			"data":       votingListItems,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateVotingListItem(c *gin.Context) {
	var (
		votingListItem structs.VotingListItem
		result         gin.H
	)

	err := c.BindJSON(&votingListItem)

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Response",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Create(&votingListItem).Error

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
			"message": "Success Creating VotingListItem",
			"data":    votingListItem,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateVotingListItem(c *gin.Context) {
	var (
		votingListItem structs.VotingListItem
		result         gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingListItem).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = c.BindJSON(&votingListItem)

	if err != nil {
		result = gin.H{
			"status":  502,
			"message": "Error While Binding JSON",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		err = idb.DB.Save(&votingListItem).Error

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
				"message":  "Success Update VotingListItem Data",
				"new_data": votingListItem,
			}
			c.JSON(http.StatusOK, result)
			return
		}
	}
}

func (idb *InDB) DeleteVotingListItem(c *gin.Context) {
	var (
		votingListItem structs.VotingListItem
		result         gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingListItem).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Delete(&votingListItem).Error

	if err != nil {
		result = gin.H{
			"status":  304,
			"message": "Error Deleting VotingListItem",
		}
		c.JSON(http.StatusNotModified, result)
		return
	} else {
		result = gin.H{
			"status":  200,
			"message": "Success Delete VotingListItem",
		}
		c.JSON(http.StatusOK, result)
		return
	}
}
