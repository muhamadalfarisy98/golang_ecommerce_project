package controllers

import (
	"net/http"
	"strings"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	Name          string `json:"nama"`
	CostPrice     int    `json:"cost_price"`
	QtyAvailable  int    `json:"qty_available"`
	QtyForecasted int    `json:"qty_forecasted"`
	ImageUrl      string `json:"image_url"`
	Keterangan    string `json:"keterangan"`
	UomID         uint   `json:"uom_id"`
	ProductTypeID uint   `json:"product_type_id"`
}

// GetAllProduct godoc
// @Summary Get all Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var product []models.Product
	db.Find(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductById godoc
// @Summary Get a Product.
// @Description Get a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProduct godoc
// @Summary Create New Products.
// @Description Creating a new Products. (Admin Only)
// @Tags Product
// @Param Body body productInput true "the body to create a new Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Product
	db := c.MustGet("db").(*gorm.DB)

	product := models.Product{Name: input.Name,
		CostPrice:       input.CostPrice,
		QtyAvailable:    input.QtyAvailable,
		QtyForecasted:   input.QtyForecasted,
		ImageUrl:        input.ImageUrl,
		Keterangan:      input.Keterangan,
		UnitOfMeasureID: input.UomID,
		ProductTypeID:   input.ProductTypeID,
	}

	// Cek Hanya Admin
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	} else {
		bearerToken = ""
	}
	// cek user from bearer token (apakah login sebagai admin?)
	var bearer_user models.User
	if err := db.Where("remember_token = ?", bearerToken).First(&bearer_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if bearer_user.Role != 1 {
		// jika bukan admin
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya Bisa dilakukan oleh admin atau User.Role = 1"})
		return
	}

	// jika admin
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update Products.
// @Description Update Product by id. (Admin Only)
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Product
// @Router /product/{id} [put]
func UpdateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Cek Hanya Admin
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	} else {
		bearerToken = ""
	}
	// cek user from bearer token (apakah login sebagai admin?)
	var bearer_user models.User
	if err := db.Where("remember_token = ?", bearerToken).First(&bearer_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if bearer_user.Role != 1 {
		// jika bukan admin
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya Bisa dilakukan oleh admin atau User.Role = 1"})
		return
	}

	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Product
	updatedInput.Name = input.Name
	updatedInput.CostPrice = input.CostPrice
	updatedInput.QtyAvailable = input.QtyAvailable
	updatedInput.QtyForecasted = input.QtyForecasted
	updatedInput.ImageUrl = input.ImageUrl
	updatedInput.Keterangan = input.Keterangan
	updatedInput.UnitOfMeasureID = input.UomID
	updatedInput.ProductTypeID = input.ProductTypeID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id. (Admin Only)
// @Tags Product
// @Produce json
// @Param id path string true "so id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Cek Hanya Admin
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	} else {
		bearerToken = ""
	}
	// cek user from bearer token (apakah login sebagai admin?)
	var bearer_user models.User
	if err := db.Where("remember_token = ?", bearerToken).First(&bearer_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if bearer_user.Role != 1 {
		// jika bukan admin
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya Bisa dilakukan oleh admin atau User.Role = 1"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
