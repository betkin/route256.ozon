package route_sql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/act-device-api/internal/database"
	"github.com/ozonmp/act-device-api/tests/internal/models"
)

type EventStorage interface {
	ByDeviceId(ctx context.Context, deviceID uint64) (*models.DeviceEvent, error)
}

type Storage struct {
	DB *sqlx.DB
}

func NewPostgres(dsn, driver string) (*Storage, error) {
	db, err := database.NewPostgres(dsn, driver)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (r Storage) ByDeviceId(ctx context.Context, deviceID uint64) (*models.DeviceEvent, error) {
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
