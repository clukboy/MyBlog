package tag

import (
	"context"

	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailTagLogic {
	return &DetailTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailTagLogic) DetailTag(req *types.TagIdBody) (resp *types.TagInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
