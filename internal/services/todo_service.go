package services

import (
	"context"
	"fmt"
	"math"

	"github.com/1cbyc/go-todo-api/internal/models"
	"github.com/1cbyc/go-todo-api/internal/repository"
	"github.com/google/uuid"
)

// TodoService defines the interface for todo business operations
type TodoService interface {
	Create(ctx context.Context, req *models.CreateTodoRequest) (*models.TodoResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.TodoResponse, error)
	GetAll(ctx context.Context, page, perPage int) (*models.TodoListResponse, error)
	Update(ctx context.Context, id uuid.UUID, req *models.UpdateTodoRequest) (*models.TodoResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Toggle(ctx context.Context, id uuid.UUID) (*models.TodoResponse, error)
}

// todoService implements TodoService
type todoService struct {
	repo repository.TodoRepository
}

// NewTodoService creates a new todo service
func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

// Create creates a new todo
func (s *todoService) Create(ctx context.Context, req *models.CreateTodoRequest) (*models.TodoResponse, error) {
	// Set default priority if not provided
	if req.Priority == "" {
		req.Priority = models.PriorityMedium
	}

	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Completed:   false,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	response := todo.ToResponse()
	return &response, nil
}

// GetByID retrieves a todo by ID
func (s *todoService) GetByID(ctx context.Context, id uuid.UUID) (*models.TodoResponse, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	response := todo.ToResponse()
	return &response, nil
}

// GetAll retrieves all todos with pagination
func (s *todoService) GetAll(ctx context.Context, page, perPage int) (*models.TodoListResponse, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}

	todos, total, err := s.repo.GetAll(ctx, page, perPage)
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	// Convert todos to responses
	responses := make([]models.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = todo.ToResponse()
	}

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	if totalPages == 0 {
		totalPages = 1
	}

	meta := models.Meta{
		Total:       total,
		Page:        page,
		PerPage:     perPage,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}

	return &models.TodoListResponse{
		Data: responses,
		Meta: meta,
	}, nil
}

// Update updates a todo
func (s *todoService) Update(ctx context.Context, id uuid.UUID, req *models.UpdateTodoRequest) (*models.TodoResponse, error) {
	// Get existing todo
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	// Update fields if provided
	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}
	if req.Priority != nil {
		todo.Priority = *req.Priority
	}
	if req.DueDate != nil {
		todo.DueDate = req.DueDate
	}

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}

	response := todo.ToResponse()
	return &response, nil
}

// Delete deletes a todo
func (s *todoService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}

// Toggle toggles the completed status of a todo
func (s *todoService) Toggle(ctx context.Context, id uuid.UUID) (*models.TodoResponse, error) {
	if err := s.repo.Toggle(ctx, id); err != nil {
		return nil, fmt.Errorf("failed to toggle todo: %w", err)
	}

	// Get updated todo
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated todo: %w", err)
	}

	response := todo.ToResponse()
	return &response, nil
}
