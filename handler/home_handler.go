package handler

import (
	"example.com/Test/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	return render(c, home.Show())
}
