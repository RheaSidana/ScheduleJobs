package slots

import (
	"ScheduleJobs/model"
	"ScheduleJobs/modules/jobs"
	"ScheduleJobs/modules/jobsSlots"
	"ScheduleJobs/modules/users"

	"errors"
	"time"

	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	BeforeCreate(slotBody model.SlotBody, timeArr []time.Time) ([]model.Slot, error)
	Create(slot model.SlotBody) (model.Slot, error)
	CreateSlot(slots []model.Slot, userJobSlot model.User_job_slot) (bool, error)
	AfterCreate(slotBody model.User_job_slot) error
	FindSlot(slot model.Slot) (model.Slot, bool)
	FindAll(jobSlots []model.User_job_slot) ([]model.Slot, error)
	FindAllCurrentSlots(userID int) ([]model.Slot, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(slotBody model.SlotBody) (model.Slot, error) {
	// start_time > now
	timeFromAPI, _ := time.Parse(time.RFC3339Nano, slotBody.Start_Time)
	if time.Now().After(timeFromAPI) {
		return model.Slot{}, errors.New("Cannot add past time")
	}

	timeArr := createRepeated(timeFromAPI)

	slots, err := r.BeforeCreate(slotBody, timeArr)
	if err != nil {
		return model.Slot{}, err
	} else if slots == nil {
		return model.Slot{}, errors.New("slot overlapping")
	}

	userJobSlot := model.User_job_slot{
		User_Id: int(slotBody.UserId),
		Job_Id:  int(slotBody.JobId),
	}

	_, err = r.CreateSlot(slots, userJobSlot)
	if err != nil {
		return model.Slot{}, err
	}

	return slots[0], nil
}

func (r *repository) FindSlot(slot model.Slot) (model.Slot, bool) {
	start := slot.Start.UTC()
	end := slot.End.UTC()

	var slotInTable model.Slot
	res := r.client.Where("slots.start=? AND slots.end=?", start, end).First(&slotInTable)

	if res.RowsAffected >= 1 {
		slot = slotInTable
		return slot, true
	}
	return slot, false
}

func (r *repository) FindAllCurrentSlots(userID int) ([]model.Slot, error) {
	jobSlots, err := jobsSlots.GetUserJobSlot(userID, jobsSlots.NewRepository(r.client))
	if err != nil {
		return []model.Slot{}, err
	}

	slots, err := r.FindAll(jobSlots)
	if err != nil {
		return []model.Slot{}, err
	}

	return currentWeek(slots), nil
}

func (r *repository) FindAll(jobSlots []model.User_job_slot) ([]model.Slot, error) {
	var slots []model.Slot
	for _, jS := range jobSlots {
		var slot model.Slot
		res := r.client.Where("id=?", jS.Slot_Id).First(&slot)
		if res.Error != nil {
			return []model.Slot{}, res.Error
		}

		slots = append(slots, slot)
	}

	return slots, nil
}

func (r *repository) BeforeCreate(slotBody model.SlotBody, timeArr []time.Time) ([]model.Slot, error) {
	// find user in users table
	_, err := users.GetUser(slotBody.UserId,users.NewRepository(r.client))
	if err != nil {
		return nil, err
	}

	// find entries in jobslot
	jobSlots, _ := jobsSlots.GetUserJobSlot(slotBody.UserId, jobsSlots.NewRepository(r.client))

	// find job in jobs table
	job, e := jobs.GetJob(slotBody.JobId, jobs.NewRepository(r.client))
	if e != nil {
		return nil, e
	}

	createdSlots := createSlots(timeArr, job.Duration)

	isUserInJobSlot := len(jobSlots)
	if isUserInJobSlot == 0 {
		return createdSlots, nil
	}

	// find all slots
	presentSlots, err := r.FindAll(jobSlots)
	if err != nil {
		return nil, err
	}

	//compare for overlapping
	overlapping := isSlotOverlapping(createdSlots, presentSlots)
	if overlapping {
		return nil, errors.New("Overlapping slot")
	}

	return createdSlots, nil
}

func (r *repository) AfterCreate(jobSlot model.User_job_slot) error {
	_, err := jobsSlots.AddJobSlot(jobSlot, jobsSlots.NewRepository(r.client))
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateSlot(slots []model.Slot, userJobSlot model.User_job_slot) (bool, error) {
	for _, slot := range slots {
		//check if slot already exists
		slotInTable, isSlotInTable := r.FindSlot(slot)
		if isSlotInTable == false {
			result := r.client.Create(&slotInTable)
			if result.Error != nil {
				return false, result.Error
			}
		}

		userJobSlot.Slot_Id = int(slotInTable.ID)

		//update jobslots
		err := r.AfterCreate(userJobSlot)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
