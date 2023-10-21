package controller

import (
	"github.com/jinzhu/gorm"
	"Mini-Challenge-Golang-6/model"
)

func CreateVariant(db *gorm.DB, variantName string, quantity uint, productID uint) (*model.Variant, error) {
	variant := model.Variant{VariantName: variantName, Quantity: quantity, ProductID: productID}
	err := db.Create(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func UpdateVariantById(db *gorm.DB, id uint, variantName string, quantity uint) (*model.Variant, error) {
	var variant model.Variant
	err := db.Where("id = ?", id).First(&variant).Error
	if err != nil {
		return nil, err
	}
	variant.VariantName = variantName
	variant.Quantity = quantity
	err = db.Save(&variant).Error
	if err != nil {
		return nil, err
	}
	return &variant, nil
}

func DeleteVariantById(db *gorm.DB, id uint) error {
	var variant model.Variant
	err := db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductWithVariant(db *gorm.DB, id uint) (*model.Product, error) {
	var product model.Product
	err := db.Preload("Variants").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}