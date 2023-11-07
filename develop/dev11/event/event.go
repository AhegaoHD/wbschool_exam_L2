package event

import (
	"encoding/json"
	"io"
	"time"
)

type Event struct {
	EventID     int       `json:"event_id" validate:"required"`
	UserID      int       `json:"user_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date" time_format:"2006-01-02" time_utc:"true"`
}

var Events = make(map[int]Event)

func (e *Event) Decode(r io.Reader) error {

	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}
	return nil
}
