package controller

import (
	"github.com/gin-gonic/gin"
	"goAPI/database"
	"goAPI/helper"
	"goAPI/models"
	"net/http"
	"time"
)

func PhotoIndex(c *gin.Context) {

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	//Display the user photos
	var photos []models.Photos
	if err := database.DB.Find(&photos, "user_id = ?", *decodedUuid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Check if the photos slice is empty
	if len(photos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No photos found for this user"})
		return
	}

	//Return the response
	c.JSON(http.StatusOK, gin.H{
		"Status":     http.StatusOK,
		"Photo Data": photos,
		"Errors":     nil,
	})

	//Created By Rafly Andrian
}

func PhotoCreate(c *gin.Context) {
	photo := new(models.Photos)
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Invalid request",
			Errors:   err.Error(),
		})
		return
	}

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	//Bind the input to the model
	photoAdd := models.Photos{
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    *decodedUuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//Created By Rafly Andrian

	//Create the photo
	result := database.DB.Create(&photoAdd)
	//Check if there is a error
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "Photo Created",
		Errors:   nil,
	})
}

func PhotoEdit(c *gin.Context) {
	//Create struct for User Edit
	var PhotoEdit struct {
		Title    string
		Caption  string
		PhotoUrl string
	}
	//Created By Rafly Andrian

	if err := c.ShouldBindJSON(&PhotoEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Parameter
	uuid := c.Param("uuid")

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	//Check if the photo is owned by the user logged in
	var photoCheck models.Photos
	checkUser := database.DB.First(&photoCheck, "uuid = ? AND user_id = ?", uuid, *decodedUuid).Error
	if checkUser != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":   http.StatusUnauthorized,
			"Messages": "You don't have permission to edit this photo",
			"Errors":   nil,
		})
		return
	}

	//Check if the photo is exists
	var photo models.Photos
	find := database.DB.First(&photo, "uuid = ?", uuid).Error
	if find != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":   http.StatusBadRequest,
			"Messages": "Photo is not exists",
			"Errors":   nil,
		})
		return
	}

	// Update the photos if it passes the check
	database.DB.Model(&photo).Updates(models.Photos{
		Title:     PhotoEdit.Title,
		Caption:   PhotoEdit.Caption,
		PhotoUrl:  PhotoEdit.PhotoUrl,
		UpdatedAt: time.Now(),
	})

	//Return the response
	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": PhotoEdit,
		"Errors":   nil,
	})
}

func PhotoDelete(c *gin.Context) {
	// Get Parameter
	uuid := c.Param("uuid")

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	//Check if the photo is owned by the user logged in
	var photoCheck models.Photos
	checkUser := database.DB.First(&photoCheck, "uuid = ? AND user_id = ?", uuid, *decodedUuid).Error
	if checkUser != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":   http.StatusUnauthorized,
			"Messages": "You don't have permission to delete this photo",
			"Errors":   nil,
		})
		return
	}

	//Check if the photo exists
	var photo models.Photos
	find := database.DB.First(&photo, "uuid = ?", uuid).Error
	if find != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":   http.StatusBadRequest,
			"Messages": "Photo is not exists",
			"Errors":   nil,
		})
		return
	}

	//Delete the photo if it passes the checks
	database.DB.Delete(&photo, "uuid = ?", uuid)

	//Return the response
	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": "Photo with caption " + photo.Caption + " has been deleted",
		"Errors":   nil,
	})
}
