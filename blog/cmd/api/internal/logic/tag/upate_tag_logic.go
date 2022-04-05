package tag

import (
	"context"

	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpateTagLogic {
	return &UpateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpateTagLogic) UpateTag(req *types.TagInfo) (resp *types.TagIdBody, err error) {
	// todo: add your logic here and delete this line

	return
}
