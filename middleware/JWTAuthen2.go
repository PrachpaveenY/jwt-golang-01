package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Check Token
func JWTAuthen2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Authorization"))
		// hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		fmt.Println("token header : ", header)
		// ตรวจ token
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		// tokenString := strings.TrimPrefix(header, "Bearer ")
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	//Don't forget to validate the alg is what you expect:
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	//hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		// 	return hmacSampleSecret, nil
		// })

		fmt.Println("token : ", tokenString)
		if err := validateToken(tokenString); err != nil {
			fmt.Println("error : ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {

		}

		// fmt.Println("token string : ", tokenString)
		// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// 	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		// 	}

		// 	return []byte("MySignature"), nil
		// })
		// fmt.Println("err : ", err)
		// // token.Valid => false
		// // token.Valid => false

		// fmt.Println("token : ", token)
		// claims, error := token.Claims.(jwt.MapClaims)
		// fmt.Println("token error : ", error)
		// fmt.Println("token check : ", claims)

		// // if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// // 	c.Set("userId", claims["userId"])
		// // 	// var users []database.User
		// // 	// database.Db.Find(&users)
		// // 	// c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
		// // 	// fmt.Println(claims["use"], claims["nbf"])
		// // } else {
		// // 	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "forbidden", "message": err.Error()})
		// // 	// return
		// // 	// fmt.Println(err)
		// // }

		// //Set example variable
		// // c.Set("example", "12345")

		// //before request
		// c.Next()
	}
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("my_secret_key"), nil
	})

	fmt.Println("ERROR:", err)

	return err
}

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// t := time.Now()

//Set example variable
// c.Set("example", "12345")

//before request
// c.Next()

//after request
// latency := time.Since(t)
// log.Print(latency)

//access the status we are sending
// status := c.Writer.Status()
// log.Println(status)
// 	}
// }

//ให้บาง request ที่เซ็ตไว้มาเรียกใช้ middleware เพื่อเช็ค Token ก่อน
