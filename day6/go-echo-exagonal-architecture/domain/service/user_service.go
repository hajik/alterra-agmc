package service

import (
	"net/http"
	"test-github/domain/orm/user"
	"test-github/util"
	"test-github/util/requests"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	orm  *user.UserOrm
	resp *util.Response
}

func (u *UserService) NewUserOrm() {
	u.resp = &util.Response{}
	u.orm = &user.UserOrm{}
	u.orm.NewUserOrm()
}

func (u *UserService) Login(c echo.Context) error {
	req := requests.UserLogin{}
	c.Bind(&req)

	user, err := u.orm.Login(&req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user.Token)
}
