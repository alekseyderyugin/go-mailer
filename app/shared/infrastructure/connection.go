package infrastructure

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection() *Connection {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_DSN") + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Connection{DB: db}
}
