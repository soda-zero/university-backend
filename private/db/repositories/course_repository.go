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

type CourseRepository struct {
	pool *pgxpool.Pool
}

func (r *CourseRepository) Close() {
	r.pool.Close()
}

func NewCourseRepository() (*CourseRepository, error) {
	pool, err := db.ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to te database: %v", err)
	}

	return &CourseRepository{
		pool: pool,
	}, nil
}

func (r *CourseRepository) GetCourses() ([]models.Course, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT id, name, year, semester, optative, course_code, career_id FROM course")
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("Failed to fetch course: %w", pgErr)
		}
		return nil, fmt.Errorf("Failed to fetch course: %w", err)
	}
	defer rows.Close()

	courses := make([]models.Course, 0)
	for rows.Next() {
		course := models.Course{}
		if err := rows.Scan(&course.ID, &course.Name, &course.Year, &course.Semester, &course.Optative, &course.CourseCode, &course.CareerID); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Failed to scan course data: %w", pgErr)
			}
			return nil, fmt.Errorf("Failed to scan course data: %w", err)
		}
		courses = append(courses, course)

		if err := rows.Err(); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("Error occurred while fetching course: %w", pgErr)
			}
			return nil, fmt.Errorf("Error occurred while fetching course: %w", err)
		}
	}
	return courses, nil
}

func (r *CourseRepository) GetCourseByID(courseID string) (models.Course, error) {
	var course models.Course
	err := r.pool.QueryRow(context.Background(), "SELECT id, name, year, semester, optative, course_code, career_id FROM course WHERE id = $1", courseID).Scan(&course.ID, &course.Name, &course.Year, &course.Semester, &course.Optative, &course.CourseCode, &course.CareerID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return models.Course{}, fmt.Errorf("Error ocurred while fetchig course: %w", pgErr)
		}
		return models.Course{}, fmt.Errorf("Error ocurred while fetchig course: %w", err)
	}
	return course, nil
}

func (r *CourseRepository) CreateCourse(courseData models.Course) error {
	query := "INSERT INTO course(name, year, semester, optative, course_code, career_id) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := r.pool.Exec(context.Background(), query, courseData.Name, courseData.Year, courseData.Semester, courseData.Optative, courseData.CourseCode, courseData.CareerID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error occurred while inserting course: %w", pgErr)
		}
		return fmt.Errorf("Error occurred while inserting course: %w", err)
	}
	return nil
}

func (r *CourseRepository) DeleteCourse(courseID string) error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM course WHERE id = $1", courseID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error ocurred while deleting course: %w", pgErr)
		}
		return fmt.Errorf("Error ocurred while deleting course: %w", err)
	}
	return nil
}

func (r *CourseRepository) UpdateCourse(courseID string, courseData models.Course) error {
	query := "UPDATE course SET name = COALESCE($1, name), year = COALESCE($2, year), semester = COALESCE($3, semester), optative = COALESCE($4, optative), course_code = COALESCE($5, course_code), career_id = COALESCE($6, career_id) WHERE id = $7"
	_, err := r.pool.Exec(context.Background(), query, courseData.Name, courseData.Year, courseData.Semester, courseData.Optative, courseData.CourseCode, courseData.CareerID, courseID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return fmt.Errorf("Error ocurred while updating course: %w", pgErr)
		}
		return fmt.Errorf("Error ocurred while updating course: %w", err)
	}
	return nil
}
