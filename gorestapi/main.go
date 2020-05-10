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

	r.GET("/api/v1/Users/", controllers.GetUsers)
	r.GET("/api/v1/Users/:id", controllers.GetUser)
	r.POST("/api/v1/Users/", controllers.CreateUser)
	r.PUT("/api/v1/Users/:id", controllers.UpdateUser)
	r.DELETE("/api/v1/Users/:id", controllers.DeleteUser)

	r.Run(":8080")
}
