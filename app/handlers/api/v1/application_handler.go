package v1

import (
	"github.com/anurag925/identity/app/views"
	"github.com/labstack/echo/v4"
)

type applicationHandler struct {
}

func NewApplicationHandler() *applicationHandler {
	return &applicationHandler{}
}

func (h *applicationHandler) Index(c echo.Context) error {
	return c.Render(200, "hello_word", views.Hello("word"))
}
