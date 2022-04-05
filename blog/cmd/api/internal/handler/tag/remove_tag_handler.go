package tag

import (
	"net/http"

	"OneBlog/blog/cmd/api/internal/logic/tag"
	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagIdBody
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewRemoveTagLogic(r.Context(), svcCtx)
		resp, err := l.RemoveTag(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
