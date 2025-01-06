package person

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/2_adapter/controller"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

func viaGRPC(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := c.Request().Context()

	// person := http_parameter.V1Person{}

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
	// 	pkg.Logging(ctx, reqPerson.Err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		reqPerson.Err,
	// 	)
	// }

	// personList, err := toController.GetPersonByCondition(
	// 	ctx,
	// 	*reqPerson,
	// )

	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)
	log.Println("== == == == == == == == == == ")

	err = toController.ViaGRPC(
		ctx,
	)

	if err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	// responseList := []http_parameter.V1Person{}
	// for _, person := range personList {
	// 	id := person.ID.Content.GetValue()
	// 	name := person.Name.Content.GetValue()
	// 	mailAddress := person.MailAddress.Content.GetValue()
	// 	responseList = append(
	// 		responseList,
	// 		http_parameter.V1Person{
	// 			ID:          &id,
	// 			Name:        &name,
	// 			MailAddress: &mailAddress,
	// 		},
	// 	)
	// }

	return c.JSON(
		http.StatusOK,
		nil,
	)

}
