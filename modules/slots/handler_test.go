package slots

import (
	"ScheduleJobs/model"
	"ScheduleJobs/modules/slots/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSlotHandlerWhenEmptySlot(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/slot", handler.CreateSlotHandler)
	newSlotBody := model.SlotBody{}
	b,_:=json.Marshal(newSlotBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/slot",body)
	respR := httptest.NewRecorder()
	expectedSlot := model.Slot{}
	repo.On("Create", newSlotBody).Return(model.Slot{},errors.New("Empty Slot JSON."))

	actualSlot,_ :=  repo.Create(newSlotBody)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusBadRequest)
	assert.Equal(t,expectedSlot, actualSlot)
}

func TestCreateSlotHandlerWhenUnableToCreateSlot(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/slot", handler.CreateSlotHandler)
	newSlotBody := model.SlotBody{
		UserId: 4,
		JobId: 3,
		Start_Time: "2023-03-18T08:00:00.000Z",
	}
	b,_:=json.Marshal(newSlotBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/slot",body)
	respR := httptest.NewRecorder()
	start,_ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end,_ := time.Parse(time.RFC3339Nano, "2023-03-18T09:00:00.000Z")
	expectedSlot := model.Slot{
		Start: start,
		End: end,
	}
	expectedSlot.ID = 1
	repo.On("Create", newSlotBody).Return(model.Slot{},errors.New("Error while creating slot"))

	actualSlot,_ :=  repo.Create(newSlotBody)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusInternalServerError)
	assert.NotEqual(t,expectedSlot, actualSlot)
}

func TestCreateSlotHandler(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/slot", handler.CreateSlotHandler)
	newSlotBody := model.SlotBody{
		UserId: 4,
		JobId: 3,
		Start_Time: "2023-03-18T08:00:00.000Z",
	}
	b,_:=json.Marshal(newSlotBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/slot",body)
	respR := httptest.NewRecorder()
	start,_ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end,_ := time.Parse(time.RFC3339Nano, "2023-03-18T09:00:00.000Z")
	expectedSlot := model.Slot{
		Start: start,
		End: end,
	}
	expectedSlot.ID = 1
	repo.On("Create", newSlotBody).Return(expectedSlot,nil)

	actualSlot,_ :=  repo.Create(newSlotBody)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusOK)
	assert.Equal(t,expectedSlot, actualSlot)
}

func TestFindSlotHandlerWhenNotFound(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.GET("/slot", handler.FindSlotHandler)
	newGetSlotReq := GetSlotRequest{
		UserID: 1,
	}
	b,_:=json.Marshal(newGetSlotReq)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodGet, "/slot",body)
	respR := httptest.NewRecorder()
	start,_ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end,_ := time.Parse(time.RFC3339Nano, "2023-03-18T09:00:00.000Z")
	expectedSlot := model.Slot{
		Start: start,
		End: end,
	}
	expectedSlot.ID = 1
	expectedSlots := []model.Slot{expectedSlot}
	expectedError := errors.New("Slots not found")
	repo.On("FindAllCurrentSlots", newGetSlotReq.UserID).Return([]model.Slot{},expectedError)

	actualSlots,err :=  repo.FindAllCurrentSlots(newGetSlotReq.UserID)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusBadRequest)
	assert.Error(t,expectedError, err)
	assert.NotEqual(t,expectedSlots, actualSlots)
}

func TestFindSlotHandler(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.GET("/slot", handler.FindSlotHandler)
	newGetSlotReq := GetSlotRequest{
		UserID: 1,
	}
	b,_:=json.Marshal(newGetSlotReq)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodGet, "/slot",body)
	respR := httptest.NewRecorder()
	start,_ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end,_ := time.Parse(time.RFC3339Nano, "2023-03-18T09:00:00.000Z")
	expectedSlot := model.Slot{
		Start: start,
		End: end,
	}
	expectedSlot.ID = 1
	expectedSlots := []model.Slot{expectedSlot}
	repo.On("FindAllCurrentSlots", newGetSlotReq.UserID).Return(expectedSlots,nil)

	actualSlots,_ :=  repo.FindAllCurrentSlots(newGetSlotReq.UserID)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusOK)
	assert.Equal(t,expectedSlots, actualSlots)
}