package repository

import (
	//"context"
	"database/sql"
	"github.com/alexrobertCQ/Sensor-monitor-api/internal/domain"
	_ "github.com/jackc/pgx/v5/stdlib" // Driver para Postgres
)

type PostgresSensorRepo struct {
	db *sql.DB
}

func NewPostgresSensorRepo(connString string) (*PostgresSensorRepo, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	return &PostgresSensorRepo{db: db}, nil
}

func (r *PostgresSensorRepo) Save(s *domain.Sensor) error {
	query := `INSERT INTO sensors (id, name, type, value, unit, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  ON CONFLICT (id) DO UPDATE 
			  SET value = $4, updated_at = $6`
	
	_, err := r.db.Exec(query, s.ID, s.Name, s.Type, s.Value, s.Unit, s.UpdatedAt)
	return err
}

// Nota: Por brevedad, implementaremos GetByID y GetAll luego
func (r *PostgresSensorRepo) GetByID(id string) (*domain.Sensor, error) { return nil, nil }
func (r *PostgresSensorRepo) GetAll() ([]*domain.Sensor, error) {
	query := `SELECT id, name, type, value, unit, updated_at FROM sensors`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*domain.Sensor
	for rows.Next() {
		s := &domain.Sensor{}
		err := rows.Scan(&s.ID, &s.Name, &s.Type, &s.Value, &s.Unit, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, s)
	}
	return sensors, nil
}