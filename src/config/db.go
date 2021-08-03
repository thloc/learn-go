package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return db
}
