package controller

import (
	"github.com/labstack/echo"

	kitchen "app/3_kitchen"
)

// HambargarController ...
type HambargarController struct {
	HambargarKitchenHandle kitchen.HambargarKitchenHandle
}

// NewHambargarController ...
func NewHambargarController() *HambargarController {
	return &HambargarController{HambargarKitchenHandle: *kitchen.NewHambargarKitchenHandle()}
}

// Index ...
func (controller *HambargarController) Index(c echo.Context) (err error) {
	hambargars, err := controller.HambargarKitchenHandle.FindAllHambargars()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, hambargars)
	return
}

// Show ...
func (controller *HambargarController) Show(c echo.Context) (err error) {
	// name, _ := strconv.Atoi(c.Param("name"))
	name := c.Param("name")

	hambargar, err := controller.HambargarKitchenHandle.FindHambargarByName(name)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, hambargar)
	return
}

// Request ...
func (controller *HambargarController) Request(c echo.Context) (err error) {
	name := c.Param("name")

	hambargar, err := controller.HambargarKitchenHandle.Cook(name)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, hambargar)
	return
}
