package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goAPI/database"
	helper "goAPI/helper"
	models "goAPI/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Errors   interface{} `json:"errors"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status   int    `json:"status"`
	Messages string `json:"messages"`
}

func GetUserFromID(c *gin.Context) {

	uuid := c.Param("uuid")

	var user models.Users
	if err := database.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Map the user data to the UserResponse struct
	userResponse := UserResponse{
		Email:    user.Email,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": userResponse,
		"Errors":   nil,
	})
}

func UserLogin(c *gin.Context) {
	var LoginRequest LoginRequest
	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check the user existence by selecting the email first
	var user models.Users
	err := database.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Status:   http.StatusUnauthorized,
			Messages: "Email or password is incorrect1",
		})
		return
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Status:   http.StatusUnauthorized,
			Messages: "Email or password is incorrect2",
		})
		return
	}

	//Generate JWT Token
	token, _ := helper.GenerateToken(user)
	//Set the JWT token to cookies
	helper.SetCookie(c, token)

	//Return Status Response
	c.JSON(http.StatusOK, LoginResponse{
		Status:   http.StatusOK,
		Messages: "Logined successfully!, Hello " + user.Username + "!",
	})
}

func UserRegister(c *gin.Context) {
	user := new(models.Users)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Invalid request",
			Errors:   err.Error(),
		})
		return
	}

	//Validate Input
	errorList := helper.ValidateUser(user)
	if errorList != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:   http.StatusBadRequest,
			Messages: "Validation Error",
			Errors:   errorList,
		})
		return
	}
	//Check Email
	var checkUser models.Users
	err := database.DB.First(&checkUser, "email = ?", user.Email).Error
	if err != nil {
		//Check if there is no record from query
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//If there is no email that user inputted exists
			//then create user
			//Hash the password
			hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			hashedPass := string(hash)
			//Bind the input
			userRegister := models.Users{
				Username:  user.Username,
				Email:     user.Email,
				Password:  hashedPass,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			//Create user
			result := database.DB.Create(&userRegister)
			if result.Error != nil {
				c.Status(400)
				return
			}
		} else {
			// Some other error occurred
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   http.StatusInternalServerError,
				"messages": "Internal server error",
			})
			return
		}
	} else {
		// Email exists, return the error response
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Status:   http.StatusUnauthorized,
			Messages: "Email already exists",
		})
		return
	}

	//Return status response
	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "User Created",
		Errors:   nil,
	})
}

func UserUpdate(c *gin.Context) {
	//Create struct for User Edit
	var UserEdit struct {
		Username string
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&UserEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get Parameter
	uuid := c.Param("uuid")

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	if uuid != *decodedUuid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":   http.StatusUnauthorized,
			"Messages": "You don't have permission to edit this user",
			"Errors":   nil,
		})
		return
	}

	// Find the data based on parameter uuid
	var user models.Users
	database.DB.First(&user, "uuid = ?", uuid)
	// Update the data
	database.DB.Model(&user).Updates(models.Users{
		Username:  UserEdit.Username,
		Email:     UserEdit.Email,
		Password:  UserEdit.Password,
		UpdatedAt: time.Now(),
	})

	//Return Status Response
	c.JSON(http.StatusOK, gin.H{
		"Status":   http.StatusOK,
		"Messages": UserEdit,
		"Errors":   nil,
	})
}

func UserDelete(c *gin.Context) {
	// Get Parameter
	uuid := c.Param("uuid")

	//Get the JWT token from cookies
	cookie, _ := c.Cookie("UserData")
	decodedUuid, _ := helper.ValidateToken(cookie)

	if uuid != *decodedUuid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":   http.StatusUnauthorized,
			"Messages": "You don't have permission to delete this user",
			"Errors":   nil,
		})
		return
	}

	//Check the user existence
	var user models.Users
	if err := database.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//If the user with the uuid exists, delete the user
	database.DB.Delete(&user, "uuid = ?", uuid)
	c.JSON(http.StatusOK, Response{
		Status:   http.StatusOK,
		Messages: "User Deleted",
		Errors:   nil,
	})
	//Logout the user if user deleted
	UserLogout(c)
}

// Created By Rafly Andrian
func UserLogout(c *gin.Context) {
	//Logout the current user by expiring the cookies
	c.SetCookie("UserData", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Message": "User Logouted Successfully",
	})
}
