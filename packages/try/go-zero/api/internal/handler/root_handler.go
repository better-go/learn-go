package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero/api/internal/logic"
	"zero/api/internal/svc"
)

func RootHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRootLogic(r.Context(), svcCtx)
		resp, err := l.Root()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
