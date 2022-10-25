package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	GormModel
	Title     string `gorm:"type:varchar(188);uniqueIndex;not null" validate:"required"`
	Caption   string `gorm:"type:varchar(188)"`
	Photo_Url string `gorm:"type:varchar(188);uniqueIndex;not null" validate:"required"`
	User_id   int
	User      User `gorm:"foreignKey:User_id"`
}

type PhotoQuery struct {
	ID        uint
	Title     string
	Caption   string
	Photo_Url string
	User_id   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Username  string
}

type CreatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_Url string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotoResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_Url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type AllPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_Url string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		return errCreate
	}
	return
}
