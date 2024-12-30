package person

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/1_framework/input/go-echo/http_parameter"
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
	ctx := pkg.GetNewContext(
		c.Request().Context(),
		c.Response().Header().Get(echo.HeaderXRequestID),
	)

	person := http_parameter.V1PersonParameter{}

	if err := c.Bind(&person); err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}
	log.Println("== rest get== == == == == == == == == ")
	log.Printf("%#v\n", person.MailAddress)
	log.Printf("%#v\n", &person.MailAddress)
	log.Printf("%#v\n", *person.MailAddress)
	log.Println("== == == ==  == == == == == == ")

	reqPerson := struct_object.NewPerson(
		&struct_object.NewPersonArgs{
			ID:          person.ID,
			Name:        person.Name,
			MailAddress: person.MailAddress,
		},
	)

	if reqPerson.Err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
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

	responseList := []http_parameter.V1PersonParameter{}
	for _, person := range personList {
		responseList = append(
			responseList,
			http_parameter.V1PersonParameter{
				ID:          &person.ID.Content.Value,
				Name:        &person.Name.Content.Value,
				MailAddress: &person.MailAddress.Content.Value,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		responseList,
	)
}
