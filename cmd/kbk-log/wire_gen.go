// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ZQCard/kbk-log/internal/biz"
	"github.com/ZQCard/kbk-log/internal/conf"
	"github.com/ZQCard/kbk-log/internal/data"
	"github.com/ZQCard/kbk-log/internal/server"
	"github.com/ZQCard/kbk-log/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(env *conf.Env, confServer *conf.Server, confData *conf.Data, bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysqlCmd(bootstrap, logger)
	client := data.NewRedisClient(confData)
	dataData, cleanup, err := data.NewData(bootstrap, db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	logRepo := data.NewLogRepo(dataData, logger)
	logUsecase := biz.NewLogUsecase(logRepo, logger)
	logService := service.NewLogService(logUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, logService, logger)
	httpServer := server.NewHTTPServer(confServer, logService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
