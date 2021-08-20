package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) isUserIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	username, err := h.authorizer.ParseTokenReturnUsername(header)
	if err != nil {
		// Return 401 - unauth
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
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
