package callback

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
)

func get(
	c echo.Context,
	// toController controller.ToController,
	_ controller.ToController,
) (
	err error,
) {

	claims := c.Get("user").(jwt.MapClaims)
	username := claims["sub"].(string) // Example: "sub" from the token claims

	return c.String(http.StatusOK, fmt.Sprintf("Welcome to the protected endpoint, %s!", username))

	// ctx := pkg.GetNewContext(
	// 	c.Request().Context(),
	// 	c.Response().Header().Get(echo.HeaderXRequestID),
	// )

	// person := http_parameter.V1PersonParameter{}

	// if err := c.Bind(&person); err != nil {
	// 	pkg.Logging(ctx, err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		err,
	// 	)
	// }

	// reqPerson := struct_object.NewPerson(
	// 	&struct_object.NewPersonArgs{
	// 		ID:          person.ID,
	// 		Name:        person.Name,
	// 		MailAddress: person.MailAddress,
	// 	},
	// )

	// if reqPerson.Err != nil {
	// 	pkg.Logging(ctx, err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		err,
	// 	)
	// }

	// personList, err := toController.GetPersonByCondition(
	// 	ctx,
	// 	*reqPerson,
	// )

	// if err != nil {
	// 	pkg.Logging(ctx, err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		err,
	// 	)
	// }

	// responseList := []http_parameter.V1PersonParameter{}
	// for _, person := range personList {
	// 	responseList = append(
	// 		responseList,
	// 		http_parameter.V1PersonParameter{
	// 			ID:          &person.ID.Content.Value,
	// 			Name:        &person.Name.Content.Value,
	// 			MailAddress: &person.MailAddress.Content.Value,
	// 		},
	// 	)
	// }

	// return c.JSON(
	// 	http.StatusOK,
	// 	responseList,
	// )
}
