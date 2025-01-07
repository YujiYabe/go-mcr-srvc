package person

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	groupObject "backend/internal/4_domain/group_object"
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

	person := httpParameter.V1Person{}

	if err := c.Bind(&person); err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	reqPerson := groupObject.NewPerson(
		ctx,
		&groupObject.NewPersonArgs{
			ID:          person.ID,
			Name:        person.Name,
			MailAddress: person.MailAddress,
		},
	)

	if reqPerson.GetError() != nil {
		pkg.Logging(ctx, reqPerson.GetError())
		return c.JSON(
			http.StatusBadRequest,
			reqPerson.GetError(),
		)
	}

	// personList, err := toController.GetPersonByCondition(
	// 	ctx,
	// 	*reqPerson,
	// )

	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)
	log.Println("== == == == == == == == == == ")

	res := toController.ViaGRPC(
		ctx,
		*reqPerson,
	)

	if res.GetError() != nil {
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
