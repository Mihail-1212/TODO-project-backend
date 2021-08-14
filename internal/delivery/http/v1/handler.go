package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mihail-1212/todo-project-backend/internal/service"
	"github.com/mihail-1212/todo-project-backend/pkg/auth"
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
	"github.com/mihail-1212/todo-project-backend/pkg/logger"
)

var log, _ = logger.NewLoggerDev()

const (
	authorizationHeader = "Authorization"
	userCtx             = "currentUser"
)

type Handler struct {
	services   *service.Services
	authorizer *auth.Authorizer
}

func NewHandler(services *service.Services, authorizer *auth.Authorizer) *Handler {
	return &Handler{
		services:   services,
		authorizer: authorizer,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {

	v1 := api.Group("/v1")
	{
		// v1.Use(h.corsMiddleware)	// ?
		h.initTaskRoutes(v1)
		h.initAuthRoutes(v1)
	}
}

func (h *Handler) isUserIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	username, err := h.authorizer.ParseTokenReturnUsername(header)
	if err != nil {
		// TODO: сделать возвращение заголовка о необходимости перехода на страницу авторизации?
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Получение пользователя
	user, err := h.services.UserService.GetUserByLogin(username)

	if err != nil {
		// TODO: сделать empty sql row
		newErrorResponse(c, http.StatusUnauthorized, "Error token")
		return
	}

	c.Set(userCtx, user)
}

func (h *Handler) getCurrentUser(c *gin.Context) (*domain.User, error) {
	userInterface, isUserFound := c.Get(userCtx)
	if !isUserFound {
		return nil, UserContextIsEmptyError{"User context was not found! User is not auth."}
	}
	user, _ := userInterface.(*domain.User)

	return user, nil
}

func (h *Handler) SignIn(login, password string) (string, error) {
	user := models.User{
		Username: login,
		Password: password,
	}
	return h.authorizer.SignInReturnToken(&user)
}

type UserContextIsEmptyError struct {
	Message string
}

func (e UserContextIsEmptyError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
