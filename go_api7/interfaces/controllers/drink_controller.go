package controllers

import (
	"strconv"

	"github.com/labstack/echo"

	"app/interfaces/database"
	"app/recipe"
)

// DrinkController ...
type DrinkController struct {
	CookDrink recipe.CookDrink
}

// NewDrinkController ...
func NewDrinkController(SQLHandler database.IFDBSQLHandler) *DrinkController {
	return &DrinkController{
		CookDrink: recipe.CookDrink{
			PrepareDrink: &database.IFDBUserRepository{
				IFDBSQLHandler: SQLHandler,
			},
		},
	}
}

// ShowAllDrinks ...
func (controller *DrinkController) ShowAllDrinks(c echo.Context) (err error) {
	users, err := controller.CookDrink.UCUIUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// ShowDetailDrink ...
func (controller *DrinkController) ShowDetailDrink(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.CookDrink.UCUIUserByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// // IFCNCreate ...
// func (controller *DrinkController) IFCNCreate(c echo.Context) (err error) {
// 	u := domain.User{}
// 	c.Bind(&u)
// 	user, err := controller.CookDrink.UCUIAdd(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// // IFCNSave ...
// func (controller *DrinkController) IFCNSave(c echo.Context) (err error) {
// 	u := domain.User{}
// 	c.Bind(&u)

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	u.ID = id

// 	user, err := controller.CookDrink.UCUIUpdate(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// // IFCNDelete ...
// func (controller *DrinkController) IFCNDelete(c echo.Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := domain.User{
// 		ID: id,
// 	}
// 	err = controller.CookDrink.UCUIDeleteByID(user)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }
