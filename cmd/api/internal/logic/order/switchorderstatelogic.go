package order

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"greet-pb/order/types/order"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
	"order/cmd/business"
	"order/cmd/dao/query"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwitchOrderStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwitchOrderStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwitchOrderStateLogic {
	return &SwitchOrderStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwitchOrderStateLogic) SwitchOrderState(req *types.SwitchOrderStateReq) (resp *types.SwitchOrderStateResp, err error) {
	query.SetDefault(l.svcCtx.DbEngine)
	orderDao := query.Order
	count, err := orderDao.WithContext(context.Background()).Where(orderDao.OrderSerialNumber.Eq(req.OrderSerialNumber)).Count()
	if count == 0 || err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewBusinessError(xerr.SetCode(xerr.ErrorNotFound), xerr.SetMsg("订单不存在")), "在查询订单数据库时错误 %+v", err)
	}
	if strings.EqualFold(req.State, "cancel") {
		_, err := l.svcCtx.OrderRpc.CancelOrder(l.ctx, &order.CancelOrderReq{OrderSerialNumber: req.OrderSerialNumber, CancelOperator: business.USER_CANCEL, CancelCause: "用户取消"})
		if err != nil {
			return nil, err
		}
	}
	if strings.EqualFold(req.State, "delete") {
		_, err := l.svcCtx.OrderRpc.DeleteOrder(l.ctx, &order.DeleteOrderReq{OrderSerialNumber: req.OrderSerialNumber})
		if err != nil {
			return nil, err
		}
	}
	return &types.SwitchOrderStateResp{OrderSerialNumber: req.OrderSerialNumber}, nil
}
