package main

import (
	"fmt"
	"gophersize/gorange/controllers"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	// "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "email"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	fmt.Println(claims)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"email":     claims[identityKey],
		"firstName": user.(*User).FirstName,
	})
}

// User demo
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	City      string `json:"city"`
	Mobile    string `json:"mobile"`
}

func CheckCredentials(useremail, userpassword string, db *gorm.DB) {
	// db := c.MustGet("db").(*gorm.DB)
	var User User
	// Store user supplied password in mem map
	var expectedpassword string
	// check if the email exists
	if err := db.Where("email = ?", useremail).First(&User).Error; err == nil {
		// User Exists...Now compare his password with our password
		expectedpassword = User.Password
		if err = bcrypt.CompareHashAndPassword([]byte(expectedpassword), []byte(userpassword)); err != nil {
			// If the two passwords don't match, return a 401 status
			log.Println("User is Not Authorized")
			// return false
		}
		fmt.Println("Err in bcrypt...", err)
		// User is AUthenticates, Now set the JWT Token
		fmt.Println("User Verified, found")
		fmt.Println(User.FirstName, User.LastName, User.Email, User.Mobile)
		// return true
	} else {
		fmt.Println("trying to find username but this err found", err)
	}
	// return false
}

// SetupModels will start the database
// func SetupModels() *gorm.DB {

// 	db, err := gorm.Open("mysql", "root:@/gotest?charset=utf8&parseTime=True&loc=Local")

// 	if err != nil {
// 		fmt.Println(err)
// 		panic("error db")
// 	}
// 	log.Println("DATABASE STARTED SUCCESSFULLY")
// 	fmt.Println("DATABASE STARTED SUCCESSFULLY")
// 	// defer db.Close()
// 	db.AutoMigrate(&User{})
// 	return db
// }

func main() {
	port := os.Getenv("PORT")
	// SetupModels()
	db, err := gorm.Open("mysql", "root:@/gotest?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("DATABASE ERROR OCCOURED ", err)
		panic("error db")
	} else {
		fmt.Println("DATABASE STARTED SUCCESSFULLY")
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password
			CheckCredentials(email, password, db)

			if (email == "admin@range.com" && password == "admin") || (email == "react@range.com" && password == "react") {
				return &User{
					Email:     email,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// Allow Users to access some special pages
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok {
				fmt.Println("AUthorization v is ", v.FirstName, v.Email)
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/whoami", controllers.GetIDFromEmail)
		// Get all Users
		auth.GET("/api/v1/Users/", controllers.GetUsers)
		// Get individual user details
		auth.GET("/api/v1/Users/:id", controllers.GetUser)
		// Signup User
		auth.POST("/api/v1/Users/", controllers.CreateUser)
		// Update User
		auth.PUT("/api/v1/Users/:id", controllers.UpdateUser)
		// Delete User
		auth.DELETE("/api/v1/Users/:id", controllers.DeleteUser)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
