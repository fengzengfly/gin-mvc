package middleware

import (
	"bytes"
	"gin-mvc/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

// RequestLogger 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 记录请求body
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新写入 body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 记录响应
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 执行下一个中间件
		c.Next()

		// 请求耗时
		latency := time.Since(startTime)

		// 日志字段
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
		}

		// 记录请求体和响应体（可选）
		if len(requestBody) > 0 {
			fields = append(fields, zap.String("request_body", string(requestBody)))
		}

		if blw.body.Len() > 0 {
			fields = append(fields, zap.String("response_body", blw.body.String()))
		}

		// 根据状态码选择日志级别
		switch {
		case c.Writer.Status() >= 500:
			log.Error("Server Error", fields...)
		case c.Writer.Status() >= 400:
			log.Warn("Client Error", fields...)
		default:
			log.Info("Request", fields...)
		}
	}
}

// 自定义 ResponseWriter 以捕获响应内容
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
