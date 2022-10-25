package models

import (
	"mygram/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" validate:"required"`
	Email    string `gorm:"type:varchar(50);uniqueIndex;not null" validate:"required"`
	Password string `gorm:"type:varchar(188);not null" validate:"required,min=6"`
	Age      int    `gorm:"not null" validate:"required,gt=8"`
}

type CreateUserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserResponse struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	UpdateAt time.Time `json:"update_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		return errCreate
	}

	u.Password, _ = utils.HashPassword(u.Password)
	return
}
