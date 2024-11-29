package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
}

var (
	AppConfig Config
	DB        *gorm.DB
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env, usando variables de entorno por defecto.")
	}

	AppConfig.Server.Port = getEnv("PORT", "8080")
	AppConfig.Database.DSN = getEnv("DATABASE_DSN", "user:password@tcp(localhost:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Local")

	// Conectar a la base de datos
	var dbErr error
	DB, dbErr = gorm.Open(mysql.Open(AppConfig.Database.DSN), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("Error conectando a la base de datos: %v", dbErr)
	}
	log.Println("Conexión a la base de datos establecida.")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
