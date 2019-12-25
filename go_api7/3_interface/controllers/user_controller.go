package controllers

import (
	"strconv"

	"github.com/labstack/echo"

	usecase "app/2_usecase"
	"app/3_interface/database"
)

// IFCNUserController ...
type IFCNUserController struct {
	UCUserInteractor usecase.UCUserInteractor
}

// NewUserController ...
func NewUserController(SQLHandler database.IFDBSQLHandler) *IFCNUserController {
	return &IFCNUserController{UCUserInteractor: *usecase.NewUCUserInteractor(SQLHandler)}
}

// IFCNShow ...
func (controller *IFCNUserController) IFCNShow(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.UCUserInteractor.UCUIUserByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// IFCNIndex ...
func (controller *IFCNUserController) IFCNIndex(c echo.Context) (err error) {
	users, err := controller.UCUserInteractor.UCUIUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// // IFCNCreate ...
// func (controller *IFCNUserController) IFCNCreate(c echo.Context) (err error) {
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
// func (controller *IFCNUserController) IFCNSave(c echo.Context) (err error) {
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
// func (controller *IFCNUserController) IFCNDelete(c echo.Context) (err error) {
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
