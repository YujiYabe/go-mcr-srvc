package logger

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"

	middlewareRequestContext "backend/internal/1_framework/middleware/request_context"
	primitiveObject "backend/internal/4_domain/primitive_object"
	typeObject "backend/internal/4_domain/type_object"
)

const (
	envLocal   = "local"
	envLCL     = "lcl"
	envProd    = "prod"
	envDefault = envLocal
)

func init() {
	zerolog.TimeFieldFormat = "15:04:05"
	zerolog.SetGlobalLevel(logLevel())
	log.SetFlags(0)
}

func Logging(
	ctx context.Context,
	data interface{},
) {
	// テスト中であればロギングしない
	if flag.Lookup("test.v") != nil {
		log.Println("run under go test")
		return
	}

	_, fullPath, line, _ := runtime.Caller(1)
	trimPath := fullPath
	if idx := strings.Index(fullPath, "internal/"); idx != -1 {
		trimPath = fullPath[idx:]
	}

	logger := zerolog.New(logWriter()).
		With().
		Timestamp().
		Str("file", fmt.Sprintf("%s:%d", trimPath, line))

	logger = appendContextFields(ctx, logger)

	event := logger.Logger()

	switch typedData := data.(type) {
	case error:
		event.Error().Err(typedData).Msg("error occurred")
	default:
		event.Info().Interface("data", data).Msg("")
	}
}

func logWriter() (
	writer io.Writer,
) {
	env := normalizedEnv()
	if env == envLocal {
		return zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFieldFormat,
		}
	}

	return os.Stdout
}

func logLevel() (
	level zerolog.Level,
) {
	env := normalizedEnv()
	if env == envProd {
		return zerolog.InfoLevel
	}

	return zerolog.DebugLevel
}

func normalizedEnv() (
	env string,
) {
	env = strings.ToLower(
		strings.TrimSpace(
			os.Getenv("ENV"),
		),
	)
	if env == "" || env == envLCL {
		return envDefault
	}

	return env
}

func appendContextFields(
	ctx context.Context,
	loggerContext zerolog.Context,
) (
	updatedLoggerContext zerolog.Context,
) {
	updatedLoggerContext = loggerContext

	requestContext := middlewareRequestContext.GetRequestContext(ctx)
	if requestContext != nil {
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"traceID",
			requestContext.TraceID().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"userID",
			requestContext.UserID().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"tenantID",
			requestContext.TenantID().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"clientIP",
			requestContext.ClientIP().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"userAgent",
			requestContext.UserAgent().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"locale",
			requestContext.Locale().GetValue(),
		)
		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			"timeZone",
			requestContext.TimeZone().GetValue(),
		)

		return updatedLoggerContext
	}

	return appendContextValueFields(ctx, updatedLoggerContext)
}

func appendContextValueFields(
	ctx context.Context,
	loggerContext zerolog.Context,
) (
	updatedLoggerContext zerolog.Context,
) {
	updatedLoggerContext = loggerContext

	contextFields := map[string]primitiveObject.ContextKey{
		"traceID":   typeObject.TraceIDContextName,
		"userID":    typeObject.UserIDContextName,
		"tenantID":  typeObject.TenantIDContextName,
		"clientIP":  typeObject.ClientIPContextName,
		"userAgent": typeObject.UserAgentContextName,
		"locale":    typeObject.LocaleContextName,
		"timeZone":  typeObject.TimeZoneContextName,
	}

	for fieldName, contextName := range contextFields {
		fieldValue, ok := ctx.Value(contextName).(string)
		if !ok {
			continue
		}

		updatedLoggerContext = appendStringField(
			updatedLoggerContext,
			fieldName,
			fieldValue,
		)
	}

	return updatedLoggerContext
}

func appendStringField(
	loggerContext zerolog.Context,
	fieldName string,
	fieldValue string,
) (
	updatedLoggerContext zerolog.Context,
) {
	if fieldValue == "" {
		return loggerContext
	}

	return loggerContext.Str(fieldName, fieldValue)
}
