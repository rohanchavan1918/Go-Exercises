package controllers

import (
	"fmt"
	"log"
	"time"

	"gophersize/gorestapi/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("range")

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

			// User is AUthenticates, Now set the JWT Token

			// Set expiry time for the Jwt TOken
			expirationTime := time.Now().Add(5 * time.Minute)

			// Create the JWT claims, which includes the username and expiry time
			claims := &Claims{
				Email: User.Email,
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: expirationTime.Unix(),
				},
			}

			// Declare the token with the algorithm used for signing, and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				// If there is an error in creating the JWT return an internal server error
				c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
				return
			}

			fmt.Println("User AUthenticated.")
			c.SetCookie("token", tokenString, 6000, "/", "", false, false)
			c.JSON(200, "WELCOME")

		}
	} else {
		c.AbortWithStatusJSON(403, gin.H{"status": false, "message": "User Doesnot Exist"})
	}
}

func Refresh(c *gin.Context) {
	usercookie, err := c.Cookie("token")
	if err != nil {
		// If jwt session cookie is not set then return as 401
		if err.Error() == "ErrNoCookie" {
			c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
			return
		}
		// If something else, return as BAD request
		c.AbortWithStatusJSON(400, gin.H{"status": false, "message": "BAD REQUEST"})
		return
	}
	// If cookie is present,get jwt string from the cookie
	tknStr := usercookie
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
			return
		}
		c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
		return
	}
	if !tkn.Valid {
		c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
		return
	}
	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "Its Not the time"})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "message": "Internal Server Error"})
		return
	}
	// Set the new token as the users `token` cookie
	c.SetCookie("token", tokenString, 6000, "/", "", false, false)

}

func Welcome(c *gin.Context) {
	usercookie, err := c.Cookie("token")
	if err != nil {
		// If jwt session cookie is not set then return as 401
		if err.Error() == "ErrNoCookie" {
			c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
			return
		}
		// If something else, return as BAD request
		c.AbortWithStatusJSON(400, gin.H{"status": false, "message": "BAD REQUEST"})
		return
	}
	// If cookie is present,get jwt string from the cookie
	tknStr := usercookie
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
			return
		}
		c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
		return
	}
	if !tkn.Valid {
		c.AbortWithStatusJSON(401, gin.H{"status": false, "message": "UNAUTHORIZED"})
		return
	}
	email := claims.Email
	var logedinuser models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("email = ?", email).First(&logedinuser).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, logedinuser)
	}
}
