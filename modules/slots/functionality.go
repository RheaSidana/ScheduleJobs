package slots

import (
	"ScheduleJobs/model"
	"time"
)

func createSlot(start time.Time, duration string) model.Slot {
	dur, _ := time.ParseDuration(duration)
	endTime := start.Add(time.Second * time.Duration(dur.Seconds()))

	return model.Slot{
		Start: start,
		End:   endTime,
	}
}

func createSlots(startTimings []time.Time, duration string) []model.Slot {
	var Slots []model.Slot
	for _, t := range startTimings {
		slot := createSlot(t, duration)

		Slots = append(Slots, slot)
	}

	return Slots
}

func isSlotOverlapping(createdSlots []model.Slot, presentSlots []model.Slot) bool {
	for _, createdSlot := range createdSlots {
		for _, presentSlot := range presentSlots {
			if createdSlot.Start.Equal(presentSlot.Start) || createdSlot.End.Equal(presentSlot.End) {
				return true
			} else if createdSlot.Start.Sub(presentSlot.End) >= 0 {
				continue
			} else if presentSlot.Start.Sub(createdSlot.End) >= 0 {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func createRepeated(x time.Time) []time.Time {
	day := int(x.Weekday())
	var daysTime []time.Time
	dur, _ := time.ParseDuration("24h")

	if day <= 5 {
		for day <= 5 {
			daysTime = append(daysTime, x)
			x = x.Add(time.Second * time.Duration(dur.Seconds()))
			day++
		}
	} else {
		for day <= 7 {
			daysTime = append(daysTime, x)
			x = x.Add(time.Second * time.Duration(dur.Seconds()))
			day++
		}
	}

	return daysTime
}

func currentWeek(slots []model.Slot) []model.Slot {
	today := time.Now()
	oneWeek := today.AddDate(0, 0, 7)
	var currentSlots []model.Slot
	for _, slot := range slots {
		if slot.Start.Before(today) && slot.End.Before(today) {
			continue
		}
		if slot.Start.After(oneWeek) {
			continue
		}
		currentSlots = append(currentSlots, slot)
	}
	return currentSlots
}
