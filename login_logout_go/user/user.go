package user

import (
	"fmt"
	"net/http"

	"github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/database"

	"github.com/gin-gonic/gin"
)

// คนที่ Login แล้วเท่านั้นถึงเรียกได้
func ReadAll(c *gin.Context) {
	fmt.Println("check : ", c.MustGet("userId"))
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success"})
	// var users []database.User
	// database.Db.Find(&users)
	// c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
	// readAll, err :=
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
}

// ส่ง Tiken(userId) เพื่อไปเรียก Profile
func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []database.User
	database.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": user})
}

// คนที่ Login แล้วเท่านั้นถึงเรียกได้
// func ReadAll(c *gin.Context) {
// 	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
// 	header := c.Request.Header.Get("Authorization")
// ตรวจ token
// tokenString := strings.Replace(header, "Bearer", "", 1)
// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//Don't forget to validate the alg is what you expect:
// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// }

//hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 	return hmacSampleSecret, nil
// })
// if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 	var users []database.User
// 	database.Db.Find(&users)
// 	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
// fmt.Println(claims["use"], claims["nbf"])
// } else {
// c.JSON(http.StatusOK, gin.H{"status": "forbidden", "message": err.Error()})
// return
// fmt.Println(err)
// }
// }
