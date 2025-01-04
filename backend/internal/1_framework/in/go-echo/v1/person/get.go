package person

import (
	"net/http"

	"github.com/labstack/echo"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/struct_object"
	"backend/pkg"
)

func get(
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

	reqPerson := struct_object.NewPerson(
		&struct_object.NewPersonArgs{
			ID:          person.ID,
			Name:        person.Name,
			MailAddress: person.MailAddress,
		},
	)

	if reqPerson.Err != nil {
		pkg.Logging(ctx, reqPerson.Err)
		return c.JSON(
			http.StatusBadRequest,
			reqPerson.Err,
		)
	}

	personList, err := toController.GetPersonByCondition(
		ctx,
		*reqPerson,
	)

	if err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	responseList := []httpParameter.V1Person{}
	for _, person := range personList {
		id := person.ID.Content.GetValue()
		name := person.Name.Content.GetValue()
		mailAddress := person.MailAddress.Content.GetValue()
		responseList = append(
			responseList,
			httpParameter.V1Person{
				ID:          &id,
				Name:        &name,
				MailAddress: &mailAddress,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		responseList,
	)

}
