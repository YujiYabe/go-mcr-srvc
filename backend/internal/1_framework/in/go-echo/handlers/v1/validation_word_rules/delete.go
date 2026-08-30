package validation_word_rules

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Delete(
	echoContext echo.Context,
	toController controller.ToController,
) (
	err error,
) {
	var request openapi.ValidationWordRuleDelete
	if err := echoContext.Bind(&request); err != nil {
		return errorJSON(echoContext, http.StatusBadRequest, fmt.Errorf("invalid request"))
	}
	if err := validateTargetTypeAndWord(request.TargetType, request.Word); err != nil {
		return errorJSON(echoContext, http.StatusBadRequest, err)
	}

	if err := toController.DeleteValidationWord(
		echoContext.Request().Context(),
		request.TargetType,
		request.IsBlacklist,
		request.Word,
	); err != nil {
		logger.Logging(echoContext.Request().Context(), err)
		return errorJSON(echoContext, http.StatusBadRequest, err)
	}

	return echoContext.NoContent(http.StatusNoContent)
}
