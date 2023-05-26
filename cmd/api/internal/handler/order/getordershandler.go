package order

import (
	"go-common/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"order/cmd/api/internal/logic/order"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
)

func GetOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrdersCollectionReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResponse(r, w, err)
			return
		}

		l := order.NewGetOrdersLogic(r.Context(), svcCtx)
		resp, err := l.GetOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(r, w, resp)
		}
	}
}
