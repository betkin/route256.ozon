package server

import (
	"context"
	"gitlab.ozon.dev/betkin/device-api/internal/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			logRequest := md.Get("log-request")
			if (len(logRequest) > 0) && (logRequest[0] == "true") {
				logger.DebugKV(
					ctx,
					"request data",
					"method", info.FullMethod,
					"metadata", md,
					"requestBody", req,
				)
			}
		}

		return handler(ctx, req)
	}
}

func ResponseLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		response, err := handler(ctx, req)

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			logResponse := md.Get("log-response")

			if (len(logResponse) > 0) && (logResponse[0] == "true") {
				logger.DebugKV(
					ctx,
					"response data",
					"method", info.FullMethod,
					"responseBody", response,
					"err", err,
				)
			}
		}

		return response, err
	}
}
