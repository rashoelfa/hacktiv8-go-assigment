package CommentsController

import (
	"mygram/app/models"
	h "mygram/database"
	"mygram/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComments(c *gin.Context) {
	db := h.GetDB()
	contentType := utils.GetContentType(c)
	Comments := models.Comments{}
	user := models.User{}
	photo := models.Photo{}
	userData := c.MustGet("user").(jwt.MapClaims)

	userID := userData["id"].(float64)
	convertedUserId := uint(userID)

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}

	db.Where("id = ?", convertedUserId).First(&user)
	db.Where("id = ?", Comments.Photo_id).First(&photo)

	Comments.User_id = int(user.GormModel.ID)
	Comments.User = user
	Comments.Photo_id = int(photo.GormModel.ID)
	Comments.Photo = photo
	Comments.Photo.User = user

	if err := db.Create(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusCreated, models.CreateCommentsResponse{
		ID:        Comments.ID,
		Message:   Comments.Message,
		Photo_id:  Comments.Photo_id,
		User_id:   Comments.User_id,
		CreatedAt: Comments.CreatedAt,
	})
}

func GetAllComments(c *gin.Context) {
	db := h.GetDB()
	Comments := models.Comments{}
	CommentsQuery := []models.CommentsQuery{}
	CommentsResponse := []models.AllCommentsResponse{}

	if err := db.Model(&Comments).Select("Comments.*, photos.title, photos.caption, photos.photo_url, users.email, users.username").Joins("join users on users.id = Comments.user_id").Joins("join photos on photos.id = Comments.photo_id").Order("Comments.id").Scan(&CommentsQuery).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	for _, v := range CommentsQuery {
		CommentsResponse = append(CommentsResponse, models.AllCommentsResponse{
			ID:        v.ID,
			Message:   v.Message,
			Photo_id:  v.Photo_id,
			User_id:   v.User_id,
			UpdatedAt: v.UpdatedAt,
			CreatedAt: v.CreatedAt,
			User: struct {
				User_id  int    `json:"id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				User_id:  v.User_id,
				Email:    v.Email,
				Username: v.Username,
			},
			Photo: struct {
				Photo_id  int    `json:"id"`
				Title     string `json:"title"`
				Caption   string `json:"caption"`
				Photo_url string `json:"photo_url"`
				User_id   int    `json:"user_id"`
			}{
				Photo_id:  v.Photo_id,
				Title:     v.Title,
				Caption:   v.Caption,
				Photo_url: v.Photo_url,
				User_id:   v.User_id,
			},
		})
	}

	c.JSON(http.StatusOK, CommentsResponse)
}

func UpdateComments(c *gin.Context) {
	db := h.GetDB()
	id := c.Param("id")
	contentType := utils.GetContentType(c)
	Comments := models.Comments{}
	updateComments := models.UpdateCommentsResponse{}

	if err := db.Where("id = ?", id).First(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Comments not found"})
		return
	}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}

	if err := db.Save(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	if err := db.Model(&Comments).Select("Comments.*, photos.title, photos.caption, photos.photo_url").Joins("join photos on photos.id = Comments.photo_id").Order("Comments.id").Scan(&updateComments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, updateComments)
}

func DeleteComments(c *gin.Context) {
	db := h.GetDB()
	id := c.Param("id")
	Comments := models.Comments{}

	if err := db.Where("id = ?", id).First(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Comments not found"})
		return
	}

	if err := db.Delete(&Comments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ValidatorErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "your Comments has been successfully deleted"})
}
