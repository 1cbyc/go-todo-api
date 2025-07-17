package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// OK sends a 200 OK response
func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 Created response
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// Forbidden sends a 403 Forbidden response
func Forbidden(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusForbidden, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// Conflict sends a 409 Conflict response
func Conflict(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusConflict, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string, err interface{}) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
