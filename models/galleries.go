package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Gallery struct {
	Id         uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountId  uint      `gorm:"NOT NULL" json:"account_id"`
	Name       string    `gorm:"type:TEXT;NOT NULL" json:"name"`
	Brief      string    `gorm:"type:TEXT;NOT NULL" json:"brief"`
	Visibility bool      `gorm:"NOT NULL" json:"visibility"`
	CreatedAt  time.Time `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt  time.Time `gorm:"NOT NULL" json:"updated_at"`
	Photos     []Photo   `json:"photos,omitempty"`
}
