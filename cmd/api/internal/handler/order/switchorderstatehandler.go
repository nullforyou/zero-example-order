package order

import (
	"go-zero-base/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"order/cmd/api/internal/logic/order"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
)

func SwitchOrderStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SwitchOrderStateReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParseParamErrResponse(r, w, err)
			return
		}

		if err := svcCtx.Validator.Validate.StructCtx(r.Context(), req); err != nil {
			response.ValidateErrResponse(r, w, err, svcCtx.Validator.Trans)
			return
		}

		l := order.NewSwitchOrderStateLogic(r.Context(), svcCtx)
		resp, err := l.SwitchOrderState(&req)
		response.Response(r, w, resp, err)
		return
	}
}
