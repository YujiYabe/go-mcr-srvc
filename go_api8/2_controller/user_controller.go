package controller

import (
	"strconv"

	"github.com/labstack/echo"

	kitchen "app/3_kitchen"
)

// ControllerUser ...
type ControllerUser struct {
	UserKitchenHandle kitchen.UserKitchenHandle
}

// NewControllerUser ...
func NewControllerUser() *ControllerUser {
	return &ControllerUser{UserKitchenHandle: *kitchen.NewUserKitchenHandle()}
}

// IFCNIndex ...
func (controller *ControllerUser) Index(c echo.Context) (err error) {
	users, err := controller.UserKitchenHandle.UCUIUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// IFCNShow ...
func (controller *ControllerUser) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.UserKitchenHandle.UCUIUserByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// // IFCNCreate ...
// func (controller *ControllerUser) IFCNCreate(c echo.Context) (err error) {
// 	u := 1_entity.User{}
// 	c.Bind(&u)
// 	user, err := controller.UserKitchenHandle.UCUIAdd(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// // IFCNSave ...
// func (controller *ControllerUser) IFCNSave(c echo.Context) (err error) {
// 	u := 1_entity.User{}
// 	c.Bind(&u)

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	u.ID = id

// 	user, err := controller.UserKitchenHandle.UCUIUpdate(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// // IFCNDelete ...
// func (controller *ControllerUser) IFCNDelete(c echo.Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := 1_entity.User{
// 		ID: id,
// 	}
// 	err = controller.UserKitchenHandle.UCUIDeleteByID(user)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }
