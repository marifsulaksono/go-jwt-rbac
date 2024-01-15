package model

import (
	"go-jwt-rbac/config"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	Phonenumber string         `json:"phonenumber"`
	Role        string         `json:"role"`
	CreateAt    time.Time      `json:"create_at"`
	UpdateAt    time.Time      `json:"update_at"`
	DeleteAt    gorm.DeletedAt `json:"-"`
}

type UserResponse struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Role        string `json:"role"`
}

func (user *User) CreateUser() error {
	user.CreateAt = time.Now()
	return config.DB.Create(&user).Error
}

func GetUserByUsername(usernmae string) (UserResponse, error) {
	var user UserResponse
	err := config.DB.Where("username = ?", usernmae).First(&user).Error
	if err != nil {
		return UserResponse{}, err
	}

	return user, nil
}

func (UserResponse) TableName() string {
	return "users"
}
