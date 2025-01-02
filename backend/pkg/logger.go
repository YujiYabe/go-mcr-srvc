package pkg

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type RowData []string

func (receiver *RowData) append(appendData string) {
	*receiver = append(*receiver, appendData)
}

func Logging(
	ctx context.Context,
	data interface{},
) {
	requestID := ""
	if val, ok := ctx.Value(RequestIDKey).(string); ok {
		requestID = val
	}
	_, fullPath, line, _ := runtime.Caller(1)

	fullPath = strings.TrimPrefix(fullPath, "/go/src")
	var rowData RowData
	rowData.append(requestID)                       // key
	rowData.append(time.Now().Format("15:04:05"))   // dateTime
	rowData.append(fmt.Sprint(fullPath, ":", line)) // ファイル名:行番号

	logContent := ""
	if errorContent, ok := data.(error); ok {
		logContent = errorContent.Error() // エラー内容
	} else {
		logContent = fmt.Sprintf("%#v", data)
	}
	rowData.append(logContent)

	row := strings.Join(rowData, "  ") // タブ区切り

	log.Printf("%s\n", row)
}
