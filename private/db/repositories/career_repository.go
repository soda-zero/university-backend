package repositories

import (
	"context"
	"fmt"
	"zeroCalSoda/university-backend/private/db"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CareerRepository struct {
	pool *pgxpool.Pool
}

func (r *CareerRepository) Close() {
	r.pool.Close()
}

func NewCareerRepository() (*CareerRepository, error) {
	pool, err := db.ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database: %v", err)
	}

	return &CareerRepository{
		pool: pool,
	}, nil
}

func (r *CareerRepository) GetCareers() ([]models.Career, error) {
	query := "SELECT id, name, duration_years, department_id, career_level_id"
	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		return nil, utils.HandlePgError(err)
	}
	defer rows.Close()

	careers := make([]models.Career, 0)
	for rows.Next() {
		career := models.Career{}
		if err := rows.Scan(&career.ID, &career.Name, &career.DurationYears, &career.DepartmentID, &career.DepartmentID, &career.CareerLevelID); err != nil {
			return nil, utils.HandlePgError(err)
		}
		careers = append(careers, career)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.HandlePgError(err)

	}

	return careers, nil
}
func (r *CareerRepository) GetCareerByID(careerID string) (models.Career, error) {
	var career models.Career
	query := "SELECT id, name, duration_years, department_id, career_level_id WHERE id = $1"
	err := r.pool.QueryRow(context.Background(), query, careerID).Scan(&career.ID, &career.Name, &career.DurationYears, &career.DepartmentID, &career.CareerLevelID)
	if err != nil {
		return models.Career{}, utils.HandlePgError(err)
	}

	return career, nil
}

func (r *CareerRepository) CreateCareer(careerData models.Career) error {
	query := "INSERT INTO career(name, duration_years, department_id, career_level_id) VALUES($1, $2, $3, $4)"
	_, err := r.pool.Exec(context.Background(), query, careerData.Name, careerData.DurationYears, careerData.DepartmentID, careerData.CareerLevelID)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}

func (r *CareerRepository) UpdateCareer(careerID string, careerData models.Career) error {
	query := "UPDATE career SET name = COALESCE($1, name), duration_years = COALESCE($2, duration_years), department_id = COALESCE($3, department_id), career_level_id = COALESCE($4, career_level_id) WHERE id = $5"
	_, err := r.pool.Exec(context.Background(), query, careerData.Name, careerData.DurationYears, careerData.DepartmentID, careerData.CareerLevelID, careerID)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}

func (r *CareerRepository) DeleteCareer(careerID string) error {
	query := "DELETE FROM career WHERE id = $1"
	_, err := r.pool.Exec(context.Background(), query, careerID)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}
