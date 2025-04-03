package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/carlosgrillet/go-api/db"
)

type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Location  string    `json:"location" binding:"required"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"userId"`
}

func (e *Event) Save() error {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return errors.New("Error encoding the event to json")
	}

	db.Put("/events/", e.ID, string(jsonData))
	return nil
}

func (e *Event) Delete() error {
	err := db.Delete("/events/", e.ID)
	if err != nil {
		return errors.New("Failed to delete event")
	}
	return nil
}

func GetEventById(id string) (Event, error) {
	eventInDb, err := db.Get("/events/", id, false)
	if err != nil {
		return Event{}, errors.New("Event not found")
	}

	var event Event
	for _, value := range eventInDb {
		json.Unmarshal([]byte(value), &event)
	}
	return event, nil
}

func GetEventByName(name string) (Event, error) {
	eventInDb, err := db.Get("/events/", "", true)
	if err != nil {
		return Event{}, errors.New("Event not found")
	}

	var event Event
  var searchEvent Event
	for _, value := range eventInDb {
		json.Unmarshal([]byte(value), &event)
    if event.Name == name {
      json.Unmarshal([]byte(value), &searchEvent)
      break
    }
	}
  if searchEvent.Name == "" {
		return Event{}, errors.New("Event not found")
  }
	return searchEvent, nil
}

func GetAllEvents() ([]Event, error) {
	events, err := db.Get("/events/", "", true)
	if err != nil {
		return nil, errors.New("Failed to get events")
	}

	eventList := []Event{}

	for _, value := range events {
		var event Event
		json.Unmarshal([]byte(value), &event)
		eventList = append(eventList, event)
	}

	return eventList, nil
}
