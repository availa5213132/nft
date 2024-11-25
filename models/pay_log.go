package models

// PayLogModel 支付日志表 结构体
//type PayLogModel struct {
//	MODEL
//	OrderID     *int           `json:"order_id" form:"order_id" gorm:"column:order_id;comment:;" binding:"required"`                                           //订单ID
//	UserID      *int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:;" binding:"required"`                                              //用户ID
//	RequestData datatypes.JSON `json:"request_data" form:"request_data" gorm:"column:request_data;comment:;type:text;" binding:"required"swaggertype:"object"` //请求数据信息
//	Amount      *float64       `json:"amount" form:"amount" gorm:"column:amount;comment:;" binding:"required"`                                                 //支付金额
//	PayMethod   string         `json:"pay_method" form:"pay_method" gorm:"column:pay_method;comment:;" binding:"required"`                                     //支付方式
//	StatusID    *int           `json:"status_id" form:"status_id" gorm:"column:status_id;comment:;" binding:"required"`                                        //支付状态
//	ErrCode     *int           `json:"err_code" form:"err_code" gorm:"column:err_code;comment:;" binding:"required"`                                           //错误码
//	ErrMessage  string         `json:"err_message" form:"err_message" gorm:"column:err_message;comment:;" binding:"required"`                                  //错误信息
//}
//
//// TableName 支付日志表 PayLogModel自定义表名 pay_log_model
//func (PayLogModel) TableName() string {
//	return "pay_log_model"
//}
