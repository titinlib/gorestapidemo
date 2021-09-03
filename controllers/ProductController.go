package controllers

import (
	"encoding/json"
	"net/http"
	"restapidemo/apiHelpers"
	"restapidemo/models"
	"restapidemo/services"

	"github.com/gorilla/mux"
)

type Product models.Product

var productService = &services.ProductService{}

func GetAllProducts(rw http.ResponseWriter, req *http.Request) {
	allProducts, err := productService.GetAllProducts()

	if err != nil {
		apiHelpers.RespondWithError(rw, http.StatusOK, "Not found")
	}

	apiHelpers.RespondWithJSON(rw, http.StatusOK, allProducts)
}

func GetProductById(rw http.ResponseWriter, req *http.Request) {
	productId := mux.Vars(req)["id"]

	product, err := productService.GetProductById(productId)

	if err != nil {
		apiHelpers.RespondWithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	apiHelpers.RespondWithJSON(rw, http.StatusOK, product)
}

func CreateProduct(rw http.ResponseWriter, req *http.Request) {
	var newProductRequest models.Product

	json.NewDecoder(req.Body).Decode(&newProductRequest)

	defer req.Body.Close()

	err := productService.CreateProduct(&newProductRequest) // pass pointer of data to Create

	if err != nil {
		apiHelpers.RespondWithError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	apiHelpers.RespondWithJSON(rw, http.StatusCreated, newProductRequest)
}

func DeleteProductById(rw http.ResponseWriter, req *http.Request) {
	productId := mux.Vars(req)["id"]

	err := productService.DeleteProductById(productId)

	if err != nil {
		apiHelpers.RespondWithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	apiHelpers.RespondWithJSON(rw, http.StatusOK, "Deletion successful")
}

func UpdateProductId(rw http.ResponseWriter, req *http.Request) {
	productId := mux.Vars(req)["id"]

	var updateProductRequest models.Product
	json.NewDecoder(req.Body).Decode(&updateProductRequest)

	updatedProduct, err := productService.UpdateProductId(productId, updateProductRequest)

	if err != nil {
		apiHelpers.RespondWithError(rw, http.StatusBadRequest, err.Error())
		return
	}

	apiHelpers.RespondWithJSON(rw, http.StatusOK, updatedProduct)
}
