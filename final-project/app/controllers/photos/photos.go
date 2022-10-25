package photosController

import (
	"mygram/app/models"
	h "mygram/database"
	"mygram/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	photo := models.Photo{}
	user := models.User{}
	userData := c.MustGet("user").(jwt.MapClaims)

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	if contentType == "application/json" {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	db.Where("id = ?", convertedUserId).First(&user)

	photo.User_id = int(user.GormModel.ID)
	photo.User = user

	if err := db.Create(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusCreated, models.CreatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Photo_Url: photo.Photo_Url,
		User_id:   photo.User_id,
		CreatedAt: photo.CreatedAt,
	})
}

func GetAllPhotos(c *gin.Context) {
	db := h.GetDB()
	photo := models.Photo{}
	photoQuery := []models.PhotoQuery{}
	photoResponse := []models.AllPhotoResponse{}

	if err := db.Model(&photo).Select("photos.*, users.email, users.username").Joins("join users on users.id = photos.user_id").Order("photos.id").Scan(&photoQuery).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	for _, v := range photoQuery {
		photoResponse = append(photoResponse, models.AllPhotoResponse{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			Photo_Url: v.Photo_Url,
			User_id:   v.User_id,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				Email:    v.Email,
				Username: v.Username,
			},
		})
	}

	c.JSON(http.StatusOK, photoResponse)
}

func UpdatePhoto(c *gin.Context) {
	db := h.GetDB()
	id := c.Param("id")
	contentType := utils.GetContentType(c)
	photo := models.Photo{}

	if err := db.Where("id = ?", id).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	if contentType == "application/json" {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	if err := db.Save(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, models.UpdatePhotoResponse{
		ID:         photo.ID,
		Title:      photo.Title,
		Caption:    photo.Caption,
		Photo_Url:  photo.Photo_Url,
		User_id:    photo.User_id,
		Updated_at: photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := h.GetDB()
	id := c.Param("id")
	photo := models.Photo{}

	if err := db.Where("id = ?", id).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	if err := db.Delete(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})
}
