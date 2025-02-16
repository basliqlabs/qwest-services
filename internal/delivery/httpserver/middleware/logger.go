package middleware

import (
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// Logger returns a middleware that logs HTTP requests using zap logger
func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// TODO: add config to check if logging is enabled
			errorMessage := ""
			operation := ""
			if v.Error != nil {
				re, ok := v.Error.(*richerror.RichError)
				if ok {
					errorMessage = re.GetMessage()
					operation = re.GetOperation()
				} else {
					errorMessage = v.Error.Error()
				}
			}

			fields := []zap.Field{
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content_length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error_message", errorMessage),
				zap.String("operation", operation),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			}

			if v.Status >= 500 {
				logger.L().Named("http-server").Error("request", fields...)
			} else {
				logger.L().Named("http-server").Info("request", fields...)
			}

			return nil
		},
	})
}
