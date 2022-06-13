package repo

import (
	"context"
	"github.com/opentracing/opentracing-go"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"gitlab.ozon.dev/betkin/device-api/internal/model"
)

// Repo is DAO for Template
type Repo interface {
	CreateDevice(ctx context.Context, device *model.Device) (uint64, error)
	DescribeDevice(ctx context.Context, deviceID uint64) (*model.Device, error)
	DescribeLastDevice(ctx context.Context) (*model.Device, error)
	ListDevices(ctx context.Context, page uint64, perPage uint64) ([]*model.Device, error)
	LogDevice(ctx context.Context, deviceID uint64) ([]*model.DeviceEvent, error)
	UpdateDevice(ctx context.Context, device *model.Device) (bool, error)
	UpdateLastDevice(ctx context.Context, device *model.Device) (bool, error)
	RemoveDevice(ctx context.Context, deviceID uint64) (bool, error)
	RemoveLastDevice(ctx context.Context, deviceID *uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) CreateDevice(ctx context.Context, device *model.Device) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.CreateDevice")
	defer span.Finish()

	query := sq.Insert("devices").PlaceholderFormat(sq.Dollar).
		Columns("user_id", "platform", "entered_at").
		Values(device.UserID, device.Platform, device.EnteredAt).
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var id uint64

	err = r.db.GetContext(ctx, &id, s, args...)

	return id, err
}

func (r *repo) DescribeDevice(ctx context.Context, deviceID uint64) (*model.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.DescribeDevice")
	defer span.Finish()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("devices").
		Where(sq.And{sq.Eq{"id": deviceID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var device model.Device

	err = r.db.GetContext(ctx, &device, s, args...)

	return &device, err
}

func (r *repo) DescribeLastDevice(ctx context.Context) (*model.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.DescribeDevice")
	defer span.Finish()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("devices").
		Where(sq.Eq{"removed": false}).
		OrderBy("id DESC").
		Limit(1)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var device model.Device

	err = r.db.GetContext(ctx, &device, s, args...)

	return &device, err
}

func (r *repo) ListDevices(ctx context.Context, page uint64, perPage uint64) ([]*model.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.ListDevices")
	defer span.Finish()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("devices").
		Where(sq.Eq{"removed": false}).
		OrderBy("created_at DESC").
		Limit(perPage).
		Offset((page - 1) * perPage)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var devices []*model.Device

	err = r.db.SelectContext(ctx, &devices, s, args...)

	return devices, err
}

func (r *repo) LogDevice(ctx context.Context, deviceID uint64) ([]*model.DeviceEvent, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.LogDevices")
	defer span.Finish()

	query := sq.Select("id, type, status, created_at, updated_at").PlaceholderFormat(sq.Dollar).
		From("devices_events").
		Where(sq.Eq{"device_id": deviceID}).
		OrderBy("id DESC")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var events []*model.DeviceEvent

	err = r.db.SelectContext(ctx, &events, s, args...)

	return events, err
}

func (r *repo) UpdateDevice(ctx context.Context, device *model.Device) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.UpdateDevice")
	defer span.Finish()

	query := sq.Update("devices").PlaceholderFormat(sq.Dollar).
		Set("platform", device.Platform).
		Set("user_id", device.UserID).
		Where(sq.And{sq.Eq{"id": device.ID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func (r *repo) UpdateLastDevice(ctx context.Context, device *model.Device) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.UpdateDevice")
	defer span.Finish()

	rows, err := r.db.Query("WITH last AS ( SELECT id FROM devices WHERE removed = 'false' ORDER BY id DESC LIMIT 1 ) "+
		"UPDATE devices SET platform=$1, user_id=$2 FROM last WHERE devices.id = last.id RETURNING devices.id;", device.Platform, device.UserID)
	if err != nil {
		return false, err
	}

	rows.Next()
	err = rows.Scan(&device.ID)
	if err != nil {
		return false, err
	}

	err = rows.Close()

	return true, err
}

func (r *repo) RemoveDevice(ctx context.Context, deviceID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.RemoveDevice")
	defer span.Finish()

	query := sq.Update("devices").PlaceholderFormat(sq.Dollar).
		Set("removed", true).
		Where(sq.And{sq.Eq{"id": deviceID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func (r *repo) RemoveLastDevice(ctx context.Context, deviceID *uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.device.RemoveDevice")
	defer span.Finish()

	rows, err := r.db.Query("WITH last AS ( SELECT id FROM devices WHERE removed = 'false' ORDER BY id DESC LIMIT 1 ) " +
		"UPDATE devices SET removed = 'true' FROM last WHERE devices.id = last.id RETURNING devices.id;")
	if err != nil {
		return false, err
	}

	rows.Next()
	err = rows.Scan(deviceID)
	if err != nil {
		return false, err
	}

	err = rows.Close()

	return true, err
}
