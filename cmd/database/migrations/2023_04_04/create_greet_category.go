package main

import (
	"go-zero-example/cmd/database"
	"gorm.io/plugin/soft_delete"
	"time"
)

type GreetCategory struct {
	ID              int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true"`
	CategoryID      int64                 `gorm:"column:category_id;type:int;not null;default:0;index:idx_category_id;comment:第三方Id;"`
	CategoryPid     int64                 `gorm:"column:category_pid;type:int;not null;default:0;index:idx_category_pid;comment:第三方父级Id;"`
	CategoryName    string                `gorm:"column:category_name;type:varchar(10);not null;default:'';comment:分类名称;"`
	CategoryPicture string                `gorm:"column:category_picture;type:varchar(200);not null;default:'';comment:分类图片;"`
	DiffHash        string                `gorm:"column:diff_hash;type:varchar(32);not null;default:'';comment:值hash;"`
	IsRecommend     int64                 `gorm:"column:is_recommend;type:tinyint(1);not null;default:0;comment:是否推荐1:已推荐 0:未推荐;"`
	CategoryStatus  int64                 `gorm:"column:category_status;type:tinyint(1);not null;default:1;default:1;comment:分类状态1：使用 0：禁用;"`
	OperationUser   int64                 `gorm:"column:operation_user;type:int;not null;default:0;comment:操作人id;"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	CreatedAt       *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt       *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (GreetCategory) TableName() string {
	return "greet_category"
}

func main()  {
	db := database.Connect()
	db.AutoMigrate(&GreetCategory{})
}