package withmiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	httpParameter "backend/internal/1_framework/parameter/http"
	"backend/internal/2_adapter/controller"
	groupObject "backend/internal/4_domain/group_object"
	"backend/internal/logger"
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
		logger.Logging(ctx, err)
		return c.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	//-------------------------
	credential := groupObject.NewCredential(
		ctx,
		&groupObject.NewCredentialArgs{
			ClientID:     v1Credential.ClientID,
			ClientSecret: v1Credential.ClientSecret,
		},
	)
	if credential.GetError() != nil {
		logger.Logging(ctx, credential.GetError())
		return c.JSON(
			http.StatusBadRequest,
			credential.GetError(),
		)
	}

	//-------------------------
	accessToken := toController.FetchAccessToken(
		ctx,
		*credential,
	)
	if accessToken.GetError() != nil {
		logger.Logging(ctx, accessToken.GetError())
		return c.JSON(
			http.StatusBadRequest,
			accessToken.GetError(),
		)
	}

	return c.JSON(
		http.StatusOK,
		accessToken.GetValue(),
	)
}
