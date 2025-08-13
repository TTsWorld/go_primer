package svc

import (
	"shorturl/rpc/internal/config"
	"shorturl/rpc/transform/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShorturlModel model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ShorturlModel: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
