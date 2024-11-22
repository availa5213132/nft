package models

import (
	"time"
)

type UserModel struct {
	MODEL
	NickName      string     `json:"nick_name" form:"nick_name" gorm:"column:nick_name;comment:;" binding:"required"` //用户昵称
	UserName      string     `json:"user_name" form:"user_name" gorm:"column:user_name;comment:;" binding:"required"` //用户名
	Avatar        string     `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`                             //头像
	Password      string     `json:"password" form:"password" gorm:"column:password;comment:;" binding:"required"`    //密码
	Email         string     `json:"email" form:"email" gorm:"column:email;comment:;" binding:"required"`             //邮箱
	Phone         string     `json:"phone" form:"phone" gorm:"column:phone;comment:;"`                                //电话号码
	Token         string     `json:"token" form:"token" gorm:"column:token;comment:;"`                                //令牌
	RoleID        *int       `json:"role_id" form:"role_id" gorm:"column:role_id;comment:;" binding:"required"`       //角色
	Status_ID     *int       `json:"status_id" form:"status_id" gorm:"column:status__id;comment:;"`                   //用户状态
	LastLogin     *time.Time `json:"last_login" form:"last_login" gorm:"column:last_login;comment:;"`                 //最后一次登录时间
	Registered_At *time.Time `json:"registered_at" form:"registered_at" gorm:"column:registered__at;comment:;"`       //注册时间
	WalletAddress string     `json:"wallet_address" form:"wallet_address" gorm:"column:wallet_address;comment:;"`     //钱包地址
	IsDelete      string     `json:"is_delete" form:"is_delete" gorm:"column:is_delete;comment:;"`                    //用户是否被软删除
}

// TableName 用户 UserModel自定义表名 user
func (UserModel) TableName() string {
	return "user"
}
