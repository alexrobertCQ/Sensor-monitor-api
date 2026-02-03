package usecase

import (
	"errors"
	"fmt"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/domain"
)

// SensorService contiene la lógica de negocio.
type SensorService struct {
	repo domain.SensorRepository
}

// NewSensorService es el constructor.
func NewSensorService(r domain.SensorRepository) *SensorService {
	return &SensorService{
		repo: r,
	}
}

// RegisterMeasurement procesa una nueva lectura y aplica lógica de alertas.
func (s *SensorService) RegisterMeasurement(sensor *domain.Sensor) error {
	// 1. Regla de Negocio: Alerta de pH crítico (Basado en tu experiencia con calderas)
	if sensor.Type == "pH" && sensor.Value > 12.0 {
		fmt.Printf("⚠️ ALERTA CRÍTICA: pH de %.2f detectado. Riesgo de corrosión/incrustación.\n", sensor.Value)
		// Aquí podrías disparar un correo o un webhook en el futuro.
	}

	// 2. Regla de Negocio: Evitar valores imposibles (Sensores fallidos)
	if sensor.Value < 0 && sensor.Type != "Temperature" {
		return errors.New("lectura de sensor inválida: el valor no puede ser negativo")
	}

	// 3. Si todo está bien, lo mandamos a guardar al repositorio.
	return s.repo.Save(sensor)
}
func (s *SensorService) GetAll() ([]*domain.Sensor, error) {
	return s.repo.GetAll()
}