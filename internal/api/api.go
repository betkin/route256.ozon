package api

import (
	"context"
	"database/sql"
	"time"

	"github.com/opentracing/opentracing-go"
	repo2 "gitlab.ozon.dev/betkin/device-api/internal/app/repo"
	"gitlab.ozon.dev/betkin/device-api/internal/model"
	"gitlab.ozon.dev/betkin/device-api/internal/pkg/logger"
	tspb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "gitlab.ozon.dev/betkin/device-api/pkg/act-device-api"
)

var (
	totalDeviceNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "act_device_api",
		Name:      "device_not_found_total",
		Help:      "Total number of devices that were not found",
	})

	cudActionsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "act_device_api",
		Name:      "cud_actions_total",
	}, []string{"action_type"})
)

type deviceAPI struct {
	pb.UnimplementedActDeviceApiServiceServer
	repo      repo2.Repo
	eventRepo repo2.EventRepo
}

// NewDeviceAPI returns api of act-device-api service
func NewDeviceAPI(r repo2.Repo, er repo2.EventRepo) pb.ActDeviceApiServiceServer {
	return &deviceAPI{repo: r, eventRepo: er}
}

func (o *deviceAPI) CreateDeviceV1(
	ctx context.Context,
	req *pb.CreateDeviceV1Request,
) (*pb.CreateDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	now := time.Now()

	device := &model.Device{
		UserID:    req.GetUserId(),
		Platform:  req.GetPlatform(),
		EnteredAt: &now,
	}

	deviceId, err := o.repo.CreateDevice(ctx, device)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	err = o.eventRepo.Add(ctx, &model.DeviceEvent{
		DeviceID: deviceId,
		Type:     model.Created,
		Status:   model.Deferred,
		Device:   device,
	})
	if err != nil {
		logger.ErrorKV(
			ctx,
			"CreateDeviceV1 -- failed record to event table",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	cudActionsTotal.WithLabelValues("create").Inc()

	logger.DebugKV(ctx, "CreateDeviceV1 -- success")

	return &pb.CreateDeviceV1Response{
		DeviceId: deviceId,
	}, nil
}

func (o *deviceAPI) DescribeDeviceV1(
	ctx context.Context,
	req *pb.DescribeDeviceV1Request,
) (*pb.DescribeDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	device, err := o.repo.DescribeDevice(ctx, req.GetDeviceId())
	if err != nil && err != sql.ErrNoRows {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if device == nil || err == sql.ErrNoRows {
		logger.DebugKV(
			ctx,
			"DescribeDeviceV1 -- device not found",
			"deviceId", req.DeviceId,
		)
		totalDeviceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "device not found")
	}

	logger.DebugKV(ctx, "DescribeDeviceV1 -- success")

	return &pb.DescribeDeviceV1Response{
		Value: &pb.Device{
			Id:        device.ID,
			Platform:  device.Platform,
			UserId:    device.UserID,
			EnteredAt: tspb.New(*device.EnteredAt),
		},
	}, nil
}

func (o *deviceAPI) DescribeLastDeviceV1(
	ctx context.Context,
	req *pb.Empty,
) (*pb.Device, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	device, err := o.repo.DescribeLastDevice(ctx)
	if err != nil && err != sql.ErrNoRows {
		logger.ErrorKV(
			ctx,
			"DescribeDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if device == nil || err == sql.ErrNoRows {
		logger.DebugKV(
			ctx,
			"DescribeDeviceV1 -- any device not found",
		)
		totalDeviceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "device not found")
	}

	logger.DebugKV(ctx, "DescribeDeviceV1 -- success")

	return &pb.Device{
		Id:        device.ID,
		Platform:  device.Platform,
		UserId:    device.UserID,
		EnteredAt: tspb.New(*device.EnteredAt),
	}, nil
}

func (o *deviceAPI) LogDeviceV1(
	ctx context.Context,
	req *pb.LogDeviceV1Request,
) (*pb.LogDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.LogDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"LogDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logs, err := o.repo.LogDevice(ctx, req.GetDeviceId())
	if err != nil {
		logger.ErrorKV(
			ctx,
			"LogDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "LogDeviceV1 -- success")

	var pbLogs []*pb.DeviceLog

	for _, log := range logs {
		pbLogs = append(pbLogs,
			&pb.DeviceLog{
				Id:        log.ID,
				Type:      uint64(log.Type),
				Status:    uint64(log.Status),
				CreatedAt: tspb.New(log.CreatedAt),
				UpdatedAt: tspb.New(log.UpdatedAt),
			},
		)
	}

	return &pb.LogDeviceV1Response{
		Items: pbLogs,
	}, nil
}

func (o *deviceAPI) ListDevicesV1(
	ctx context.Context,
	req *pb.ListDevicesV1Request,
) (*pb.ListDevicesV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.ListDevicesV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"ListDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	devices, err := o.repo.ListDevices(ctx, req.GetPage(), req.GetPerPage())
	if err != nil {
		logger.ErrorKV(
			ctx,
			"ListDevicesV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "ListDevicesV1 -- success")

	var pbDevices []*pb.Device

	for _, device := range devices {
		pbDevices = append(pbDevices,
			&pb.Device{
				Id:        device.ID,
				Platform:  device.Platform,
				UserId:    device.UserID,
				EnteredAt: tspb.New(*device.EnteredAt),
			},
		)
	}

	return &pb.ListDevicesV1Response{
		Items: pbDevices,
	}, nil
}

func (o *deviceAPI) UpdateDeviceV1(
	ctx context.Context,
	req *pb.UpdateDeviceV1Request,
) (*pb.UpdateDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.UpdateDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	deviceId := req.GetDeviceId()

	device := &model.Device{
		ID:       deviceId,
		UserID:   req.GetUserId(),
		Platform: req.GetPlatform(),
	}

	success, err := o.repo.UpdateDevice(ctx, device)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if success {
		cudActionsTotal.WithLabelValues("update").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: deviceId,
			Type:     model.Updated,
			Status:   model.Deferred,
			Device:   device,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"UpdateDeviceV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "UpdateDeviceV1 -- success")

	return &pb.UpdateDeviceV1Response{
		Success: success,
	}, nil
}

func (o *deviceAPI) UpdateLastDeviceV1(
	ctx context.Context,
	req *pb.UpdateLastDeviceV1Request,
) (*pb.UpdateDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.UpdateDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	device := &model.Device{
		UserID:   req.GetUserId(),
		Platform: req.GetPlatform(),
	}

	success, err := o.repo.UpdateLastDevice(ctx, device)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"UpdateDeviceV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if success {
		cudActionsTotal.WithLabelValues("update").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: device.ID,
			Type:     model.Updated,
			Status:   model.Deferred,
			Device:   device,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"UpdateDeviceV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "UpdateDeviceV1 -- success")

	return &pb.UpdateDeviceV1Response{
		Success: success,
	}, nil
}

func (o *deviceAPI) RemoveDeviceV1(
	ctx context.Context,
	req *pb.RemoveDeviceV1Request,
) (*pb.RemoveDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	deviceId := req.GetDeviceId()

	found, err := o.repo.RemoveDevice(ctx, deviceId)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		totalDeviceNotFound.Inc()
	} else {
		cudActionsTotal.WithLabelValues("remove").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: deviceId,
			Type:     model.Removed,
			Status:   model.Deferred,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"RemoveDevicesV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "RemoveDevicesV1 -- success")

	return &pb.RemoveDeviceV1Response{
		Found: found,
	}, nil
}

func (o *deviceAPI) RemoveLastDeviceV1(
	ctx context.Context,
	req *pb.Empty,
) (*pb.RemoveDeviceV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveDeviceV1")
	defer span.Finish()

	ctx = logger.LogLevelFromContext(ctx)

	if err := req.Validate(); err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var id uint64
	found, err := o.repo.RemoveLastDevice(ctx, &id)
	if err != nil {
		logger.ErrorKV(
			ctx,
			"RemoveDevicesV1 -- failed",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		totalDeviceNotFound.Inc()
	} else {
		cudActionsTotal.WithLabelValues("remove").Inc()

		err = o.eventRepo.Add(ctx, &model.DeviceEvent{
			DeviceID: id,
			Type:     model.Removed,
			Status:   model.Deferred,
		})
		if err != nil {
			logger.ErrorKV(
				ctx,
				"RemoveDevicesV1 -- failed record to event table",
				"err", err,
			)

			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	logger.DebugKV(ctx, "RemoveDevicesV1 -- success")

	return &pb.RemoveDeviceV1Response{
		Found: found,
	}, nil
}