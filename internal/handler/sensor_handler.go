package handler

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/domain"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/usecase"
)

type SensorHandler struct {
	service *usecase.SensorService
}

func NewSensorHandler(s *usecase.SensorService) *SensorHandler {
	return &SensorHandler{service: s}
}

// RegisterMeasurement es el "puerto" que recibe el JSON
func (h *SensorHandler) RegisterMeasurement(c *gin.Context) {
	var sensor domain.Sensor

	// 1. Validar que el JSON recibido sea correcto
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// 2. Pasar los datos al "cerebro" (Usecase)
	if err := h.service.RegisterMeasurement(&sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Responder que todo salió bien
	c.JSON(http.StatusOK, gin.H{"message": "Lectura recibida correctamente"})
}
// GetAllSensors devuelve la lista de todos los sensores registrados
func (h *SensorHandler) GetAllSensors(c *gin.Context) {
    sensors, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // --- NUEVO: Log visual para el programador ---
    fmt.Println("\n=== REPORTE DE ESTADO DE CALDERA ===")
    fmt.Printf("%-15s | %-10s | %-8s | %-10s\n", "ID", "TIPO", "VALOR", "ESTADO")
    fmt.Println("---------------------------------------------------------")
    
    for _, s := range sensors {
        estado := "OK"
        if s.Type == "pH" && s.Value > 12.0 {
            estado = "CRÍTICO ⚠️"
        }
        fmt.Printf("%-15s | %-10s | %-8.2f%-2s | %-10s\n", s.ID, s.Type, s.Value, s.Unit, estado)
    }
    fmt.Println("=========================================================\n")

    c.JSON(http.StatusOK, sensors)
}