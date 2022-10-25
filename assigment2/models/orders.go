package models

import (
	"time"
)

type Items struct {
	Item_id     uint   `gorm:"primaryKey" json:"itemId"`
	Item_code   string `gorm:"not null;type:varchar(191)" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(191)" json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint   `gorm:"foreignKey:OrderRefer"`
}

type Orders struct {
	Order_id      uint      `gorm:"primaryKey" json:"orderId"`
	Customer_name string    `gorm:"not null;type:varchar(191)" json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
}
