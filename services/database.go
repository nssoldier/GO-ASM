package services

import (
	"gallery/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB(connection string) (err error) {

	// Create connection
	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		return
	}

	err = DB.AutoMigrate(
		&models.Account{},
		&models.Gallery{},
		&models.Photo{},
		&models.Reaction{},
	).Error
	if err != nil {
		return
	}
	DB.Model(&models.Gallery{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Photo{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Photo{}).AddForeignKey(
		"gallery_id", "galleries(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Reaction{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Reaction{}).AddForeignKey(
		"photo_id", "photos(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Reaction{}).AddUniqueIndex("idx_acc_photo", "account_id", "photo_id")
	return
}
