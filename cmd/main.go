package main

import (
	"gin_aws/config"
	"gin_aws/routes"
	"log"
)

func main() {
	// Cargar configuración
	config.LoadConfig()

	// Inicializar rutas
	router := routes.SetupRouter()

	// Iniciar servidor
	port := config.AppConfig.Server.Port
	log.Printf("Servidor ejecutándose en el puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
