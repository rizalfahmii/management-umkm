package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stocklog struct {
	gorm.Model
	ID        uint
	ProductId int
	Product   Product `json:"product"`
	Type      string  `gorm:"type:enum('IN','OUT');not null" json:"type"`
	Qty       int
	Note      string
	CreatedAt time.Time
}
