package model

import (
	"time"
)

// Device describes SQL structure
type Device struct {
	ID        uint64     `db:"id"         json:"id,omitempty"`
	Platform  string     `db:"platform"   json:"platform,omitempty"`
	UserID    uint64     `db:"user_id"    json:"user_id,omitempty"`
	EnteredAt *time.Time `db:"entered_at" json:"entered_at,omitempty"`
	Removed   bool       `db:"removed"    json:"removed,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// EventType describes type for event constants
type EventType uint8

// EventType constants
const (
	Created EventType = iota + 1
	Updated
	Removed
)

// EventStatus describes type for event constants
type EventStatus uint8

// EvenStatus constants
const (
	Deferred EventStatus = iota + 1
	Processed
)

// DeviceEvent describes SQL structure
type DeviceEvent struct {
	ID        uint64      `db:"id"`
	DeviceID  uint64      `db:"device_id"`
	Type      EventType   `db:"type"`
	Status    EventStatus `db:"status"`
	Device    *Device     `db:"payload"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}
