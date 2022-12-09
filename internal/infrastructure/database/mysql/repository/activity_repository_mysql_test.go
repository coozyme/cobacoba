package repository

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type activityRepositoryMysqlTest struct {
	suite.Suite
	mock          sqlmock.Sqlmock
	activityMYSQL repository.ActivityRepository
}

func (t *activityRepositoryMysqlTest) TestGetActivity() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	query := `SELECT id, email, title FROM activities`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"email",
			"title",
		}).AddRow(
			activity.ID,
			activity.Email,
			activity.Title,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		actualActivity, err := t.activityMYSQL.GetActivity()

		t.NotNil(actualActivity)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualActivity, err := t.activityMYSQL.GetActivity()

		t.Nil(actualActivity)
		t.Error(err)
	})
}

func (t *activityRepositoryMysqlTest) TestGetActivityByID() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	query := `SELECT id, email, title FROM activities WHERE id = ?`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"email",
			"title",
		}).AddRow(
			activity.ID,
			activity.Email,
			activity.Title,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(activity.ID).WillReturnRows(rows)

		actualActivity, err := t.activityMYSQL.GetActivityByID(activity.ID)

		log.Println(actualActivity)
		t.NotNil(actualActivity)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(activity.ID).WillReturnError(errors.New(""))

		actualActivity, err := t.activityMYSQL.GetActivityByID(activity.ID)

		log.Println(actualActivity)
		t.Nil(actualActivity)
		t.Error(err)
	})

}

func (t *activityRepositoryMysqlTest) TestCreate() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")

	t.Run("success", func() {
		t.mock.ExpectExec("INSERT INTO activities").WithArgs(
			activity.Email,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		t.NoError(t.activityMYSQL.CreateActivity(activity.Email, activity.Title))
	})

	t.Run("failed created", func() {
		t.mock.ExpectExec("INSERT INTO activities").WithArgs(
			activity.Email,
		).WillReturnError(errors.New(""))
		t.Error(t.activityMYSQL.CreateActivity(activity.Email, activity.Title))
	})

}

func TestActivityRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &activityRepositoryMysqlTest{
		mock:          mock,
		activityMYSQL: NewMysqlActivityRepository(db),
	})
}
