package main

//Created By Rafly Andrian
import (
	"github.com/gin-gonic/gin"
	database "goAPI/database"
	"goAPI/migrate"
	router "goAPI/router"
	"os"
)

func init() {
	database.LoadEnvVariables()
	database.DBConnect()
	migrate.SyncDB()
}

func main() {
	r := gin.Default()
	usersApi := r.Group("/users")
	photosApi := r.Group("/photos")
	router.UserRoutes(usersApi)
	router.PhotoRoutes(photosApi)

	r.Run("localhost:" + os.Getenv("PORT"))
}

//Created By Rafly Andrian
