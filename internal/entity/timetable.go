package entity

import (
	"github.com/google/uuid"
	"time"
)

type Timetable struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	StartDate   time.Time  `json:"start_date"`
	StartTime   time.Time  `json:"start_time"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	EndTime     time.Time  `json:"end_time"`
	Description string     `json:"description"`
	Stock       int32      `json:"stock"`
	EventID     uuid.UUID  `json:"-"`
	Event       Event      `json:"event"`
	Auditable
}

func NewTimetable(eventId uuid.UUID, name string, startDate time.Time, endDate *time.Time, startTime, endTime time.Time, description string, stock int32) *Timetable {
	return &Timetable{
		ID:          uuid.New(),
		EventID:     eventId,
		Name:        name,
		StartDate:   startDate,
		EndDate:     endDate,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
		Stock:       stock,
	}
}

func EditTimetable(id, eventId uuid.UUID, name string, startDate time.Time, endDate *time.Time, startTime, endTime time.Time, description string) *Timetable {
	return &Timetable{
		ID:          id,
		EventID:     eventId,
		Name:        name,
		StartDate:   startDate,
		EndDate:     endDate,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
	}
}

func UpdateStock(id uuid.UUID, stock int32) *Timetable {
	return &Timetable{
		ID:    id,
		Stock: stock,
	}
}
