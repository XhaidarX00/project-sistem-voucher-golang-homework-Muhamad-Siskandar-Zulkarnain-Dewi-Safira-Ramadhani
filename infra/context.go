package infra

import (
	"project-voucher-team3/config"
	"project-voucher-team3/controller"
	"project-voucher-team3/database"
	"project-voucher-team3/log"
	"project-voucher-team3/repository"
	"project-voucher-team3/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.Controller
	Log *zap.Logger
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := controller.NewController(service, log)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log}, nil
}
