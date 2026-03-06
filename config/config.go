package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort          string
	DBHost           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBPort           string
	JWTSecret        string
	JWTExpireMinutes string
	JWTRefreshToken  string
	JWTExpire        string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}

	AppConfig = &Config{
		AppPort:          getEnv("PORT", "3015"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", "123456"),
		DBName:           getEnv("DB_NAME", "project_management"),
		DBPort:           getEnv("DB_PORT", "5432"),
		JWTSecret:        getEnv("JWT_SECRET", "secret"),
		JWTExpireMinutes: getEnv("JWT_EXPIRE_MINUTES", "60"),
		JWTRefreshToken:  getEnv("JWT_REFRESH_TOKEN", "refresh_token"),
		JWTExpire:        getEnv("JWT_EXPIRE", "24h"),
	}
}

func getEnv(key string, fallback string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	} else {
		return fallback
	}
}

func ConnectDB() {
	cfg := AppConfig
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	fmt.Println("Database connected successfully")
}
