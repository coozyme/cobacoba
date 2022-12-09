package repository

import (
	"database/sql"
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/infrastructure/database/mysql/model"
	"ddd-to-do-list/internal/repository"
	"errors"

	"github.com/sirupsen/logrus"
)

type activityRepositoryMySQL struct {
	db *sql.DB
}

func NewMysqlActivityRepository(Conn *sql.DB) repository.ActivityRepository {
	return &activityRepositoryMySQL{Conn}
}

func (m *activityRepositoryMySQL) fetch(query string, args ...interface{}) (aggregate.Activities, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()
	activityDTOs := []model.ActivityDTO{}
	for rows.Next() {
		t := model.ActivityDTO{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Title,
		)

		if err != nil {
			return nil, err
		}

		activityDTOs = append(activityDTOs, t)
	}

	activities := aggregate.Activities{}
	for _, activityDTO := range activityDTOs {
		activities = append(activities, aggregate.RebuildActivity(
			activityDTO.ID,
			activityDTO.Email,
			activityDTO.Title,
		))
	}

	return activities, nil
}

func (m *activityRepositoryMySQL) GetActivity() (res aggregate.Activities, err error) {
	query := `SELECT id, email, title FROM activities`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *activityRepositoryMySQL) GetActivityByID(id uint64) (res aggregate.Activities, err error) {
	query := `SELECT id, email, title FROM activities WHERE id = ? LIMIT 1`

	res, err = m.fetch(query, id)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *activityRepositoryMySQL) CreateActivity(email, title string) error {
	query := "INSERT INTO activities (email, title) VALUES(?, ?)"
	_, err := m.db.Exec(
		query,
		email,
		title,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *activityRepositoryMySQL) UpdateActivity(id uint64, email, title string) error {
	query := "UPDATE activities email = ?, title = ? WHERE id = ?"
	_, err := m.db.Exec(
		query,
		email,
		title,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityRepositoryMySQL) DeleteActivity(id uint64) error {
	query := "DELETE activities WHERE id = ?"
	_, err := s.db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
