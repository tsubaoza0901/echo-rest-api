package rdb

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB InitDB関数
func InitDB() *gorm.DB {
	// Heroku用に一時的に追加
	if os.Getenv("DATABASE_URL") != "" {
		dbURL := os.Getenv("DATABASE_URL")

		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		return db
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
