package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"
)

func TestTransCtx(t *testing.T) {
	model := NewArticleCategoryModel(nil, nil)
	model.TransCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		model.Insert(ctx, session, &ArticleCategory{})
		return nil
	})
}
