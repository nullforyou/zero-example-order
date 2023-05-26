package main

import (
	"go-zero-example/cmd/database"
	"gorm.io/plugin/soft_delete"
	"time"
)

type GreetAddress struct {
	ID              int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:Id;"`
	MemberID        int64                 `gorm:"column:member_id;type:int;not null;index:idx_member_id;default:0;comment:用户id;"`
	ContactName     string                `gorm:"column:contact_name;type:varchar(10);not null;default:'';comment:联系人姓名;"`
	ContactMobile   string                `gorm:"column:contact_mobile;type:char(11);not null;default:'';comment:联系人手机号;"`
	ProvinceCode    string                `gorm:"column:province_code;type:varchar(6);not null;default:'';comment:地区编码 省;"`
	CityCode        string                `gorm:"column:city_code;type:varchar(6);not null;default:'';comment:地区编码 市;"`
	CountyCode      string                `gorm:"column:county_code;type:varchar(6);not null;default:'';comment:地区编码 县;"`
	ProvinceName    string                `gorm:"column:province_name;type:varchar(20);not null;default:'';comment:地区名称 省;"`
	CityName        string                `gorm:"column:city_name;type:varchar(20);not null;default:'';comment:地区名称 市;"`
	CountyName      string                `gorm:"column:county_name;type:varchar(20);not null;default:'';comment:地区名称 县;"`
	DetailedAddress string                `gorm:"column:detailed_address;type:varchar(100);not null;default:'';comment:详细地址;"`
	FirstBoot       int64                 `gorm:"column:first_boot;type:tinyint(1);not null;default:0;comment:默认地址1：是；0：否；;"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp"`
	CreatedAt       *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt       *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (GreetAddress) TableName() string {
	return "greet_address"
}

func main()  {
	db := database.Connect()
	db.AutoMigrate(&GreetAddress{})
}