package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
)

type activityUsecaseTest struct {
	suite.Suite
	activity *aggregate.Activity
	repo     *repository.ActivityMock
	usecase  ActivityUsecase
}

func (t *activityUsecaseTest) SetupSuite() {
	t.activity = aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	t.usecase = NewActivityUsecase(t.repo)
}

func (t *activityUsecaseTest) SetupTest() {
	t.repo.On("GetActivity").Return(t.activity, nil)
}

func (t *activityUsecaseTest) TestGetActivity() {
	t.Run("success", func() {
		activity, err := t.usecase.GetActivity()
		log.Println(activity)
		t.NotNil(activity)
		t.NoError(err)
	})
}

func TestActivityUsecase(t *testing.T) {
	suite.Run(t, new(activityUsecaseTest))
}
