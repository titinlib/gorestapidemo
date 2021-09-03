package services

import (
	"errors"
	"restapidemo/models"

	"gorm.io/gorm"
)

type Product models.Product

type ProductService struct{}

func (productService *ProductService) GetAllProducts() ([]Product, error) {
	var allProducts []Product

	db.Find(&allProducts)

	return allProducts, nil

}

func (productService *ProductService) GetProductById(productId string) (Product, error) {

	var product Product

	result := db.First(&product, productId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return product, errors.New("no records found")
	}

	return product, nil
}

func (productService *ProductService) CreateProduct(newProductRequest *models.Product) error {

	result := db.Create(newProductRequest) // pass pointer of data to Create

	if result.Error != nil {
		return errors.New("Product Creation failed")
	}
	return nil
}

func (productService *ProductService) DeleteProductById(productId string) error {

	// check for existence of product id
	var deleteProduct Product
	findResult := db.First(&deleteProduct, productId)

	if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("Product Id not Found")
	}

	db.Delete(&Product{}, productId)

	return nil
}

func (productService *ProductService) UpdateProductId(
	productId string,
	updateProductRequest models.Product,
) (Product, error) {

	// check for existence of product id
	var updateProduct Product
	findResult := db.First(&updateProduct, productId)

	if errors.Is(findResult.Error, gorm.ErrRecordNotFound) {
		return updateProduct, errors.New("Product Id not Found")
	}

	// making sure user can't change ID during update
	updateProductRequest.ID = 0

	updateResult := db.Model(&updateProduct).Updates(updateProductRequest)

	if updateResult.Error != nil {
		return updateProduct, errors.New("updation failed")
	}

	return updateProduct, nil
}
