package grpc

import (
	"context"
	"net/http"
	"tmpl-go-vercel/app/grpc/addons/endpoints"
	"tmpl-go-vercel/app/grpc/addons/server"
	"tmpl-go-vercel/app/grpc/addons/server/options"
	"tmpl-go-vercel/app/services/hello"
	hellopb "tmpl-go-vercel/gen/go/api/hello/v1"

	// "tmpl-go-vercel/app/services/hello"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc_cors "tmpl-go-vercel/app/grpc/addons/cors"
	grpc_security "tmpl-go-vercel/app/grpc/addons/security"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

var (
	GRPCEndpoints = endpoints.GRPCRegistry{}
	zapLogger     *zap.Logger
)

type Handler http.Handler

type ServerOptions struct {
	RecommendedOptions *options.RecommendedOptions
	LogRPC             bool
}

func NewServerOptions(logRpc bool, logger *zap.Logger) *ServerOptions {
	o := &ServerOptions{
		RecommendedOptions: options.NewRecommendedOptions(),
		LogRPC:             logRpc,
	}

	zapLogger = logger
	return o
}

func (o ServerOptions) Config() (*server.Config, error) {
	config := server.NewConfig()
	if err := o.RecommendedOptions.ApplyTo(config); err != nil {
		return nil, err
	}

	GRPCEndpoints.Register(hellopb.RegisterHelloServiceServer, &hello.Service{})
	// GRPCEndpoints.Register(healthcheckpb.RegisterHelloServiceServer, &healthcheck.Service{})
	config.SetGRPCRegistry(GRPCEndpoints)

	opts := []grpc_zap.Option{
		grpc_zap.WithDecider(func(methodFullName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			if err == nil && methodFullName == "/api.healthcheck.v1.StatusService/Status" {
				return false
			}

			// by default you will log all calls
			return o.LogRPC
		}),
	}
	payloadDecider := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		// will not log gRPC calls if it was a call to healthcheck and no error was raised
		if fullMethodName == "/api.healthcheck.v1.StatusService/Status" {
			return false
		}

		// by default you will log all calls
		return o.LogRPC
	}

	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)

	config.GRPCServerOption(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			// grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, opts...),
			grpc_zap.PayloadUnaryServerInterceptor(zapLogger, payloadDecider),
			grpc_cors.UnaryServerInterceptor(grpc_cors.OriginHost(config.CORSOriginHost), grpc_cors.AllowSubdomain(config.CORSAllowSubdomain)),
			grpc_security.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			// grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger, opts...),
			grpc_zap.PayloadStreamServerInterceptor(zapLogger, payloadDecider),
			grpc_cors.StreamServerInterceptor(grpc_cors.OriginHost(config.CORSOriginHost), grpc_cors.AllowSubdomain(config.CORSAllowSubdomain)),
			grpc_security.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)

	return config, nil
}

func (o ServerOptions) New() (Handler, error) {

	config, err := o.Config()
	if err != nil {
		return nil, err
	}

	server, err := config.New()
	if err != nil {
		return nil, err
	}

	s := server.NewGRPCServer(false)
	// hellopb.RegisterHelloServiceServer(s, &hello.Service{})
	// healthcheckpb.RegisterStatusServiceServer(s, &healthcheck.Service{})
	// return grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(origin string) bool {
	// 	// Allow all origins, DO NOT do this in production
	// 	return true
	// }), grpcweb.WithCorsForRegisteredEndpointsOnly(false), grpcweb.WithAllowedRequestHeaders([]string{"*"})), nil
	return grpcweb.WrapServer(s), nil
}
