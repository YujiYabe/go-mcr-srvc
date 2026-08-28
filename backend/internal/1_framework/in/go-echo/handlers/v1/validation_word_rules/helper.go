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
) error {
	if err := validateTargetType(targetType); err != nil {
		return err
	}
	return validateWord(word)
}

func validateTargetType(
	targetType string,
) error {
	if strings.TrimSpace(targetType) == "" {
		return fmt.Errorf("targetType is required")
	}
	return nil
}

func validateWord(
	word string,
) error {
	if strings.TrimSpace(word) == "" {
		return fmt.Errorf("word is required")
	}
	return nil
}

func errorJSON(
	echoContext echo.Context,
	status int,
	err error,
) error {
	return echoContext.JSON(status, openapi.Error{
		Code:    status,
		Message: err.Error(),
	})
}
