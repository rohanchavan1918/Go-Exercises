package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")
	var User User
	d := db.Where("id = ?", id).Delete(&User)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
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
	db := c.MustGet("db").(*gorm.DB)
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
	db := c.MustGet("db").(*gorm.DB)
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
	db := c.MustGet("db").(*gorm.DB)
	var people []User
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}
