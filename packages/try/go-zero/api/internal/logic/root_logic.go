package logic

import (
	"context"

	"zero/api/internal/svc"
	"zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RootLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRootLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RootLogic {
	return &RootLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RootLogic) Root() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return &types.Response{
		Message: "Hello go-zero, home page!",
	}, nil
}
