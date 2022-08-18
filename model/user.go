package model

type User struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Token    string `json:"token"`
}

func (User) TableName() string {
	return "user"
}
