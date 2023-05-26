package main

import (
	"go-zero-example/cmd/database"
	"gorm.io/plugin/soft_delete"
	"time"
)

type GreetMember struct {
	ID              int64                 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:Id;"`
	Mobile   		string                `gorm:"column:mobile;type:char(11);index:uk_mobile;not null;default:'';comment:登陆手机号;"`
	Nickname   		string                `gorm:"column:nickname;type:varchar(50);not null;default:'';comment:用户昵称;"`
	Password	    string                `gorm:"column:password;type:char(32);not null;default:'';comment:密码;"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null"`
	CreatedAt       *time.Time            `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt       *time.Time            `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
}

func (GreetMember) TableName() string {
	return "greet_member"
}

func main()  {
	db := database.Connect()
	db.AutoMigrate(&GreetMember{})
}