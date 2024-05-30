package postgres

import (
	"fmt"
	"log"

	"github.com/ilhamgepe/prakerja-s7/config"
	"github.com/ilhamgepe/prakerja-s7/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Get.DB_HOST, config.Get.DB_USER, config.Get.DB_PASS, config.Get.DB_NAME, config.Get.DB_PORT)

	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.Users{}, &models.Products{})
	return db
}
