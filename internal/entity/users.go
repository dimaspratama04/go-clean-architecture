package entity

import "time"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MasterUsers struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username" validate:"required" require:"true"`
	Email     string    `gorm:"unique" json:"email" validate:"required" require:"true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	UserID int
	Token  string
}
