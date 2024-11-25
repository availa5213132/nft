package models

import "time"

// 操作日志表 结构体  OperationLog
type OperationLog struct {
	MODEL
	UserID        *int       `json:"user_id" form:"user_id" gorm:"column:user_id;comment:;" binding:"required"`                      //用户ID
	OperationTime *time.Time `json:"operation-time" form:"operation-time" gorm:"column:operation_time;comment:;" binding:"required"` //操作时间
	OperationType string     `json:"operation_type" form:"operation_type" gorm:"column:operation_type;comment:;" binding:"required"` //操作类型
	Decription    string     `json:"decription" form:"decription" gorm:"column:decription;comment:;" binding:"required"`             //操作描述
}

// TableName 操作日志表 OperationLog自定义表名 operation_log
func (OperationLog) TableName() string {
	return "operation_log"
}
