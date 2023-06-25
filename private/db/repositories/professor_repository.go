package repositories

import (
	"context"
	"fmt"
	"zeroCalSoda/university-backend/private/db"
	"zeroCalSoda/university-backend/private/db/models"
	"zeroCalSoda/university-backend/private/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfessorRepository struct {
	pool *pgxpool.Pool
}

func (r *ProfessorRepository) Close() {
	r.pool.Close()
}

func NewProfessorRepository() (*ProfessorRepository, error) {
	pool, err := db.ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Unable to conenct to the database: %v", err)
	}
	return &ProfessorRepository{
		pool: pool,
	}, nil
}

func (r *ProfessorRepository) GetProfessors() ([]models.Professor, error) {
	query := "SELECT id, name, last_name, email, phone_number FROM professor"
	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		utils.HandlePgError(err)
	}
	defer rows.Close()

	professors := make([]models.Professor, 0)
	for rows.Next() {
		professor := models.Professor{}
		if err := rows.Scan(&professor.ID, &professor.Name, &professor.Email, &professor.PhoneNumber); err != nil {
			return []models.Professor{}, utils.HandlePgError(err)
		}
		professors = append(professors, professor)

		if err := rows.Err(); err != nil {
			return []models.Professor{}, utils.HandlePgError(err)
		}
	}
	return professors, nil
}

func (r *ProfessorRepository) GetProfessorByID(professorID string) (models.Professor, error) {
	var professor models.Professor
	query := "SELECT id, name, last_name, email, phone_number FROM professor WHERE id = $1"
	err := r.pool.QueryRow(context.Background(), query, professorID).Scan(&professor.ID, &professor.Name, &professor.LastName, &professor.Email, &professor.PhoneNumber)
	if err != nil {
		return models.Professor{}, utils.HandlePgError(err)
	}
	return professor, nil
}

func (r *ProfessorRepository) CreateProfessor(professorData models.Professor) error {
	query := "INSERT INTO professor(name, last_name, email, phone_number) VALUES ($1, $2, $3, $4)"
	_, err := r.pool.Exec(context.Background(), query, professorData)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}
func (r *ProfessorRepository) DeleteProfessor(professorID string) error {
	query := "DELETE FROM professor WHERE id = $1"
	_, err := r.pool.Exec(context.Background(), query, professorID)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}

func (r *ProfessorRepository) UpdateProfessor(professorID string, professorData models.Professor) error {
	query := "UPDATE professors SET name = COALESCE($1, name), last_name = COALESCE($2 last_name), email = COALESCE($3, email), phone_number = COALESCE($4, phone_number) WHERE id = $5"
	_, err := r.pool.Exec(context.Background(), query, professorData.Name, professorData.LastName, professorData.Email, professorData.PhoneNumber, professorID)
	if err != nil {
		return utils.HandlePgError(err)
	}
	return nil
}
