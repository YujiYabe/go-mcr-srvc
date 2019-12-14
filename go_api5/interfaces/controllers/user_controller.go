package controllers

import (
	"strconv"

	"app/domain"
	"app/interfaces/database"
	"app/usecase"

	"github.com/labstack/echo"
)

// UserController ...
type UserController struct {
	UCUserInteractor usecase.UCUserInteractor
}

// NewUserController ...
func NewUserController(SQLHandler database.IFDBSQLHandler) *UserController {
	return &UserController{
		UCUserInteractor: usecase.UCUserInteractor{
			UCUserRepository: &database.IFDBUserRepository{
				IFDBSQLHandler: SQLHandler,
			},
		},
	}
}

// IFCNShow ...
func (controller *UserController) IFCNShow(c echo.Context) (err error) {
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
func (controller *UserController) IFCNIndex(c echo.Context) (err error) {
	users, err := controller.UCUserInteractor.UCUIUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// IFCNCreate ...
func (controller *UserController) IFCNCreate(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.UCUserInteractor.UCUIAdd(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// IFCNSave ...
func (controller *UserController) IFCNSave(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.UCUserInteractor.UCUIUpdate(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// IFCNDelete ...
func (controller *UserController) IFCNDelete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		ID: id,
	}
	err = controller.UCUserInteractor.UCUIDeleteByID(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
