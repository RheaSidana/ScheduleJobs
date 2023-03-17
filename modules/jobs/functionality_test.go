package jobs

import (
	"ScheduleJobs/model"
	"ScheduleJobs/modules/jobs/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJob(t *testing.T){
	repo := new(mocks.Repository)
	myJob := model.Job{
		Name: "test",
		Duration: "5h2m4s",
	}
	expectedJob := myJob
	expectedJob.ID = 1
	repo.On("Find", 1).Return(expectedJob,nil)

	job, _ := GetJob(1, repo)

	assert.Equal(t,expectedJob, job)
}

func TestGetJobNotFound(t *testing.T){
	repo := new(mocks.Repository)
	expectedJob := model.Job{}
	expectedError := errors.New("Record not found")
	repo.On("Find", 1).Return(model.Job{},expectedError)

	job, err := GetJob(1, repo)

	assert.Equal(t,expectedJob, job)
	assert.Error(t,expectedError, err)
}
