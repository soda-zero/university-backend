package db

import (
	"context"
	"errors"
	"fmt"
	"zeroCalSoda/university-backend/private/db/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	pool *pgxpool.Pool
}

func (r *DepartmentRepository) Close() {
	r.pool.Close()
}

func NewDepartmentRepository() (*DepartmentRepository, error) {
	pool, err := ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database: %v", err)
	}

	return &DepartmentRepository{
		pool: pool,
	}, nil
}

func (r *DepartmentRepository) GetAllDepartments() ([]models.Department, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT id, name FROM department")
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("Failed to fetch departments: %w", pgErr)
		}
		return nil, fmt.Errorf("Failed to fetch departments: %w", err)
	}
	defer rows.Close()

	departments := make([]models.Department, 0)
	for rows.Next() {
		dept := models.Department{}
		if err := rows.Scan(&dept.ID, &dept.Name); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Failed to scan department data: %w", pgErr)
			}
			return nil, fmt.Errorf("Failed to scan department data: %w", err)
		}
		departments = append(departments, dept)

		if err := rows.Err(); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Error occurred while fetching departments: %w", pgErr)
			}
			return nil, fmt.Errorf("Error occurred while fetching departments: %w", err)
		}
	}
	return departments, nil
}

func (r *DepartmentRepository) GetDepartmentByID(departmentID string) (models.Department, error) {
	var department models.Department
	err := r.pool.QueryRow(context.Background(), "SELECT id, name FROM department WHERE id = $1", departmentID).Scan(&department.ID, &department.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return models.Department{}, fmt.Errorf("Error occurred while fetching department: %w", pgErr)
		}
		return models.Department{}, fmt.Errorf("Error occurred while fetching department: %w", err)
	}

	return department, nil
}

func (r *DepartmentRepository) CreateDepartment(name string) error {
	_, err := r.pool.Exec(context.Background(), "INSERT INTO department(name) VALUES($1)", name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while inserting department: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while inserting department: %w", err)
	}
	return nil
}

func (r *DepartmentRepository) DeleteDepartment(departmentID string) error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM department WHERE id = $1", departmentID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while deleting department: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while deleting department: %w", err)
	}
	return nil
}

func (r *DepartmentRepository) UpdateDepartment(departmentID int, newName string) error {
	query := "UPDATE department SET name = $1 WHERE id = $2"
	_, err := r.pool.Exec(context.Background(), query, newName, departmentID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("error occurred while updating department: %w", pgErr)
		}
		return fmt.Errorf("error occurred while updating department: %w", err)
	}
	return nil
}
