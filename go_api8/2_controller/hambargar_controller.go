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

// // IFCNCreate ...
// func (controller *HambargarController) IFCNCreate(c echo.Context) (err error) {
// 	u := 1_entity.Hambargar{}
// 	c.Bind(&u)
// 	hambargar, err := controller.HambargarKitchenHandle.UCUIAdd(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, hambargar)
// 	return
// }

// // IFCNSave ...
// func (controller *HambargarController) IFCNSave(c echo.Context) (err error) {
// 	u := 1_entity.Hambargar{}
// 	c.Bind(&u)

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	u.ID = id

// 	hambargar, err := controller.HambargarKitchenHandle.UCUIUpdate(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, hambargar)
// 	return
// }

// // IFCNDelete ...
// func (controller *HambargarController) IFCNDelete(c echo.Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	hambargar := 1_entity.Hambargar{
// 		ID: id,
// 	}
// 	err = controller.HambargarKitchenHandle.UCUIDeleteByID(hambargar)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, hambargar)
// 	return
// }
