package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/database"
	AuthController "github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/login_logout_go/auth"
	UserController "github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/login_logout_go/user"
	"github.com/PrachpaveenY/golang-vuetify-crud-dashboard-csv/middleware"
)

func main() {
	fmt.Println("Welcome to Server")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	database.InitDB()

	r := gin.Default()
	r.Use(cors.Default()) //ทำให้ข้างนอกเรียก API เราได้ "github.com/gin-contrib/cors"
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/users", middleware.JWTAuthen()) //สร้างตัวแปรให้ authorized.GET("/users/readall" และ Group
	authorized.GET("/users/readall", UserController.ReadAll)
	authorized.GET("/profile", UserController.Profile)
	// authorized.PUT("/update/{id}", UserController.Update)
	// authorized.DELETE("/delete/{id}", UserController.Delete)
	r.Run("localhost:8083")

	// shutdown := make(chan os.Signal, 1)
	// signal.Notify(shutdown, os.Interrupt)
	// <-shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	// defer cancel()
	// if err := r.Shutdown(ctx); err != nil {
	// 	r.Logger.Fatal(err)
	// }
}

// database.InitDB()

// e := echo.New()

// e.Use(middleware.Logger())
// e.Use(middleware.Recover())

// e.POST("/expenses", database.CreateExpensesAllHandler)
// e.GET("/expenses", database.GetExpensesHandler)
// e.GET("/expenses/:id", database.GetExpensesIDHandler)
// e.PUT("/expenses/:id", database.UpdateAllExpensesHandler)
// e.PATCH("/expenses/:id", database.UpdateExpensesHandler)
// e.DELETE("/expenses/:id", database.DeleteExpensesHandler)

// type User struct {
// 	gorm.Model
// 	Username string
// 	Password string
// 	Fullname string
// 	Avatar   string
// }

// dsn := "root:1234@tcp(127.0.0.1:3308)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	panic("failed to connect database")
// }

// Migrate the schema
// db.AutoMigrate(&User{}) //db.AutoMigrate(&User{}, &Register{})

// r.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
