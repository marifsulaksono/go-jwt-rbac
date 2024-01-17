package model

import "go-jwt-rbac/config"

type Cart struct {
	ID        int
	UserId    int
	ProductId int
	Qty       int
}

func GetCartByUser(userId int) ([]Cart, error) {
	var carts []Cart
	err := config.DB.Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (cart *Cart) CreateCart() error {
	return config.DB.Create(&cart).Error
}
