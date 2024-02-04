package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"time"

	"plramos.win/ptscheduler/internal/datamodel"
	"plramos.win/ptscheduler/internal/sqlc"
)

func AddDemoData(ctx context.Context, db *sql.DB) error {
	log.SetFlags(log.Llongfile)
	testtable := []struct {
		datamodel.Trainee
		username     string
		password     string
		sessionDates []time.Time
		availability []datamodel.AvailabilityPeriod
	}{
		{
			datamodel.Trainee{"Pedro Ramos", 3, 0},
			"pedramos", "pedramos",
			[]time.Time{time.Now(), time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
			[]datamodel.AvailabilityPeriod{
				datamodel.AvailabilityPeriod{Start: time.Now(), End: time.Now().Add(time.Hour)},
			},
		},
		{
			datamodel.Trainee{"Maria Gato", 2, 10},
			"marigato", "marigato",
			[]time.Time{time.Now()},
			[]datamodel.AvailabilityPeriod{
				datamodel.AvailabilityPeriod{Start: time.Now(), End: time.Now().Add(time.Hour)},
			},
		},
	}

	queries := sqlc.New(db)

	for _, test := range testtable {
		traineeid, err := queries.NewTrainee(ctx, sqlc.NewTraineeParams{
			Name:    sql.NullString{test.Trainee.Name, true},
			Perweek: sql.NullInt64{test.Trainee.PerWeek, true},
			Late:    sql.NullInt64{test.Trainee.Late, true},
		})
		if err != nil {
			return fmt.Errorf("Failed to insert trainee '%v': %v", test.Trainee.Name, err)
		}
		userhash := sha256.Sum256([]byte(test.username))
		passhash := sha256.Sum256([]byte(test.password))
		err = queries.NewUsername(ctx, sqlc.NewUsernameParams{
			Traineeid: sql.NullInt64{traineeid, true},
			Username:  userhash[:],
			Password:  passhash[:],
		})
		if err != nil {
			return fmt.Errorf("Failed to insert username %v for userid %d: %v", test.username, traineeid, err)
		}
		for _, session := range test.sessionDates {
			date, err := session.MarshalText()
			if err != nil {
				return fmt.Errorf("failed to convert date '%s' to string: %v", session.String(), err)
			}
			err = queries.NewSession(ctx, sqlc.NewSessionParams{
				Traineeid: sql.NullInt64{traineeid, true},
				Date:      sql.NullString{string(date), true},
			})
			if err != nil {
				return fmt.Errorf("failed to insert sessions %v for userid %d: %v", session, traineeid, err)
			}
		}
		for _, avail := range test.availability {
			start, err := avail.Start.MarshalText()
			if err != nil {
				return fmt.Errorf("failed to convert date '%s' to string: %v", avail.Start.String(), err)
			}
			end, err := avail.End.MarshalText()
			if err != nil {
				return fmt.Errorf("failed to convert date '%s' to string: %v", avail.End.String(), err)
			}
			err = queries.AddAvailability(ctx, sqlc.AddAvailabilityParams{
				Traineeid: sql.NullInt64{traineeid, true},
				Startdate: sql.NullString{string(start), true},
				Enddate:   sql.NullString{string(end), true},
			})
			if err != nil {
				return fmt.Errorf("failed to insert sessions %#v for userid %d: %v", avail, traineeid, err)
			}
		}
	}
	return nil
}
