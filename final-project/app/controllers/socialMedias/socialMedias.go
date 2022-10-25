package socialMediasController

import (
	"mygram/app/models"
	h "mygram/database"
	"mygram/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	SocialMedia := models.SocialMedias{}
	user := models.User{}
	userData := c.MustGet("user").(jwt.MapClaims)

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	db.Where("id = ?", convertedUserId).First(&user)

	SocialMedia.User_id = int(user.GormModel.ID)
	SocialMedia.User = user

	if err := db.Create(&SocialMedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusCreated, models.CreateSocialMediaResponse{
		ID:               SocialMedia.ID,
		Name:             SocialMedia.Name,
		Social_Media_Url: SocialMedia.Social_Media_Url,
		User_id:          SocialMedia.User_id,
		CreatedAt:        SocialMedia.GormModel.CreatedAt,
	})
}

func GetAllSocialMedias(c *gin.Context) {
	db := h.GetDB()
	socialMedias := models.SocialMedias{}
	socialMediaQuery := []models.SocialMediaQuery{}
	socialMediasResponse := []models.AllSocialMediasResponse{}

	if err := db.Model(&socialMedias).Select("social_medias.*, users.username").Joins("join users on users.id = social_medias.user_id").Order("social_medias.id").Scan(&socialMediaQuery).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	for _, socialMedias := range socialMediaQuery {
		socialMediasResponse = append(socialMediasResponse, models.AllSocialMediasResponse{
			ID:               socialMedias.ID,
			Name:             socialMedias.Name,
			Social_Media_Url: socialMedias.Social_Media_Url,
			User_id:          socialMedias.User_id,
			CreatedAt:        socialMedias.CreatedAt,
			UpdatedAt:        socialMedias.UpdatedAt,
			User: struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
			}{
				ID:       socialMedias.User_id,
				Username: socialMedias.Username,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": socialMediasResponse,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	id := c.Param("id")
	SocialMedia := models.SocialMedias{}
	SocialMediaQuery := models.SocialMediaQuery{}

	if err := db.Where("id = ?", id).First(&SocialMedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "social media not found"})
		return
	}

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	if err := db.Model(&SocialMedia).Where("id = ?", SocialMedia.ID).Save(&SocialMedia).Scan(&SocialMediaQuery).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, models.UpdateSocialMediaResponse{
		ID:               SocialMediaQuery.ID,
		Name:             SocialMediaQuery.Name,
		Social_Media_Url: SocialMediaQuery.Social_Media_Url,
		User_id:          SocialMediaQuery.User_id,
		UpdatedAt:        SocialMediaQuery.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := h.GetDB()
	SocialMedia := models.SocialMedias{}
	user := models.User{}
	userData := c.MustGet("user").(jwt.MapClaims)

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	db.Where("id = ?", convertedUserId).First(&user)

	SocialMedia.User_id = int(user.GormModel.ID)
	SocialMedia.User = user

	if err := db.Where("id = ?", c.Param("id")).Delete(&SocialMedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
