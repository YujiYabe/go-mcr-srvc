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
	traceID := valueObject.GetTraceID(ctx)
	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

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

	resPersonList := toController.ViaGRPC(
		ctx,
		*reqPerson,
	)

	if resPersonList.GetError() != nil {
		pkg.Logging(ctx, resPersonList.GetError())
		return c.JSON(
			http.StatusBadRequest,
			resPersonList.GetError(),
		)
	}

	responseList := []httpParameter.V1Person{}
	for _, person := range resPersonList.Content {
		id := person.ID.GetValue()
		name := person.Name.GetValue()
		mailAddress := person.MailAddress.GetValue()
		responseList = append(
			responseList,
			httpParameter.V1Person{
				ID:          &id,
				Name:        &name,
				MailAddress: &mailAddress,
			},
		)
	}

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, traceID)

	return c.JSON(
		http.StatusOK,
		responseList,
	)

}
