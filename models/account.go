package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	Id        uint       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"type:VARCHAR(256);NOT NULL" json:"name"`
	Avatar    string     `gorm:"type:VARCHAR(256);NOT NULL" json:"avatar"`
	Email     string     `gorm:"type:VARCHAR(256);NOT NULL" json:"email"`
	Phone     string     `gorm:"type:VARCHAR(20);NOT NULL" json:"phone"`
	Address   string     `gorm:"type:TEXT;NOT NULL" json:"address"`
	Password  string     `gorm:"type:VARCHAR(64);NOT NULL" json:"-"`
	CreatedAt time.Time  `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt time.Time  `gorm:"NOT NULL" json:"updated_at"`
	Galleries []Gallery  `json:"galleries,omitempty"`
	Photos    []Photo    `json:"photos,omitempty"`
	Reactions []Reaction `json:"reactions,omitempty"`
}
