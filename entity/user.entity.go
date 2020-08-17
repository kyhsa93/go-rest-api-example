package entity

import "time"

// UserEntity entity for user data
type UserEntity struct {
	ID        string     `gorm:"unique; not null; primary_key; index;"`
	Email     string     `gorm:"unique; not null; index;"`
	Password  string     `gorm:"not null;"`
	CreatedAt time.Time  `gorm:"not null;"`
	UpdatedAt time.Time  `gorm:"not null;"`
	DeletedAt *time.Time `gorm:"index;"`
}
