package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"size:100" bson:"name"`
	Email    string `json:"email" form:"email" bson:"email"`
	Password string `json:"password" form:"password" bson:"password"`
	Token    string `json:"token,omitempty" form:"token"`
}
