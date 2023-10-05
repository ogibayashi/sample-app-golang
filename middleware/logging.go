package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func JSONFormatter(params gin.LogFormatterParams) string {
	entry := map[string]interface{}{
		"timestamp":   params.TimeStamp.UnixMilli(),
		"status":      params.StatusCode,
		"latency":     params.Latency / time.Millisecond,
		"remote_addr": params.ClientIP,
		"method":      params.Method,
		"path":        params.Path,
		"error":       params.ErrorMessage,
	}

	// JSON形式でログを出力
	logJSON, _ := json.Marshal(entry)
	return fmt.Sprintln(string(logJSON))

}

func AccessLogMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(JSONFormatter)
}
