package withmiddleware

import (
	"net/http"

	"github.com/labstack/echo"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	structObject "backend/internal/4_domain/struct_object"
	"backend/pkg"
)

func fetchAccessToken(
	c echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	ctx := c.Request().Context()

	v1Credential := httpParameter.V1Credential{}

	if err := c.Bind(&v1Credential); err != nil {
		pkg.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	credential := structObject.NewCredential(
		&structObject.NewCredentialArgs{
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

	return c.JSON(
		http.StatusOK,
		accessToken.Content.GetValue(),
	)
}
