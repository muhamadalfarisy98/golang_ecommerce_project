package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type saleOrderInput struct {
	Status          string `json:"status"`
	ShippingAddress string `json:"shipping_address"`
	UserID          uint   `json:"users_id"`
	PaymentID       uint   `json:"payment_id"`
}

// GetAllSaleOrder godoc
// @Summary Get all Sale Orders.
// @Description Get a list of saleorder.
// @Tags Sale Order
// @Produce json
// @Success 200 {object} []models.SaleOrder
// @Router /sale-order [get]
func GetAllSaleOrder(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var so []models.SaleOrder
	db.Find(&so)

	c.JSON(http.StatusOK, gin.H{"data": so})
}

// CreateSaleOrder godoc
// @Summary Create New Sale Orders.
// @Description Creating a new Sale orders.
// @Tags Sale Order
// @Param Body body saleOrderInput true "the body to create a new SaleOrder"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.SaleOrder
// @Router /sale-order [post]
func CreateSaleOrder(c *gin.Context) {
	// Validate input
	var input saleOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	invoice := "INV/"
	str_year := strconv.Itoa(year)
	invoice += str_year

	// Create SaleOrder
	db := c.MustGet("db").(*gorm.DB)
	var get_so []models.SaleOrder
	db.Find(&get_so)
	len_so := strconv.Itoa(len(get_so) + 1)
	invoice = invoice + "/" + len_so
	so := models.SaleOrder{Code: invoice,
		Status:          input.Status,
		OrderDate:       time.Now(),
		ShippingAddress: input.ShippingAddress,
		UserID:          input.UserID,
		PaymentID:       input.PaymentID}

	db.Create(&so)

	c.JSON(http.StatusOK, gin.H{"data": so})
}

// UpdateSaleOrder godoc
// @Summary Update Sale orders.
// @Description Update Sale Order by id.
// @Tags Sale Order
// @Produce json
// @Param id path string true "SaleOrder id"
// @Param Body body saleOrderInput true "the body to update Sale Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.SaleOrder
// @Router /sale-order/{id} [put]
func UpdateSaleOrder(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var so models.SaleOrder
	if err := db.Where("id = ?", c.Param("id")).First(&so).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input saleOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.SaleOrder
	// updatedInput.Code = input.Code
	updatedInput.ShippingAddress = input.ShippingAddress
	updatedInput.Status = input.Status
	// updatedInput.OrderDate = input.OrderDate
	updatedInput.UpdatedAt = time.Now()

	db.Model(&so).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": so})
}

// DeleteSaleOrder godoc
// @Summary Delete one Sale Orders.
// @Description Delete a SO by id.
// @Tags Sale Order
// @Produce json
// @Param id path string true "so id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
//@Router /sale-order/{id} [delete]
func DeleteSaleOrder(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var so models.SaleOrder
	if err := db.Where("id = ?", c.Param("id")).First(&so).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&so)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
