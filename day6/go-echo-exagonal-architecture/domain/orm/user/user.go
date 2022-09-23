package user

import (
	"test-github/domain/entity/user"
	_g "test-github/domain/repository/gorm"
	"test-github/middleware"
	"test-github/util/requests"

	"github.com/jinzhu/gorm"
)

type UserOrm struct {
	db *gorm.DB
}

func (u *UserOrm) NewUserOrm() {
	g := _g.Gorm{}
	u.db = g.NewDb()[0]
}

func (u *UserOrm) Login(req *requests.UserLogin) (*user.User, error) {
	var (
		err  error
		user user.User
	)
	if err = u.db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
