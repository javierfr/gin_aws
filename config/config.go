package config

import (
	"fmt"
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
		Host     string
		User     string
		Password string
		Name     string
		Port     string
	}
}

var (
	AppConfig Config
	DB        *gorm.DB
)

func LoadConfig() {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env, usando variables de entorno por defecto.")
	}

	// Cargar configuración del servidor
	AppConfig.Server.Port = getEnv("PORT", "8080")

	// Cargar configuración de la base de datos
	AppConfig.Database.Host = getEnv("DB_HOST", "localhost")
	AppConfig.Database.User = getEnv("DB_USER", "root")
	AppConfig.Database.Password = getEnv("DB_PASSWORD", "root")
	AppConfig.Database.Name = getEnv("DB_NAME", "gin_aws")
	AppConfig.Database.Port = getEnv("DB_PORT", "3306")

	// Construir la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name,
	)

	// Conectar a la base de datos
	var dbErr error
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
