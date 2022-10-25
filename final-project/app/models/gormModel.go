package models

import (
	"time"
)

type GormModel struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
