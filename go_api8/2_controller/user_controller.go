package controller

import (
	"strconv"

	"github.com/labstack/echo"

	kitchen "app/3_kitchen"
)

// ControllerUser ...
type ControllerUser struct {
	UCUserInteractor kitchen.UCUserInteractor
}

// NewControllerUser ...
func NewControllerUser() *ControllerUser {
	return &ControllerUser{UCUserInteractor: *kitchen.NewUCUserInteractor()}
}

// IFCNIndex ...
func (controller *ControllerUser) Index(c echo.Context) (err error) {
	users, err := controller.UCUserInteractor.UCUIUsers()
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
	user, err := controller.UCUserInteractor.UCUIUserByID(id)
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
// 	user, err := controller.UCUserInteractor.UCUIAdd(u)
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

// 	user, err := controller.UCUserInteractor.UCUIUpdate(u)
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
// 	err = controller.UCUserInteractor.UCUIDeleteByID(user)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }
