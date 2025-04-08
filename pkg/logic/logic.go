package logic

import (
	"errors"
	"sync"
	"taskmanager/internal/generated/models"
	"taskmanager/internal/generated/restapi/operations/tasks"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

var (
	// taskMap stores tasks with UUID as key
	taskMap = make(map[strfmt.UUID]*models.Task)
	// mutex for concurrent access to taskMap
	mutex = &sync.RWMutex{}
)

// ErrTaskNotFound is returned when a task with the given ID doesn't exist
var ErrTaskNotFound = errors.New("task not found")

// CreateTask creates a new task with the provided details
func CreateTask(params tasks.CreateTaskParams) middleware.Responder {
	mutex.Lock()
	defer mutex.Unlock()

	// Generate a new UUID for the task
	id := strfmt.UUID(uuid.New().String())
	
	now := strfmt.DateTime(time.Now())
	
	// Set default status to "pending" if not provided
	status := "pending"
	if params.Task.Status != nil {
		status = *params.Task.Status
	}

	// Set default priority to "medium" if not provided
	priority := "medium"
	if params.Task.Priority != nil {
		priority = *params.Task.Priority
	}

	// Create new task from the request
	task := &models.Task{
		ID:          &id,
		Title:       params.Task.Title,
		Description: params.Task.Description,
		Status:      &status,
		Priority:    priority,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		DueDate:     params.Task.DueDate,
		Tags:        params.Task.Tags,
	}

	// Store the task
	taskMap[id] = task

	return tasks.NewCreateTaskCreated().WithPayload(&tasks.CreateTaskCreatedBody{
		Data: task,
	})
}

// DeleteTask deletes a task by ID
func DeleteTask(params tasks.DeleteTaskParams) middleware.Responder {
	mutex.Lock()
	defer mutex.Unlock()

	// Check if task exists
	if _, exists := taskMap[params.TaskID]; !exists {
		code := "404"
		msg := "Task not found"
		return tasks.NewDeleteTaskNotFound().WithPayload(&models.Error{
			Code:    &code,
			Message: &msg,
		})
	}

	// Delete the task
	delete(taskMap, params.TaskID)

	return tasks.NewDeleteTaskNoContent()
}

// GetTask retrieves a task by ID
func GetTask(params tasks.GetTaskParams)middleware.Responder {
	mutex.RLock()
	defer mutex.RUnlock()

	// Check if task exists
	task, exists := taskMap[params.TaskID]
	if !exists {
		code := "404"
		msg := "Task not found"
		return tasks.NewGetTaskNotFound().WithPayload(&models.Error{
			Code:    &code,
			Message: &msg,
		})
	}

	return tasks.NewGetTaskOK().WithPayload(&tasks.GetTaskOKBody{
		Data: task,
	})
}

// ListTasks returns all tasks, with optional filtering
func ListTasks(params tasks.ListTasksParams) middleware.Responder {
	mutex.RLock()
	defer mutex.RUnlock()

	result := []*models.Task{}

	// Apply filters if they exist
	for _, task := range taskMap {
		// Filter by status if provided
		if params.Status != nil && *task.Status != *params.Status {
			continue
		}
		

		// Add task to result list
		result = append(result, task)
	}

	return tasks.NewListTasksOK().WithPayload(&tasks.ListTasksOKBody{
		Data: result,
	})
}

// UpdateTask updates an existing task
func UpdateTask(params tasks.UpdateTaskParams) middleware.Responder {
	mutex.Lock()
	defer mutex.Unlock()

	// Check if task exists
	task, exists := taskMap[params.TaskID]
	if !exists {
		code := "404"
		msg := "Task not found"
		return tasks.NewUpdateTaskNotFound().WithPayload(&models.Error{
			Code:    &code,
			Message: &msg,
		})
	}

	// Update the task fields
	if params.Task.Title != "" {
		task.Title = &params.Task.Title
	}

	if params.Task.Description != "" {
		task.Description = params.Task.Description
	}

	if params.Task.Status != "" {
		status := params.Task.Status
		task.Status = &status
	}

	if params.Task.Priority != "" {
		task.Priority = params.Task.Priority
	}

	if !params.Task.DueDate.IsZero() {
		task.DueDate = params.Task.DueDate
	}

	if params.Task.Tags != nil {
		task.Tags = params.Task.Tags
	}

	// Update the UpdatedAt timestamp
	now := strfmt.DateTime(time.Now())
	task.UpdatedAt = &now

	return tasks.NewUpdateTaskOK().WithPayload(&tasks.UpdateTaskOKBody{
		Data: task,
	})
}

