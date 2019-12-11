package controllers

import (
    "app/domain"
    "app/interfaces/database"
    "app/usecase"
    "strconv"
    "errors"
)


type UserController stpackage controllers

import (
    "app/domain"
    "app/interfaces/database"
    "app/usecase"
    "strconv"
    "errors"
)


type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Create(c echo.Context) (err error) {
    u := domain.User{}
    c.Bind(&u)
    user, err := controller.Interactor.Add(u)
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(201, user)
    return
}


func (controller *UserController) Index(c echo.Context) {
    users, err := controller.Interactor.Users()
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(200, users)
}

func (controller *UserController) Show(c echo.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := controller.Interactor.UserById(id)
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(200, user)
}


func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Create(c echo.Context) {
    u := domain.User{}
    c.Bind(&u)
    err := controller.Interactor.Add(u)
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(int(201),nil)
}

func (controller *UserController) Index(c echo.Context) {
    users, err := controller.Interactor.Users()
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(200, users)
}

func (controller *UserController) Show(c echo.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := controller.Interactor.UserById(id)
    if err != nil {
        c.JSON(500, errors.New("err"))
        return
    }
    c.JSON(200, user)
}
