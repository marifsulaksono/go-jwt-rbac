package model

import (
	"go-jwt-rbac/config"

	"gorm.io/gorm"
)

type Product struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	Sold       int            `json:"sold"`
	Desc       string         `json:"desc"`
	Status     string         `json:"status"`
	CategoryId int            `json:"category_id"`
	StoreId    int            `json:"store_id"`
	DeleteAt   gorm.DeletedAt `json:"-"`
}

type Category struct {
	ID   int
	Name string
}

func GetProducts() ([]Product, error) {
	var products []Product
	err := config.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(id int) (Product, error) {
	var product Product
	err := config.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func GetProductByCategory(categoryId int) ([]Product, error) {
	var products []Product
	err := config.DB.Where("category_id = ?", categoryId).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductCategories() ([]Category, error) {
	var categories []Category
	err := config.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (product *Product) CreateProduct() error {
	if product.Stock < 1 {
		product.Status = "soldout"
	}

	product.Status = "sale"
	return config.DB.Create(&product).Error
}

func (category *Category) CreateCategory() error {
	return config.DB.Create(&category).Error
}
