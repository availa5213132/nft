package models

import "time"

// 订单表 结构体  OrderModel
type OrderModel struct {
	MODEL
	UserID         *int       `json:"user_id" form:"user_id" gorm:"column:user_id;comment:;" binding:"required"`                       //订单所属哪个用户
	GoodsID        *int       `json:"goods_id" form:"goods_id" gorm:"column:goods_id;comment:;" binding:"required"`                    //商品ID
	Price          *int       `json:"price" form:"price" gorm:"column:price;comment:;" binding:"required"`                             //商品单价
	Quantity       *int       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;" binding:"required"`                    //购买商品的数量
	TotalAmount    *float64   `json:"total_amount" form:"total_amount" gorm:"column:total_amount;comment:;" binding:"required"`        //订单总金额
	OrderNumber    *int       `json:"order_number" form:"order_number" gorm:"column:order_number;comment:;" binding:"required"`        //订单号
	StatusID       *int       `json:"status_id" form:"status_id" gorm:"column:status_id;comment:;" binding:"required"`                 //支付状态
	OrderStatus    string     `json:"order_status" form:"order_status" gorm:"column:order_status;comment:;" binding:"required"`        //订单状态
	LockExpiration *time.Time `json:"lockExpiration" form:"lockExpiration" gorm:"column:lock_expiration;comment:;" binding:"required"` //锁单时间
}

// TableName 订单表 OrderModel自定义表名 order_model
func (OrderModel) TableName() string {
	return "order_model"
}
