package main

import (
	"gorm.io/plugin/soft_delete"
	"time"
	"workbench/cmd/database"
)

type GreetGoods struct {
	ID                int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true"`
	OriginGoodsID     int64                 `gorm:"column:origin_goods_id;type:int;not null;default:0;index:idx_origin_goods_id;comment:第三方商品id;"`
	CategoryID        int64                 `gorm:"column:category_id;type:int;not null;default:0;index:idx_category_id;comment:商品分类id;"`
	GoodsName         string                `gorm:"column:goods_name;type:varchar(50);not null;default:'';comment:商品名称;"`
	GoodsPicture      string                `gorm:"column:goods_picture;type:varchar(200);not null;default:'';comment:商品图片;"`
	GoodsPrice        float64               `gorm:"column:goods_price;type:decimal(10,2);not null;default:0.00;default:0;comment:商品价格;"`
	DiffHash          string                `gorm:"column:diff_hash;type:varchar(32);not null;default:'';comment:值hash;"`
	GoodsStatus       int64                 `gorm:"column:goods_status;type:tinyint(1);not null;default:1;comment:商品状态1:上架 0:下架;"`
	GoodsIsRecommend  int64                 `gorm:"column:goods_is_recommend;type:tinyint(1);not null;default:0;comment:商品是否推荐1:已推荐 0:未推荐;"`
	GoodsSort         int64                 `gorm:"column:goods_sort;type:smallint;not null;default:1;default:0;comment:商品排序;"`
	GoodsSalesVolume  int64                 `gorm:"column:goods_sales_volume;type:int;not null;default:1;comment:商品销量;"`
	ActualSalesVolume int64                 `gorm:"column:actual_sales_volume;type:int;not null;default:0;comment:商品实际销量;"`
	DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	CreatedAt         *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt         *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (GreetGoods) TableName() string {
	return "greet_goods"
}

func main()  {
	db := database.Connect()
	db.AutoMigrate(&GreetGoods{})
}