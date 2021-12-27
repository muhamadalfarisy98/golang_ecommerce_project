package controllers

import (
	"net/http"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type uomInput struct {
	Name string `json:"nama_unit"`
}

// GetAllUom godoc
// @Summary Get all Unit of measures.
// @Description Get a list of Unit of measures.
// @Tags UoM
// @Produce json
// @Success 200 {object} []models.UnitOfMeasure
// @Router /uom [get]
func GetAllUom(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var uom []models.UnitOfMeasure
	db.Find(&uom)

	c.JSON(http.StatusOK, gin.H{"data": uom})
}

// CreateUom godoc
// @Summary Create New Unit of measures.
// @Description Creating a new Unit of measures.
// @Tags UoM
// @Param Body body uomInput true "the body to create a new UnitOfMeasure"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.UnitOfMeasure
// @Router /uom [post]
func CreateUom(c *gin.Context) {
	// Validate input
	var input uomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create UoM
	uom := models.UnitOfMeasure{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&uom)

	c.JSON(http.StatusOK, gin.H{"data": uom})
}

// UpdateUoM godoc
// @Summary Update Unit of measures.
// @Description Update Unit of measures by id.
// @Tags UoM
// @Produce json
// @Param id path string true "UnitOfMeasure id"
// @Param Body body uomInput true "the body to update Unit of Measure"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.UnitOfMeasure
// @Router /uom/{id} [put]
func UpdateUom(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var uom models.UnitOfMeasure
	if err := db.Where("id = ?", c.Param("id")).First(&uom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input uomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.UnitOfMeasure
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&uom).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": uom})
}

// DeleteUom godoc
// @Summary Delete one Unit of Measures.
// @Description Delete a UoM by id.
// @Tags UoM
// @Produce json
// @Param id path string true "uom id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /uom/{id} [delete]
func DeleteUom(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var uom models.UnitOfMeasure
	if err := db.Where("id = ?", c.Param("id")).First(&uom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&uom)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
