package slots

import (
	"ScheduleJobs/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSlot(t *testing.T) {
	start, _ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end, _ := time.Parse(time.RFC3339Nano, "2023-03-18T10:05:00.000Z")
	duration := "2h5m0s"
	expectedSlot := model.Slot{
		Start: start,
		End:   end,
	}

	actualSlot := createSlot(start, duration)

	assert.Equal(t, expectedSlot, actualSlot)
}

func TestCreateSlots(t *testing.T) {
	start, _ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end, _ := time.Parse(time.RFC3339Nano, "2023-03-18T10:05:00.000Z")
	end2, _ := time.Parse(time.RFC3339Nano, "2023-03-18T12:10:00.000Z")
	startTimings := []time.Time{start, end}
	duration := "2h5m0s"
	expectedSlots := []model.Slot{
		{
			Start: start,
			End:   end,
		},
		{
			Start: end,
			End:   end2,
		},
	}

	actualSlots := createSlots(startTimings, duration)

	assert.Equal(t, expectedSlots, actualSlots)
}

func TestIsSlotOverlapping(t *testing.T) {
	start, _ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	end, _ := time.Parse(time.RFC3339Nano, "2023-03-18T10:05:00.000Z")
	end2, _ := time.Parse(time.RFC3339Nano, "2023-03-18T12:10:00.000Z")
	startTimings := []time.Time{start, end}
	duration := "2h5m0s"
	expectedSlots := []model.Slot{
		{
			Start: start,
			End:   end,
		},
		{
			Start: end,
			End:   end2,
		},
	}

	actualSlots := createSlots(startTimings, duration)

	assert.Equal(t, expectedSlots, actualSlots)
}

func TestCreateReapeated(t *testing.T) {
	start, _ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	ele1, _ := time.Parse(time.RFC3339Nano, "2023-03-18T08:00:00.000Z")
	ele2, _ := time.Parse(time.RFC3339Nano, "2023-03-19T08:00:00.000Z")

	expectedSlots := []time.Time{ele1, ele2};

	actualSlots := createRepeated(start);
	assert.Equal(t, expectedSlots, actualSlots);
}

func TestCurrentWeek(t *testing.T) {
	start, _ := time.Parse(time.RFC3339Nano, "2023-03-31T18:00:00.000Z")
	end, _ := time.Parse(time.RFC3339Nano, "2023-03-31T19:05:00.000Z")
	paramSlots := []model.Slot{
		{
			Start: start,
			End: end,
		},
		{
			Start: start.AddDate(0,0,1),
			End: end.AddDate(0,0,1),
		},
		{
			Start: start.AddDate(0,0,9),
			End: end.AddDate(0,0,9),
		},
	}
	expectedSlots := []model.Slot{
		{
			Start: start,
			End: end,
		},
		{
			Start: start.AddDate(0,0,1),
			End: end.AddDate(0,0,1),
		},
	}

	actualSlots := currentWeek(paramSlots);

	assert.Equal(t, expectedSlots, actualSlots);
}

