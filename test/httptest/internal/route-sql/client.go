package route_sql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"gitlab.ozon.dev/betkin/device-api/internal/database"
	"gitlab.ozon.dev/betkin/device-api/test/httptest/internal/models"
)

// EventStorage describes SQL methods for DB object
type EventStorage interface {
	ByDeviceID(ctx context.Context, deviceID uint64) (*models.DeviceEvent, error)
}

// Storage describes pointer to DB object
type Storage struct {
	DB *sqlx.DB
}

// NewPostgres makes new point for connection to PostgresSQL
func NewPostgres(dsn, driver string) (*Storage, error) {
	db, err := database.NewPostgres(dsn, driver)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

// ByDeviceID returns last row in Events table for input deviceID
func (r Storage) ByDeviceID(ctx context.Context, deviceID uint64) (*models.DeviceEvent, error) {
	var (
		event models.DeviceEvent
	)
	query := sq.Select("id", "device_id", "type", "status", "payload", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From("devices_events").
		Where(sq.Eq{"device_id": deviceID}).
		OrderBy("id DESC")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &event, s, args...)

	return &event, err
}

// GetCountDevices returns count of devices in device table
func (r Storage) GetCountDevices(ctx context.Context, existsOnly bool) (*models.DevicesCount, error) {
	var (
		data models.DevicesCount
	)
	query := sq.Select("COUNT(id) as count").
		PlaceholderFormat(sq.Dollar).
		From("devices")
	if existsOnly {
		query.Where(sq.Eq{"removed": false})
	}

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &data, s, args...)

	return &data, err
}

// GetDBTime returns current time on SQL server
func (r Storage) GetDBTime(ctx context.Context) (*models.TimeDB, error) {
	var (
		data models.TimeDB
	)
	query := sq.Select("NOW() as time_db")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &data, s, args...)

	return &data, err
}
