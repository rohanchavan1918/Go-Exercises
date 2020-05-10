package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"gophersize/gorestapi/controllers"
	"gophersize/gorestapi/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	City      string `json:"city"`
	Mobile    string `json:"mobile"`
}

func main() {
	db := models.SetupModels()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Get all users
	r.GET("/api/v1/Users/", controllers.GetUsers)
	// Get individual user details
	r.GET("/api/v1/Users/:id", controllers.GetUser)
	// Signup User
	r.POST("/api/v1/Users/", controllers.CreateUser)
	// Update User
	r.PUT("/api/v1/Users/:id", controllers.UpdateUser)
	// Delete User
	r.DELETE("/api/v1/Users/:id", controllers.DeleteUser)
	// User signin
	r.POST("/api/v1/auth/SignIn", controllers.SignIn)
	r.POST("/api/v1/auth/Refresh", controllers.Refresh)
	r.GET("/api/v1/auth/Welcome", controllers.Welcome)
	r.Run(":8080")
}
