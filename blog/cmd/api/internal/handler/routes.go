// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	tag "OneBlog/blog/cmd/api/internal/handler/tag"
	"OneBlog/blog/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: tag.CreateTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: tag.UpateTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/remove/:id",
				Handler: tag.RemoveTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id",
				Handler: tag.DetailTagHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: tag.ListTagHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/tag/v1"),
	)
}
