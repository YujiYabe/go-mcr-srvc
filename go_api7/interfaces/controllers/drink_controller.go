package controllers

import (
	"github.com/labstack/echo"

	"app/interfaces/supplier"
	"app/recipe"
)

// DrinkController ...
type DrinkController struct {
	CookDrink recipe.CookDrink
}

// NewDrinkController ...
func NewDrinkController(DrinkStocker supplier.ExtractDrink) *DrinkController {
	return &DrinkController{
		CookDrink: recipe.CookDrink{
			PrepareDrink: &supplier.SupplyDrink{
				ExtractDrink: DrinkStocker,
			},
		},
	}
}

// RequestCoffee ...
func (controller *DrinkController) RequestCoffee(c echo.Context) (err error) {
	users, err := controller.CookDrink.CookCoffee()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// // ShowDetailDrink ...
// func (controller *DrinkController) ShowDetailDrink(c echo.Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user, err := controller.CookDrink.UCUIUserByID(id)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }

// // IFCNCreate ...
// func (controller *DrinkController) IFCNCreate(c echo.Context) (err error) {
// 	u := menu.User{}
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
// 	u := menu.User{}
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
// 	user := menu.User{
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
