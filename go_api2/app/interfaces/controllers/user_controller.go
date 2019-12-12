package controllers

import (
	"app/domain"
	"app/interfaces/database"
	"app/usecase"
	"errors"
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
func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, errors.New("err"))
		return
	}
	c.JSON(201, nil)
}

// Index ...
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, errors.New("err"))
		return
	}
	c.JSON(200, users)
}

// Show ...
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, errors.New("err"))
		return
	}
	c.JSON(200, user)
}
