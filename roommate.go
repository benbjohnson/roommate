package roommate

import (
	"time"
)

// User represents a user on the system.
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Room represents a room that can be scheduled.
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Reservation represents a time block scheduled for a room.
type Reservation struct {
	ID       int           `json:"id"`
	RoomID   int           `json:"roomID"`
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
}
