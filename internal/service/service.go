package service

import (
	"Aoi/global"
	"Aoi/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context) *Service {
	return &Service{ctx: ctx, dao: dao.NewDB(global.Db)}
}
