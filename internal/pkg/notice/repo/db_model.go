package repo

import "time"

// Notice struct definition
type Notice struct {
	ID        int64 `gorm:"primary_key"`
	Title     string
	Content   string    `gorm:"type:text"`
	UserTypes *string   `gorm:"type:varchar(20)"`
	Lang      string    `gorm:"type:varchar(10)"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
