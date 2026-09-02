package validation_word_rules

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	"backend/internal/1_framework/in/go-echo/openapi"
)

func validateTargetTypeAndWord(
	targetType string,
	word string,
) (
	err error,
) {
	if returnedErr := validateTargetType(targetType); returnedErr != nil {
		err = returnedErr
		return
	}
	err = validateWord(word)
	return
}

func validateTargetType(
	targetType string,
) (
	err error,
) {
	if strings.TrimSpace(targetType) == "" {
		err = fmt.Errorf("targetType is required")
		return
	}
	err = nil
	return
}

func validateWord(
	word string,
) (
	err error,
) {
	if strings.TrimSpace(word) == "" {
		err = fmt.Errorf("word is required")
		return
	}
	err = nil
	return
}

func errorJSON(
	echoContext echo.Context,
	status int,
	err error,
) (
	errResult error,
) {
	errResult = echoContext.JSON(status, openapi.Error{
		Code:    status,
		Message: err.Error(),
	})
	return
}
