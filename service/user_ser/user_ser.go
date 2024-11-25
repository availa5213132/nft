package user_ser

import (
	"gorm.io/gorm"
	"log"
	"nft/server/global"
	"nft/server/models"
)

// GetUserByOpenID 根据 OpenID 获取用户
func GetUserByOpenID(openID string) (*models.UserModel, error) {
	var user models.UserModel
	err := global.DB.Where("open_id = ?", openID).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error fetching user by OpenID: %v", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil // 用户不存在
	}
	return &user, nil
}

// RegisterUser 注册新用户
func RegisterUser(user *models.UserModel) error {
	err := global.DB.Create(user).Error
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	}
	return nil
}
