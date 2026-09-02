package validation_word_rules

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
	"backend/internal/2_adapter/controller"
	"backend/internal/logger"
)

func Get(
	echoContext echo.Context,
	toController controller.ToController,
	params openapi.V1ValidationWordRulesGetParams,
) (
	err error,
) {
	if returnedErr := validateTargetType(params.TargetType); returnedErr != nil {
		err = errorJSON(echoContext, http.StatusBadRequest, returnedErr)
		return
	}

	words, returnedErr := toController.GetValidationWords(
		echoContext.Request().Context(),
		params.TargetType,
		params.IsBlacklist,
	)
	if returnedErr != nil {
		logger.Logging(echoContext.Request().Context(), returnedErr)
		err = errorJSON(echoContext, http.StatusBadRequest, returnedErr)
		return
	}

	response := make([]openapi.ValidationWordRule, 0, len(words))
	for _, word := range words {
		response = append(response, openapi.ValidationWordRule{
			TargetType:  params.TargetType,
			IsBlacklist: params.IsBlacklist,
			Word:        word,
		})
	}

	err = echoContext.JSON(http.StatusOK, response)
	return
}
