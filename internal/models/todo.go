package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Todo represents a todo item
type Todo struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	Title       string         `json:"title" gorm:"not null;size:255" validate:"required,min=1,max=255"`
	Description string         `json:"description" gorm:"size:1000"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	Priority    Priority       `json:"priority" gorm:"default:medium"`
	DueDate     *time.Time     `json:"due_date,omitempty"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// Priority represents the priority level of a todo
type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

// TableName specifies the table name for Todo
func (Todo) TableName() string {
	return "todos"
}

// BeforeCreate is called before creating a new todo
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description" validate:"max=1000"`
	Priority    Priority   `json:"priority" validate:"oneof=low medium high urgent"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// UpdateTodoRequest represents the request body for updating a todo
type UpdateTodoRequest struct {
	Title       *string    `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string    `json:"description,omitempty" validate:"omitempty,max=1000"`
	Completed   *bool      `json:"completed,omitempty"`
	Priority    *Priority  `json:"priority,omitempty" validate:"omitempty,oneof=low medium high urgent"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// TodoResponse represents the response body for todo operations
type TodoResponse struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Priority    Priority   `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ToResponse converts a Todo to TodoResponse
func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		Priority:    t.Priority,
		DueDate:     t.DueDate,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// TodoListResponse represents the response for listing todos
type TodoListResponse struct {
	Data []TodoResponse `json:"data"`
	Meta Meta           `json:"meta"`
}

// Meta represents metadata for paginated responses
type Meta struct {
	Total       int64 `json:"total"`
	Page        int   `json:"page"`
	PerPage     int   `json:"per_page"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}
