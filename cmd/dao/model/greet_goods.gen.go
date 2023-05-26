// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameGoods = "greet_goods"

// Goods mapped from table <greet_goods>
type Goods struct {
	ID                int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CategoryID        int64                 `gorm:"column:category_id;type:int;not null" json:"category_id"`                        // 商品分类id
	GoodsName         string                `gorm:"column:goods_name;type:varchar(50);not null" json:"goods_name"`                  // 商品名称
	GoodsPicture      string                `gorm:"column:goods_picture;type:varchar(200);not null" json:"goods_picture"`           // 商品图片
	GoodsPrice        float64               `gorm:"column:goods_price;type:decimal(10,2);not null;default:0.00" json:"goods_price"` // 商品价格
	GoodsStatus       int64                 `gorm:"column:goods_status;type:tinyint(1);not null;default:1" json:"goods_status"`     // 商品状态1:上架;0:下架;
	GoodsIsRecommend  int64                 `gorm:"column:goods_is_recommend;type:tinyint(1);not null" json:"goods_is_recommend"`   // 商品是否推荐1:已推荐;0:未推荐;
	GoodsSort         int64                 `gorm:"column:goods_sort;type:smallint;not null;default:1" json:"goods_sort"`           // 商品排序
	GoodsSalesVolume  int64                 `gorm:"column:goods_sales_volume;type:int;not null" json:"goods_sales_volume"`          // 商品销量
	ActualSalesVolume int64                 `gorm:"column:actual_sales_volume;type:int;not null" json:"actual_sales_volume"`        // 商品实际销量
	DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null" json:"deleted_at"`
	CreatedAt         *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt         *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime" json:"updated_at"`

	Category Category
}

// TableName Goods's table name
func (*Goods) TableName() string {
	return TableNameGoods
}
