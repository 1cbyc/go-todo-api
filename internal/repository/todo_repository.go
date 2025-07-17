package repository

import (
	"context"
	"fmt"

	"github.com/1cbyc/go-todo-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TodoRepository defines the interface for todo data operations
type TodoRepository interface {
	Create(ctx context.Context, todo *models.Todo) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Todo, error)
	GetAll(ctx context.Context, page, perPage int) ([]models.Todo, int64, error)
	Update(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id uuid.UUID) error
	Toggle(ctx context.Context, id uuid.UUID) error
}

// todoRepository implements TodoRepository
type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new todo repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

// Create creates a new todo
func (r *todoRepository) Create(ctx context.Context, todo *models.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

// GetByID retrieves a todo by ID
func (r *todoRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("todo not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}
	return &todo, nil
}

// GetAll retrieves all todos with pagination
func (r *todoRepository) GetAll(ctx context.Context, page, perPage int) ([]models.Todo, int64, error) {
	var todos []models.Todo
	var total int64

	// Get total count
	if err := r.db.WithContext(ctx).Model(&models.Todo{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count todos: %w", err)
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Get todos with pagination
	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(perPage).
		Find(&todos).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get todos: %w", err)
	}

	return todos, total, nil
}

// Update updates a todo
func (r *todoRepository) Update(ctx context.Context, todo *models.Todo) error {
	result := r.db.WithContext(ctx).Save(todo)
	if result.Error != nil {
		return fmt.Errorf("failed to update todo: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}

// Delete deletes a todo by ID
func (r *todoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Todo{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete todo: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}

// Toggle toggles the completed status of a todo
func (r *todoRepository) Toggle(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Model(&models.Todo{}).
		Where("id = ?", id).
		Update("completed", gorm.Expr("NOT completed"))

	if result.Error != nil {
		return fmt.Errorf("failed to toggle todo: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}
