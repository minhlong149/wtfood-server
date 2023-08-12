package handler

import (
	"gorm.io/gorm"

	"wtfood/pkg/service"
	"wtfood/pkg/util"
)

type Handler struct {
	Config  util.Config
	Logger  util.Logger
	Service service.Service
}

func SetupHandler(logger util.Logger, config util.Config, dataSource *gorm.DB) *Handler {
	return &Handler{
		Logger:  logger,
		Config:  config,
		Service: &service.Repository{Db: dataSource},
	}
}
