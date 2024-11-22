package models

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
=======
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
import "nft/server/models/ctype"

type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36;comment:昵称" json:"nick_name,select(c|info)"`                    // 昵称
	UserName   string           `gorm:"size:36;comment:用户名" json:"user_name"`                                  // 用户名
	Password   string           `gorm:"size:128;comment:密码" json:"-"`                                          // 密码
	Avatar     string           `gorm:"size:256;comment:头像" json:"avatar,select(c)"`                           // 头像
	Email      string           `gorm:"size:128;comment:邮箱" json:"email,select(info)"`                         // 邮箱
	Tel        string           `gorm:"size:18;comment:手机号" json:"tel"`                                        // 手机号
	Addr       string           `gorm:"size:64;comment:地址" json:"addr,select(c|info)"`                         // 地址
	Token      string           `gorm:"size:64;comment:其他平台的唯一id" json:"token"`                                // 其他平台的唯一id
	IP         string           `gorm:"size:20;comment:ip" json:"ip,select(c)"`                                // ip地址
	Role       ctype.Role       `gorm:"size:4;default:1;comment:权限，1管理员，2普通用户，3游客" json:"role,select(info)"`   // 权限  1 管理员  2 普通用户  3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6);comment:注册来源，1qq，3邮箱" json:"sign_status,select(info)"` // 注册来源
	Integral   int              `gorm:"default:0;comment:我的积分" json:"integral,select(info)"`                   // 我的积分
	Scope      int              `gorm:"default:0;comment:我的积分" json:"scope,select(info)"`                      // 我的积分
	Sign       string           `gorm:"size:128;comment:我的签名" json:"sign,select(info)"`                        // 我的签名
	Link       string           `gorm:"size:128;comment:我的链接地址" json:"link,select(info)"`                      // 我的链接地址
<<<<<<< HEAD
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
=======
>>>>>>> b28096b5c385046dde09c48bd6e0c0be1de76153
}
