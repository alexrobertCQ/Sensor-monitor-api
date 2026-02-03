package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/handler"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/repository"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/usecase"
)

func main() {
	// 1. Cadena de conexión (los datos que pusimos en el docker-compose.yml)
	// Formato: postgres://usuario:password@host:puerto/nombre_bd?sslmode=disable
	//connStr := "postgres://user_mecatronic:password123@localhost:5432/industrial_monitor?sslmode=disable"
	connStr := os.Getenv("DB_URL")
    if connStr == "" {
        connStr = "postgres://user_mecatronic:password123@localhost:5432/industrial_monitor?sslmode=disable"
    }
	// 2. Inicializamos el repositorio REAL (Postgres) en lugar del de memoria
	repo, err := repository.NewPostgresSensorRepo(connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// 3. El resto del sistema sigue IGUAL (Inyección de dependencias)
	service := usecase.NewSensorService(repo)
	sensorHandler := handler.NewSensorHandler(service)

	r := gin.Default()
	// Rutas
	r.POST("/sensors/measurement", sensorHandler.RegisterMeasurement)
	r.GET("/sensors", sensorHandler.GetAllSensors) // <-- Nueva ruta
	log.Println("🚀 Servidor industrial conectado a PostgreSQL en el puerto 8080...")
	r.Run(":8080")
}