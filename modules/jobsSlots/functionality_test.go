package jobsSlots

import (
	"ScheduleJobs/model"
	"ScheduleJobs/modules/jobsSlots/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserJobSlot(t *testing.T) {
	repo := new(mocks.Repository)
	myUserJobSlots := []model.User_job_slot{
		{
			User_Id: 1,
			Job_Id:  1,
			Slot_Id: 1,
		},
	}
	expectedUserJobSlots := myUserJobSlots
	expectedUserJobSlots[0].ID = 1
	repo.On("FindAll", myUserJobSlots[0].User_Id).Return(expectedUserJobSlots, nil)

	userJobSlots, _ := GetUserJobSlot(myUserJobSlots[0].User_Id, repo)

	assert.Equal(t, expectedUserJobSlots, userJobSlots)

}

func TestGetUserJobSlotNotFound(t *testing.T) {
	repo := new(mocks.Repository)
	expectedUserJobSlots := []model.User_job_slot{}
	expectedError := errors.New("Record not found")
	repo.On("FindAll", 1).Return([]model.User_job_slot{}, expectedError)

	userJobSlots, err := GetUserJobSlot(1, repo)

	assert.Equal(t, expectedUserJobSlots, userJobSlots)
	assert.Error(t, expectedError, err)

}

func TestAddJobSlot(t *testing.T) {
	repo := new(mocks.Repository)
	myUserJobSlot := model.User_job_slot{
		User_Id: 1,
		Job_Id:  1,
		Slot_Id: 1,
	}
	expectedUserJobSlot := myUserJobSlot
	expectedUserJobSlot.ID = 1
	repo.On("Create", myUserJobSlot).Return(expectedUserJobSlot, nil)

	userJobSlot, _ := AddJobSlot(myUserJobSlot, repo)

	assert.Equal(t, expectedUserJobSlot, userJobSlot)
}

func TestAddJobSlotWhenEmptyJobSlot(t *testing.T) {
	repo := new(mocks.Repository)
	expectedUserJobSlots := model.User_job_slot{}
	expectedError := errors.New("Record not found")
	repo.On("Create", model.User_job_slot{}).Return(model.User_job_slot{}, expectedError)

	userJobSlots, err := AddJobSlot(model.User_job_slot{}, repo)

	assert.Equal(t, expectedUserJobSlots, userJobSlots)
	assert.Error(t, expectedError, err)
}
