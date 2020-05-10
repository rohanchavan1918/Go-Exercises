package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	db, _ = gorm.Open("mysql", "root:@/gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/api/v1/Users/", GetUsers)
	r.GET("/api/v1/Users/:id", GetUser)
	r.POST("/api/v1/Users/", CreateUser)
	r.PUT("/api/v1/Users/:id", UpdateUser)
	r.DELETE("/api/v1/Users/:id", DeleteUser)

	r.Run(":8080")
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var User User
	d := db.Where("id = ?", id).Delete(&User)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateUser(c *gin.Context) {

	var User User
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&User)

	db.Save(&User)
	c.JSON(200, User)

}

func CreateUser(c *gin.Context) {
	var User User
	c.BindJSON(&User)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), 8)
	if err != nil {
		log.Fatal("Error in Hashing")
	}
	User.Password = string(hashedPassword)
	db.Create(&User)
	c.JSON(200, User)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var User User
	if err := db.Where("id = ?", id).First(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, User)
	}
}
func GetUsers(c *gin.Context) {
	var people []User
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}
