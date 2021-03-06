// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"tencent/internal/biz"
	"tencent/internal/conf"
	"tencent/internal/data"
	"tencent/internal/server"
	"tencent/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	instanceRepo := data.NewInstanceRepo(dataData, logger)
	instanceUsecase := biz.NewInstanceUseCase(instanceRepo, logger)
	regionRepo := data.NewRegionRepo(dataData, logger)
	regionUsecase := biz.NewRegionUseCase(regionRepo, logger)
	imageRepo := data.NewImageRepo(dataData, logger)
	imageUsecase := biz.NewImageUseCase(imageRepo, logger)
	instanceTypeRepo := data.NewInstanceTypeRepo(dataData, logger)
	instanceTypeUsecase := biz.NewInstanceTypeUseCase(instanceTypeRepo, logger)
	instanceService := service.NewInstanceService(instanceUsecase, regionUsecase, imageUsecase, instanceTypeUsecase)
	httpServer := server.NewHTTPServer(confServer, instanceService, logger)
	grpcServer := server.NewGRPCServer(confServer, instanceService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
