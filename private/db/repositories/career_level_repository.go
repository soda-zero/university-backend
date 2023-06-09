package repositories

import (
	"context"
	"errors"
	"fmt"
	"zeroCalSoda/university-backend/private/db"
	"zeroCalSoda/university-backend/private/db/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CareerLevelRepository struct {
	pool *pgxpool.Pool
}

func (r *CareerLevelRepository) Close() {
	r.pool.Close()
}

func NewCareerLevelRepository() (*CareerLevelRepository, error) {
	pool, err := db.ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database: %v", err)
	}

	return &CareerLevelRepository{
		pool: pool,
	}, nil
}

func (r *CareerLevelRepository) GetCareerLevels() ([]models.CareerLevel, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT id, name FROM career_level")
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("Failed to fetch career_level: %w", pgErr)
		}
		return nil, fmt.Errorf("Failed to fetch career_level: %w", err)
	}
	defer rows.Close()

	careerLevels := make([]models.CareerLevel, 0)
	for rows.Next() {
		career := models.CareerLevel{}
		if err := rows.Scan(&career.ID, &career.Name); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Failed to scan career_level data: %w", pgErr)
			}
			return nil, fmt.Errorf("Failed to scan career_level data: %w", err)
		}
		careerLevels = append(careerLevels, career)
		if err := rows.Err(); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Error occurred while fetching career_level: %w", pgErr)
			}
			return nil, fmt.Errorf("Error occurred while fetching career_level: %w", err)
		}
	}
	return careerLevels, nil
}

func (r *CareerLevelRepository) GetCareerLevelByID(careerLevelID string) (models.CareerLevel, error) {
	var careerLevel models.CareerLevel
	err := r.pool.QueryRow(context.Background(), "SELECT id, name FROM career_level WHERE id = $1", careerLevelID).Scan(&careerLevel.ID, &careerLevel.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return models.CareerLevel{}, fmt.Errorf("Error occurred while fetching career_level: %w", pgErr)
		}
		return models.CareerLevel{}, fmt.Errorf("Error occurred while fetching career_level: %w", err)
	}
	return careerLevel, nil
}

func (r *CareerLevelRepository) CreateCareerLevel(name string) error {
	_, err := r.pool.Exec(context.Background(), "INSERT INTO career_level(name) VALUES($1)", name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while inserting career_level: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while inserting career_level: %w", err)
	}
	return nil
}

func (r *CareerLevelRepository) DeleteCareerLevel(careerID string) error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM career_level WHERE id = $1", careerID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while deleting career_level: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while deleting career_level: %w", err)
	}
	return nil
}

func (r *CareerLevelRepository) UpdateCareerLevel(careerLevelID string, newName string) error {
	query := "UPDATE career_level SET name = $1 WHERE id = $2"
	_, err := r.pool.Exec(context.Background(), query, newName, careerLevelID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while updating career_level: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while updating career_level: %w", err)
	}
	return nil
}
