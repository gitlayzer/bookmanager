package model

type BookUser struct {
	UserID int64 `gorm:"primary_key"`
	BookID int64 `gorm:"primary_key"`
}
