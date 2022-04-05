package tag

import (
	"net/http"

	"OneBlog/blog/cmd/api/internal/logic/tag"
	"OneBlog/blog/cmd/api/internal/svc"
	"OneBlog/blog/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tag.NewUpateTagLogic(r.Context(), svcCtx)
		resp, err := l.UpateTag(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
