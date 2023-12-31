package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/babakkamali/note-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error){
	
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}

	dbhost := os.Getenv("MYSQL_HOST")
	dbportStr := os.Getenv("MYSQL_PORT")
	dbport, err := strconv.Atoi(dbportStr)
	dbname := os.Getenv("MYSQL_DBNAME")
	dbuser := os.Getenv("MYSQL_USER")
	dbpass := os.Getenv("MYSQL_PASSWORD")

	if err != nil {
		panic("Failed to convert MYSQL_PORT to integer: "+ err.Error())
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection failed")
	}

	fmt.Println(" db connected successfully")

	AutoMigrate(db)
	return db, nil
}

func AutoMigrate(connection *gorm.DB){
	connection.Debug().AutoMigrate(
		&models.User{},
		&models.AuthToken{},
		&models.Note{},
	)
}