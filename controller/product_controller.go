package controller

import (
	"github.com/jinzhu/gorm"
	"Mini-Challenge-Golang-6/model"
)

func CreateProduct(db *gorm.DB, name string) (*model.Product, error) {
	product := model.Product{Name: name}
	err := db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(db *gorm.DB, id uint, name string) (*model.Product, error) {
	var product model.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	product.Name = name
	err = db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProductById(db *gorm.DB, id uint) (*model.Product, error) {
	var product model.Product
	err := db.Preload("Variants").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}