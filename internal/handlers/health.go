package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HealthCheck handles GET /health
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "Todo API is running",
		"version":   "1.0.0",
		"timestamp": time.Now().UTC(),
		"uptime":    time.Since(startTime).String(),
		"memory": gin.H{
			"alloc":       m.Alloc,
			"total_alloc": m.TotalAlloc,
			"sys":         m.Sys,
			"num_gc":      m.NumGC,
		},
		"goroutines": runtime.NumGoroutine(),
	})
}

// Metrics handles GET /api/v1/metrics
// @Summary Prometheus metrics
// @Description Get Prometheus metrics
// @Tags metrics
// @Accept json
// @Produce text/plain
// @Success 200 {string} string
// @Router /metrics [get]
func Metrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

// SwaggerHandler handles Swagger documentation
func SwaggerHandler(c *gin.Context) {
	// This will be handled by gin-swagger middleware
	c.Next()
}

var startTime = time.Now()
