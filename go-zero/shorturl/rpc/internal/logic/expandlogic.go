package logic

import (
	"context"

	"shorturl/rpc/internal/svc"
	"shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	shorturl, err := l.svcCtx.ShorturlModel.FindOne(l.ctx, in.Shorten)
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{
		Url: shorturl.Url,
	}, nil
}
