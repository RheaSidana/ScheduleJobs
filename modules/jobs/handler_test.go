package jobs

import (
	"ScheduleJobs/model"
	"ScheduleJobs/modules/jobs/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateJobHandlerWhenEmptyJob(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/job", handler.CreateJobHandler)
	newJob := model.Job{}
	b,_:=json.Marshal(newJob)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/job",body)
	respR := httptest.NewRecorder()
	expectedJob := newJob
	repo.On("Create", newJob).Return(model.Job{},errors.New("Empty Job JSON."))

	actualJob,_ :=  repo.Create(newJob)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusBadRequest)
	assert.Equal(t,expectedJob, actualJob)
}

func TestCreateJobHandlerWhenUnableToCreateJob(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/job", handler.CreateJobHandler)
	newJob := model.Job{
		Name: "test",
		Duration: "3h20m10s",
	}
	b,_:=json.Marshal(newJob)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/job",body)
	respR := httptest.NewRecorder()
	expectedJob := newJob
	expectedJob.ID = 1
	repo.On("Create", newJob).Return(model.Job{},errors.New("Error while creating job"))

	actualJob,_ :=  repo.Create(newJob)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusInternalServerError)
	assert.NotEqual(t,expectedJob, actualJob)
}

func TestCreateJobHandler(t *testing.T){
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
    r.POST("/job", handler.CreateJobHandler)
	newJob := model.Job{
		Name: "test",
		Duration: "3h20m10s",
	}
	b,_:=json.Marshal(newJob)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/job",body)
	respR := httptest.NewRecorder()
	expectedJob := newJob
	expectedJob.ID = 1
	repo.On("Create", newJob).Return(expectedJob,nil)

	actualJob,_ :=  repo.Create(newJob)
	r.ServeHTTP(respR, req)

	assert.Equal(t,respR.Code, http.StatusOK)
	assert.Equal(t,expectedJob, actualJob)
}
