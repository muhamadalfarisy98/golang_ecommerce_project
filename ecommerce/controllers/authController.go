package controllers

import (
	"ecommerce/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordInput struct {
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type dataInput struct {
	Name   string `json:"nama"`
	NoHP   string `json:"no_hp"`
	Alamat string `json:"alamat"`
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags User
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	// update token ke data current user
	var current_user models.User
	if err := db.Where("username = ?", u.Username).First(&current_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var updatedInput models.User
	updatedInput.RememberToken = token
	db.Model(&current_user).Updates(updatedInput)

	user := map[string]string{
		"username": u.Username,
		"email":    u.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// GetAllUser godoc
// @Summary Get all Users.
// @Description Get a list of User. (Only Admin Role 1)
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.User
// @Router /user [get]
func GetAllUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya Bisa dilakukan oleh admin atau Role = 1"})
		return
	}

	// Jika Admin Get All data
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserByUserName godoc
// @Summary Get a User.
// @Description Get a User by userName.
// @Tags User
// @Produce json
// @Param username path string true "username"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.User
// @Router /user/{username} [get]
func GetUserByUserName(c *gin.Context) { // Get model if exist
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	} else {
		bearerToken = ""
	}
	// Validasi Token
	if user.RememberToken != bearerToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bearer Token tidak sama dengan yang tersimpan didatabase, harap login akun anda terlebih dahulu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// ChangeUserData godoc
// @Summary Update Change Current User Data
// @Description Only user with the same token or only them can change their own Data
// @Tags User
// @Produce json
// @Param username path string true "Username"
// @Param Body body dataInput true "the body to update Data"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.User
// @Router /user/profil/{username} [put]
func ChangeUserData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var current_user models.User
	if err := db.Where("username = ?", c.Param("username")).First(&current_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
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
		// Validasi Token
		if current_user.RememberToken != bearerToken {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bearer Token tidak sama dengan yang tersimpan didatabase, harap login akun anda terlebih dahulu"})
			return
		}
	}

	// Validate input
	var input dataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Name = input.Name
	updatedInput.NoHP = input.NoHP
	updatedInput.Alamat = input.Alamat
	updatedInput.UpdatedAt = time.Now()

	db.Model(&current_user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": current_user})
}

// ChangeUserPassword godoc
// @Summary Update Change Current User Password
// @Description Only user with the same token or only them can change their own password
// @Tags User
// @Produce json
// @Param username path string true "Username"
// @Param Body body PasswordInput true "the body to update Password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.User
// @Router /user/{username} [put]
func ChangeUserPassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var current_user models.User
	if err := db.Where("username = ?", c.Param("username")).First(&current_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		bearerToken = strings.Split(bearerToken, " ")[1]
	} else {
		bearerToken = ""
	}
	// Validasi Token
	if current_user.RememberToken != bearerToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bearer Token tidak sama dengan yang tersimpan didatabase, harap login akun anda terlebih dahulu"})
		return
	}
	// Validate input
	var input PasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Password = input.Password
	updatedInput.UpdatedAt = time.Now()

	db.Model(&current_user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": current_user})
}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access. Default Role values is 0 (customer); for admin is Role = 1
// @Tags User
// @Param Body body RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	/*
		Role 0 = customer
		Role 1 = Admin
		default 0
	*/
	u.Role = 0

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
		"email":    input.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// DeleteUser godoc
// @Summary Delete one User.
// @Description Delete a User by username. (Admin Role Only)
// @Tags User
// @Produce json
// @Param username path string true "username"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /user/{username} [delete]
func DeleteUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var current_user models.User
	if err := db.Where("username = ?", c.Param("username")).First(&current_user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya Bisa dilakukan oleh admin atau Role = 1"})
		return
	}
	db.Delete(&current_user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
