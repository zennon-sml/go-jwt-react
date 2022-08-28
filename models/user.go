package models

type User struct {
	ID       uint   `json:"id" form:"id" gorm:"primarykey:auto_increment"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"passowrd" form:"password"`
}
