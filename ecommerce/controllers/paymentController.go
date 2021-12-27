package controllers

import (
	"net/http"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/swag/example/celler/httputil"
	"gorm.io/gorm"
)

type paymentInput struct {
	Name string `json:"provider"`
}

// GetPayment godoc
// @Summary Get all payments.
// @Description Get a list of payment.
// @Tags Payment
// @Produce json
// @Success 200 {object} []models.Payment
// @Router /payment [get]
func GetAllProvider(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var payment []models.Payment
	db.Find(&payment)
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// CreatePayment godoc
// @Summary Create New Payments.
// @Description Creating a new Payment.
// @Tags Payment
// @Param Body body paymentInput true "the body to create a new Payment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Payment
// @Router /payment [post]
func CreateProvider(c *gin.Context) {
	// Validate input
	var input paymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create UoM
	payment := models.Payment{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&payment)

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// UpdatePayment godoc
// @Summary Update Payments.
// @Description Update Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Param Body body paymentInput true "the body to update Payment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Payment
// @Router /payment/{id} [put]
func UpdateProvider(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var payment models.Payment
	if err := db.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input paymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Payment
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&payment).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DeletePayment godoc
// @Summary Delete one Payments.
// @Description Delete a Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "payment id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /payment/{id} [delete]
func DeleteProvider(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var payment models.Payment
	if err := db.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&payment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
