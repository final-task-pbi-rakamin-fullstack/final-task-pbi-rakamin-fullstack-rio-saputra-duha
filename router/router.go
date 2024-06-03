package router

import (
	"github.com/gin-gonic/gin"
	controller "goAPI/controller"
	"goAPI/middleware"
	"net/http"
)

func UserRoutes(api *gin.RouterGroup) {

	api.GET("/home", Hello)
	//Users
	//api.GET("/:uuid", controller.GetUserFromID)
	api.POST("/register", controller.UserRegister)
	api.POST("/login", controller.UserLogin)
	api.GET("/logout", controller.UserLogout)
	api.PUT("/edit/:uuid", controller.UserUpdate)
	api.DELETE("/delete/:uuid", controller.UserDelete)

}

func PhotoRoutes(api *gin.RouterGroup) {
	//Photos
	api.Use(middleware.AuthMiddleware())
	api.GET("/", controller.PhotoIndex)
	api.POST("/add", controller.PhotoCreate)
	api.PUT("/edit/:uuid", controller.PhotoEdit)
	api.DELETE("delete/:uuid", controller.PhotoDelete)

}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello Welcome to home page")
}
