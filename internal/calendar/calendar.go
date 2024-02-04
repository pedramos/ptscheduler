package calendar

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"net/http"
//	"time"
//
//	"plramos.win/ptscheduler/internal/datamodel"
//
//	"github.com/emersion/go-ical"
//	"github.com/emersion/go-webdav/caldav"
//)
//
//// go doc github.com/emersion/go-webdav/caldav.Backend
//type TraineeBackend struct {
//	user string
//
//	calendars []caldav.Calendar
//	objectMap map[string][]caldav.CalendarObject
//}
//
//func (b TraineeBackend) Calendar(ctx context.Context) (*caldav.Calendar, error) {
//	return &b.calendars[0], nil
//}
//
//func (b TraineeBackend) ListCalendars(ctx context.Context) ([]caldav.Calendar, error) {
//	return b.calendars, nil
//}
//
//func (b TraineeBackend) GetCalendar(ctx context.Context, path string) (*caldav.Calendar, error) {
//	for _, cal := range b.calendars {
//		if cal.Path == path {
//			return &cal, nil
//		}
//	}
//	return nil, fmt.Errorf("Calendar for path: %s not found", path)
//}
//
//func (b TraineeBackend) CalendarHomeSetPath(ctx context.Context) (string, error) {
//	return fmt.Sprintf("/%s/calendars/", b.user), nil
//}
//
//func (b TraineeBackend) CurrentUserPrincipal(ctx context.Context) (string, error) {
//	return fmt.Sprintf("/%s/", b.user), nil
//}
//
//func (b TraineeBackend) DeleteCalendarObject(ctx context.Context, path string) error {
//	return nil
//}
//
//func (b TraineeBackend) GetCalendarObject(ctx context.Context, path string, req *caldav.CalendarCompRequest) (*caldav.CalendarObject, error) {
//	for _, objs := range b.objectMap {
//		for _, obj := range objs {
//			if obj.Path == path {
//				return &obj, nil
//			}
//		}
//	}
//	return nil, fmt.Errorf("Couldn't find calendar object at: %s", path)
//}
//
//func (b TraineeBackend) PutCalendarObject(ctx context.Context, path string, calendar *ical.Calendar, opts *caldav.PutCalendarObjectOptions) (string, error) {
//	return "", nil
//}
//
//func (b TraineeBackend) ListCalendarObjects(ctx context.Context, path string, req *caldav.CalendarCompRequest) ([]caldav.CalendarObject, error) {
//	return b.objectMap[path], nil
//}
//
//func (b TraineeBackend) QueryCalendarObjects(ctx context.Context, query *caldav.CalendarQuery) ([]caldav.CalendarObject, error) {
//	return nil, nil
//}
//
//func ServeUserCalendar(username string, sessions ...datamodel.Session) {
//	sessionsCal := caldav.Calendar{
//		Path:                  fmt.Sprintf("/%s/calendars/sessions", username),
//		SupportedComponentSet: []string{ical.CompEvent},
//	}
//
//	calendars := []caldav.Calendar{
//		sessionsCal,
//	}
//	cal := ical.NewCalendar()
//	cal.Props.SetText(ical.PropVersion, "2.0")
//	cal.Props.SetText(ical.PropProductID, "-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN")
//
//	for _, session := range sessions {
//
//		eventSummary := "Treino"
//		event := ical.NewEvent()
//		event.Name = ical.CompEvent
//		event.Props.SetText(ical.PropUID, "46bbf47a-1861-41a3-ae06-8d8268c6d41e")
//		eventTime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", session.Date)
//		event.Props.SetDateTime(ical.PropDateTimeStamp, eventTime)
//		event.Props.SetText(ical.PropSummary, eventSummary)
//
//		cal.Children = []*ical.Component{
//			event.Component,
//		}
//
//	}
//	object := caldav.CalendarObject{
//		Path: fmt.Sprintf("/%s/calendars/sessions/test.ics", username),
//		Data: cal,
//	}
//
//	handler := caldav.Handler{Backend: TraineeBackend{
//		calendars: calendars,
//		objectMap: map[string][]caldav.CalendarObject{
//			sessionsCal.Path: []caldav.CalendarObject{object},
//		},
//	}}
//	s := &http.Server{
//		Addr:           ":8080",
//		Handler:        handler,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	// http.Handle("/", handler.ServeHTTP)
//	log.Fatal(s.ListenAndServe(":8080", nil))
//}
