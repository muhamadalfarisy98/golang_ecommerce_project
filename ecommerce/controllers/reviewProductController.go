package controllers

import (
	"net/http"
	"time"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewProductInput struct {
	Rating    int    `json:"rating"`
	Deskripsi string `json:"deskripsi"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
}

// GetAllReviewProduct godoc
// @Summary Get all Review Products.
// @Description Get a list of ReviewProduct.
// @Tags Review Product
// @Produce json
// @Success 200 {object} []models.ReviewProduct
// @Router /review-product [get]
func GetAllReviewProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var review_product []models.ReviewProduct
	db.Find(&review_product)

	c.JSON(http.StatusOK, gin.H{"data": review_product})
}

// GetReviewProductBySaleOrder godoc
// @Summary Get a list of ReviewProducts.
// @Description Get all ReviewProduct by product id
// @Tags Review Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} []models.ReviewProduct
// @Router /product/{id}/review-product [get]
func GetReviewProductByProductId(c *gin.Context) { // Get model if exist
	var rp []models.ReviewProduct

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("product_id = ?", c.Param("id")).Find(&rp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rp})
}

// GetReviewProductByUser godoc
// @Summary Get a list of ReviewProducts.
// @Description Get all ReviewProduct by user id
// @Tags User
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} []models.ReviewProduct
// @Router /users/{id}/review-product [get]
func GetReviewProductByUserId(c *gin.Context) { // Get model if exist
	var rp []models.ReviewProduct

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("users_id = ?", c.Param("id")).Find(&rp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rp})
}

// CreateReviewProduct godoc
// @Summary Create New Review Products.
// @Description Creating a new ReviewProducts.
// @Tags Review Product
// @Param Body body reviewProductInput true "the body to create a new ReviewProduct"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.ReviewProduct
// @Router /review-product [post]
func CreateReviewProduct(c *gin.Context) {
	// Validate input
	var input reviewProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Product
	db := c.MustGet("db").(*gorm.DB)

	review_product := models.ReviewProduct{Rating: input.Rating,
		Deskripsi: input.Deskripsi,
		UserID:    input.UserID,
		ProductID: input.ProductID,
	}

	db.Create(&review_product)

	c.JSON(http.StatusOK, gin.H{"data": review_product})
}

// UpdateReviewProduct godoc
// @Summary Update ReviewProducts.
// @Description Update ReviewProduct by id.
// @Tags Review Product
// @Produce json
// @Param id path string true "Review Product id"
// @Param Body body reviewProductInput true "the body to update Review Product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.ReviewProduct
// @Router /review-product/{id} [put]
func UpdateReviewProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var review_product models.ReviewProduct
	if err := db.Where("id = ?", c.Param("id")).First(&review_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input reviewProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ReviewProduct
	updatedInput.Rating = input.Rating
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.UserID = input.UserID
	updatedInput.ProductID = input.ProductID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&review_product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review_product})
}

// DeleteReviewProduct godoc
// @Summary Delete one ReviewProduct.
// @Description Delete a ReviewProduct by id.
// @Tags Review Product
// @Produce json
// @Param id path string true "reviewProduct id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /review-product/{id} [delete]
func DeleteReviewProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review_product models.ReviewProduct
	if err := db.Where("id = ?", c.Param("id")).First(&review_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review_product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
