package validation_word_rules

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Put(
	echoContext echo.Context,
	toController controller.ToController,
) error {
	var request openapi.ValidationWordRuleUpdate
	if err := echoContext.Bind(&request); err != nil {
		return errorJSON(echoContext, http.StatusBadRequest, fmt.Errorf("invalid request"))
	}
	if err := validateTargetTypeAndWord(request.TargetType, request.OldWord); err != nil {
		return errorJSON(echoContext, http.StatusBadRequest, err)
	}
	if err := validateWord(request.NewWord); err != nil {
		return errorJSON(echoContext, http.StatusBadRequest, err)
	}

	if err := toController.UpdateValidationWord(
		echoContext.Request().Context(),
		request.TargetType,
		request.IsBlacklist,
		request.OldWord,
		request.NewWord,
	); err != nil {
		logger.Logging(echoContext.Request().Context(), err)
		return errorJSON(echoContext, http.StatusBadRequest, err)
	}

	return echoContext.NoContent(http.StatusNoContent)
}
