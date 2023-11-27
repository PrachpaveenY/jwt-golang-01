package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB //ตัวใหญ่ขึ้นต้นเพื่อให้ตัวอื่นใช้ได้
var err error

func InitDB() {
	dsn := os.Getenv("MYSQL_DNS")
	// dsn := "root:1234@tcp(127.0.0.1:3308)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&User{}) //db.AutoMigrate(&User{}, &Register{})
}

// close connection
// var db *sql.DB

// Get database url from environment variable, Create Table
// func InitDB() {
// 	var err error
// 	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal("Connect to database error", err)
// 	}

// 	createTb := `
// 	CREATE TABLE IF NOT EXISTS expenses (
// 		id SERIAL PRIMARY KEY,
// 		title TEXT,
// 		amount FLOAT,
// 		note TEXT,
// 		tags TEXT[] );
// 	`
// 	_, err = db.Exec(createTb)

// 	if err != nil {
// 		log.Fatal("can't create table", err)
// 	}

// }
