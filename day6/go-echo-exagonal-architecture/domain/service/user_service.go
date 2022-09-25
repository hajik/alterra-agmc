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

func (u *UserService) ListAll(c echo.Context) error {
	req := requests.UserListAll{}
	c.Bind(&req)

	result, err := u.orm.ListAll(&req)
	if err != nil {
		return err
	}

	u.resp.Status = "Success"
	u.resp.Data = result

	return c.JSON(http.StatusOK, u.resp)
}

func (u *UserService) StoreOne(c echo.Context) error {
	req := requests.UserStoreOne{}
	c.Bind(&req)
	err := u.orm.StoreOne(&req)
	if err != nil {
		return err
	}

	u.resp.Status = "Success"
	u.resp.Data = req
	return c.JSON(http.StatusOK, u.resp)
}

func (u *UserService) Update(c echo.Context) error {
	req := requests.UserUpdate{}
	c.Bind(&req)
	err := u.orm.Update(&req)
	if err != nil {
		return err
	}
	u.resp.Status = "Success"
	u.resp.Data = req
	return c.JSON(http.StatusOK, u.resp)
}

func (u *UserService) Delete(c echo.Context) error {
	req := requests.UserDelete{}
	c.Bind(&req)
	err := u.orm.Delete(req.Name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Success")
}
