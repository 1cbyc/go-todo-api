package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	requestid "github.com/gin-contrib/requestid"

	gintimeout "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Logger middleware for structured logging
func Logger(logger zerolog.Logger) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.Info().
			Str("method", param.Method).
			Str("path", param.Path).
			Int("status", param.StatusCode).
			Dur("latency", param.Latency).
			Str("client_ip", param.ClientIP).
			Str("user_agent", param.Request.UserAgent()).
			Msg("HTTP Request")
		return ""
	})
}

// Recovery middleware with structured logging
func Recovery(logger zerolog.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			logger.Error().Str("error", err).Msg("Panic recovered")
		}
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Internal server error",
		})
	})
}

// CORS middleware configuration
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Request-ID"}
	config.ExposeHeaders = []string{"Content-Length", "X-Request-ID"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return cors.New(config)
}

// RequestID middleware
func RequestID() gin.HandlerFunc {
	return requestid.New()
}

// Timeout middleware
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return gintimeout.New(
		gintimeout.WithTimeout(timeout),
		gintimeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		gintimeout.WithResponse(func(c *gin.Context) {
			c.JSON(408, gin.H{
				"success": false,
				"message": "Request timeout",
			})
		}),
	)
}
