package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Comments struct {
	GormModel
	User_id  int
	User     User `gorm:"foreignKey:User_id"`
	Photo_id int
	Photo    Photo  `gorm:"foreignKey:Photo_id"`
	Message  string `gorm:"type:varchar(188);uniqueIndex;not null" validate:"required"`
}

type CommentsQuery struct {
	ID            uint
	Message       string
	Photo_id      int
	User_id       int
	UpdatedAt     time.Time
	CreatedAt     time.Time
	Email         string
	Username      string
	Title         string
	Caption       string
	Photo_url     string
}

type CommentsCreateRequest struct {
	Message  string `json:"message" validate:"required"`
	Photo_id int    `json:"photo_id" validate:"required"`
}

type CreateCommentsResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AllCommentsResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		User_id  int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		Photo_id  int    `json:"id"`
		Title     string `json:"title"`
		Caption   string `json:"caption"`
		Photo_url string `json:"photo_url"`
		User_id   int    `json:"user_id"`
	} `json:"photo"`
}

type UpdateCommentsResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Comments) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		return errCreate
	}
	return
}
