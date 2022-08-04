package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("log : %v  [%v] %v %v %v %v \n",
	params.ClientIP,
	params.TimeStamp.Format(time.RFC822),
	params.Method,
	params.Path,
	params.StatusCode,
	params.Latency,
)
	})
}