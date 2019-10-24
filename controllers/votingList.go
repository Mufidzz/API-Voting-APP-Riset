package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetVotingList(c *gin.Context) {
	var (
		votingList structs.VotingList
		result     gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingList).Error

	if err != nil {
		result = gin.H{
			"status":  403,
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving VotingList Data",
			"data_count": 1,
			"data":       votingList,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetVotingLists(c *gin.Context) {
	var (
		votingLists []structs.VotingList
		result      gin.H
	)

	idb.DB.Find(&votingLists)

	if len(votingLists) <= 0 {
		result = gin.H{
			"status":  404,
			"message": "VotingList is null",
		}
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving VotingList Data",
			"data_count": len(votingLists),
			"data":       votingLists,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateVotingList(c *gin.Context) {
	var (
		votingList structs.VotingList
		result     gin.H
	)

	err := c.BindJSON(&votingList)

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Response",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Create(&votingList).Error

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
			"message": "Success Creating VotingList",
			"data":    votingList,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateVotingList(c *gin.Context) {
	var (
		votingList structs.VotingList
		result     gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingList).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = c.BindJSON(&votingList)

	if err != nil {
		result = gin.H{
			"status":  502,
			"message": "Error While Binding JSON",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		err = idb.DB.Save(&votingList).Error

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
				"message":  "Success Update VotingList Data",
				"new_data": votingList,
			}
			c.JSON(http.StatusOK, result)
			return
		}
	}
}

func (idb *InDB) DeleteVotingList(c *gin.Context) {
	var (
		votingList structs.VotingList
		result     gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&votingList).Error

	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = idb.DB.Delete(&votingList).Error

	if err != nil {
		result = gin.H{
			"status":  304,
			"message": "Error Deleting VotingList",
		}
		c.JSON(http.StatusNotModified, result)
		return
	} else {
		result = gin.H{
			"status":  200,
			"message": "Success Delete VotingList",
		}
		c.JSON(http.StatusOK, result)
		return
	}
}
