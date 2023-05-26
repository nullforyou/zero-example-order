package business

const (
	YY = "2006"
	YYMM = "2006-01"
	YYMMDD = "2006-01-02"
	YYMMDDHH = "2006-01-02 15"
	YYMMDDHHMM = "2006-01-02 15:04"
	YYMMDDHHMMSS = "2006-01-02 15:04:05"
)

const (
	PAYMENT_LIMIT_TIME = "15m" //订单支付过期时间
)

const (
	/* 订单状态 begin */
	CANCELLED_STATE = -10 //已取消
	WAIT_PAYMENT_STATE = 10 //新创建待付款
	PAID_STATE = 20 //已支付待完成
	FINISHED_STATE = 40 //已完成待结算
	SETTLED_STATE = 50 //已结算
	CLOSED = -50 //已关闭 (此状态只用于筛选和展示,不需要保存到数据库)
	APPEND_WAIT_PAYMENT_STATE = 11 //补差价 (此状态只用于筛选和展示,不需要保存到数据库)
	/* 订单状态 end */
)

// OrderStateZhCN 订单状态文字描述
func OrderStateZhCN(code int64) string {
	var orderStateZhCN = map[int64] string {
		CANCELLED_STATE : "已取消",
		WAIT_PAYMENT_STATE : "待支付",
		PAID_STATE : "已支付",
		FINISHED_STATE : "已完成",
		SETTLED_STATE : "已结算",
		CLOSED : "已关闭",
		APPEND_WAIT_PAYMENT_STATE : "补差价",
	}
	return orderStateZhCN[code]
}

