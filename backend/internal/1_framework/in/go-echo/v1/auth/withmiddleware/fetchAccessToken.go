package withmiddleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"backend/internal/1_framework/http_parameter"
	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/struct_object"
	"backend/pkg"
)

func fetchAccessToken(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := pkg.GetNewContext(
		c.Request().Context(),
		c.Response().Header().Get(echo.HeaderXRequestID),
	)

	v1Credential := http_parameter.V1Credential{}

	if err := c.Bind(&v1Credential); err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	credential := struct_object.NewCredential(
		&struct_object.NewCredentialArgs{
			ClientID:     v1Credential.ClientID,
			ClientSecret: v1Credential.ClientSecret,
		},
	)

	if credential.Err != nil {
		pkg.Logging(ctx, credential.Err)
		return c.JSON(
			http.StatusBadRequest,
			credential.Err,
		)
	}

	accessToken, err := toController.FetchAccessToken(
		ctx,
		*credential,
	)
	if err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, accessToken.Content.GetValue())
	log.Println("== == == == == == == == == == ")

	// if err != nil {
	// 	pkg.Logging(ctx, err)
	// 	return c.JSON(
	// 		http.StatusBadRequest,
	// 		err,
	// 	)
	// }

	// responseList := []http_parameter.V1CredentialParameter{}
	// for _, Credential := range CredentialList {
	// 	responseList = append(
	// 		responseList,
	// 		http_parameter.V1CredentialParameter{
	// 			ID:          &Credential.ID.Content.Value,
	// 			Name:        &Credential.Name.Content.Value,
	// 			MailAddress: &Credential.MailAddress.Content.Value,
	// 		},
	// 	)
	// }

	return c.JSON(
		http.StatusOK,
		nil,
	)
}
