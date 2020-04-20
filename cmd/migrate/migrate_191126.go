package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

var addDeviceTable = &gormigrate.Migration{
	ID: "2019-11-26 21:00:00",
	Migrate: func(tx *gorm.DB) error {
		type Device struct {
			ID        int64     `gorm:"primary_key"`
			AndroidID string    `gorm:"type:varchar(64);unique_index;not null"`
			FCMToken  string    `gorm:"type:varchar(512);not null"`
			Email     string    `gorm:"type:varchar(100);index;not null"`
			CreatedAt time.Time `gorm:"not null"`
			UpdatedAt time.Time `gorm:"not null"`
		}

		return tx.AutoMigrate(&Device{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("devices").Error
	},
}

var createTables = &gormigrate.Migration{
	ID: "2019-08-05 09:47:51",
	Migrate: func(tx *gorm.DB) error {

		type Notice struct {
			ID        int64     `gorm:"primary_key"`
			Title     string    `gorm:"not null"`
			Content   string    `gorm:"type:text;not null"`
			UserTypes *string   `gorm:"type:varchar(20)"`
			Lang      string    `gorm:"type:varchar(10);not null"`
			CreatedAt time.Time `gorm:"not null"`
			UpdatedAt time.Time `gorm:"not null"`
		}

		return tx.AutoMigrate(&Notice{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("videos", "notices", "device_preferences").Error
	},
}
