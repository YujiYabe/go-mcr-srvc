package controller

import (
	"strconv"

	"github.com/labstack/echo"

	kitchen "app/3_kitchen"
)

// UserController ...
type UserController struct {
	UserKitchenHandle kitchen.UserKitchenHandle
}

// NewUserController ...
func NewUserController() *UserController {
	return &UserController{UserKitchenHandle: *kitchen.NewUserKitchenHandle()}
}

// Index ...
func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.UserKitchenHandle.FindAllUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// Show ...
func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.UserKitchenHandle.FindUserByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// // IFCNCreate ...
// func (controller *UserController) IFCNCreate(c echo.Context) (err error) {
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
// func (controller *UserController) IFCNSave(c echo.Context) (err error) {
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
// func (controller *UserController) IFCNDelete(c echo.Context) (err error) {
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
