package main

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"Mini-Challenge-Golang-6/controller"
	"Mini-Challenge-Golang-6/model"
	"os"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Database connection failed")
	}
	defer db.Close()

	// AutoMigrate models
	db.AutoMigrate(&model.Product{}, &model.Variant{})

	// Example of using controller functions
	product, err := controller.CreateProduct(db, "SABUN")
	if err != nil {
		log.Fatal("Error creating product:", err)
	}
	fmt.Println("Created Product:", product)

	// product, err = controller.UpdateProduct(db, product.ID, "SABUNn")
	// if err != nil {
	// 	log.Fatal("Error updating product:", err)
	// }
	// fmt.Println("Updated Product:", product)

	product, err = controller.GetProductById(db, product.ID)
	if err != nil {
		log.Fatal("Error fetching product:", err)
	}
	fmt.Println("Fetched Product:", product)

	variant, err := controller.CreateVariant(db, "SABUN MANDI", 10, product.ID)
	if err != nil {
		log.Fatal("Error creating variant:", err)
	}
	fmt.Println("Created Variant:", variant)

	// variant, err = controller.UpdateVariantById(db, variant.ID, "Updated Variant Name", 20)
	// if err != nil {
	// 	log.Fatal("Error updating variant:", err)
	// }
	// fmt.Println("Updated Variant:", variant)

	// err = controller.DeleteVariantById(db, variant.ID)
	// if err != nil {
	// 	log.Fatal("Error deleting variant:", err)
	// }
	// fmt.Println("Deleted Variant")

	productWithVariant, err := controller.GetProductWithVariant(db, product.ID)
	if err != nil {
		log.Fatal("Error fetching product with variants:", err)
	}
	fmt.Println("Fetched Product with Variants:", productWithVariant)
}
