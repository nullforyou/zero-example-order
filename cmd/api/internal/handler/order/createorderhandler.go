package order

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-base/utils/response"
	"net/http"
	"order/cmd/api/internal/logic/order"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
)

func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ValidateErrOrResponse(r, w, err, svcCtx.Validator.Trans)
			return
		}

		l := order.NewCreateOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrder(&req)
		response.Response(r, w, resp, err)
		return
	}
}
