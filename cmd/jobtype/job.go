package jobtype

// DeferCloseOrderJob 延时关闭订单
const DeferCloseOrderJob = "defer:order:close"

// DeferCloseOrderPayload 延时关闭订单有效载荷
type DeferCloseOrderPayload struct {
	OrderSerialNumber string
}
