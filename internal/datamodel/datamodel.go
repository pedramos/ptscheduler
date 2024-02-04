package datamodel

import "time"

type Session struct {
	TraineeID int64
	Date      string
}

type Trainee struct {
	Name    string
	PerWeek int64
	Late    int64
}

type AvailabilityPeriod struct {
	Start time.Time
	End   time.Time
}
