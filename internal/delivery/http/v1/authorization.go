package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}
}

// SignIn godoc
// @Summary Return JWT
// @Accept json
// @Produce json
// @Param input body SignInForm true "login and password"
// @Success 200 {object} domain.Token
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	// TODO: настроить 500 - no sql rows (если нет такого пароля и логина)
	var input SignInForm

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user := models.User{
		Username: input.Login,
		Password: input.Password,
	}
	token, err := h.authorizer.SignInReturnToken(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenObj := domain.Token{
		AuthToken: token,
	}

	c.JSON(http.StatusOK, tokenObj)
}

type SignInForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
