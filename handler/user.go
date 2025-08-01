package handler

import (
	"example.com/Test/model"
	"example.com/Test/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "a@ggyayayyayaya.com",
	}
	return render(c, user.Show(u))
}
