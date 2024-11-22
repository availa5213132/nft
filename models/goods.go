package models

type GoodsModel struct {
	MODEL
	GoodsName     string   `json:"goods_name" form:"goods_name" gorm:"column:goods_name;comment:;" binding:"required"`             //商品名称
	Uuid          *int     `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;" binding:"required"`                               //商品uuid
	Description   string   `json:"description" form:"description" gorm:"column:description;comment:;"`                             //商品简介
	BannerID      *int     `json:"banner_id" form:"banner_id" gorm:"column:banner_id;comment:;" binding:"required"`                //商品图片
	Price         *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`                                               //商品价格
	CategoryID    *int     `json:"category_id" form:"category_id" gorm:"column:category_id;comment:;" binding:"required"`          //商品类型
	StockQuantity *int     `json:"stock_quantity" form:"stock_quantity" gorm:"column:stock_quantity;comment:;" binding:"required"` //库存数量
	StatusID      *int     `json:"status_id" form:"status_id" gorm:"column:status_id;comment:;" binding:"required"`                //商品状态
}

// TableName 商品 GoodsModel自定义表名 goods_model
func (GoodsModel) TableName() string {
	return "goods_model"
}
