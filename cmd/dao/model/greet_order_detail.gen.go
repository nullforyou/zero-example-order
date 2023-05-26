// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameOrderDetail = "greet_order_detail"

// OrderDetail mapped from table <greet_order_detail>
type OrderDetail struct {
	ID                   int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	OrderID              int64                 `gorm:"column:order_id;type:int;not null" json:"order_id"`                                      // 订单id
	AppointmentStartTime time.Time             `gorm:"column:appointment_start_time;type:datetime;not null" json:"appointment_start_time"`     // 预约开始时间
	AppointmentEndTime   time.Time             `gorm:"column:appointment_end_time;type:datetime;not null" json:"appointment_end_time"`         // 预约结束时间
	Remark               string                `gorm:"column:remark;type:varchar(255);not null" json:"remark"`                                 // 给工厂的备注
	ExpressRemark        string                `gorm:"column:express_remark;type:varchar(255);not null" json:"express_remark"`                 // 给快递备注
	SenderName           string                `gorm:"column:sender_name;type:varchar(20);not null" json:"sender_name"`                        // 发件人姓名
	SenderMobile         string                `gorm:"column:sender_mobile;type:char(11);not null" json:"sender_mobile"`                       // 发件人手机号
	SenderProvince       string                `gorm:"column:sender_province;type:varchar(20);not null" json:"sender_province"`                // 地区名称 省
	SenderCity           string                `gorm:"column:sender_city;type:varchar(20);not null" json:"sender_city"`                        // 地区名称 市
	SenderCounty         string                `gorm:"column:sender_county;type:varchar(20);not null" json:"sender_county"`                    // 地区名称 县
	SenderAddress        string                `gorm:"column:sender_address;type:varchar(100);not null" json:"sender_address"`                 // 详细地址
	ReceiveName          string                `gorm:"column:receive_name;type:varchar(20);not null" json:"receive_name"`                      // 收件人姓名
	ReceiveMobile        string                `gorm:"column:receive_mobile;type:char(11);not null" json:"receive_mobile"`                     // 收件人手机号
	ReceiveProvince      string                `gorm:"column:receive_province;type:varchar(20);not null" json:"receive_province"`              // 地区名称 省
	ReceiveCity          string                `gorm:"column:receive_city;type:varchar(20);not null" json:"receive_city"`                      // 地区名称 市
	ReceiveCounty        string                `gorm:"column:receive_county;type:varchar(20);not null" json:"receive_county"`                  // 地区名称 县
	ReceiveAddress       string                `gorm:"column:receive_address;type:varchar(100);not null" json:"receive_address"`               // 详细地址
	IsPacking            int64                 `gorm:"column:is_packing;type:tinyint(1);not null" json:"is_packing"`                           // 是否需要包装 0不需要，1需要
	PackFee              float64               `gorm:"column:pack_fee;type:decimal(5,2);not null;default:0.00" json:"pack_fee"`                // 包装费
	OutTotalPrice        float64               `gorm:"column:out_total_price;type:decimal(10,2);not null;default:0.00" json:"out_total_price"` // 第三方订单总价（同步第三方订单时返回）
	SendDeliveryID       string                `gorm:"column:send_delivery_id;type:varchar(15);not null" json:"send_delivery_id"`              // 送程快递信息（同步第三方订单时返回）
	ReceiveDeliveryID    string                `gorm:"column:receive_delivery_id;type:varchar(15);not null" json:"receive_delivery_id"`        // 返程快递信息（同步第三方订单时返回）
	RefundDeliveryID     string                `gorm:"column:refund_delivery_id;type:varchar(15);not null" json:"refund_delivery_id"`          // 退款快递信息（同步第三方订单时返回）
	CreatedAt            *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt            *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime" json:"updated_at"`
	DeletedAt            soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null" json:"deleted_at"`
	CityName             *string               `gorm:"column:city_name;type:varchar(122)" json:"city_name"`
}

// TableName OrderDetail's table name
func (*OrderDetail) TableName() string {
	return TableNameOrderDetail
}