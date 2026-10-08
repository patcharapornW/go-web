package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "stock_db"
)

func main() {
	// สร้างตัวแปรเก็บข้อมูลการเชื่อมต่อ
    dsn := "root:@tcp(localhost:3306)/stock_db?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
  		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
  		logger.Config{
    		SlowThreshold:              time.Second,   // Slow SQL threshold
    		LogLevel:                   logger.Silent, // Log level
    		Colorful:                  	true,          // Disable color
  		},
	)

	// สร้าง database connection
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Connection successful")

	// ทำ Migration ตัว Model กับ Database
	db.AutoMigrate(&Product{})
	fmt.Println("Migration success!")
}

