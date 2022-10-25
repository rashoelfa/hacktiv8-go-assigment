package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedias struct {
	GormModel
	Name             string `gorm:"type:varchar(188);uniqueIndex;not null" validate:"required"`
	Social_Media_Url string `gorm:"type:varchar(188);uniqueIndex;not null" validate:"required"`
	User_id          int
	User             User `gorm:"foreignKey:User_id"`
}

type SocialMediaQuery struct {
	ID               uint
	Name             string
	Social_Media_Url string
	User_id          int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Username         string
}

type CreateSocialMediaResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

type AllSocialMediasResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	User             struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

type UpdateSocialMediaResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (u *SocialMedias) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		return errCreate
	}
	return
}
