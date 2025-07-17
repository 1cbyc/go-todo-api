package handlers

import (
	"strconv"

	"github.com/1cbyc/go-todo-api/internal/models"
	"github.com/1cbyc/go-todo-api/internal/services"
	"github.com/1cbyc/go-todo-api/pkg/response"
	"github.com/1cbyc/go-todo-api/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// TodoHandler handles HTTP requests for todo operations
type TodoHandler struct {
	service services.TodoService
	logger  zerolog.Logger
}

// NewTodoHandler creates a new todo handler
func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
		logger:  zerolog.Logger{},
	}
}

// Create handles POST /api/v1/todos
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.CreateTodoRequest true "Todo to create"
// @Success 201 {object} response.SuccessResponse{data=models.TodoResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var req models.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err)
		return
	}

	// Validate request
	if err := validator.Validate.Struct(req); err != nil {
		response.BadRequest(c, "Validation failed", err)
		return
	}

	todo, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to create todo")
		response.InternalServerError(c, "Failed to create todo", err)
		return
	}

	response.Created(c, "Todo created successfully", todo)
}

// GetByID handles GET /api/v1/todos/:id
// @Summary Get a todo by ID
// @Description Get a specific todo item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID" format(uuid)
// @Success 200 {object} response.SuccessResponse{data=models.TodoResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", err)
		return
	}

	todo, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error().Err(err).Str("id", idStr).Msg("Failed to get todo")
		response.NotFound(c, "Todo not found", err)
		return
	}

	response.OK(c, "Todo retrieved successfully", todo)
}

// GetAll handles GET /api/v1/todos
// @Summary Get all todos
// @Description Get all todo items with pagination
// @Tags todos
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param per_page query int false "Items per page" default(20)
// @Success 200 {object} response.SuccessResponse{data=models.TodoListResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos [get]
func (h *TodoHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	todos, err := h.service.GetAll(c.Request.Context(), page, perPage)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get todos")
		response.InternalServerError(c, "Failed to get todos", err)
		return
	}

	response.OK(c, "Todos retrieved successfully", todos)
}

// Update handles PUT /api/v1/todos/:id
// @Summary Update a todo
// @Description Update an existing todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID" format(uuid)
// @Param todo body models.UpdateTodoRequest true "Todo updates"
// @Success 200 {object} response.SuccessResponse{data=models.TodoResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos/{id} [put]
func (h *TodoHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", err)
		return
	}

	var req models.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err)
		return
	}

	// Validate request
	if err := validator.Validate.Struct(req); err != nil {
		response.BadRequest(c, "Validation failed", err)
		return
	}

	todo, err := h.service.Update(c.Request.Context(), id, &req)
	if err != nil {
		h.logger.Error().Err(err).Str("id", idStr).Msg("Failed to update todo")
		response.NotFound(c, "Todo not found", err)
		return
	}

	response.OK(c, "Todo updated successfully", todo)
}

// Delete handles DELETE /api/v1/todos/:id
// @Summary Delete a todo
// @Description Delete a todo item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID" format(uuid)
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		h.logger.Error().Err(err).Str("id", idStr).Msg("Failed to delete todo")
		response.NotFound(c, "Todo not found", err)
		return
	}

	response.OK(c, "Todo deleted successfully", nil)
}

// Toggle handles PATCH /api/v1/todos/:id/toggle
// @Summary Toggle todo completion status
// @Description Toggle the completed status of a todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID" format(uuid)
// @Success 200 {object} response.SuccessResponse{data=models.TodoResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /todos/{id}/toggle [patch]
func (h *TodoHandler) Toggle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.BadRequest(c, "Invalid todo ID", err)
		return
	}

	todo, err := h.service.Toggle(c.Request.Context(), id)
	if err != nil {
		h.logger.Error().Err(err).Str("id", idStr).Msg("Failed to toggle todo")
		response.NotFound(c, "Todo not found", err)
		return
	}

	response.OK(c, "Todo toggled successfully", todo)
}
