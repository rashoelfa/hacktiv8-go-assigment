package usersController

import (
	"mygram/app/models"
	h "mygram/database"
	"mygram/utils"
	"net/http"
	"net/mail"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	user := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	_, errEmail := mail.ParseAddress(user.Email)
	if errEmail != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"email": "Invalid email"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusCreated, models.CreateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		CreatedAt: user.GormModel.CreatedAt,
	})
}

func UserLogin(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	user := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password := user.Password

	if err := db.Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "email or username not found"})
		return
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"password": "Invalid password"})
		return
	}

	token, _ := utils.GenerateNewAccessToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": "jwt " + token,
	})
}

func UpdateUser(c *gin.Context) {
	db := h.GetDB()
	userData := c.MustGet("user").(jwt.MapClaims)
	contentType := utils.GetContentType(c)
	user := models.User{}

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	if err := db.Where("id = ?", convertedUserId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "email or username not found"})
		return
	}

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	_, errEmail := mail.ParseAddress(user.Email)
	if errEmail != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"email": "Invalid email"})
		return
	}

	if err := db.Where("id = ?", convertedUserId).Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, models.UpdateUserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Age:      user.Age,
		UpdateAt: user.GormModel.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	db := h.GetDB()
	userData := c.MustGet("user").(jwt.MapClaims)
	user := models.User{}

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	if err := db.Delete(&user, convertedUserId).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
