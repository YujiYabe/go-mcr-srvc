package validation_word_rules

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Post(
	echoContext echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	var request openapi.ValidationWordRuleCreate
	if returnedErr := echoContext.Bind(&request); returnedErr != nil {
		err = errorJSON(echoContext, http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}
	if returnedErr := validateTargetTypeAndWord(request.TargetType, request.Word); returnedErr != nil {
		err = errorJSON(echoContext, http.StatusBadRequest, returnedErr)
		return
	}

	if returnedErr := toController.AddValidationWord(
		echoContext.Request().Context(),
		request.TargetType,
		request.IsBlacklist,
		request.Word,
	); returnedErr != nil {
		logger.Logging(echoContext.Request().Context(), returnedErr)
		err = errorJSON(echoContext, http.StatusBadRequest, returnedErr)
		return
	}

	err = echoContext.JSON(
		http.StatusCreated,
		openapi.ValidationWordRule(request),
	)
	return
}
