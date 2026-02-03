package domain

import "time"

// Sensor representa la entidad principal de nuestro sistema industrial.
type Sensor struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`  // Ej: "Presión Caldera Principal"
	Type      string    `json:"type"`  // Ej: "Pressure", "Temperature", "pH"
	Value     float64   `json:"value"` // El valor actual medido
	Unit      string    `json:"unit"`  // Ej: "PSI", "°C", "pH"
	UpdatedAt time.Time `json:"updated_at"`
}

// SensorRepository es la INTERFAZ (el contrato).
// Define qué acciones podemos hacer con los sensores sin decir CÓMO se hacen.
type SensorRepository interface {
	Save(sensor *Sensor) error
	GetByID(id string) (*Sensor, error)
	GetAll() ([]*Sensor, error)
}
