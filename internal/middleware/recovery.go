package middleware

import (
	"gin-mvc/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查是否是网络断开错误
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				// 记录错误日志
				log.Error("Panic Recovery",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				// 返回错误响应
				if brokenPipe {
					_ = c.Error(err.(error))
					c.Abort()
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "Internal Server Error",
					})
				}
			}
		}()
		c.Next()
	}
}
