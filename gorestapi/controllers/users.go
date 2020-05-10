package controllers

import (
	"fmt"
	"log"

	"gophersize/gorestapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var User models.User
	d := db.Where("id = ?", id).Delete(&User)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var User models.User
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
	db := c.MustGet("db").(*gorm.DB)
	var User models.User
	c.BindJSON(&User)
	// Check if the user already exists
	var email string = User.Email
	if err := db.Where("email = ?", email).First(&User).Error; err == nil {
		// Means user exists
		c.AbortWithStatusJSON(403, gin.H{"status": false, "message": "User already Exist"})
		fmt.Println(err)
	} else {
		// User doesnot exists proceed
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), 8)
		if err != nil {
			log.Fatal("Error in Hashing")
		}
		User.Password = string(hashedPassword)
		db.Create(&User)
		c.JSON(200, User)
	}
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var User models.User
	if err := db.Where("id = ?", id).First(&User).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, User)
	}
}
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var people []models.User
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}

// JWT AUth
func SignIn(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type Credentials struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var Creds Credentials
	var User models.User
	// Store user supplied password in mem map
	c.BindJSON(&Creds)
	email := Creds.Email
	userpassword := Creds.Password
	var expectedpassword string
	// check if the email exists
	if err := db.Where("email = ?", email).First(&User).Error; err == nil {
		// User Exists...Now compare his password with our password
		expectedpassword = User.Password
		if err = bcrypt.CompareHashAndPassword([]byte(expectedpassword), []byte(userpassword)); err != nil {
			// If the two passwords don't match, return a 401 status
			c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
		} else {
			fmt.Println("User AUthenticated.")
			c.JSON(200, "WELCOME")
		}
	} else {
		c.AbortWithStatusJSON(403, gin.H{"status": false, "message": "User Doesnot Exist"})
	}
}

func Refresh(c *gin.Context) {
	fmt.Println("Refresh")
}

func Welcome(c *gin.Context) {
	fmt.Println("Welcome")
}
