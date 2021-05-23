package controllers

import (
	"app/domain"
	"app/interfaces/database"
	"app/usecase"
	"errors"
	"github.com/labstack/echo"
	"strconv"
)

// UserController ...
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController ...
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Create ...

func (controller *UserController) Create(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	err = controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, errors.New("err"))
		return
	}

	return c.JSON(201, nil)
}

// Index ...
func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, errors.New("err"))
		return
	}
	return c.JSON(200, users)
}

// Show ...
func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		return c.JSON(500, errors.New("err"))

	}
	return c.JSON(200, user)
}
