package models

// 库存表 结构体  InventoryModel
type InventoryModel struct {
	MODEL
	GoodsID  *int `json:"goods_id" form:"goods_id" gorm:"column:goods_id;comment:;" binding:"required"` //商品ID
	Quantity *int `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;" binding:"required"` //库存数量
}

// TableName 库存表 InventoryModel自定义表名 inventory_model
func (InventoryModel) TableName() string {
	return "inventory_model"
}
