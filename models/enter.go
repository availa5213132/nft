package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`  // 创建时间
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`                        // 更新时间
}
