package controllers

import (
	"net/http"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productTypeInput struct {
	Name string `json:"type"`
}

// GetProductType godoc
// @Summary Get all Product Types.
// @Description Get a list of Product Type.
// @Tags Product Type
// @Produce json
// @Success 200 {object} []models.ProductType
// @Router /product-type [get]
func GetAllProductType(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var product_type []models.ProductType
	db.Find(&product_type)

	c.JSON(http.StatusOK, gin.H{"data": product_type})
}

// CreateProductType godoc
// @Summary Create New Product Types.
// @Description Creating a new Product Type.
// @Tags Product Type
// @Param Body body productTypeInput true "the body to create a new ProductType"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.ProductType
// @Router /product-type [post]
func CreateProductType(c *gin.Context) {
	// Validate input
	var input productTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create UoM
	uom := models.ProductType{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&uom)

	c.JSON(http.StatusOK, gin.H{"data": uom})
}

// UpdateProductType godoc
// @Summary Update Product Types.
// @Description Update Product Type by id.
// @Tags Product Type
// @Produce json
// @Param id path string true "ProductType id"
// @Param Body body productTypeInput true "the body to update ProductType"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.ProductType
// @Router /product-type/{id} [put]
func UpdateProductType(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product_type models.ProductType
	if err := db.Where("id = ?", c.Param("id")).First(&product_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ProductType
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product_type).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product_type})
}

// DeleteProductType godoc
// @Summary Delete one Product Types.
// @Description Delete a Product Type by id.
// @Tags Product Type
// @Produce json
// @Param id path string true "product type id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /product-type/{id} [delete]
func DeleteProductType(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product_type models.ProductType
	if err := db.Where("id = ?", c.Param("id")).First(&product_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product_type)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
