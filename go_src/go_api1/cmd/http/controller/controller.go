package controller

import (
    // "go_api1/internal/service1"
    "go_api1/internal/service1"

    // "github.com/labstack/echo"
    "github.com/labstack/echo"
)

// Controller ...
type Controller struct{}

// HandleMessage ...
func (ctrl *Controller) HandleMessage(c echo.Context) error {
    msg := c.Param("message")
    if msg == "" {
        msg = "Hello, from http!"
    }

    arg := service1.AppCoreLogicIn{
        From:    "http",
        Message: msg,
    }

    service1.AppCoreLogic(c.Request().Context(), arg)
    return nil
}
