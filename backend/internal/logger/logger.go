package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	valueObject "backend/internal/4_domain/value_object"
)

type CustomJSONFormatter struct {
	logrus.JSONFormatter
}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Data == nil {
		entry.Data = make(logrus.Fields)
	}

	// Create a new map to control field order
	orderedData := logrus.Fields{
		"time": entry.Time.Format("15:04:05"), // Add timestamp first
	}

	// Add level and message explicitly
	orderedData["level"] = entry.Level.String()
	orderedData["msg"] = entry.Message

	// Copy other fields into the ordered map
	for key, value := range entry.Data {
		orderedData[key] = value
	}

	// Use the custom map for JSON formatting
	buffer := &strings.Builder{}
	encoder := f.JSONFormatter
	entry.Data = orderedData
	encoded, err := encoder.Format(entry)
	if err != nil {
		return nil, err
	}

	buffer.Write(encoded)
	return []byte(buffer.String()), nil
}

func init() {
	logrus.SetFormatter(&CustomJSONFormatter{
		JSONFormatter: logrus.JSONFormatter{
			DisableTimestamp: true, // Disable default timestamp
		},
	})
}

func Logging(
	ctx context.Context,
	data interface{},
) {
	_, fullPath, line, _ := runtime.Caller(1)
	trimPath := strings.TrimPrefix(fullPath, "/go/src/backend/")

	logger := logrus.WithFields(logrus.Fields{
		"file": fmt.Sprintf("%s:%d", trimPath, line),
	})

	traceID, ok := ctx.Value(valueObject.TraceIDContextName).(string)
	if ok {
		logger = logger.WithField("traceID", traceID)
	}

	switch v := data.(type) {
	case error:
		logger.Error(v)
	default:
		logger.Info(data)
	}
}
