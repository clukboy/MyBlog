package tag

import (
	"context"

	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTagLogic {
	return &RemoveTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveTagLogic) RemoveTag(req *types.TagIdBody) (resp *types.TagIdBody, err error) {
	// todo: add your logic here and delete this line

	return
}
