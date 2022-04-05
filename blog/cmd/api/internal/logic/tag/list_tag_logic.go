package tag

import (
	"context"

	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagLogic {
	return &ListTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagLogic) ListTag(req *types.TagPageReq) (resp *types.TagInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
