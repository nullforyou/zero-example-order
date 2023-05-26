package order

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"go-common/ctxdata"
	"go-common/utils/xerr"
	"gorm.io/gorm"
	"greet-pb/pb/types/user"
	"order/cmd/api/business"
	"order/cmd/dao/model"
	"order/cmd/dao/query"
	"strconv"
	"time"

	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	userId, err := ctxdata.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	query.SetDefault(l.svcCtx.DbEngine)
	ctx := context.Background()
	orderDAL := query.Goods
	//所选商品的总价格
	goodsTotalAmount := decimal.NewFromFloat(0)

	order := model.Order{}
	for _, createGoods := range req.OrderInfo.Goods {
		goodsModel, err := orderDAL.WithContext(ctx).Where(orderDAL.ID.Eq(createGoods.GoodsId)).First()
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerr.NewBusinessError(xerr.SetCode("ErrorGoodsNotExists"))
		}

		orderGoods := model.OrderGoods{
			CategoryID: goodsModel.Category.CategoryID,
			CategoryName: goodsModel.Category.CategoryName,
			GoodsID: goodsModel.ID,
			GoodsName: goodsModel.GoodsName,
			GoodsPicture: goodsModel.GoodsPicture,
			GoodsPrice: goodsModel.GoodsPrice,
			GoodsNum: createGoods.Num,
		}
		order.Goods = append(order.Goods, orderGoods)

		goodsTotalAmount = goodsTotalAmount.Add(decimal.NewFromFloat(goodsModel.GoodsPrice).Mul(decimal.NewFromInt(createGoods.Num)))
	}

	//验证订单金额
	if !goodsTotalAmount.Equal(decimal.NewFromFloat(req.TotalPrice)) {
		return nil, xerr.NewBusinessError(xerr.SetCode("ErrorOrderPrice"))
	}

	order.MemberID = 1
	order.Client = req.OrderInfo.Client
	order.OrderStatus = business.WAIT_PAYMENT_STATE
	order.OrderStatusName = business.OrderStateZhCN(order.OrderStatus)
	order.OrderAmount, _ = strconv.ParseFloat(goodsTotalAmount.String(), 64)
	order.OrderSerialNumber = business.GenerateOrderNumber(time.Now())

	// 使用user rpc
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	order.MemberMobile = userInfo.Mobile
	order.MemberNickname = userInfo.Nikename

	limitTime, _ := time.ParseDuration(business.PAYMENT_LIMIT_TIME)
	paymentLimitTime := time.Now().Add(limitTime)
	order.PaymentLimitTime = &paymentLimitTime

	orderDetail := model.OrderDetail{}

	appointmentStartTime, _ := time.ParseInLocation(business.YYMMDDHHMM, req.OrderInfo.AppointmentStartTime, time.Local)
	orderDetail.AppointmentStartTime = appointmentStartTime

	appointmentEndTime, _ := time.ParseInLocation(business.YYMMDDHHMM, req.OrderInfo.AppointmentEndTime, time.Local)
	orderDetail.AppointmentEndTime = appointmentEndTime

	orderDetail.Remark = req.OrderInfo.Remark

	addressDal := query.Address
	senderAddressModel, err := addressDal.WithContext(ctx).Where(addressDal.ID.Eq(req.OrderInfo.SenderAddressId)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode("ErrorAddressNotExists"))
	}
	orderDetail.SenderName = senderAddressModel.ContactName
	orderDetail.SenderMobile = senderAddressModel.ContactMobile
	orderDetail.SenderProvince = senderAddressModel.ProvinceName
	orderDetail.SenderCity = senderAddressModel.CityName
	orderDetail.SenderCounty = senderAddressModel.CountyName
	orderDetail.SenderAddress = senderAddressModel.DetailedAddress



	receiveAddressModel, err := addressDal.WithContext(ctx).Where(addressDal.ID.Eq(req.OrderInfo.ReceiveAddressId)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode("ErrorAddressNotExists"))
	}
	orderDetail.ReceiveName = receiveAddressModel.ContactName
	orderDetail.ReceiveMobile = receiveAddressModel.ContactMobile
	orderDetail.ReceiveProvince = receiveAddressModel.ProvinceName
	orderDetail.ReceiveCity = receiveAddressModel.CityName
	orderDetail.ReceiveCounty = receiveAddressModel.CountyName
	orderDetail.ReceiveAddress = receiveAddressModel.DetailedAddress
	orderDetail.PackFee = 0
	orderDetail.IsPacking = 0

	order.OrderDetail = orderDetail

	result := l.svcCtx.DbEngine.Create(&order)

	return &types.CreateOrderResp{Id: order.ID, OrderSerialNumber: order.OrderSerialNumber},result.Error
}
