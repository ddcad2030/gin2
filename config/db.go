package config

import (
	"github.com/ddcad2030/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=db_postgres user=root password=root dbname=db_test port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.User{})
	DB = db
}
