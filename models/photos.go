package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Photo struct {
	Id            uint       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	AccountId     uint       `gorm:"NOT NULL" json:"account_id"`
	GalleryId     uint       `gorm:"NOT NULL" json:"gallery_id"`
	Name          string     `gorm:"type:TEXT;NOT NULL" json:"name"`
	Description   string     `gorm:"type:TEXT;NOT NULL" json:"desciption"`
	Path          string     `gorm:"type:TEXT;NOT NULL" json:"path"`
	W1920Path     string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w1920_path"`
	W1600Path     string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w1600_path"`
	W1280Path     string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w1280_path"`
	W1024Path     string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w1024_path"`
	W800Path      string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w800_path"`
	W256Path      string     `gorm:"type:VARCHAR(256);NOT NULL" json:"w256_path"`
	Size          int64      `gorm:"type:DOUBLE;NOT NULL" json:"size"`
	CreatedAt     time.Time  `gorm:"NOT NULL" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"NOT NULL" json:"updated_at"`
	Reactions     []Reaction `json:"reactions,omitempty"`
	ReactionCount int        `gorm:"-"`
}
