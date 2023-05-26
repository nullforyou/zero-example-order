package main

import (
	"gorm.io/plugin/soft_delete"
	"time"
	"workbench/cmd/database"
)

type GreetOrder struct {
	ID                       int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true"`
	OrderSerialNumber        string                `gorm:"column:order_serial_number;type:varchar(30);not null;default '';comment:订单号;"`
	MemberID                 int64                 `gorm:"column:member_id;type:int;not null;default 0;comment:用户id;"`
	Client                   int64                 `gorm:"column:client;type:tinyint;not null;default 0;comment:下单源 10：安卓 11：IOS 12：WAP 13：小程序 20：PC;"`
	OrderStatus              int64                 `gorm:"column:order_status;type:smallint;not null;default:1;comment:订单状态 -30:支付异常 -10:取消订单 10:新订单待支付 20:已支付 40:已完成待结算 50:已结算;"`
	OrderStatusName          string                `gorm:"column:order_status_name;type:varchar(20);not null;default '';comment:订单状态名称(不需要手动赋值);"`
	OrderAmount              float64               `gorm:"column:order_amount;type:decimal(10,2);not null;default:0.00;default '';comment:订单总金额（生成订单时的总价）;"`
	TechnicalServicesFee     float64               `gorm:"column:technical_services_fee;type:decimal(10,2);not null;default:0.00;comment:技术服务费,order_amount字段会加上技术服务费;"`
	OrderExpressFee          float64               `gorm:"column:order_express_fee;type:decimal(10,2);not null;default:0.00;comment:运费;"`
	GoodsNum                 int64                 `gorm:"column:goods_num;type:smallint;not null;default 0;comment:商品总数量(不需要手动赋值);"`
	OrderRebateRate          float64               `gorm:"column:order_rebate_rate;type:decimal(4,2);not null;default:0.00;comment:订单第三方返点(百分比);"`
	OrderProfit              float64               `gorm:"column:order_profit;type:decimal(10,2);not null;default:0.00;comment:订单利润;"`
	MemberNickname           string                `gorm:"column:member_nickname;type:varchar(50);not null;default '';comment:会员昵称;"`
	MemberMobile             string                `gorm:"column:member_mobile;type:char(11);not null;default '';comment:会员手机号;"`
	OutOrderID               int64                 `gorm:"column:out_order_id;type:int;not null;default 0;comment:第三方订单ID（调用第三方创建订单时返回）;"`
	OutTradeNo               string                `gorm:"column:out_trade_no;type:varchar(30);not null;default '';comment:第三方好涤交易号（是好涤交易号）;"`
	OutState                 int64                 `gorm:"column:out_state;type:smallint;not null;default 0;comment:第三方订单状态（调用第三方创建订单时返回）;"`
	OutStateName             string                `gorm:"column:out_state_name;type:varchar(20);not null;default '';comment:第三方订单状态(不需要手动赋值);"`
	PaymentAmount            float64               `gorm:"column:payment_amount;type:decimal(10,2);not null;default:0.00;comment:应该支付总金额（生成订单时免去一切优惠券、积分抵扣、微币抵扣后需要支付的价格）;"`
	PaymentSn                string                `gorm:"column:payment_sn;type:varchar(32);not null;default '';comment:支付订单号（调用支付服务创建支付时返回）;"`
	PaymentType              int64                 `gorm:"column:payment_type;type:tinyint;not null;default 0;comment:支付方式： 1：银联支付 2：支付宝支付 3：微信支付 4.个人余额 5.小巴余额;"`
	PaymentTime              *time.Time            `gorm:"column:payment_time;type:datetime;default null;comment:支付时间;"`
	PaymentStatus            int64                 `gorm:"column:payment_status;type:smallint;not null;default 0;comment:订单支付状态0:未支付 1:已支付等待支付结果 2:支付成功 3:支付失败;"`
	DeductCoin               float64               `gorm:"column:deduct_coin;type:decimal(10,2);not null;default:0.00;comment:微币抵扣金额;"`
	ReturnCredits            float64               `gorm:"column:return_credits;type:decimal(10,2);not null;default:0.00;comment:赠送积分（订单积分）;"`
	ReturnCreditsStatus      int64                 `gorm:"column:return_credits_status;type:tinyint;not null;default 0;comment:赠送积分状态：-1:赠送失败 0：未赠送 1：已赠送 2:赠送中;"`
	ReturnCreditsFailedCause *string               `gorm:"column:return_credits_failed_cause;type:varchar(100);default null;comment:积分领取失败原因;"`
	IsAppendPrice            int64                 `gorm:"column:is_append_price;type:tinyint(1);not null;default 0;comment:是否补差价 默认0=不需要 1=需要;"`
	AppendPriceStatus        int64                 `gorm:"column:append_price_status;type:tinyint;not null;default 0;comment:补差价状态 0:未支付 1:已支付等待支付结果 2:支付成功3:支付失败;"`
	IsAfterSales             int64                 `gorm:"column:is_after_sales;type:smallint;not null;default 0;comment:介入退款流程 支付服务申请退款=1 支付服务退款完成=2 支付服务退款失败=3;"`
	CancelOperator           *int64                `gorm:"column:cancel_operator;type:tinyint;default null;comment:取消操作：1 用户取消 2 系统取消(未付款自动取消)；3：平台管理取消；4：第三方取消；;"`
	CancelTime               *time.Time            `gorm:"column:cancel_time;type:datetime;comment:取消时间，取消成功才有值(不需要手动赋值);"`
	CancelCause              *string               `gorm:"column:cancel_cause;type:varchar(255);default null;comment:取消原因（如果时自动取消也需要注明为超时自动取消）;"`
	IsUserDelete             int64                 `gorm:"column:is_user_delete;type:tinyint(1);not null;default 0;comment:用户删除(0：未删除； 1：已删除);"`
	UserDeleteTime           *time.Time            `gorm:"column:user_delete_time;type:datetime;default null;comment:用户删除时间，删除成功才有值(不需要手动赋值);"`
	PaymentLimitTime         *time.Time            `gorm:"column:payment_limit_time;type:datetime;default null;comment:支付时限，一般下单后有15分钟支付时限 ，必须在此时间前发起支付出票，超过这个时间订单将自动取消;"`
	SettleTime               *time.Time            `gorm:"column:settle_time;type:datetime;default null;comment:结算时间，结算成功才有值(不需要手动赋值);"`
	CompleteTime             *time.Time            `gorm:"column:complete_time;type:datetime;default null;comment:完成时间，完成才有值(不需要手动赋值);"`
	InvoiceStatus            int64                 `gorm:"column:invoice_status;type:tinyint;not null;default:1;comment:发票状态 1：未开具 2：申请中 3：已开;"`
	AdminRemark              *string               `gorm:"column:admin_remark;type:varchar(255);default null;comment:备注（后台订单备注）;"`
	IsDeductCoin             *int64                `gorm:"column:is_deduct_coin;type:tinyint;default null;comment:是否使用抵扣 默认0=不使用 1使用红包抵扣 2余额抵扣;"`
	DeductBalance            float64               `gorm:"column:deduct_balance;type:decimal(10,2);not null;default:0.00;comment:余额抵扣金额;"`
	DeductType               int64                 `gorm:"column:deduct_type;type:tinyint;not null;default 0;comment:抵扣类型:0-未抵扣 1-红包抵扣 2-个人余额抵扣 3-小巴余额;"`
	DeletedAt                soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	CreatedAt                *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt                *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (GreetOrder) TableName() string {
	return "greet_order"
}

func main()  {
	db := database.Connect()
	db.AutoMigrate(&GreetOrder{})
}