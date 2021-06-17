package dao

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type BasePmapLog struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Model string `json:"model"`
	FileName string `json:"fileName"`
	EndTime time.Time `json:"endTime"`
	EndTimeStamp int64 `json:"endTimeStamp"`
	StartTime time.Time `json:"startTime"`
	StartTimeStamp int64 `json:"startTimeStamp"`
	Size float32 `json:"size"`
	SN string `json:"sn"`
	isEnd int `json:"isEnd" validate:"min=0,max=1"`

}

// UserBaseModel User represents a registered user.
type UserBaseModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username  string    `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password  string    `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Phone     int64     `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Sex       int       `gorm:"column:sex" json:"sex"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Validate the fields.
func (u *UserBaseModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// TableName 表名
func (p *BasePmapLog) TableName() string {
	return "P1904"
}

