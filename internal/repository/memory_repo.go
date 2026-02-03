package repository

import (
	"errors"
	"sync"

	"github.com/alexrobertCQ/Sensor-monitor-api/internal/domain"
)

// MemorySensorRepo usa un map para guardar sensores en RAM.
// Usamos sync.Mutex para que sea seguro usarlo con Goroutines (concurrencia).
type MemorySensorRepo struct {
	sync.Mutex
	sensors map[string]*domain.Sensor
}

func NewMemorySensorRepo() *MemorySensorRepo {
	return &MemorySensorRepo{
		sensors: make(map[string]*domain.Sensor),
	}
}

func (r *MemorySensorRepo) Save(s *domain.Sensor) error {
	r.Lock()
	defer r.Unlock()
	r.sensors[s.ID] = s
	return nil
}

func (r *MemorySensorRepo) GetByID(id string) (*domain.Sensor, error) {
	r.Lock()
	defer r.Unlock()
	sensor, ok := r.sensors[id]
	if !ok {
		return nil, errors.New("sensor no encontrado")
	}
	return sensor, nil
}

func (r *MemorySensorRepo) GetAll() ([]*domain.Sensor, error) {
	r.Lock()
	defer r.Unlock()
	list := make([]*domain.Sensor, 0, len(r.sensors))
	for _, s := range r.sensors {
		list = append(list, s)
	}
	return list, nil
}