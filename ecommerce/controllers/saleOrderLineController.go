package controllers

import (
	"net/http"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type saleOrderLineInput struct {
	JumlahBarang int  `json:"jumlah_barang"`
	ProductID    uint `json:"product_id"`
	SaleOrderID  uint `json:"sale_order_id"`
}

// GetAllSaleOrderLine godoc
// @Summary Get all Sale order line.
// @Description Get a list of SaleOrderLine.
// @Tags Sale Order Line
// @Produce json
// @Success 200 {object} []models.SaleOrderLine
// @Router /sale-order-line [get]
func GetAllSaleOrderLine(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var sol []models.SaleOrderLine
	db.Find(&sol)

	c.JSON(http.StatusOK, gin.H{"data": sol})
}

// GetSaleOrderLineByProductId godoc
// @Summary Get a list of SaleOrderLines.
// @Description Get all SaleOrderLine by Product id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} []models.SaleOrderLine
// @Router /product/{id}/sale-order-line [get]
func GetSaleOrderLineByProductId(c *gin.Context) { // Get model if exist
	var sol []models.SaleOrderLine

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("product_id = ?", c.Param("id")).Find(&sol).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sol})
}

// GetSaleOrderLineBySaleOrder godoc
// @Summary Get a list of SaleOrderLines.
// @Description Get all SaleOrderLine by SaleOrder id
// @Tags Sale Order
// @Produce json
// @Param id path string true "SaleOrder id"
// @Success 200 {object} []models.SaleOrderLine
// @Router /sale-order/{id}/sale-order-line [get]
func GetSaleOrderLineBySaleOrderId(c *gin.Context) { // Get model if exist
	var so []models.SaleOrderLine

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("sale_order_id = ?", c.Param("id")).Find(&so).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": so})
}

// CreateSaleOrderLine godoc
// @Summary Create New Review SaleOrderLine.
// @Description Creating a new SaleOrderLine.
// @Tags Sale Order Line
// @Param Body body saleOrderLineInput true "the body to create a new SaleOrderLine"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.SaleOrderLine
// @Router /sale-order-line [post]
func CreateSaleOrderLine(c *gin.Context) {
	// Validate input
	var input saleOrderLineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Product
	db := c.MustGet("db").(*gorm.DB)

	// query product
	var get_product []models.Product
	db.Find(&get_product)
	jumlah_harga := 0
	for _, data := range get_product {
		if data.ID == input.ProductID {
			jumlah_harga = input.JumlahBarang * data.CostPrice
		}
	}

	sol := models.SaleOrderLine{JumlahBarang: input.JumlahBarang,
		JumlahHarga: jumlah_harga,
		ProductID:   input.ProductID,
		SaleOrderID: input.SaleOrderID,
	}

	db.Create(&sol)

	c.JSON(http.StatusOK, gin.H{"data": sol})
}

// UpdateSaleOrderLine godoc
// @Summary Update SaleOrderLines.
// @Description Update SaleOrderLine by id.
// @Tags Sale Order Line
// @Produce json
// @Param id path string true "SaleOrderLine id"
// @Param Body body saleOrderLineInput true "the body to update SaleOrderLine"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.SaleOrderLine
// @Router /sale-order-line/{id} [put]
func UpdateSaleOrderLine(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sol models.SaleOrderLine
	if err := db.Where("id = ?", c.Param("id")).First(&sol).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input saleOrderLineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// query product : untuk menghitung jumlah harga
	var get_product []models.Product
	db.Find(&get_product)
	jumlah_harga := 0
	for _, data := range get_product {
		if data.ID == input.ProductID {
			jumlah_harga = input.JumlahBarang * data.CostPrice
		}
	}
	var updatedInput models.SaleOrderLine
	updatedInput.JumlahBarang = input.JumlahBarang
	updatedInput.JumlahHarga = jumlah_harga
	updatedInput.SaleOrderID = input.SaleOrderID
	updatedInput.ProductID = input.ProductID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&sol).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": sol})
}

// DeleteSaleOrderLine godoc
// @Summary Delete one SaleOrderLine.
// @Description Delete a SaleOrderLine by id.
// @Tags Sale Order Line
// @Produce json
// @Param id path string true "SaleOrderLine id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /sale-order-line/{id} [delete]
func DeleteSaleOrderLine(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var sol models.SaleOrderLine
	if err := db.Where("id = ?", c.Param("id")).First(&sol).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&sol)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
