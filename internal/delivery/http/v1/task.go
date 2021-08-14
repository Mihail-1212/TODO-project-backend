package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

func (h *Handler) initTaskRoutes(api *gin.RouterGroup) {
	tasks := api.Group("/tasks", h.isUserIdentify)
	{
		tasks.GET("/", h.getTaskList)
		tasks.POST("/", h.createTask)
		tasks.GET("/:taskId", h.getTaskByID)
		tasks.PUT("/:taskId", h.updateTask)
		tasks.DELETE("/:taskId", h.deleteTask)
	}
}

// GetTaskList godoc
// @Summary Return task list of current user
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} Tasks
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tasks/ [get]
func (h *Handler) getTaskList(c *gin.Context) {

	var tasks Tasks
	var err error

	user, _ := h.getCurrentUser(c)

	tasks.Tasks, err = h.services.TaskService.GetTaskListOfUser(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary Create task and return it
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body domain.CreateTask true "Create task form"
// @Success 200 {object} domain.Task
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tasks/ [post]
func (h *Handler) createTask(c *gin.Context) {
	var input domain.CreateTask
	user, _ := h.getCurrentUser(c)

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	task := domain.Task{
		Name:        input.Name,
		Description: input.Description,
		DateStart:   input.DateStart,
		UserId:      user.Id,
	}

	newTask, err := h.services.TaskService.CreateTask(&task)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newTask)
}

// GetTaskByID godoc
// @Summary Return task by url id
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param taskId path int true "id of task"
// @Success 200 {object} domain.Task
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tasks/{taskId} [get]
func (h *Handler) getTaskByID(c *gin.Context) {
	user, _ := h.getCurrentUser(c)

	taskId, _ := strconv.Atoi(c.Param("taskId"))

	task, err := h.services.TaskService.GetUserTaskByID(taskId, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Create task and return it
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param taskId path int true "id of task"
// @Param input body domain.UpdateTask true "Update task form"
// @Success 200 {object} domain.Task
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tasks/{taskId} [put]
func (h *Handler) updateTask(c *gin.Context) {
	var input domain.UpdateTask
	user, _ := h.getCurrentUser(c)

	taskId, _ := strconv.Atoi(c.Param("taskId"))

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if taskId != input.Id {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task := domain.Task{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		DateStart:   input.DateStart,
		UserId:      user.Id,
	}

	err = h.services.TaskService.UpdateTaskOfUser(&task)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary Delete task and return bool value
// @Security ApiKeyAuth
// @Param taskId path int true "id of task"
// @Success 200
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/tasks/{taskId} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	user, _ := h.getCurrentUser(c)

	taskId, _ := strconv.Atoi(c.Param("taskId"))

	err := h.services.TaskService.DeleteUserTask(taskId, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

type Tasks struct {
	Tasks []domain.Task `json:"tasks"`
}
