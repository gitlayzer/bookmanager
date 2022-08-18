package model

type Book struct {
	Id    int64  `gorm:"primary_key" json:"id"`
	Name  string `json:"name" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
	Users []User `gorm:"many2many:book_users"`
}

func (Book) TableName() string {
	return "book"
}
