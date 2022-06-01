package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

func (a Device) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Device) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// EventType describes type for event constants
type EventType uint8

const (
	Created EventType = iota + 1
	Updated
	Removed
)

// EventStatus describes type for event constants
type EventStatus uint8

const (
	Deferred EventStatus = iota + 1
	Processed
)

// DeviceEvent describes SQL structure
type DeviceEvent struct {
	ID        uint64      `db:"id"`
	DeviceId  uint64      `db:"device_id"`
	Type      EventType   `db:"type"`
	Status    EventStatus `db:"status"`
	Device    *Device     `db:"payload"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}

// DevicesCount describes SQL structure for count request
type DevicesCount struct {
	Count uint64 `db:"count"`
}

// TimeDB describes SQL structure for time request
type TimeDB struct {
	Time time.Time `db:"time_db"`
}
