package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

// Binding from JSON
type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	// FirstName string `json:"firstname" binding:"required"`
	// LastName string `json:"lastname" binding:"required"`
	// Email string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	// Check ว่ากรอกช่อง User ครบไหม
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check User Exists
	var userExist database.User
	database.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "massage": "User Exists"})
		return
	}

	// Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := database.User{Username: json.Username, Password: string(encryptedPassword),
		Fullname: json.Fullname, Avatar: json.Avatar}
	database.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"massage": "User Create Success",
			"userId":  user.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"massage": "User Create Failed",
		})
	}
}

// Binding from JSON
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	// Check ว่ากรอกช่อง Login ครบไหม
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User(Login) Exists
	var userExist database.User
	database.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "massage": "User(Login) Does Not Exists"})
		return
	}
	// ตรวจ bcrypt Password
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,                            //ฝัง UserID เข้าไปใน Token
			"exp":    time.Now().Add(time.Minute * 30).Unix(), //Token มีอายุ 1 นาที=Minute * 1
			// "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "massage": "Login Success", "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "massage": "Login Failed"})
	}
}
