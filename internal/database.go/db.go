package database

import (
	"finbot/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	log.Println("✅ Успешное подключение к PostgreSQL")
}

func AutoMigrate() error {
	err := DB.AutoMigrate(&User{}, &Ballance{}, &Operation{})
	if err != nil {
		return fmt.Errorf("Ошибка автомиграции базы данных:%w", err)
	}
	log.Println("автомиграция прошла успешно")

	return nil
}
