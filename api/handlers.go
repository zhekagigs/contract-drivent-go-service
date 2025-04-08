package api 

import (
	"github.com/go-openapi/runtime/middleware"
	"taskmanager/internal/generated/restapi/operations"
	"taskmanager/internal/generated/restapi/operations/tasks"
	"taskmanager/pkg/logic"
)

// RegisterHandlers connects our business logic to the generated API
func RegisterHandlers(api *operations.TaskmanagerAPI) {
	// Implement each operation
	api.TasksCreateTaskHandler = tasks.CreateTaskHandlerFunc(CreateTask)
	api.TasksDeleteTaskHandler = tasks.DeleteTaskHandlerFunc(DeleteTask)
	api.TasksGetTaskHandler = tasks.GetTaskHandlerFunc(GetTask)
	api.TasksListTasksHandler = tasks.ListTasksHandlerFunc(ListTasks)
	api.TasksUpdateTaskHandler = tasks.UpdateTaskHandlerFunc(UpdateTask)
}

// CreateTask creates a new task
func CreateTask(params tasks.CreateTaskParams) middleware.Responder {
	// Delegate to logic package
	return logic.CreateTask(params)
}

// DeleteTask deletes a task by ID
func DeleteTask(params tasks.DeleteTaskParams) middleware.Responder {
	// Delegate to logic package
	return logic.DeleteTask(params)
}

// GetTask returns a task by ID
func GetTask(params tasks.GetTaskParams) middleware.Responder {
	// Delegate to logic package
	return logic.GetTask(params)
}

// ListTasks returns a list of tasks
func ListTasks(params tasks.ListTasksParams) middleware.Responder {
	// Delegate to logic package
	return logic.ListTasks(params)
}

// UpdateTask updates a task
func UpdateTask(params tasks.UpdateTaskParams) middleware.Responder {
	// Delegate to logic package
	return logic.UpdateTask(params)
}
