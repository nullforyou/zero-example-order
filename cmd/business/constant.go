package business

const (
	YY           = "2006"
	YYMM         = "2006-01"
	YYMMDD       = "2006-01-02"
	YYMMDDHH     = "2006-01-02 15"
	YYMMDDHHMM   = "2006-01-02 15:04"
	YYMMDDHHMMSS = "2006-01-02 15:04:05"
)

const CLOSE_ORDER_TIME_MINUTES = 30 //defer close order time

const (
	PAYMENT_LIMIT_TIME = "15m" //订单支付过期时间
)

// 操作状态
const (
	IS_SURE = 1 //表示肯定
	IS_DENY = 0 //表示否定
)

const (
	USER_CANCEL   = 1 //用户取消
	SYSTEM_CANCEL = 2 //系统取消
	STAFF_CANCEL  = 3 //平台管理取消
)

const (
	/* 订单状态 begin */
	CANCELLED_STATE    = -10 //已取消
	WAIT_PAYMENT_STATE = 10  //新创建待付款
	PAID_STATE         = 20  //已支付待完成
	FINISHED_STATE     = 40  //已完成待结算
	SETTLED_STATE      = 50  //已结算
	IS_CLOSED          = -20 //已关闭
	/* 订单状态 end */
)

// OrderStateZhCN 订单状态文字描述
func OrderStateZhCN(code int64) string {
	var orderStateZhCN = map[int64]string{
		CANCELLED_STATE:    "已取消",
		WAIT_PAYMENT_STATE: "待支付",
		PAID_STATE:         "已支付",
		FINISHED_STATE:     "已完成",
		SETTLED_STATE:      "已结算",
		IS_CLOSED:          "已关闭",
	}
	return orderStateZhCN[code]
}
