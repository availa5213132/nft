package models

// 购物车 结构体  CartModel
type CartModel struct {
	MODEL
	UserID   *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:;" binding:"required"`       //用户ID
	StatusID *int `json:"status_id" form:"status_id" gorm:"column:status_id;comment:;" binding:"required"` //购物车状态
}

// TableName 购物车 CartModel自定义表名 cart_model
func (CartModel) TableName() string {
	return "cart_model"
}
