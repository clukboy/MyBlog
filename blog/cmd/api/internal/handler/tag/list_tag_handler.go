package tag

import (
	"net/http"

	"OneBlog/blog/cmd/api/internal/logic/tag"
	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewListTagLogic(r.Context(), svcCtx)
		resp, err := l.ListTag(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
